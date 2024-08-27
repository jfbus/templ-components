package inline

import (
	"github.com/a-h/templ"
	"github.com/jfbus/templ-components/size"
)

// D is the definition for inline edits.
type D struct {
	// Value defines the current value.
	Value string
	// IconSize defines the size of the icon.
	IconSize size.Size
	// Edit is the component used to edit the value.
	Edit templ.Component
}
