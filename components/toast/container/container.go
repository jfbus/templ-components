package container

import "github.com/jfbus/templ-components/components/style"

var Defaults = style.Defaults{
	style.StyleDefault: {
		"Class": {
			style.Class("z-10 fixed top-0 left-0 m-2 space-y-2"),
		},
	},
	StyleTopLeft: {
		"Class": {
			style.Class("z-10 fixed top-0 left-0 m-2 space-y-2"),
		},
	},
	StyleTopRight: {
		"Class": {
			style.Class("z-10 fixed top-0 right-0 m-2 space-y-2"),
		},
	},
	StyleBottomLeft: {
		"Class": {
			style.Class("z-10 fixed bottom-0 left-0 m-2 space-y-2"),
		},
	},
	StyleBottomRight: {
		"Class": {
			style.Class("z-10 fixed bottom-0 right-0 m-2 space-y-2"),
		},
	},
}

const (
	StyleTopLeft     style.Style = 1 << 8
	StyleTopRight    style.Style = 1 << 9
	StyleBottomLeft  style.Style = 1 << 10
	StyleBottomRight style.Style = 1 << 11
)

type D struct {
	ID    string
	Style style.Style
	Class style.D
}

func (def D) id() string {
	return def.ID
}

func (def D) class() string {
	return def.Class.CSSClass(Defaults, def.Style, "Class")
}
