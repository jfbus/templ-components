package style

import (
	"strings"

	"github.com/jfbus/templ-components/components/helper"
)

type D struct {
	Class string
	Color string
}

func (def D) CSSClass(defaults D) string {
	class := helper.IfEmpty(def.Class, defaults.Class)
	class += " " + helper.IfEmpty(def.Color, defaults.Color)
	return class
}

type Style int

const (
	StyleDefault Style = 0

	AddStyleMask Style = 0b111100000
)

type Defaults map[Style]map[string]D

func Default(defaults Defaults, style Style, k string) D {
	res := defaults[StyleDefault][k]
	for dstyle, ddefaults := range defaults {
		if dstyle&AddStyleMask != 0 || dstyle&style == 0 {
			continue
		}
		res.Class = helper.IfEmpty(ddefaults[k].Class, res.Class)
		res.Color = helper.IfEmpty(ddefaults[k].Color, res.Color)
	}
	for dstyle, ddefaults := range defaults {
		if dstyle&AddStyleMask == 0 || dstyle&style == 0 {
			continue
		}
		res.Class = strings.TrimSpace(ddefaults[k].Class + " " + res.Class)
		res.Color = strings.TrimSpace(ddefaults[k].Color + " " + res.Color)
	}
	return res
}
