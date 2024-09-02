// Package buttongroup implements button groups.
package buttongroup

import (
	"github.com/a-h/templ"
	"github.com/jfbus/templ-components/components/button"
	"github.com/jfbus/templ-components/components/helper"
	"github.com/jfbus/templ-components/components/size"
	"github.com/jfbus/templ-components/components/style"
)

// Defaults defines the default Color/Class. They are overriden by D.Color/D.Class.
var Defaults = D{
	Class: style.D{
		Color: "text-gray-900 bg-white border-gray-300 hover:bg-gray-100 focus:ring-gray-100 dark:bg-gray-800 dark:text-white dark:border-gray-600 dark:hover:bg-gray-700 dark:hover:border-gray-600 dark:focus:ring-gray-700",
		Class: "inline-flex rounded-md shadow-sm",
	},
}

type D struct {
	// Buttons is the list of buttons to display.
	Buttons []button.D
	// Size defines the buttons size (shortcut for Buttons.Size).
	Size size.Size
	// Class defines additional CSS class for the button.
	Class style.D
	// Attributes stores additional attributes (e.g. HTMX attributes).
	Attributes templ.Attributes
}

func (def D) buttons() []button.D {
	bs := make([]button.D, len(def.Buttons))
	for i := range def.Buttons {
		bs[i] = def.Buttons[i]
		bs[i].Style = button.StyleNoBorder
		if def.Size != size.Inherit {
			bs[i].Size = def.Size
		}
		bs[i].Class.Color = helper.IfEmpty(bs[i].Class.Color, def.Class.Color, Defaults.Class.Color)
		switch i {
		case 0:
			bs[i].Class.Class += " border rounded-s-lg"
		case len(bs) - 1:
			bs[i].Class.Class += " border rounded-e-lg"
		default:
			bs[i].Class.Class += " border-t border-b"
		}
	}
	return bs
}

func (def D) groupClass() string {
	class := helper.IfEmpty(def.Class.Class, Defaults.Class.Class)
	class += " " + helper.IfEmpty(def.Class.Color, Defaults.Class.Color)
	return class
}
