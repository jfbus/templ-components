package radiogroup

import (
	"github.com/jfbus/templ-components/components/form/validation/message"
	"github.com/jfbus/templ-components/components/radio"
	"github.com/jfbus/templ-components/components/style"
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
		"radiogroup": {
			StyleHorizontal: {
				style.Set("flex flex-col sm:flex-row sm:gap-4"),
			},
			StyleBordered: {
				style.Set("flex flex-col sm:flex-row gap-4"),
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
		"radiogroup/radio": {
			style.StyleDefault: {
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
		"radiogroup/radio/input": {
			StyleLabelOnly: {
				style.Set("hidden peer"),
			},
		},
		"radiogroup/radio/label": {
			style.StyleDefault: {
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
	// Style is the radiogroup style.
	Style style.Style
	// Radios is the list of radios in the group.
	//playground:import:github.com/jfbus/templ-components/components/radio
	//playground:default:[]radio.D{{Name: "foo", Value: "1", Label: "Choice 1"},{Name: "foo", Value: "2", Label:"Choice 2"}}
	Radios []radio.D
	// Message adds a validation message below the field.
	// Just add &message.D{} to add automatic validation.
	//playground:import:github.com/jfbus/templ-components/components/form/validation/message
	//playground:default:&message.D{Message: "Validation message"}
	Message *message.D
	// CustomStyle defines a custom style.
	// 	style.Custom{
	// 		"radiogroup":       style.D{style.Add("...")},
	// 		"radiogroup/radio":       style.D{style.Add("...")},
	// 		"radiogroup/radio/input": style.D{style.Add("...")},
	// 		"radiogroup/radio/label": style.D{style.Add("...")},
	//	}
	CustomStyle style.Custom
}

func (def D) class() string {
	return style.CSSClass(def.Style, "radiogroup", def.CustomStyle)
}

func (def D) radios() []radio.D {
	cc := style.Custom{
		"radio":       style.Compute(def.Style, "radiogroup/radio", def.CustomStyle),
		"radio/input": style.Compute(def.Style, "radiogroup/radio/input", def.CustomStyle),
		"radio/label": style.Compute(def.Style, "radiogroup/radio/label", def.CustomStyle),
	}
	for i := range def.Radios {
		def.Radios[i].ID = def.Name + "-" + def.Radios[i].Value
		def.Radios[i].Name = def.Name
		def.Radios[i].CustomStyle = def.Radios[i].CustomStyle.AddBefore(cc)
	}
	return def.Radios
}

func (def D) message() message.D {
	m := *def.Message
	m.InputName = def.Name
	m.Style = def.Style
	m.CustomStyle = def.CustomStyle
	return m
}
