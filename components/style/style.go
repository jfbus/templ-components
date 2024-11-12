package style

import (
	"regexp"
	"strings"
	"unicode"

	"github.com/jfbus/templui/components/size"
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
	defs := defaults[k][Default]
	for _, dstyle := range slices.Sorted(maps.Keys(defaults[k])) {
		if dstyle&style == 0 || dstyle&style != dstyle {
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
// - 1 up to 1 << 5 are reserved for common styles
// - 1 << 6 up to 1 << 15 are reserved for component styles
// - 1 << 16 up to 1 << 29 are reserved for size styles
type Style int

const (
	// Default is the default style.
	Default Style = 0
	// Disabled is automatically added when a component has a set Disabled attribute.
	Disabled Style = 1
	// Valid is supported by input attributes. Text/background/border switch to green.
	Valid Style = 1 << 1
	// Invalid is supported by input attributes. Text/background/border switch to red.
	Invalid Style = 1 << 2

	SizeXS      Style = 1 << 16
	SizeS       Style = 1 << 17
	SizeNormal  Style = 1 << 18
	SizeL       Style = 1 << 19
	SizeXL      Style = 1 << 20
	SizeTwoXL   Style = 1 << 21
	SizeThreeXL Style = 1 << 22
	SizeFourXL  Style = 1 << 23
	SizeFiveXL  Style = 1 << 24
	SizeSixXL   Style = 1 << 25
	SizeSevenXL Style = 1 << 26
	SizeEightXL Style = 1 << 27
	SizeNineXL  Style = 1 << 28
	SizeFull    Style = 1 << 29
)

// Defaults defines the default styles for a component.
type Defaults map[string]map[Style]D

var defaults = Defaults{}

func SetDefaults(d Defaults) {
	for c, cdefs := range d {
		if defaults[c] == nil {
			defaults[c] = make(map[Style]D, len(cdefs))
		}
		for style, v := range cdefs {
			defaults[c][style] = v
		}
	}
}

func CopyDefaults(src, dst string) {
	if defaults[dst] == nil {
		defaults[dst] = make(map[Style]D, len(defaults[src]))
	}
	for k, v := range defaults[src] {
		defaults[dst][k] = slices.Clone(v)
	}
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

var sizes = map[size.Size]Style{
	size.XS:      SizeXS,
	size.S:       SizeS,
	size.Normal:  SizeNormal,
	size.L:       SizeL,
	size.XL:      SizeXL,
	size.TwoXL:   SizeTwoXL,
	size.ThreeXL: SizeThreeXL,
	size.FourXL:  SizeFourXL,
	size.FiveXL:  SizeFiveXL,
	size.SixXL:   SizeSixXL,
	size.SevenXL: SizeSevenXL,
	size.EightXL: SizeEightXL,
	size.NineXL:  SizeNineXL,
	size.Full:    SizeFull,
}

func Size(sz size.Size) Style {
	return sizes[sz]
}
