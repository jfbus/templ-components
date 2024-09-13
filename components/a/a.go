package a

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
	Href       string
	Text       string
	Class      style.D
	Attributes templ.Attributes
}

func (def D) class() string {
	return def.Class.CSSClass(Defaults, style.StyleDefault, "Class")
}
