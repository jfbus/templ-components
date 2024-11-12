package sidebar

import (
	"github.com/a-h/templ"
	"github.com/jfbus/templ-components/components/style"
)

var DefaultID = "sidebar"

func init() {
	style.SetDefaults(style.Defaults{
		"sidebar": {
			style.Default: {
				style.Set("fixed top-0 left-0 z-20 w-64 h-screen transition-transform sm:translate-x-0"),
			},
		},
		"sidebar/content": {
			style.Default: {
				style.Set("h-full px-3 py-4 overflow-y-auto"),
			},
		},
	})
}

type D struct {
	ID string
	//playground:import:github.com/jfbus/templ-components/components/button
	//playground:default:button.C(button.D{Label:"Button"})
	Content templ.Component
	// CustomStyle defines a custom style.
	// 	style.Custom{
	// 		"sidebar":         style.D{style.Add("...")},
	// 		"sidebar/content": style.D{style.Add("...")},
	//	}
	CustomStyle style.Custom
}

func (def D) id() string {
	if def.ID != "" {
		return def.ID
	}
	return DefaultID
}

func (def D) class(k string) string {
	return style.CSSClass(style.Default, k, def.CustomStyle)
}
