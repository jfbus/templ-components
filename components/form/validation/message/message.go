package message

import (
	"github.com/jfbus/templ-components/components/size"
	"github.com/jfbus/templ-components/components/style"
)

var Defaults = style.Defaults{
	style.StyleDefault: {
		"Class": {
			style.Class("mt-2"),
			style.Color("text-red-600 dark:text-red-500"),
		},
	},
	style.StyleValid: {
		"Class": {
			style.Color("text-green-600 dark:text-green-500"),
		},
	},
	style.StyleInvalid: {
		"Class": {
			style.Color("text-red-600 dark:text-red-500"),
		},
	},
}

type D struct {
	InputName string
	Style     style.Style
	Message   string
	Size      size.Size
	Class     style.D
}

func (def D) class() string {
	class := def.Class.CSSClass(Defaults, def.Style, "Class")
	switch def.Size {
	case size.S:
		class += " text-xs"
	case size.L:
		class += " text-base"
	default:
		class += " text-sm"
	}
	return class
}
