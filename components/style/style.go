package style

import (
	"regexp"
	"strings"
	"unicode"

	"maps"
	"slices"
)

const customSelf = "self"

type Opt func(string) string

// Set sets/replace the class attribute.
func Set(set string) Opt {
	return func(_ string) string {
		return set
	}
}

// Add adds CSS classes.
func Add(add string) Opt {
	return func(d string) string {
		if add == "" {
			return d
		}
		if d != "" {
			d += " "
		}
		return d + add
	}
}

// ReplaceVariants replace a Tailwind class and its variants from the class attribute.
//
//	// replaces border-[...], [...]:border-[...]
//	style.ReplaceVariants("border", "border-none")
func ReplaceVariants(pattern, replace string) Opt {
	rePattern := regexp.MustCompile(`\b(?:[^ ]+:)?` + pattern + `(?:-[^ ]+)?\b`)
	return func(d string) string {
		d = strings.TrimSpace(rePattern.ReplaceAllString(d, ""))
		d = trimSpaces(d)
		if d != "" {
			d += " "
		}
		d += replace
		return d
	}
}

// Replace replaces a CSS class.
func Replace(old, new string) Opt {
	return func(d string) string {
		d = strings.ReplaceAll(d, old, new)
		return d
	}
}

// Remove removes.
func Remove(remove ...string) Opt {
	return func(d string) string {
		for _, r := range remove {
			d = strings.ReplaceAll(d, r, "")
		}
		return d
	}
}

func trimSpaces(s string) string {
	prevSpace := true
	var ns []rune
	for _, r := range s {
		switch {
		case !unicode.IsSpace(r):
			ns = append(ns, r)
			prevSpace = false
		case prevSpace:
			continue
		default:
			ns = append(ns, r)
			prevSpace = true
		}
	}
	return string(ns)
}

// D defines a style as a list of Opt (Class/Add/ReplaceXXX calls).
type D []Opt

func (def D) apply(d string) string {
	for _, opt := range def {
		d = opt(d)
	}
	return d
}

func defs(defaults Defaults, style Style, k string) D {
	defs := defaults[k][StyleDefault]
	for _, dstyle := range slices.Sorted(maps.Keys(defaults[k])) {
		if dstyle&style == 0 {
			continue
		}
		defs = append(defs, defaults[k][dstyle]...)
	}
	return defs
}

func (def D) String() string {
	return def.apply("")
}

func Compute(style Style, k string, custom Custom) D {
	d := defs(defaults, style, k)
	s := defs(skin, style, k).String()
	d = append(d, Add(s))
	if !strings.Contains(k, "/") {
		d = append(d, custom[customSelf]...)
	}
	return append(d, custom[k]...)
}

func CSSClass(style Style, k string, custom Custom) string {
	return Compute(style, k, custom).String()
}

// Delta returns all classes added by dst.
func Delta(src, dst Style, k string, custom Custom) string {
	ssrc := strings.Split(CSSClass(src, k, custom), " ")
	sdst := strings.Split(CSSClass(dst, k, custom), " ")
	delta := make([]string, 0, len(sdst))
	for _, s := range sdst {
		if !slices.Contains(ssrc, s) {
			delta = append(delta, s)
		}
	}
	return strings.Join(delta, " ")
}

// Style defines a style.
type Style int

const (
	// StyleDefault is the default style.
	StyleDefault Style = 0
	// StyleDisabled is automatically added when a component has a set Disabled attribute.
	StyleDisabled Style = 1
	// StyleValid is supported by input attributes. Text/background/border switch to green.
	StyleValid Style = 1 << 1
	// StyleInvalid is supported by input attributes. Text/background/border switch to red.
	StyleInvalid Style = 1 << 2
)

// Defaults defines the default styles for a component.
type Defaults map[string]map[Style]D

var defaults = Defaults{}

func SetDefaults(d Defaults) {
	for c, cdefs := range d {
		if defaults[c] == nil {
			defaults[c] = map[Style]D{}
		}
		for style, v := range cdefs {
			defaults[c][style] = v
		}
	}
}

func CopyDefaults(src, dst string) {
	defaults[dst] = defaults[src]
}

var skin = Defaults{}

func SetSkin(d Defaults) {
	skin = d
}

// Custom defines custom styles for a component.
type Custom map[string]D

func (c Custom) AddBefore(cc Custom) Custom {
	if cc == nil {
		return c
	}
	if c == nil {
		c = make(Custom, len(cc))
	}
	for k, v := range cc {
		c[k] = append(c[k], v...)
	}
	return c
}

// CustomSet sets css classes as a CustomStyle for the current component.
// It is the equivalent of style.Custom{"component name":{style.Set(class)})
func CustomSet(class string) Custom {
	return Custom{
		customSelf: {Set(class)},
	}
}

// CustomAdd adds css classes as a CustomStyle for the current component.
// It is the equivalent of style.Custom{"component name":{style.Add(class)})
func CustomAdd(class string) Custom {
	return Custom{
		customSelf: {Add(class)},
	}
}

// CustomReplace replaces css classes as a CustomStyle for the current component.
// It is the equivalent of style.Custom{"component name":{style.Replace(old, new)})
func CustomReplace(old, new string) Custom {
	return Custom{
		customSelf: {Replace(old, new)},
	}
}

// CustomReplaceVariants adds css classes as a CustomStyle for the current component.
// It is the equivalent of style.Custom{"component name":{style.ReplaceVariants(class)})
func CustomReplaceVariants(pattern, replace string) Custom {
	return Custom{
		customSelf: {ReplaceVariants(pattern, replace)},
	}
}
