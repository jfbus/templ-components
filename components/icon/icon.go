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
	StyleBorder style.Style = 1 << 8
)

func init() {
	style.SetDefaults(style.Defaults{
		"icon": {
			StyleBorder: {
				style.Set("p-1 rounded-md border"),
			},
		},
	})
}

// D is the definition of icons.
// Usage: @icon.C(icon.D{Icon:icon.Flower}) or @icon.C(icon.Flower) or &icon.C("flower")
type D struct {
	// Icon is the SVG string (e.g. icon.Flower).
	Icon  string
	Style style.Style
	// Size is the icon size.
	Size size.Size
	// CustomStyle defines a custom style.
	// 	style.Custom{
	// 		"icon": style.D{style.Add("...")},
	//	}
	CustomStyle style.Custom
	// Attributes adds custom attributes.
	Attributes templ.Attributes
}

func (def D) class() string {
	return style.CSSClass(def.Style, "icon", def.CustomStyle) +
		" " + def.Size.Class()
}
