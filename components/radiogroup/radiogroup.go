package radiogroup

import (
	"github.com/jfbus/templ-components/components/label"
	"github.com/jfbus/templ-components/components/radio"
	"github.com/jfbus/templ-components/components/style"
)

const (
	StyleBordered        style.Style = 2
	StyleGrouped         style.Style = 4
	StyleGroupedVertical style.Style = 8
	StyleLabelOnly       style.Style = 16
)

var Defaults = style.Defaults{
	style.StyleDefault: {
		"RadioInputClass": {},
		"ContainerClass":  {},
	},
	StyleBordered: {
		"ContainerClass": {
			style.Class("space-x-4"),
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
	Radios []radio.D
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
