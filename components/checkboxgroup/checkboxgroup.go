package checkboxgroup

import (
	"github.com/jfbus/templui/components/checkbox"
	"github.com/jfbus/templui/components/form/validation/message"
	"github.com/jfbus/templui/components/style"
)

const (
	StyleHorizontal        style.Style = 1 << 8
	StyleBordered          style.Style = 1 << 9
	StyleGrouped           style.Style = 1 << 10
	StyleGroupedHorizontal style.Style = 1 << 11
	StyleLabelOnly         style.Style = 1 << 12
)

func init() {
	style.SetDefaults(style.Defaults{
		"checkboxgroup": {
			StyleHorizontal: {
				style.Set("flex flex-col sm:flex-row sm:gap-4"),
			},
			StyleBordered: {
				style.Set("flex flex-col sm:flex-row sm:gap-4"),
			},
			StyleGrouped: {
				style.Set("border rounded-lg"),
			},
			StyleGroupedHorizontal: {
				style.Set("sm:flex border rounded-lg"),
			},
			StyleLabelOnly: {
				style.Set("inline-flex items-center justify-between"),
			},
		},
		"checkboxgroup/checkbox": {
			style.Default: {
				style.Set("flex items-center"),
			},
			StyleBordered: {
				style.Set("flex items-center px-4 border rounded w-full"),
			},
			StyleGrouped: {
				style.Set("flex items-center border-b last:border-b-0 px-4"),
			},
			StyleGroupedHorizontal: {
				style.Set("flex items-center border-b sm:border-b-0 sm:border-r last:border-0 px-4 sm:w-full"),
			},
		},
		"checkboxgroup/checkbox/input": {
			StyleLabelOnly: {
				style.Set("hidden peer"),
			},
		},
		"checkboxgroup/checkbox/label": {
			style.Default: {
				style.Add("py-3"),
			},
			StyleLabelOnly: {
				style.Set("border p-2 rounded-lg cursor-pointer"),
			},
		},
	})
}

type D struct {
	// Name is the Name of all inputs.
	Name string
	// Style is the checkboxgroup style.
	Style style.Style
	// Checkboxes is the list of checkboxes in the group.
	//playground:import:github.com/jfbus/templui/components/checkbox
	//playground:default:[]checkbox.D{{Name: "foo", Value: "1", Label: "Choice 1"},{Name: "foo", Value: "2", Label:"Choice 2"}}
	Checkboxes []checkbox.D
	// Message adds a validation message below the field.
	// Just add &message.D{} to add automatic validation.
	//playground:import:github.com/jfbus/templui/components/form/validation/message
	//playground:default:&message.D{Message: "Validation message"}
	Message *message.D
	// CustomStyle defines a custom style.
	// 	style.Custom{
	// 		"checkboxgroup":       style.D{style.Add("...")},
	// 		"checkboxgroup/checkbox":       style.D{style.Add("...")},
	// 		"checkboxgroup/checkbox/input": style.D{style.Add("...")},
	// 		"checkboxgroup/checkbox/label": style.D{style.Add("...")},
	//	}
	CustomStyle style.Custom
}

func (def D) class() string {
	return style.CSSClass(def.Style, "checkboxgroup", def.CustomStyle)
}

func (def D) radios() []checkbox.D {
	cc := style.Custom{
		"checkbox":       style.Compute(def.Style, "checkboxgroup/checkbox", def.CustomStyle),
		"checkbox/input": style.Compute(def.Style, "checkboxgroup/checkbox/input", def.CustomStyle),
		"checkbox/label": style.Compute(def.Style, "checkboxgroup/checkbox/label", def.CustomStyle),
	}
	for i := range def.Checkboxes {
		def.Checkboxes[i].ID = def.Name + "-" + def.Checkboxes[i].Value
		def.Checkboxes[i].Name = def.Name
		def.Checkboxes[i].CustomStyle = def.Checkboxes[i].CustomStyle.AddBefore(cc)
	}
	return def.Checkboxes
}

func (def D) message() message.D {
	m := *def.Message
	m.InputName = def.Name
	m.Style = def.Style
	m.CustomStyle = def.CustomStyle
	return m
}
