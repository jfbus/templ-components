// Package buttongroup implements button groups.
package buttongroup

import (
	"github.com/a-h/templ"
	"github.com/jfbus/templ-components/components/button"
	"github.com/jfbus/templ-components/components/size"
	"github.com/jfbus/templ-components/components/style"
)

func init() {
	style.SetDefaults(style.Defaults{
		"buttongroup": {
			style.StyleDefault: {
				style.Set("inline-flex rounded-md shadow-sm"),
			},
		},
		"buttongroup/button": {
			style.StyleDefault: {
				style.Set("border-t border-b border-e first:border first:rounded-s-lg last:border-t last:border-b last:border-e last:rounded-e-lg"),
			},
		},
	})
}

type D struct {
	// Buttons is the list of buttons to display.
	//playground:import:github.com/jfbus/templ-components/components/button
	//playground:default:[]button.D{{Label:"A"},{Label:"B"},{Label:"C"},{Label:"D"}}
	Buttons []button.D
	// Size defines the buttons size (shortcut for Buttons.Size).
	Size size.Size
	// CustomStyle defines a custom style.
	// 	style.Custom{
	// 		"buttongroup":        style.D{style.Add("...")},
	//		"buttongroup/button": style.D{style.Add("...")},
	//	}
	CustomStyle style.Custom
	// Attributes stores additional attributes (e.g. HTMX attributes).
	Attributes templ.Attributes
}

func (def D) buttons() []button.D {
	bcs := style.Custom{
		"button": style.Compute(style.StyleDefault, "buttongroup/button", def.CustomStyle),
	}
	bs := make([]button.D, len(def.Buttons))
	for i := range def.Buttons {
		bs[i] = def.Buttons[i]
		if def.Size != size.Inherit {
			bs[i].Size = def.Size
		}
		bs[i].CustomStyle = bs[i].CustomStyle.AddBefore(bcs)
	}
	return bs
}

func (def D) class(k string) string {
	return style.CSSClass(style.StyleDefault, k, def.CustomStyle)
}
