// Package buttongroup implements button groups.
package buttongroup

import (
	"github.com/a-h/templ"
	"github.com/jfbus/templ-components/components/button"
	"github.com/jfbus/templ-components/components/size"
	"github.com/jfbus/templ-components/components/style"
)

// Defaults defines the default Color/Class. They are overriden by D.Color/D.Class.
var Defaults = style.Defaults{
	style.StyleDefault: {
		"Class": {
			style.Color("text-gray-900 bg-white border-gray-300 hover:bg-gray-100 focus:ring-gray-100 dark:bg-gray-800 dark:text-white dark:border-gray-600 dark:hover:bg-gray-700 dark:hover:border-gray-600 dark:focus:ring-gray-700"),
			style.Class("inline-flex rounded-md shadow-sm"),
		},
		"ButtonClass": {
			style.Color("text-gray-900 bg-white border-gray-300 hover:bg-gray-100 focus:ring-gray-100 dark:bg-gray-800 dark:text-white dark:border-gray-600 dark:hover:bg-gray-700 dark:hover:border-gray-600 dark:focus:ring-gray-700"),
			style.Class("border-t border-b border-e first:border first:rounded-s-lg last:border-t last:border-b last:border-e last:rounded-e-lg"),
		},
	},
}

type D struct {
	// Buttons is the list of buttons to display.
	//playground:import:github.com/jfbus/templ-components/components/button
	//playground:default:[]button.D{{Label:"A"},{Label:"B"},{Label:"C"},{Label:"D"}}
	Buttons []button.D
	// Size defines the buttons size (shortcut for Buttons.Size).
	Size size.Size
	// Class defines additional CSS class for the buttongroup.
	Class style.D
	// Class defines the default CSS class for the buttons.
	ButtonClass style.D
	// Attributes stores additional attributes (e.g. HTMX attributes).
	Attributes templ.Attributes
}

func (def D) buttons() []button.D {
	bs := make([]button.D, len(def.Buttons))
	for i := range def.Buttons {
		bs[i] = def.Buttons[i]
		if def.Size != size.Inherit {
			bs[i].Size = def.Size
		}
		class := append(style.Default(Defaults, style.StyleDefault, "ButtonClass"), def.ButtonClass...)
		bs[i].Class = append(class, bs[i].Class...)
	}
	return bs
}

func (def D) groupClass() string {
	return def.Class.CSSClass(Defaults, style.StyleDefault, "Class")
}
