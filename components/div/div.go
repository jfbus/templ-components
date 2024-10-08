package div

import (
	"github.com/a-h/templ"
	"github.com/jfbus/templ-components/components/style"
)

type D struct {
	//playground:import:github.com/jfbus/templ-components/components/button
	//playground:default:button.C(button.D{Label:"Button"})
	Content templ.Component
	Class   style.D
}

func (def D) class() string {
	return def.Class.CSSClass(style.StyleDefault, "div")
}
