package cell

import "github.com/jfbus/templui/components/style"

type D struct {
	Header  bool
	Content any
	ColSpan string
	// CustomStyle defines a custom style.
	// 	style.Custom{
	// 		"table/cell":     style.D{style.Add("...")},
	//	}
	CustomStyle style.Custom
}

func (def D) class() string {
	return style.CSSClass(style.Default, "table/cell", def.CustomStyle)
}
