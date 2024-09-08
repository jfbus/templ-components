// Package icon implements SVG icons.
//
//go:generate go run internal/build/generate.go
package icon

import (
	"github.com/a-h/templ"
	"github.com/jfbus/templ-components/components/size"
	"github.com/jfbus/templ-components/components/style"
)

const (
	StyleBorder style.Style = 1
)

var Defaults = style.Defaults{
	style.StyleDefault: {
		"Class": {},
	},
	StyleBorder: {
		"Class": {
			style.Color("bg-white dark:bg-gray-800 border-gray-300 dark:border-gray-600"),
			style.Class("p-1 rounded-md border"),
		},
	},
}

// D is the definition of icons.
// Usage: @icon.C(icon.D{Icon:icon.Flower}) or @icon.C(icon.Flower) or &icon.C("flower")
type D struct {
	// Icon is the SVG string (e.g. icon.Flower).
	Icon  string
	Style style.Style
	// Class overrides the default class..
	Class style.D
	// Size is the icon size.
	Size size.Size
	// Attributes adds custom attributes.
	Attributes templ.Attributes
}

func (def D) class() string {
	return def.Class.CSSClass(Defaults, def.Style, "Class") +
		" " + def.Size.Class()
}
