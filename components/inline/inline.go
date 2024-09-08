// Package inline implements inline edits.
package inline

import (
	"github.com/a-h/templ"
	"github.com/jfbus/templ-components/components/size"
)

// D is the definition for inline edits.
type D struct {
	// Value defines the current value.
	Value string
	// IconSize defines the size of the icon.
	IconSize size.Size
	// Edit is the component used to edit the value.
	//playground:import:github.com/jfbus/templ-components/components/input
	//playground:default:input.C(input.D{Name:"edit"})
	Edit templ.Component
	// DefaultEdit defines if the default display is the Edit component or the text value.
	DefaultEdit bool
	// Focus defines the input field that needs to have the focus after. You'll need to add a "x-ref" attribute to this input.
	Focus string
}

func (def D) defaultState() string {
	if def.DefaultEdit {
		return "{edit: true}"
	}
	return "{edit: false}"
}

func (def D) click() string {
	if def.Focus == "" {
		return "edit = true"
	}
	return "edit = true; $nextTick(() => { $refs." + def.Focus + ".focus(); });"
}
