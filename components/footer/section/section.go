package section

import (
	"github.com/jfbus/templ-components/components/a"
	"github.com/jfbus/templ-components/components/style"
)

func init() {
	style.SetDefaults(style.Defaults{
		"footer/section/title": {
			style.StyleDefault: {
				style.Set("mb-6 text-sm font-semibold uppercase"),
			},
		},
		"footer/section/link": {
			style.StyleDefault: {
				style.Set("font-medium mb-4"),
			},
		},
	})
}

type D struct {
	Title string
	Links []a.D
}

func (def D) titleClass() string {
	return style.D{}.CSSClass(style.StyleDefault, "footer/section/title")
}

func (def D) linkClass() string {
	return style.D{}.CSSClass(style.StyleDefault, "footer/section/link")
}
