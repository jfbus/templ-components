package style

import (
	"regexp"
	"strings"
	"unicode"

	"slices"
)

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

// Replace replaces both in class and color.
func Replace(old, new string) Opt {
	return func(d string) string {
		d = strings.ReplaceAll(d, old, new)
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
	for dstyle, ddefaults := range defaults[k] {
		if dstyle&style == 0 {
			continue
		}
		defs = append(defs, ddefaults...)
	}
	return defs
}

func (def D) String() string {
	return def.apply("")
}

// WithDefault returns the D including defaults.
func (def D) WithDefault(style Style, k string) D {
	d := defs(defaults, style, k)
	d = append(d, def...)
	s := defs(skin, style, k).String()
	return append(d, Add(s))
}

// CSSClass returns the CSS class string.
func (def D) CSSClass(style Style, k string) string {
	d := defs(defaults, style, k).apply("")
	s := defs(skin, style, k).apply("")
	if d != "" && s != "" {
		return def.apply(d + " " + s)
	}
	return def.apply(d + s)
}

// Delta returns all classes added by dst.
func (def D) Delta(src, dst Style, k string) string {
	ssrc := strings.Split(def.CSSClass(src, k), " ")
	sdst := strings.Split(def.CSSClass(dst, k), " ")
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

// Defaults define the default styles for a component.
type Defaults map[string]map[Style]D

var defaults = Defaults{}
var skin = Defaults{}

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

func SetSkin(d Defaults) {
	skin = d
}
