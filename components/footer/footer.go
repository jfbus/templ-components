package footer

import (
	"github.com/a-h/templ"
	"github.com/jfbus/templui/components/a"
	"github.com/jfbus/templui/components/footer/section"
	"github.com/jfbus/templui/components/social"
	"github.com/jfbus/templui/components/style"
)

const (
	StyleSticky style.Style = 1 << 8
	StyleBorder style.Style = 1 << 9
)

func init() {
	style.SetDefaults(style.Defaults{
		"footer": {
			style.Default: {
				style.Set("mx-auto w-full"),
			},
			StyleBorder: {
				style.Set("mx-auto w-full rounded-lg shadow"),
			},
			StyleSticky: {
				style.Set("fixed bottom-0 left-0 z-20 w-full p-4 border-t shadow md:flex md:items-center md:justify-between md:p-6"),
			},
		},
		"footer/div": {
			style.Default: {
				style.Set("px-4 py-6"),
			},
		},
		"footer/content": {
			style.Default: {
				style.Set("md:flex md:items-center md:justify-between"),
			},
		},
		"footer/separator": {
			style.Default: {
				style.Set("my-6 sm:mx-auto"),
			},
		},
		"footer/copyright": {
			style.Default: {
				style.Set("block text-sm sm:text-center"),
			},
		},
		"footer/social": {
			style.Default: {
				style.Set("flex mt-4 sm:justify-center sm:mt-0"),
			},
		},
	})
}

// D defines a navbar component.
// To add multiple components (buttons, dropdowns, ...) to a section, use div.C/components.Append.
type D struct {
	Style style.Style
	//playground:default:templ.Raw("© 2024 Jean-François Bustarret. All Rights Reserved.")
	Copyright templ.Component
	Brand     templ.Component
	// Sections can either be a []section.D or a []a.D slice.
	//playground:import:github.com/jfbus/templui/components/a
	//playground:import:github.com/jfbus/templui/components/footer/section
	//playground:default:[]section.D{{Title:"Section 1", Links: []a.D{{Text: "Link 1"},{Text: "Link 2"}}}, {Title:"Section 2", Links: []a.D{{Text: "Link 1"},{Text: "Link 2"}}}, {Title:"Section 3", Links: []a.D{{Text: "Link 1"},{Text: "Link 2"}}}}
	Sections any
	Social   []social.D
	// CustomStyle defines a custom style.
	// 	style.Custom{
	// 		"footer":               style.D{style.Add("...")},
	// 		"footer/div":               style.D{style.Add("...")},
	// 		"footer/content":       style.D{style.Add("...")},
	// 		"footer/social":        style.D{style.Add("...")},
	// 		"footer/separator":     style.D{style.Add("...")},
	// 		"footer/copyright":     style.D{style.Add("...")},
	// 		"footer/section/link":  style.D{style.Add("...")},
	// 		"footer/section/title": style.D{style.Add("...")},
	//	}
	CustomStyle style.Custom
}

func (def D) class(k string) string {
	return style.CSSClass(def.Style, k, def.CustomStyle)
}

func (def D) sections() any {
	switch secs := def.Sections.(type) {
	case []a.D:
		cc := style.Custom{
			"a": style.Compute(def.Style, "footer/section/link", def.CustomStyle),
		}
		for i := range secs {
			secs[i].CustomStyle = secs[i].CustomStyle.AddBefore(cc)
		}
		return secs
	case []section.D:
		cc := style.Custom{
			"footer/section/link":  style.Compute(def.Style, "footer/section/link", def.CustomStyle),
			"footer/section/title": style.Compute(def.Style, "footer/section/title", def.CustomStyle),
		}
		for i := range secs {
			secs[i].CustomStyle = secs[i].CustomStyle.AddBefore(cc)
		}
		return secs
	}
	return def.Sections
}
