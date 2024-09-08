package style

import (
	"regexp"
	"strings"
	"unicode"
)

type Opt func(d) d

func Color(color string) Opt {
	return func(d d) d {
		d.Color = color
		return d
	}
}

func Class(class string) Opt {
	return func(d d) d {
		d.Class = class
		return d
	}
}

func Add(add string) Opt {
	return func(d d) d {
		if d.Add != "" {
			d.Add += " "
		}
		d.Add += add
		return d
	}
}

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

type Style int

const (
	StyleDefault Style = 0
)

type Defaults map[Style]map[string]D
