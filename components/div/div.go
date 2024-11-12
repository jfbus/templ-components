package div

import (
	"github.com/a-h/templ"
	"github.com/jfbus/templui/components/style"
)

type D struct {
	//playground:import:github.com/jfbus/templui/components/button
	//playground:default:button.C(button.D{Label:"Button"})
	Content     templ.Component
	CustomStyle style.Custom
}

func (def D) class() string {
	return style.CSSClass(style.Default, "div", def.CustomStyle)
}
