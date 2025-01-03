// Package textarea implements textarea fields.
package textarea

import (
	"github.com/a-h/templ"
	"github.com/jfbus/templui/components/form/validation/message"
	_ "github.com/jfbus/templui/components/input"
	"github.com/jfbus/templui/components/label"
	"github.com/jfbus/templui/components/position"
	"github.com/jfbus/templui/components/size"
	"github.com/jfbus/templui/components/style"
	"github.com/jfbus/templui/components/tooltip"
)

func init() {
	style.CopyDefaults("input/input", "textarea/input")
	style.SetDefaults(style.Defaults{
		"textarea/icon": {
			style.Default: {
				style.Set("absolute inset-y-0 flex items-top pointer-events-none"),
			},
			style.SizeS: {
				style.Add("pt-2"),
			},
			style.SizeNormal: {
				style.Add("pt-3.5"),
			},
			style.SizeL: {
				style.Add("pt-4"),
			},
		},
	})
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
	//playground:import:github.com/jfbus/templui/components/form/validation/message
	//playground:default:&message.D{Message: "Validation message"}
	Message *message.D
	// Tooltip adds a tooltip to the input.
	Tooltip *tooltip.D
	// CustomStyle defines a custom style.
	// 	style.Custom{
	// 		"textarea":       style.D{style.Add("...")},
	// 		"textarea/input": style.D{style.Add("...")},
	// 		"textarea/icon":  style.D{style.Add("...")},
	// 		"textarea/label":  style.D{style.Add("...")},
	//	}
	CustomStyle style.Custom
	Attributes  templ.Attributes
}

func (def D) size() size.Size {
	if def.Size == 0 {
		return size.Normal
	}
	return def.Size
}

func (def D) style() style.Style {
	st := def.Style | style.Size(def.size())
	if def.Disabled {
		return st | style.Disabled
	}
	return st
}

func (def D) class(k string) string {
	return style.CSSClass(def.style(), k, def.CustomStyle)
}

func (def D) iconClass() string {
	class := style.CSSClass(def.style(), "textarea/icon", def.CustomStyle)

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
	class := style.CSSClass(def.style(), "textarea/input", def.CustomStyle)
	if def.Tooltip != nil {
		class += " " + def.Tooltip.Class()
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
	cc := style.Custom{
		"label": style.Compute(def.Style, "textarea/label", def.CustomStyle),
	}
	switch l := def.Label.(type) {
	case string:
		return label.D{
			InputID:     def.id(),
			Label:       l,
			Style:       def.style(),
			CustomStyle: cc,
		}
	case label.D:
		l.InputID = def.id()
		if l.Style == style.Default {
			l.Style = def.style()
		}
		l.CustomStyle = l.CustomStyle.AddBefore(cc)
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
	m.CustomStyle = def.CustomStyle
	return m
}
