package checkboxgroup

import (
	"github.com/jfbus/templ-components/components/checkbox"
	"github.com/jfbus/templ-components/components/form/validation/message"
	"github.com/jfbus/templ-components/components/style"
)

const (
	StyleHorizontal        style.Style = 1 << 8
	StyleBordered          style.Style = 1 << 9
	StyleGrouped           style.Style = 1 << 10
	StyleGroupedHorizontal style.Style = 1 << 11
	StyleLabelOnly         style.Style = 1 << 12
)

var Defaults = style.Defaults{
	style.StyleDefault: {
		"CheckboxContainerClass": {
			style.Class("flex items-center"),
		},
		"CheckboxLabelClass": {
			style.Add("py-3"),
		},
	},
	StyleHorizontal: {
		"ContainerClass": {
			style.Class("flex flex-col sm:flex-row sm:gap-4"),
		},
	},
	StyleBordered: {
		"CheckboxContainerClass": {
			style.Class("flex items-center px-4 border rounded w-full"),
			style.Color("border-gray-200 dark:border-gray-700"),
		},
		"ContainerClass": {
			style.Class("flex flex-col sm:flex-row sm:gap-4"),
		},
	},
	StyleGrouped: {
		"ContainerClass": {
			style.Class("border rounded-lg"),
			style.Color("border-gray-200 dark:bg-gray-700 dark:border-gray-600"),
		},
		"CheckboxContainerClass": {
			style.Class("flex items-center border-b last:border-b-0 px-4"),
		},
	},
	StyleGroupedHorizontal: {
		"ContainerClass": {
			style.Class("sm:flex border rounded-lg"),
			style.Color("border-gray-200 dark:bg-gray-700 dark:border-gray-600"),
		},
		"CheckboxContainerClass": {
			style.Class("flex items-center border-b sm:border-b-0 sm:border-r last:border-0 px-4 sm:w-full"),
		},
	},
	StyleLabelOnly: {
		"CheckboxContainerClass": {
			style.Class("inline-flex items-center justify-between"),
		},
		"CheckboxInputClass": {
			style.Color(""),
			style.Class("hidden peer"),
		},
		"CheckboxLabelClass": {
			style.Class("border p-2 rounded-lg cursor-pointer"),
			style.Color("text-gray-500 bg-white border-gray-200 dark:hover:text-gray-300 dark:border-gray-700 dark:peer-checked:text-blue-500 peer-checked:border-blue-600 peer-checked:text-blue-600 hover:text-gray-600 hover:bg-gray-100 dark:text-gray-400 dark:bg-gray-800 dark:hover:bg-gray-700"),
		},
	},
}

type D struct {
	// Name is the Name of all inputs.
	Name string
	// Style is the checkboxgroup style.
	Style style.Style
	// Checkboxes is the list of checkboxes in the group.
	//playground:import:github.com/jfbus/templ-components/components/checkbox
	//playground:default:[]checkbox.D{{Name: "foo", Value: "1", Label: "Choice 1"},{Name: "foo", Value: "2", Label:"Choice 2"}}
	Checkboxes []checkbox.D
	// Message adds a validation message below the field.
	// Just add &message.D{} to add automatic validation.
	//playground:import:github.com/jfbus/templ-components/components/form/validation/message
	//playground:default:&message.D{Message: "Validation message"}
	Message *message.D
	// ContainerClass overrides the class of the div container.
	ContainerClass style.D
	// CheckboxContainerClass overrides the class of each radio div container.
	CheckboxContainerClass style.D
	// CheckboxContainerClass overrides the class of each radio input.
	CheckboxInputClass style.D
	// CheckboxContainerClass overrides the class of each radio label.
	CheckboxLabelClass style.D
}

func (def D) containerClass() string {
	return def.ContainerClass.CSSClass(Defaults, def.Style, "ContainerClass")
}

func (def D) radios() []checkbox.D {
	for i := range def.Checkboxes {
		def.Checkboxes[i].ID = def.Name + "-" + def.Checkboxes[i].Value
		def.Checkboxes[i].Name = def.Name
		def.Checkboxes[i].InputClass = append(def.CheckboxInputClass.WithDefault(Defaults, def.Style, "CheckboxInputClass"), def.Checkboxes[i].InputClass...)
		def.Checkboxes[i].ContainerClass = append(def.CheckboxContainerClass.WithDefault(Defaults, def.Style, "CheckboxContainerClass"), def.Checkboxes[i].ContainerClass...)
		def.Checkboxes[i].LabelClass = append(def.CheckboxLabelClass.WithDefault(Defaults, def.Style, "CheckboxLabelClass"), def.Checkboxes[i].LabelClass...)
	}
	return def.Checkboxes
}

func (def D) message() message.D {
	m := *def.Message
	m.InputName = def.Name
	m.Style = def.Style
	return m
}
