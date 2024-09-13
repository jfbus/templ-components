package form

import (
	"github.com/a-h/templ"
)

type D struct {
	// ID is the tag ID, used by validation.Retarget.
	ID string
	// Attributes lists the custom attributes to add to the form tag.
	Attributes templ.Attributes
}
