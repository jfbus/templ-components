package navbar

import (
	"github.com/a-h/templ"
	"github.com/jfbus/templui/components/style"
)

const (
	StyleSticky style.Style = 1 << 8
)

func init() {
	style.SetDefaults(style.Defaults{
		"navbar": {
			style.Default: {
				style.Set("w-full"),
			},
			StyleSticky: {
				style.Set("fixed w-full z-40 top-0 start-0 border-b"),
			},
		},
		"navbar/content": {
			style.Default: {
				style.Set("max-w-screen-xl flex flex-wrap items-center justify-between mx-auto p-4"),
			},
			StyleSticky: {
				style.Set("max-w-screen-xl flex flex-wrap items-center justify-between mx-auto p-4"),
			},
		},
	})
}

// D defines a navbar component.
// To add multiple components (buttons, dropdowns, ...) to a section, use div.C/components.Append.
type D struct {
	Style style.Style
	//playground:import:github.com/jfbus/templui/components/a
	//playground:import:github.com/jfbus/templui/components/button
	//playground:import:github.com/jfbus/templui/components/dropdown
	//playground:import:github.com/jfbus/templui/components/icon
	//playground:import:github.com/jfbus/templui/components/size
	//playground:default:[]templ.Component{a.C(a.D{Text:"Your Brand"}), dropdown.C(dropdown.D{Button: button.D{Style: button.StyleNoBorder, Icon: icon.CircleUserRound, Size: size.XL}, Links: [][]a.D{{{Text: "Profile"}}, {{Text: "Logout"}}}})}
	Sections []templ.Component
	// CustomStyle defines a custom style.
	// 	style.Custom{
	// 		"navbar":         style.D{style.Add("...")},
	// 		"navbar/content": style.D{style.Add("...")},
	//	}
	CustomStyle style.Custom
}

func (def D) class(k string) string {
	return style.CSSClass(def.Style, k, def.CustomStyle)
}
