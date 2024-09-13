package modal

import (
	"context"
	"net/http"

	"github.com/a-h/templ"
	"github.com/jfbus/templ-components/components/button"
	"github.com/jfbus/templ-components/components/form"
	"github.com/jfbus/templ-components/components/style"
)

const (
	StyleButtonsRight style.Style = 1
)

var Defaults = style.Defaults{
	style.StyleDefault: {
		"Class": {},
		"ButtonContainerClass": {
			style.Class("flex items-center p-4 md:p-5 border-t rounded-b gap-3"),
			style.Color("border-gray-200 dark:border-gray-600"),
		},
	},
	StyleButtonsRight: {
		"ButtonContainerClass": {
			style.Add("justify-end"),
		},
	},
}

type D struct {
	// ID is the modal ID. It is mandatory.
	ID string
	// Style defines the modal style.
	Style style.Style
	// Title is the modal title.
	Title string
	// Form defines the optional form tag that will enclose both content and buttons.
	Form *form.D
	// Content is the modal content (as an alternative, you can add the content as the component children).
	Content any
	// Close is the close button. It closes the modal without submitting the form.
	//playground:import:github.com/jfbus/templ-components/components/button
	//playground:default:&button.D{Label:"Cancel", Style: button.StyleOutline}
	Close *button.D
	// Confirm is the confirm button. It submits the modal form but does not close the modal.
	// Add modal.Close() to your success response to close the modal.
	//playground:default:&button.D{Label:"OK"}
	Confirm *button.D
	// Class adds custom CSS classes.
	Class style.D
	// ButtonContainerClass defines the class for the buttons containers.
	ButtonContainerClass style.D
}

func (def D) id() string {
	return def.ID
}

func (def D) class() string {
	return def.Class.CSSClass(Defaults, def.Style, "Class")
}

func (def D) buttonContainerClass() string {
	return def.Class.CSSClass(Defaults, def.Style, "ButtonContainerClass")
}

func (def D) close() button.D {
	if def.Close.Attributes == nil {
		def.Close.Attributes = templ.Attributes{}
	}
	def.Close.Attributes["@click.stop"] = "$refs." + def.id() + ".remove()"
	return *def.Close
}

func (def D) confirm() button.D {
	return *def.Confirm
}

func (def D) form() form.D {
	f := *def.Form
	f.ID = def.id()
	if f.Attributes == nil {
		f.Attributes = templ.Attributes{}
	}
	f.Attributes["class"] = "w-full"
	return f
}

// Retarget returns a HTMX response that retargets the response to display the response
// at the end of the body, ignoring the initial target.
// When using error status codes, do not forget to add your error codes to the list of
// codes for which HTMX swaps the content - https://htmx.org/docs/#requests
func Retarget(ctx context.Context, def D, w http.ResponseWriter, statusCode int) error {
	w.Header().Set("Content-Type", "text/html, charset=UTF-8")
	w.Header().Set("HX-Retarget", "body")
	w.Header().Set("HX-Reswap", "beforeend")
	w.WriteHeader(statusCode)
	return C(def).Render(ctx, w)
}
