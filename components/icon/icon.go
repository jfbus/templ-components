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
			style.SizeXS: {
				style.Add("h-3 w-3"),
			},
			style.SizeS: {
				style.Add("h-3.5 w-3.5"),
			},
			style.SizeNormal: {
				style.Add("h-4 w-4"),
			},
			style.SizeL: {
				style.Add("h-[18px] w-[18px]"),
			},
			style.SizeXL: {
				style.Add("h-5 w-5"),
			},
			style.SizeTwoXL: {
				style.Add("h-6 w-6"),
			},
			style.SizeThreeXL: {
				style.Add("h-[30px] w-[30px]"),
			},
			style.SizeFourXL: {
				style.Add("h-9 w-9"),
			},
			style.SizeFiveXL: {
				style.Add("h-12 w-12"),
			},
			style.SizeSixXL: {
				style.Add("h-[60px] w-[60px]"),
			},
			style.SizeSevenXL: {
				style.Add("h-[72px] w-[72px]"),
			},
			style.SizeEightXL: {
				style.Add("h-24 w-24"),
			},
			style.SizeNineXL: {
				style.Add("h-32 w-32"),
			},
			style.SizeFull: {
				style.Add("h-full w-full"),
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

func (def D) size() size.Size {
	if def.Size == 0 {
		return size.Normal
	}
	return def.Size
}

func (def D) class() string {
	return style.CSSClass(def.Style|style.Size(def.size()), "icon", def.CustomStyle)
}
