package body

import (
	"strconv"

	"github.com/a-h/templ"
	"github.com/jfbus/templui/components/footer"
	"github.com/jfbus/templui/components/navbar"
	"github.com/jfbus/templui/components/sidebar"
	"github.com/jfbus/templui/components/style"
	"github.com/jfbus/templui/components/toast/container"
)

const (
	// StyleMediumWidth defines a md width (768px).
	StyleMediumWidth style.Style = 1 << 8
	// StyleLargeWidth defines a lg width (1024px) (default).
	StyleLargeWidth style.Style = 1 << 9
	// StyleXLargeWidth defines a lg width (1280px).
	StyleXLargeWidth style.Style = 1 << 10
	// StyleXXLargeWidth defines a lg width (1536px).
	StyleXXLargeWidth style.Style = 1 << 11
	// StyleFullWidthBody removes width from the main content, keeping it only for header/footer.
	StyleFullWidthBody style.Style = 1 << 12
	// StyleWithNavbar is automatically added when a navbar is present.
	StyleWithNavbar style.Style = 1 << 13
	// StyleWithSidebar is automatically added when a sidebar is present.
	StyleWithSidebar style.Style = 1 << 14
)

func init() {
	style.SetDefaults(style.Defaults{
		"body": {
			style.Default: {
				style.Set("relative group block text-sm font-medium"),
			},
		},
		"body/div": {
			style.Default: {
				style.Set("mx-auto"),
			},
			StyleMediumWidth: {
				style.Add("max-w-screen-md"),
			},
			StyleLargeWidth: {
				style.Add("max-w-screen-lg"),
			},
			StyleXLargeWidth: {
				style.Add("max-w-screen-xl"),
			},
			StyleXXLargeWidth: {
				style.Add("max-w-screen-2xl"),
			},
			StyleFullWidthBody: {
				style.RemoveVariants("max-w-screen"),
			},
			StyleWithNavbar: {
				style.Add("pt-16"),
			},
			StyleWithSidebar: {
				style.Add("sm:pl-64"),
			},
		},
		"body/navbar/content": {
			StyleMediumWidth: {
				style.Add("max-w-screen-md"),
			},
			StyleLargeWidth: {
				style.Add("max-w-screen-lg"),
			},
			StyleXLargeWidth: {
				style.Add("max-w-screen-xl"),
			},
			StyleXXLargeWidth: {
				style.Add("max-w-screen-2xl"),
			},
		},
		"body/sidebar": {
			StyleMediumWidth: {
				style.Add("sm:left-[calc(50%-384px)]"),
			},
			StyleLargeWidth: {
				style.Add("lg:left-[calc(50%-512px)]"),
			},
			StyleXLargeWidth: {
				style.Add("xl:left-[calc(50%-640px)]"),
			},
			StyleXXLargeWidth: {
				style.Add("2xl:left-[calc(50%-768px)]"),
			},
			StyleWithNavbar: {
				style.Add("pt-16"),
			},
		},
		"body/footer": {
			style.Default: {
				style.Set("mx-auto"),
			},
			StyleMediumWidth: {
				style.Add("max-w-screen-md"),
			},
			StyleLargeWidth: {
				style.Add("max-w-screen-lg"),
			},
			StyleXLargeWidth: {
				style.Add("max-w-screen-xl"),
			},
			StyleXXLargeWidth: {
				style.Add("max-w-screen-2xl"),
			},
			StyleWithSidebar: {
				style.Add("sm:pl-64"),
			},
		},
		"body/toast/container": {
			StyleWithNavbar: {
				style.Add("pt-16"),
			},
		},
	})
}

type D struct {
	// Style defines the style (max width).
	Style style.Style
	// Navbar defines an optional navbar.
	Navbar *navbar.D
	// Sidebar defines an optional sidebar.
	Sidebar *sidebar.D
	// Toasts defines a custom toast container (a default one will be added otherwise).
	Toasts *container.D
	// Footer defines an optional footer.
	Footer *footer.D
	// State defines a default Alpine.js state.
	State map[string]string
	// CustomStyle defines a custom style.
	// 	style.Custom{
	// 		"body":     style.D{style.Add("...")},
	// 		"body/div": style.D{style.Add("...")},
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

func (def D) style() style.Style {
	st := def.Style
	if def.Style == style.Default {
		st = StyleLargeWidth
	}
	if def.Navbar != nil {
		st |= StyleWithNavbar
	}
	if def.Sidebar != nil {
		st |= StyleWithSidebar
	}
	return st
}

func (def D) navbar() navbar.D {
	d := *def.Navbar
	d.CustomStyle = d.CustomStyle.AddBefore(style.Custom{
		"navbar/content": style.Compute(def.style(), "body/navbar/content", def.CustomStyle),
	})
	return d
}

func (def D) sidebar() sidebar.D {
	d := *def.Sidebar
	d.CustomStyle = d.CustomStyle.AddBefore(style.Custom{
		"sidebar": style.Compute(def.style(), "body/sidebar", def.CustomStyle),
	})
	return d
}

func (def D) footer() footer.D {
	d := *def.Footer
	d.CustomStyle = d.CustomStyle.AddBefore(style.Custom{
		"footer": style.Compute(def.style(), "body/footer", def.CustomStyle),
	})
	return d
}

func (def D) toasts() container.D {
	if def.Toasts == nil {
		return container.D{
			CustomStyle: style.Custom{
				"toast/container": style.Compute(def.style(), "body/toast/container", def.CustomStyle),
			},
		}
	}
	d := *def.Toasts
	d.CustomStyle = d.CustomStyle.AddBefore(style.Custom{
		"toast/container": style.Compute(def.style(), "body/toast/container", def.CustomStyle),
	})
	return d
}

func (def D) class(k string) string {
	return style.CSSClass(def.style(), k, def.CustomStyle)
}
