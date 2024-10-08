package navbar

import (
	"github.com/a-h/templ"
	"github.com/jfbus/templ-components/components/style"
)

const (
	StyleSticky style.Style = 1 << 8
)

func init() {
	style.SetDefaults(style.Defaults{
		"navbar": {
			style.StyleDefault: {
				style.Set("w-full"),
			},
			StyleSticky: {
				style.Set("fixed w-full z-40 top-0 start-0 border-b"),
			},
		},
		"navbar/div": {
			style.StyleDefault: {
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
	//playground:import:github.com/jfbus/templ-components/components/a
	//playground:import:github.com/jfbus/templ-components/components/button
	//playground:import:github.com/jfbus/templ-components/components/dropdown
	//playground:import:github.com/jfbus/templ-components/components/icon
	//playground:import:github.com/jfbus/templ-components/components/size
	//playground:default:[]templ.Component{a.C(a.D{Text:"Your Brand"}), dropdown.C(dropdown.D{Button: button.D{Style: button.StyleNoBorder, Icon: icon.CircleUserRound, Size: size.XL}, Links: [][]a.D{{{Text: "Profile"}}, {{Text: "Logout"}}}})}
	Sections       []templ.Component
	ContainerClass style.D
	Class          style.D
}

func (def D) containerClass() string {
	return def.ContainerClass.CSSClass(def.Style, "navbar")
}

func (def D) class() string {
	return def.Class.CSSClass(def.Style, "navbar/div")
}
