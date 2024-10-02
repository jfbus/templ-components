package sidebar

import (
	"github.com/a-h/templ"
	"github.com/jfbus/templ-components/components/style"
)

var DefaultID = "sidebar"

var Defaults = style.Defaults{
	style.StyleDefault: {
		"ContainerClass": {
			style.Class("fixed top-0 left-0 z-20 w-64 h-screen transition-transform sm:translate-x-0"),
		},
		"Class": {
			style.Class("h-full px-3 py-4 overflow-y-auto"),
			style.Color("bg-gray-50 dark:bg-gray-800"),
		},
	},
}

type D struct {
	ID string
	//playground:import:github.com/jfbus/templ-components/components/button
	//playground:default:button.C(button.D{Label:"Button"})
	Content        templ.Component
	ContainerClass style.D
	Class          style.D
}

func (def D) id() string {
	if def.ID != "" {
		return def.ID
	}
	return DefaultID
}

func (def D) containerClass() string {
	return def.ContainerClass.CSSClass(Defaults, style.StyleDefault, "ContainerClass")
}

func (def D) class() string {
	return def.Class.CSSClass(Defaults, style.StyleDefault, "Class")
}
