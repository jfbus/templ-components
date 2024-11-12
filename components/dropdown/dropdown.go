package dropdown

import (
	"github.com/a-h/templ"
	"github.com/jfbus/templui/components/a"
	"github.com/jfbus/templui/components/button"
	"github.com/jfbus/templui/components/style"
)

func init() {
	style.SetDefaults(style.Defaults{
		"dropdown": {
			style.Default: {
				style.Set("absolute z-10 divide-y rounded-lg border shadow w-44 mt-2"),
			},
		},
		"dropdown/link": {
			style.Default: {
				style.Set("block px-4 py-2"),
			},
		},
	})
}

type D struct {
	//playground:import:github.com/jfbus/templui/components/button
	//playground:default:button.D{Label:"Open dropdown"}
	Button button.D
	//playground:default:templ.Raw(`<div class="font-medium text-normal">Lorem ipsum</div><div class="text-sm text-gray-500">dolor sit amet</div>`)
	Header templ.Component
	//playground:import:github.com/jfbus/templui/components/a
	//playground:default:[][]a.D{{{Href:"#", Text: "Section 1 link 1"},{Href:"#", Text: "Section 1 link 2"}},{{Href:"#", Text: "Section 2 link 1"}}}
	Links [][]a.D
	// CustomStyle defines a custom style.
	// 	style.Custom{
	// 		"dropdown": style.D{style.Add("...")},
	// 		"dropdown/link": style.D{style.Add("...")},
	//	}
	CustomStyle style.Custom
}

func (def D) button() button.D {
	if def.Button.Attributes == nil {
		def.Button.Attributes = templ.Attributes{}
	}
	def.Button.Attributes["@click"] = "toggle()"
	return def.Button
}

func (def D) class() string {
	return style.CSSClass(style.Default, "dropdown", def.CustomStyle)
}

func (def D) links() [][]a.D {
	cc := style.Custom{
		"a": style.Compute(style.Default, "dropdown/link", def.CustomStyle),
	}
	for i := range def.Links {
		for j := range def.Links[i] {
			def.Links[i][j].CustomStyle = def.Links[i][j].CustomStyle.AddBefore(cc)
		}
	}
	return def.Links
}
