package footer

import (
	"github.com/a-h/templ"
	"github.com/jfbus/templ-components/components/social"
	"github.com/jfbus/templ-components/components/style"
)

const (
	StyleSticky style.Style = 1 << 8
	StyleBorder style.Style = 1 << 9
)

func init() {
	style.SetDefaults(style.Defaults{
		"footer": {
			style.StyleDefault: {
				style.Set("mx-auto w-full max-w-screen-xl p-4 py-6 lg:py-8"),
			},
			StyleBorder: {
				style.Set("mx-auto w-full max-w-screen-xl p-4 rounded-lg shadow"),
			},
			StyleSticky: {
				style.Set("fixed bottom-0 left-0 z-20 w-full p-4 border-t shadow md:flex md:items-center md:justify-between md:p-6"),
			},
		},
		"footer/content": {
			style.StyleDefault: {
				style.Set("md:flex md:items-center md:justify-between"),
			},
		},
		"footer/separator": {
			style.StyleDefault: {
				style.Set("my-6 sm:mx-auto lg:my-8"),
			},
		},
		"footer/copyright": {
			style.StyleDefault: {
				style.Set("block text-sm sm:text-center"),
			},
		},
		"footer/footer": {
			style.StyleDefault: {
				style.Set("sm:flex sm:items-center sm:justify-between"),
			},
		},
		"footer/social": {
			style.StyleDefault: {
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
	//playground:import:github.com/jfbus/templ-components/components/a
	//playground:import:github.com/jfbus/templ-components/components/footer/section
	//playground:default:[]section.D{{Title:"Section 1", Links: []a.D{{Text: "Link 1"},{Text: "Link 2"}}}, {Title:"Section 2", Links: []a.D{{Text: "Link 1"},{Text: "Link 2"}}}, {Title:"Section 3", Links: []a.D{{Text: "Link 1"},{Text: "Link 2"}}}}
	Sections       any
	Social         []social.D
	ContainerClass style.D
}

func (def D) containerClass() string {
	return def.ContainerClass.CSSClass(def.Style, "footer")
}

func (def D) class() string {
	return style.D{}.CSSClass(def.Style, "footer/content")
}

func (def D) separatorClass() string {
	return style.D{}.CSSClass(def.Style, "footer/separator")
}

func (def D) copyrightClass() string {
	return style.D{}.CSSClass(def.Style, "footer/copyright")
}

func (def D) footerClass() string {
	return style.D{}.CSSClass(def.Style, "footer/footer")
}

func (def D) socialClass() string {
	return style.D{}.CSSClass(def.Style, "social/social")
}
