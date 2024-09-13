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
}

// D is the button definition.
type D struct {
	// ID is the button id.
	ID string
	// Type is the button type (button, submit, ...).
	Type Type
	// Label is the button label.
	Label string
	// HideLabel defines if the label is only available to screenreaders.
	HideLabel bool
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

func (def D) buttonType() string {
	if def.Type == "" {
		return string(Button)
	}
	return string(def.Type)
}

func (def D) buttonClass() string {
	class := def.Class.CSSClass(Defaults, def.Style, "Button")
	switch {
	case def.noLabel() && def.Size >= size.Normal:
		class += " p-2.5 text-sm"
	case def.noLabel() && def.Size == size.XS:
		class += " p-1 text-sm"
	case def.noLabel():
		class += " p-1.5 text-sm"
	case def.Size == size.XS:
		class += " px-3 py-2 text-xs"
	case def.Size == size.S:
		class += " px-3 py-2 text-sm"
	case def.Size == size.L:
		class += " px-5 py-3 text-base"
	case def.Size == size.XL:
		class += " px-6 py-3.5 text-base"
	default:
		class += " px-5 py-2.5 text-sm"
	}
	if def.Icon != "" || def.Loader {
		class += " inline-flex items-center justify-center"
	}
	return class
}

func (def D) noLabel() bool {
	return def.HideLabel || def.Label == ""
}

func (def D) iconSize() size.Size {
	if def.HideLabel || def.Label == "" {
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
	if def.IconPosition == position.End {
		return style.D{style.Class("ms-2")}
	}
	return style.D{style.Class("me-2")}
}
