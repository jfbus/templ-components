package navbar

import (
	"github.com/a-h/templ"
	"github.com/jfbus/templ-components/components/style"
)

const (
	StyleSticky style.Style = 1 << 8
)

var Defaults = style.Defaults{
	style.StyleDefault: {
		"ContainerClass": {
			style.Class("w-full"),
		},
		"Class": {
			style.Class("max-w-screen-xl flex flex-wrap items-center justify-between mx-auto p-4"),
		},
	},
	StyleSticky: {
		"ContainerClass": {
			style.Class("fixed w-full z-40 top-0 start-0 border-b"),
			style.Color("bg-white dark:bg-gray-900 border-gray-200 dark:border-gray-600"),
		},
		"Class": {
			style.Class("max-w-screen-xl flex flex-wrap items-center justify-between mx-auto p-4"),
		},
	},
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
	return def.ContainerClass.CSSClass(Defaults, def.Style, "ContainerClass")
}

func (def D) class() string {
	return def.Class.CSSClass(Defaults, def.Style, "Class")
}
