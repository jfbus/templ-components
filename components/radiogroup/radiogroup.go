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
		"InputClass":     {},
		"ContainerClass": {},
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
	Inputs              []radio.D
	ContainerClass      style.D
	InputContainerClass style.D
	InputClass          style.D
	InputLabelClass     style.D
}

func (def D) containerClass() string {
	return def.ContainerClass.CSSClass(style.Default(Defaults, def.Style, "ContainerClass"))
}

func (def D) inputs() []radio.D {
	for i := range def.Inputs {
		def.Inputs[i].ID = def.Name + "-" + def.Inputs[i].Value
		def.Inputs[i].Name = def.Name
		def.Inputs[i].Style = def.Style
		def.Inputs[i].InputClass.Class = helper.IfEmpty(def.Inputs[i].InputClass.Class, def.InputClass.Class)
		def.Inputs[i].InputClass.Color = helper.IfEmpty(def.Inputs[i].InputClass.Color, def.InputClass.Color)
		def.Inputs[i].ContainerClass.Class = helper.IfEmpty(def.Inputs[i].ContainerClass.Class, def.InputContainerClass.Class)
		def.Inputs[i].ContainerClass.Color = helper.IfEmpty(def.Inputs[i].ContainerClass.Color, def.InputContainerClass.Color)
		switch l := def.Inputs[i].Label.(type) {
		case label.D:
			l.Class.Color = helper.IfEmpty(l.Class.Color, def.InputLabelClass.Color)
			l.Class.Class = helper.IfEmpty(l.Class.Class, def.InputLabelClass.Class)
			def.Inputs[i].Label = l
		default:
			def.Inputs[i].Label = label.D{
				Label: l,
				Class: def.InputLabelClass,
			}
		}
	}
	return def.Inputs
}

func (def D) inputContainerClass() string {
	return def.InputContainerClass.CSSClass(style.Default(Defaults, def.Style, "InputContainerClass"))
}
