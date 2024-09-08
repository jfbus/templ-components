package element

import (
	"github.com/a-h/templ"
	"github.com/jfbus/templ-components/components/style"
)

var Defaults = style.Defaults{
	style.StyleDefault: {
		"TitleClass": {
			style.Color("font-medium text-gray-500 border-gray-200 focus:ring-gray-200 dark:focus:ring-gray-800 dark:border-gray-700 dark:text-gray-400 hover:bg-gray-100 dark:hover:bg-gray-800 "),
			style.Class("flex items-center justify-between w-full p-5 rtl:text-right border first:rounded-t-xl border-b-0 last:border-b-1 focus:ring-4 gap-3 cursor-pointer"),
		},
		"ContentClass": {
			style.Color("border-gray-200 dark:border-gray-700"),
			style.Class("p-5 border border-b-0"),
		},
	},
}

type D struct {
	ID           string
	Open         bool
	Title        string
	Content      templ.Component
	TitleClass   style.D
	ContentClass style.D
	Attributes   templ.Attributes
}

func (def D) titleClass() string {
	return def.TitleClass.CSSClass(Defaults, style.StyleDefault, "TitleClass")
}

func (def D) contentClass() string {
	return def.TitleClass.CSSClass(Defaults, style.StyleDefault, "ContentClass")
}
