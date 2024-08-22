package form

import "github.com/a-h/templ"

// InputFieldDefinition is the definition for input fields.
// Usage: @form.InputField(form.InputFieldDefinition{})
type InputFieldDefinition struct {
	// Name is the input name.
	Name string
	// Type is the input type (text, password, ...).
	Type string
	// Label is the input label.
	Label string
	// Value is the input value.
	Value string
	// Placeholder is the placeholder text displayed when no value is set.
	Placeholder string
	// Loader displays a spinning loader when an HTMX action is triggered bu the input.
	Loader bool
	// Icon displays an icon on the left side.
	Icon string
	// Attributes stores additional attributes (e.g. HTMX attributes).
	Attributes templ.Attributes
}
