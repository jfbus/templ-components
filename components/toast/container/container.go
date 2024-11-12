package container

import "github.com/jfbus/templ-components/components/style"

func init() {
	style.SetDefaults(style.Defaults{
		"toast/container": {
			style.Default: {
				style.Set("z-50 fixed top-0 right-0 m-2 space-y-2"),
			},
			StyleTopLeft: {
				style.Set("z-50 fixed top-0 left-0 m-2 space-y-2"),
			},
			StyleTopRight: {
				style.Set("z-50 fixed top-0 right-0 m-2 space-y-2"),
			},
			StyleBottomLeft: {
				style.Set("z-50 fixed bottom-0 left-0 m-2 space-y-2"),
			},
			StyleBottomRight: {
				style.Set("z-50 fixed bottom-0 right-0 m-2 space-y-2"),
			},
		},
	})
}

const (
	StyleTopLeft     style.Style = 1 << 8
	StyleTopRight    style.Style = 1 << 9
	StyleBottomLeft  style.Style = 1 << 10
	StyleBottomRight style.Style = 1 << 11
)

var DefaultID = "toasts"

type D struct {
	ID    string
	Style style.Style
	// CustomStyle defines a custom style.
	// 	style.Custom{
	// 		"toast/container": style.D{style.Add("...")},
	//	}
	CustomStyle style.Custom
}

func (def D) id() string {
	if def.ID != "" {
		return def.ID
	}
	return DefaultID
}

func (def D) class() string {
	return style.CSSClass(def.Style, "toast/container", def.CustomStyle)
}
