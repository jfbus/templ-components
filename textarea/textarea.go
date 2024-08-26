package textarea

import (
	"github.com/a-h/templ"
	"github.com/jfbus/templ-components/position"
	"github.com/jfbus/templ-components/size"
)

// D is the definition for input fields.
// Usage: @textarea.C(textarea.D{})
type D struct {
	// Name is the input name.
	Name string
	// Type is the input type (text, password, ...).
	Label string
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
	// Class defines additional CSS class for the container.
	Class string
	// Attributes stores additional attributes (e.g. HTMX attributes).
	Attributes templ.Attributes
}

func (def D) iconClass() string {
	class := "absolute inset-y-0 flex items-top pointer-events-none"
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

func (def D) inputClass() string {
	class := "block w-full bg-gray-50 border border-gray-300 text-gray-900 rounded-lg focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
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
