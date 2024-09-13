package style

import (
	"regexp"
	"strings"
	"unicode"

	"slices"
)

type Opt func(d) d

// Color sets/replace the color attribute.
func Color(color string) Opt {
	return func(d d) d {
		d.Color = color
		return d
	}
}

// Class sets/replace the class attribute.
func Class(class string) Opt {
	return func(d d) d {
		d.Class = class
		return d
	}
}

// Add adds CSS classes.
func Add(add string) Opt {
	return func(d d) d {
		if d.Add != "" {
			d.Add += " "
		}
		d.Add += add
		return d
	}
}

// ReplaceClass replace a Tailwind class and its variants from the class attribute.
//
//	// replaces border-[...], [...]:border-[...]
//	style.ReplaceClass("border", "border-none")
func ReplaceClass(pattern, replace string) Opt {
	rePattern := regexp.MustCompile(`\b(?:[^ ]+:)?` + pattern + `(?:-[^ ]+)?\b`)
	return func(d d) d {
		d.Class = strings.TrimSpace(rePattern.ReplaceAllString(d.Class, ""))
		d.Class = trimSpaces(d.Class)
		if d.Class != "" {
			d.Class += " "
		}
		d.Class += replace
		return d
	}
}

// ReplaceColor replace a Tailwind class and its variants from the color attribute.
//
//	// replaces bg-color-[...], [...]:bg-color-[...]
//	style.ReplaceColor("bg-color", "bg-color-gray-50 hover:bg-color-gray-100")
func ReplaceColor(pattern, replace string) Opt {
	rePattern := regexp.MustCompile(`\b(?:[^ ]+:)?` + pattern + `(?:-[^ ]+)?\b`)
	return func(d d) d {
		d.Color = strings.TrimSpace(rePattern.ReplaceAllString(d.Color, ""))
		d.Color = trimSpaces(d.Color)
		if d.Color != "" {
			d.Color += " "
		}
		d.Color += replace
		return d
	}
}

// Replace replaces both in class and color.
func Replace(pattern, replace string) Opt {
	rcl := ReplaceClass(pattern, replace)
	rco := ReplaceColor(pattern, replace)
	return func(d d) d {
		return rco(rcl(d))
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

// D defines a style as a list of Opt (Color/Class/Add/ReplaceXXX calls).
type D []Opt

func (ds D) apply(d d) d {
	for _, opt := range ds {
		d = opt(d)
	}
	return d
}

type d struct {
	Class string
	Color string
	Add   string
}

// Default finds the default style from Defaults.
func Default(defaults Defaults, style Style, k string) D {
	defs := defaults[StyleDefault][k]
	for dstyle, ddefaults := range defaults {
		if dstyle&style == 0 {
			continue
		}
		defs = append(defs, ddefaults[k]...)
	}
	return defs
}

// CSSClass returns the CSS class string.
func (def D) CSSClass(defaults Defaults, style Style, k string) string {
	d := Default(defaults, style, k).apply(d{})
	d = def.apply(d)
	class := d.Class
	if class != "" && d.Color != "" {
		class += " "
	}
	class += d.Color
	if class != "" && d.Add != "" {
		class += " "
	}
	class += d.Add
	return class
}

// Delta returns all classes added by dst.
func (def D) Delta(defaults Defaults, src, dst Style, k string) string {
	ssrc := strings.Split(def.CSSClass(defaults, src, k), " ")
	sdst := strings.Split(def.CSSClass(defaults, dst, k), " ")
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
type Defaults map[Style]map[string]D
