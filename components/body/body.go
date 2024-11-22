package body

import (
	"strconv"

	"github.com/a-h/templ"
	"github.com/jfbus/templui/components/footer"
	"github.com/jfbus/templui/components/navbar"
	"github.com/jfbus/templui/components/sidebar"
	"github.com/jfbus/templui/components/size"
	"github.com/jfbus/templui/components/style"
	"github.com/jfbus/templui/components/toast/container"
)

type NavbarHeight size.Size

func (h NavbarHeight) paddingTop() string {
	return "pt-14"
}

type D struct {
	// Navbar defines an optional navbar.
	Navbar *navbar.D
	// Sidebar defines an optional sidebar.
	Sidebar *sidebar.D
	// Toasts defines a custom toast container (a default one will be added otherwise).
	Toasts *container.D
	// Footer defines an optional footer.
	Footer *footer.D
	// State defines a default Alpine.js state.
	State        map[string]string
	NavbarHeight NavbarHeight
	// CustomStyle defines a custom style.
	// 	style.Custom{
	// 		"layout":       style.D{style.Add("...")},
	//	}
	CustomStyle style.Custom
	// Attributes adds additional attributes (e.g. HTMX attributes).
	Attributes templ.Attributes
}

func (def D) state() string {
	st := `{
  sidebar: false`
	if def.State != nil {
		for k, v := range def.State {
			st += `,
  ` + k + ": " + strconv.Quote(v)
		}
	}
	return st + "}"
}

func (def D) sidebar() sidebar.D {
	d := *def.Sidebar
	d.CustomStyle = d.CustomStyle.AddBefore(style.Custom{"sidebar": {style.Add(def.NavbarHeight.paddingTop())}})
	return d
}

func (def D) toasts() container.D {
	if def.Toasts == nil {
		return container.D{
			CustomStyle: style.Custom{"sidebar": {style.Add(def.NavbarHeight.paddingTop())}},
		}
	}
	d := *def.Toasts
	d.CustomStyle = d.CustomStyle.AddBefore(style.Custom{"toast/container": {style.Add(def.NavbarHeight.paddingTop())}})
	return d
}

func (def D) footer() footer.D {
	d := *def.Footer
	return d
}

func (def D) class() string {
	return style.CSSClass(style.Default, "layout", def.CustomStyle)
}
