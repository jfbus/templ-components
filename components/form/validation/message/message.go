package message

import (
	"github.com/jfbus/templ-components/components/size"
	"github.com/jfbus/templ-components/components/style"
)

func init() {
	style.SetDefaults(style.Defaults{
		"form/validation/message": {
			style.StyleDefault: {
				style.Set("mt-2"),
			},
		},
	})
}

type D struct {
	InputName string
	Style     style.Style
	Message   string
	Size      size.Size
	Class     style.D
}

func (def D) class() string {
	class := def.Class.CSSClass(def.Style, "form/validation/message")
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
