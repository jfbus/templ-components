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
	// ContainerClass overrides the class of the div container.
	ContainerClass style.D
	// RadioContainerClass overrides the class of each radio div container.
	RadioContainerClass style.D
	// RadioContainerClass overrides the class of each radio input.
	RadioInputClass style.D
	// RadioContainerClass overrides the class of each radio label.
	RadioLabelClass style.D
}

func (def D) containerClass() string {
	return def.ContainerClass.CSSClass(def.Style, "radiogroup")
}

func (def D) radios() []radio.D {
	ric := def.RadioInputClass.WithDefault(def.Style, "radiogroup/radio/input")
	rcc := def.RadioContainerClass.WithDefault(def.Style, "radiogroup/radio")
	rlc := def.RadioLabelClass.WithDefault(def.Style, "radiogroup/radio/label")
	for i := range def.Radios {
		def.Radios[i].ID = def.Name + "-" + def.Radios[i].Value
		def.Radios[i].Name = def.Name
		def.Radios[i].InputClass = append(ric, def.Radios[i].InputClass...)
		def.Radios[i].ContainerClass = append(rcc, def.Radios[i].ContainerClass...)
		def.Radios[i].LabelClass = append(rlc, def.Radios[i].LabelClass...)
	}
	return def.Radios
}

func (def D) message() message.D {
	m := *def.Message
	m.InputName = def.Name
	m.Style = def.Style
	return m
}
