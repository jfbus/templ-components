package form

import (
	"github.com/a-h/templ"
	"github.com/jfbus/templ-components/size"
)

// InputFieldDefinition is the definition for input fields.
// Usage: @form.InlineEdit(form.InlineEditDefinition{Value: "foo", Edit: form.Input(...)})
type InlineEditDefinition struct {
	Value    string
	IconSize size.Size
	Edit     templ.Component
}
