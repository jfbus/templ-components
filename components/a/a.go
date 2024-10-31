package a

import (
	"github.com/a-h/templ"
	"github.com/jfbus/templ-components/components/style"
)

func init() {
	style.SetDefaults(style.Defaults{
		"a": {
			style.StyleDefault: {
				style.Set("hover:underline"),
			},
		},
	})
}

type D struct {
	Href       string
	Text       string
	Class      style.D
	Attributes templ.Attributes
}

func (def D) class() string {
	return def.Class.CSSClass(style.StyleDefault, "a")
}
