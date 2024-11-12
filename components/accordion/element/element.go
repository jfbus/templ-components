package element

import (
	"github.com/a-h/templ"
	"github.com/jfbus/templ-components/components/style"
)

func init() {
	style.SetDefaults(style.Defaults{
		"accordion/element/title": {
			style.Default: {
				style.Set("font-medium flex items-center justify-between w-full p-5 rtl:text-right border first:rounded-t-xl border-b-0 last:border-b-1 focus:ring-4 gap-3 cursor-pointer"),
			},
		},
		"accordion/element/content": {
			style.Default: {
				style.Set("p-5 border border-b-0"),
			},
		},
	})
}

type D struct {
	ID          string
	Open        bool
	Title       string
	Content     templ.Component
	CustomStyle style.Custom
	Attributes  templ.Attributes
}

func (def D) class(k string) string {
	return style.CSSClass(style.Default, k, def.CustomStyle)
}
