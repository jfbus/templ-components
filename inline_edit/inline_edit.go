package inline_edit

import (
	"github.com/a-h/templ"
	"github.com/jfbus/templ-components/size"
)

// D is the definition for inline edits.
// Usage: @inline_edit.C(inline_edit.D{Value: "foo", Edit: input_field.C(...)})
type D struct {
	// Value defines the current value.
	Value string
	// IconSize defines the size of the icon.
	IconSize size.Size
	// Edit is the component used to edit the value.
	Edit templ.Component
}
