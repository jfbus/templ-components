package radiogroup

import (
	"github.com/jfbus/templ-components/components/helper"
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
			Class: "space-x-4",
		},
	},
	StyleGrouped: {
		"ContainerClass": {
			Class: "border rounded-lg",
			Color: "border-gray-200 dark:bg-gray-700 dark:border-gray-600",
		},
	},
	StyleGroupedVertical: {
		"ContainerClass": {
			Class: "border rounded-lg",
			Color: "border-gray-200 dark:bg-gray-700 dark:border-gray-600",
		},
	},
}

type D struct {
	// Name is the Name of all inputs.
	Name string
	//
	Style               style.Style
	Radios              []radio.D
	ContainerClass      style.D
	RadioContainerClass style.D
	RadioInputClass     style.D
	RadioLabelClass     style.D
}

func (def D) containerClass() string {
	return def.ContainerClass.CSSClass(style.Default(Defaults, def.Style, "ContainerClass"))
}

func (def D) radios() []radio.D {
	for i := range def.Radios {
		def.Radios[i].ID = def.Name + "-" + def.Radios[i].Value
		def.Radios[i].Name = def.Name
		def.Radios[i].Style = def.Style
		def.Radios[i].InputClass.Class = helper.IfEmpty(def.Radios[i].InputClass.Class, def.RadioInputClass.Class)
		def.Radios[i].InputClass.Color = helper.IfEmpty(def.Radios[i].InputClass.Color, def.RadioInputClass.Color)
		def.Radios[i].ContainerClass.Class = helper.IfEmpty(def.Radios[i].ContainerClass.Class, def.RadioContainerClass.Class)
		def.Radios[i].ContainerClass.Color = helper.IfEmpty(def.Radios[i].ContainerClass.Color, def.RadioContainerClass.Color)
		switch l := def.Radios[i].Label.(type) {
		case label.D:
			l.Class.Color = helper.IfEmpty(l.Class.Color, def.RadioLabelClass.Color)
			l.Class.Class = helper.IfEmpty(l.Class.Class, def.RadioLabelClass.Class)
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

func (def D) inputContainerClass() string {
	return def.RadioContainerClass.CSSClass(style.Default(Defaults, def.Style, "RadioContainerClass"))
}
