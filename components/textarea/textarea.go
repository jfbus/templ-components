// Package textarea implements textarea fields.
package textarea

import (
	"github.com/a-h/templ"
	"github.com/jfbus/templ-components/components/form/validation/message"
	_ "github.com/jfbus/templ-components/components/input"
	"github.com/jfbus/templ-components/components/label"
	"github.com/jfbus/templ-components/components/position"
	"github.com/jfbus/templ-components/components/size"
	"github.com/jfbus/templ-components/components/style"
)

func init() {
	style.CopyDefaults("input/input", "textarea/input")
	style.CopyDefaults("input/icon", "textarea/icon")
}

// D is the definition for textarea fields.
type D struct {
	ID string
	// Name is the input name.
	Name  string
	Style style.Style
	// Type is the input type (text, password, ...).
	Label any
	// Value is the input value.
	Value string
	// Placeholder is the placeholder text displayed when no value is set.
	Placeholder string
	// Rows defines the number of rows
	Rows int
	// Disabled disables the input.
	Disabled bool
	// Size defines the input size (size.S, size.Normal (default) or size.L).
	Size size.Size
	// Loader displays a spinning loader when an HTMX action is triggered by the input.
	Loader bool
	// Icon displays an icon on the left side.
	Icon string
	// IconPosition can be position.Start (default) or position.End.
	IconPosition position.Position
	// Message adds a validation message below the field.
	// Just add &message.D{} to add automatic validation.
	//playground:import:github.com/jfbus/templ-components/components/form/validation/message
	//playground:default:&message.D{Message: "Validation message"}
	Message *message.D
	// Class overrides the default CSS class for the textarea.
	Class style.D
	// Class overrides the default CSS class for the icon.
	IconClass style.D
	// Attributes stores additional attributes (e.g. HTMX attributes).
	Attributes templ.Attributes
}

func (def D) style() style.Style {
	if def.Disabled {
		return def.Style | style.StyleDisabled
	}
	return def.Style
}

func (def D) iconClass() string {
	def.IconClass = append(def.IconClass, style.Replace("items-center", "items-top"))
	class := def.IconClass.CSSClass(def.style(), "textarea/icon")
	switch def.Size {
	case size.S:
		class += " pt-2"
	case size.L:
		class += " pt-4"
	default:
		class += " pt-3.5"
	}
	switch {
	case def.Icon == "":
		return ""
	case def.IconPosition == position.End:
		return class + " end-0 pe-3.5"
	default:
		return class + " start-0 ps-3.5"
	}
}

func (def D) iconSize() size.Size {
	if def.Size == size.Inherit {
		return size.Normal
	}
	return def.Size
}

func (def D) inputClass() string {
	class := def.Class.CSSClass(def.style(), "textarea/input")
	switch def.Size {
	case size.S:
		class += " p-2 text-xs"
	case size.L:
		class += " p-4 text-base"
	default:
		class += " p-2.5 text-sm"
	}
	switch {
	case def.Icon == "":
		return class
	case def.IconPosition == position.End:
		return class + " pe-10"
	default:
		return class + " ps-10"
	}
}

func (def D) id() string {
	if def.ID != "" {
		return def.ID
	}
	return def.Name
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
		if l.Style == style.StyleDefault {
			l.Style = def.style()
		}
		return l
	default:
		return label.D{}
	}
}

func (def D) message() message.D {
	m := *def.Message
	m.InputName = def.Name
	m.Size = def.Size
	m.Style = def.style()
	return m
}
