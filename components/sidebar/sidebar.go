package sidebar

import (
	"github.com/a-h/templ"
	"github.com/jfbus/templ-components/components/style"
)

var DefaultID = "sidebar"

func init() {
	style.SetDefaults(style.Defaults{
		"sidebar": {
			style.StyleDefault: {
				style.Set("fixed top-0 left-0 z-20 w-64 h-screen transition-transform sm:translate-x-0"),
			},
		},
		"sidebar/div": {
			style.StyleDefault: {
				style.Set("h-full px-3 py-4 overflow-y-auto"),
			},
		},
	})
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
	return def.ContainerClass.CSSClass(style.StyleDefault, "sidebar")
}

func (def D) class() string {
	return def.Class.CSSClass(style.StyleDefault, "sidebar/div")
}
