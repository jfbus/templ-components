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
	StyleButtonsRight style.Style = 1 << 8
)

func init() {
	style.SetDefaults(style.Defaults{
		"modal/background": {
			style.Default: {
				style.Set("fixed top-0 left-0 z-40 flex items-center justify-center w-full h-full px-4 py-5 bg-black/40"),
			},
		},
		"modal": {
			style.Default: {
				style.Set("w-full max-w-screen-md max-h-full rounded-[10px] bg-white dark:bg-dark-2 overflow-y-auto"),
			},
		},
		"modal/title": {
			style.Default: {
				style.Set("flex items-center justify-between p-4 md:p-5 border-b rounded-t dark:border-gray-600"),
			},
		},
		"modal/title/title": {
			style.Default: {
				style.Set("text-xl font-semibold text-gray-900 dark:text-white"),
			},
		},
		"modal/buttons": {
			style.Default: {
				style.Set("flex items-center p-4 md:p-5 border-t rounded-b gap-3"),
			},
			StyleButtonsRight: {
				style.Add("justify-end"),
			},
		},
		"modal/close": {
			style.Default: {
				style.Set("rounded-lg focus:ring-2 focus:ring-gray-300 p-1.5 inline-flex items-center justify-center"),
			},
		},
	})
}

type MaxWidth string

const (
	MaxWidthScreenMD MaxWidth = "max-w-screen-md"
	MaxWidthScreenSM MaxWidth = "max-w-screen-sm"
)

type D struct {
	// ID is the modal ID. It is mandatory.
	ID string
	// Style defines the modal style.
	Style style.Style
	// MaxWidth defines the modal max width ("max-w-screen-md" by default).
	MaxWidth MaxWidth
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
	// CustomStyle defines a custom style.
	// 	style.Custom{
	// 		"modal/background":  style.D{style.Add("...")},
	// 		"modal":             style.D{style.Add("...")},
	// 		"modal/title":       style.D{style.Add("...")},
	// 		"modal/title/title": style.D{style.Add("...")},
	// 		"modal/close":       style.D{style.Add("...")},
	// 		"modal/buttons":     style.D{style.Add("...")},
	//	}
	CustomStyle style.Custom
}

func (def D) id() string {
	return def.ID
}

func (def D) class(k string) string {
	return style.CSSClass(def.Style, k, def.CustomStyle)
}

func (def D) containerClass() string {
	s := style.Compute(def.Style, "modal", def.CustomStyle)
	if def.MaxWidth != "" {
		s = append(s, style.Replace(string(MaxWidthScreenMD), string(def.MaxWidth)))
	}
	return s.String()
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
