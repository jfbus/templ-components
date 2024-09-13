package dropdown

import (
	"github.com/a-h/templ"
	"github.com/jfbus/templ-components/components/a"
	"github.com/jfbus/templ-components/components/button"
	"github.com/jfbus/templ-components/components/style"
)

var Defaults = style.Defaults{
	style.StyleDefault: {
		"Class": {
			style.Class("absolute z-10 divide-y rounded-lg shadow w-44 mt-2"),
			style.Color("bg-white divide-gray-100 dark:bg-gray-700 dark:divide-gray-600 mt-2"),
		},
		"LinkClass": {
			style.Class("block px-4 py-2"),
			style.Color("hover:bg-gray-100 dark:hover:bg-gray-600 dark:hover:text-white"),
		},
	},
}

type D struct {
	//playground:import:github.com/jfbus/templ-components/components/button
	//playground:default:button.D{Label:"Open dropdown"}
	Button button.D
	//playground:default:templ.Raw(`<div class="font-medium text-normal">Lorem ipsum</div><div class="text-sm text-gray-500">dolor sit amet</div>`)
	Header templ.Component
	//playground:import:github.com/jfbus/templ-components/components/a
	//playground:default:[][]a.D{{{Href:"#", Text: "Section 1 link 1"},{Href:"#", Text: "Section 1 link 2"}},{{Href:"#", Text: "Section 2 link 1"}}}
	Links     [][]a.D
	Class     style.D
	LinkClass style.D
}

func (def D) button() button.D {
	if def.Button.Attributes == nil {
		def.Button.Attributes = templ.Attributes{}
	}
	def.Button.Attributes["@click"] = "toggle()"
	return def.Button
}

func (def D) class() string {
	return def.Class.CSSClass(Defaults, style.StyleDefault, "Class")
}

func (def D) links() [][]a.D {
	lc := append(Defaults[style.StyleDefault]["LinkClass"], def.LinkClass...)
	for i := range def.Links {
		for j := range def.Links[i] {
			def.Links[i][j].Class = append(lc, def.Links[i][j].Class...)
		}
	}
	return def.Links
}
