package radiogroup

import (
	"github.com/jfbus/templ-components/components/form/validation/message"
	"github.com/jfbus/templ-components/components/label"
	"github.com/jfbus/templ-components/components/radio"
	"github.com/jfbus/templ-components/components/style"
)

const (
	StyleBordered        style.Style = 1 << 8
	StyleGrouped         style.Style = 1 << 9
	StyleGroupedVertical style.Style = 1 << 10
	StyleLabelOnly       style.Style = 1 << 11
)

var Defaults = style.Defaults{
	style.StyleDefault: {
		"RadioInputClass": {},
		"ContainerClass":  {},
	},
	StyleBordered: {
		"ContainerClass": {
			style.Class("gap-4"),
		},
	},
	StyleGrouped: {
		"ContainerClass": {
			style.Class("border rounded-lg"),
			style.Color("border-gray-200 dark:bg-gray-700 dark:border-gray-600"),
		},
	},
	StyleGroupedVertical: {
		"ContainerClass": {
			style.Class("border rounded-lg"),
			style.Color("border-gray-200 dark:bg-gray-700 dark:border-gray-600"),
		},
	},
}

type D struct {
	// Name is the Name of all inputs.
	Name string
	// Style is the radiogroup style.
	Style style.Style
	// Radios is the list of radios in the group.
	//playground:import:github.com/jfbus/templ-components/components/radio
	//playground:default:[]radio.D{{Name: "foo", Value: 1, Label: "Choice 1"},{Name: "foo", Value: 2, Label:"Choice 2"}}
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
	return def.ContainerClass.CSSClass(Defaults, def.Style, "ContainerClass")
}

func (def D) radios() []radio.D {
	for i := range def.Radios {
		def.Radios[i].ID = def.Name + "-" + def.Radios[i].Value
		def.Radios[i].Name = def.Name
		def.Radios[i].Style = def.Style
		def.Radios[i].InputClass = append(def.RadioInputClass, def.Radios[i].InputClass...)
		def.Radios[i].ContainerClass = append(def.RadioContainerClass, def.Radios[i].ContainerClass...)
		switch l := def.Radios[i].Label.(type) {
		case label.D:
			l.Class = append(def.RadioLabelClass, l.Class...)
			def.Radios[i].Label = l
		default:
			def.Radios[i].Label = label.D{
				Label: l,
				Class: def.RadioLabelClass,
			}
		}
	}
	return def.Radios
}

func (def D) message() message.D {
	m := *def.Message
	m.InputName = def.Name
	m.Style = def.Style
	return m
}
