// Package button implements buttons.
package button

import (
	"github.com/a-h/templ"
	"github.com/jfbus/templui/components/position"
	"github.com/jfbus/templui/components/size"
	"github.com/jfbus/templui/components/style"
	"github.com/jfbus/templui/components/tooltip"
)

type Type string

const (
	Button Type = "button"
	Submit Type = "submit"
	A      Type = "a"
)

const (
	StylePill            style.Style = 1 << 8
	StyleAlternative     style.Style = 1 << 9
	StyleOutline         style.Style = 1 << 10
	StyleNoBorder        style.Style = 1 << 11
	StyleFullWidth       style.Style = 1 << 12
	StyleHideLabelAlways style.Style = 1 << 13
	StyleHideLabelSmall  style.Style = 1 << 14
	StyleHideLabelMedium style.Style = 1 << 15
)

func init() {
	style.SetDefaults(style.Defaults{
		"button": {
			style.Default: {
				style.Set("rounded-lg font-medium focus:ring-4 focus:outline-none"),
			},
			StylePill: {
				style.Replace("rounded-lg", "rounded-full"),
			},
			StyleAlternative: {
				style.Add("border"),
			},
			StyleOutline: {
				style.Add("border"),
			},
			StyleFullWidth: {
				style.Add("block w-full text-center"),
			},
			style.Disabled: {
				style.Add("cursor-not-allowed"),
			},
			style.SizeXS: {
				style.Add("p-2 text-xs"),
			},
			style.SizeS: {
				style.Add("p-2 text-sm"),
			},
			style.SizeNormal: {
				style.Add("p-2.5 text-sm"),
			},
			style.SizeL: {
				style.Add("p-3 text-base"),
			},
			style.SizeXL: {
				style.Add("p-3.5 text-base"),
			},
			StyleHideLabelAlways | style.SizeXS: {
				style.Replace("p-2", "p-1"),
			},
			StyleHideLabelAlways | style.SizeS: {
				style.Replace("p-2", "p-1.5"),
			},
			StyleHideLabelAlways | style.SizeNormal: {
				style.Replace("p-2.5", "p-2"),
			},
			StyleHideLabelAlways | style.SizeL: {
				style.Replace("p-3", "p-2.5"),
			},
			StyleHideLabelAlways | style.SizeXL: {
				style.Replace("p-3.5", "p-3"),
			},
		},
		"button/label": {
			StyleHideLabelAlways: {
				style.Set("sr-only"),
			},
			StyleHideLabelSmall: {
				style.Set("sr-only sm:not-sr-only"),
			},
			StyleHideLabelMedium: {
				style.Set("sr-only md:not-sr-only"),
			},
		},
	})
}

// D is the button definition.
type D struct {
	// ID is the button id.
	ID string
	// Type is the button type (button, submit, ...).
	Type Type
	// Label is the button label.
	Label string
	// Style is the button style (Normal, Pill, Outline ou Pill|Outline).
	Style style.Style
	// Size defines the input size (size.XS, size.S, size.Normal (default) size.L or size.XL).
	//playground:values:size.Normal,size.XS,size.S,size.L,size.XL
	Size size.Size
	// Icon displays an icon on the left side.
	//playground:values:icon.Flower
	Icon string
	// IconPosition can be position.Start (default) or position.End.
	//playground:values:position.Start,position.End
	IconPosition position.Position
	// Disabled disables the input.
	Disabled bool
	// Loader displays a spinning loader when an HTMX action is triggered by the input.
	Loader bool
	// Tooltip adds a tooltip to the button.
	Tooltip *tooltip.D
	// StyleKey defines the style key to use ("button" by default).
	StyleKey string
	// CustomStyle defines a custom style.
	// 	style.Custom{
	// 		"button": style.D{style.Add("text-sm")},
	//	}
	CustomStyle style.Custom
	// Attributes stores additional attributes (e.g. HTMX attributes).
	Attributes templ.Attributes
}

func (def D) size() size.Size {
	if def.Size == 0 {
		return size.Normal
	}
	return def.Size
}

func (def D) style() style.Style {
	st := def.Style | style.Size(def.size())
	if def.Disabled {
		return st | style.Disabled
	}
	return st
}

func (def D) styleKey() string {
	if def.StyleKey == "" {
		return "button"
	}
	return def.StyleKey
}

func (def D) buttonType() string {
	if def.Type == "" {
		return string(Button)
	}
	return string(def.Type)
}

func (def D) buttonClass() string {
	class := style.CSSClass(def.style(), def.styleKey(), def.CustomStyle)
	if def.Icon != "" || def.Loader {
		class += " inline-flex items-center justify-center"
	}
	if def.Tooltip != nil {
		class += " " + def.Tooltip.Class()
	}
	return class
}

func (def D) noLabel() bool {
	return def.Style&StyleHideLabelAlways != 0 || def.Label == ""
}

func (def D) labelClass() string {
	return style.CSSClass(def.style(), def.styleKey()+"/label", def.CustomStyle)
}

func (def D) iconSize() size.Size {
	if def.noLabel() {
		return def.size() + 1
	}
	switch def.size() {
	case size.XS:
		return size.XS
	case size.L, size.XL:
		return size.Normal
	default:
		return size.S
	}
}

func (def D) iconCustomStyle() style.Custom {
	if def.noLabel() {
		return style.Custom{}
	}
	switch {
	case def.Style&StyleHideLabelSmall != 0 && def.IconPosition == position.End:
		return style.CustomAdd("sm:ms-2")
	case def.IconPosition == position.End:
		return style.CustomAdd("ms-2")
	case def.Style&StyleHideLabelSmall != 0:
		return style.CustomAdd("sm:me-2")
	default:
		return style.CustomAdd("me-2")
	}
}
