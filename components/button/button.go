// Package button implements buttons.
package button

import (
	"github.com/a-h/templ"
	"github.com/jfbus/templ-components/components/helper"
	"github.com/jfbus/templ-components/components/position"
	"github.com/jfbus/templ-components/components/size"
)

type Type string

const (
	Button Type = "button"
	Submit Type = "submit"
)

type Style int

const (
	// Normal is the normal style
	Normal   Style = 0
	Outline  Style = 1
	Pill     Style = 2
	NoBorder Style = 4
)

// Defaults defines the default Color/Class values, and may be changed. They are overriden by D.Color/D.Class.
var (
	Defaults = map[Style]D{
		Normal: {
			Color: "text-white bg-blue-700 hover:bg-blue-800 focus:ring-blue-300 dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800",
			Class: "font-medium focus:ring-4 focus:outline-none",
		},
		Outline: {
			Color: "text-gray-900 bg-white border-gray-300 hover:bg-gray-100 focus:ring-gray-100 dark:bg-gray-800 dark:text-white dark:border-gray-600 dark:hover:bg-gray-700 dark:hover:border-gray-600 dark:focus:ring-gray-700",
			Class: "font-medium focus:ring-4 focus:outline-none border",
		},
		NoBorder: {
			Color: "text-gray-900 bg-white hover:bg-gray-100 focus:ring-gray-100 dark:bg-gray-800 dark:text-white dark:hover:bg-gray-700 dark:focus:ring-gray-700",
		},
	}
)

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
	Style Style
	// Color overrides the default color CSS classes
	// (Defaults[Outline].Color if Outline, Defaults[Normal].Color otherwise)
	Color string
	// Size defines the input size (size.XS, size.S, size.Normal (default) size.L or size.XL).
	Size size.Size
	// Icon displays an icon on the left side.
	Icon string
	// IconPosition can be position.Start (default) or position.End.
	IconPosition position.Position
	// Disabled disables the input.
	Disabled bool
	// Loader displays a spinning loader when an HTMX action is triggered by the input.
	Loader bool
	// Class overrides the default CSS class for the button.
	Class string
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
	var class string
	switch {
	case def.Style&Outline != 0:
		class = helper.IfEmpty(def.Class, Defaults[Outline].Class)
		class += " " + helper.IfEmpty(def.Color, Defaults[Outline].Color)
	case def.Style&NoBorder != 0:
		class = helper.IfEmpty(def.Class, Defaults[NoBorder].Class)
		class += " " + helper.IfEmpty(def.Color, Defaults[NoBorder].Color)
	default:
		class = helper.IfEmpty(def.Class, Defaults[Normal].Class)
		class += " " + helper.IfEmpty(def.Color, Defaults[Normal].Color)
	}

	if def.Style&Pill != 0 {
		class += " rounded-full"
	} else if def.Style&NoBorder == 0 {
		class += " rounded-lg"
	}
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
		class += " inline-flex items-center"
	}
	if def.Color != "" {
		class += " " + def.Color
	}
	if def.Class != "" {
		class += " " + def.Class
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

func (def D) iconClass() string {
	if def.noLabel() {
		return ""
	}
	if def.IconPosition == position.End {
		return "ms-2"
	}
	return "me-2"
}
