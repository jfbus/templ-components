// Package button implements buttons.
package button

import (
	"github.com/a-h/templ"
	"github.com/jfbus/templ-components/components/position"
	"github.com/jfbus/templ-components/components/size"
	"github.com/jfbus/templ-components/components/style"
)

type Type string

const (
	Button Type = "button"
	Submit Type = "submit"
	A      Type = "a"
)

const (
	StylePill        style.Style = 1 << 8
	StyleOutline     style.Style = 1 << 9
	StyleOutlinePill style.Style = 1 << 10
	StyleNoBorder    style.Style = 1 << 11
	StyleFullWidth   style.Style = 1 << 12
)

type HideLabel int

const (
	HideLabelNever = iota
	HideLabelAlways
	HideLabelSmall
)

// Defaults defines the default Color/Class values, and may be changed. They are overriden by D.Color/D.Class.
var Defaults = style.Defaults{
	style.StyleDefault: map[string]style.D{
		"Button": {
			style.Color("text-white bg-blue-700 hover:bg-blue-800 focus:ring-blue-300 dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"),
			style.Class("rounded-lg font-medium focus:ring-4 focus:outline-none"),
		},
	},
	StylePill: map[string]style.D{
		"Button": {
			style.Color("text-white bg-blue-700 hover:bg-blue-800 focus:ring-blue-300 dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"),
			style.Class("rounded-full font-medium focus:ring-4 focus:outline-none"),
		},
	},
	StyleOutline: map[string]style.D{
		"Button": {
			style.Color("text-gray-900 bg-white border-gray-300 hover:bg-gray-100 focus:ring-gray-100 dark:bg-gray-800 dark:text-white dark:border-gray-600 dark:hover:bg-gray-700 dark:hover:border-gray-600 dark:focus:ring-gray-700"),
			style.Class("rounded-lg font-medium focus:ring-4 focus:outline-none border"),
		},
	},
	StyleOutlinePill: map[string]style.D{
		"Button": {
			style.Color("text-gray-900 bg-white border-gray-300 hover:bg-gray-100 focus:ring-gray-100 dark:bg-gray-800 dark:text-white dark:border-gray-600 dark:hover:bg-gray-700 dark:hover:border-gray-600 dark:focus:ring-gray-700"),
			style.Class("rounded-full font-medium focus:ring-4 focus:outline-none border"),
		},
	},
	StyleNoBorder: map[string]style.D{
		"Button": {
			style.Color("text-gray-900 bg-white hover:bg-gray-100 focus:ring-gray-100 dark:bg-gray-800 dark:text-white dark:hover:bg-gray-700 dark:focus:ring-gray-700"),
		},
	},
	StyleFullWidth: map[string]style.D{
		"Button": {
			style.Add("block w-full text-center"),
		},
	},
	style.StyleDisabled: map[string]style.D{
		"Button": {
			style.ReplaceColor("bg", "bg-blue-400 dark:bg-blue-500"),
			style.Add("cursor-not-allowed"),
		},
	},
}

// D is the button definition.
type D struct {
	// ID is the button id.
	ID string
	// Type is the button type (button, submit, ...).
	Type Type
	// Label is the button label.
	Label string
	// HideLabel defines if the label is displayed or only available to screenreaders.
	HideLabel HideLabel
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
	// Class overrides the default CSS class for the button.
	Class style.D
	// Attributes stores additional attributes (e.g. HTMX attributes).
	Attributes templ.Attributes
}

func (def D) style() style.Style {
	if def.Disabled {
		return def.Style | style.StyleDisabled
	}
	return def.Style
}

func (def D) buttonType() string {
	if def.Type == "" {
		return string(Button)
	}
	return string(def.Type)
}

func (def D) buttonClass() string {
	class := def.Class.CSSClass(Defaults, def.style(), "Button")
	switch {
	case def.noLabel() && def.Size >= size.Normal:
		class += " p-2.5 text-sm"
	case def.noLabel() && def.Size == size.XS:
		class += " p-1 text-sm"
	case def.noLabel():
		class += " p-1.5 text-sm"
	case def.Size == size.XS:
		class += " p-2 text-xs"
	case def.Size == size.S:
		class += " p-2 text-sm"
	case def.Size == size.L:
		class += " p-3 text-base"
	case def.Size == size.XL:
		class += " p-3.5 text-base"
	default:
		class += " p-2.5 text-sm"
	}
	if def.Icon != "" || def.Loader {
		class += " inline-flex items-center justify-center"
	}
	return class
}

func (def D) noLabel() bool {
	return def.HideLabel == HideLabelAlways || def.Label == ""
}

func (def D) labelClass() string {
	switch def.HideLabel {
	case HideLabelAlways:
		return "sr-only"
	case HideLabelSmall:
		return "hidden sm:block"
	default:
		return ""
	}
}

func (def D) iconSize() size.Size {
	if def.HideLabel == HideLabelAlways || def.Label == "" {
		return def.Size + 1
	}
	switch def.Size {
	case size.XS:
		return size.XS
	case size.L, size.XL:
		return size.Normal
	default:
		return size.S
	}
}

func (def D) iconClass() style.D {
	if def.noLabel() {
		return style.D{}
	}
	switch {
	case def.HideLabel == HideLabelSmall && def.IconPosition == position.End:
		return style.D{style.Class("sm:ms-2")}
	case def.IconPosition == position.End:
		return style.D{style.Class("ms-2")}
	case def.HideLabel == HideLabelSmall:
		return style.D{style.Class("sm:me-2")}
	default:
		return style.D{style.Class("me-2")}

	}
}
