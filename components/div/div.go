package div

import (
	"github.com/a-h/templ"
	"github.com/jfbus/templ-components/components/style"
)

var Defaults = style.Defaults{
	style.StyleDefault: {
		"Class": {},
	},
}

type D struct {
	//playground:import:github.com/jfbus/templ-components/components/button
	//playground:default:button.C(button.D{Label:"Button"})
	Content templ.Component
	Class   style.D
}

func (def D) class() string {
	return style.Default(Defaults, style.StyleDefault, "Class").String()
}
