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
	// CustomStyle defines a custom style.
	// 	style.Custom{
	// 		"form/validation/message": style.D{style.Add("...")},
	//	}
	CustomStyle style.Custom
}

func (def D) class() string {
	class := style.CSSClass(def.Style, "form/validation/message", def.CustomStyle)
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
