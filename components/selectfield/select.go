// Package selectfield implements select fields.
// `select` being a reserved keyword, the package cannot be named select.
package selectfield

import (
	"github.com/a-h/templ"
	"github.com/jfbus/templ-components/components/form/validation/message"
	_ "github.com/jfbus/templ-components/components/input"
	"github.com/jfbus/templ-components/components/label"
	"github.com/jfbus/templ-components/components/selectfield/option"
	"github.com/jfbus/templ-components/components/size"
	"github.com/jfbus/templ-components/components/style"
)

func init() {
	style.CopyDefaults("input/input", "select/input")
}

// D is the select definition.
type D struct {
	// ID is the tag id.
	ID string
	// Name is the tag name.
	Name string
	// Style is the style (style.StyleDefault, style.StyleValid or style.StyleInvalid).
	Style style.Style
	// Label is the label (either a string or a label.D).
	Label any
	// Options is the list of options.
	//playground:import:github.com/jfbus/templ-components/components/selectfield/option
	//playground:default:[]option.D{{Label:"Select a value"},{Label:"Option 1"},{Label:"Option 2"}}
	Options []option.D
	// Select is the selected value.
	Selected string
	// Disabled disables the select.
	Disabled bool
	// Size is the size.
	Size size.Size
	// Message adds a validation message below the field.
	// Just add &message.D{} to add automatic validation.
	//playground:import:github.com/jfbus/templ-components/components/form/validation/message
	//playground:default:&message.D{Message: "Validation message"}
	Message *message.D
	// CustomStyle defines a custom style.
	// 	style.Custom{
	// 		"select":       style.D{style.Add("...")},
	// 		"select/input": style.D{style.Add("...")},
	// 		"select/label": style.D{style.Add("...")},
	//	}
	CustomStyle style.Custom
	// Attributes add additional attributes.
	Attributes templ.Attributes
}

func (def D) id() string {
	if def.ID != "" {
		return def.ID
	}
	return def.Name
}

func (def D) style() style.Style {
	if def.Disabled {
		return def.Style | style.StyleDisabled
	}
	return def.Style
}

func (def D) label() label.D {
	switch l := def.Label.(type) {
	case string:
		return label.D{
			InputID: def.id(),
			Label:   l,
			Style:   def.style(),
		}
	case label.D:
		l.InputID = def.id()
		return l
	default:
		return label.D{}
	}
}

func (def D) class(k string) string {
	return style.CSSClass(def.style(), k, def.CustomStyle)
}

func (def D) inputClass() string {
	class := style.CSSClass(def.style(), "select/input", def.CustomStyle)
	switch def.Size {
	case size.S:
		class += " p-2 text-xs"
	case size.L:
		class += " p-4 text-base"
	default:
		class += " p-2.5 text-sm"
	}
	return class
}

func (def D) message() message.D {
	m := *def.Message
	m.InputName = def.Name
	m.Size = def.Size
	m.Style = def.style()
	m.CustomStyle = def.CustomStyle
	return m
}
