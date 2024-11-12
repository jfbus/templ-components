package row

import (
	"github.com/a-h/templ"
	"github.com/jfbus/templui/components/style"
	"github.com/jfbus/templui/components/table/cell"
)

// D defines a table row (tr).
type D struct {
	// Header defines if it is a Header row or not.
	Header bool
	// Cells is the slice of cell contents
	// (either []string, []cell.D or []templ.Component or []any mixing the previous types).
	Cells any
	// CustomStyle defines a custom style.
	// 	style.Custom{
	// 		"table/row": style.D{style.Add("...")},
	// 		"table/cell":     style.D{style.Add("...")},
	//	}
	CustomStyle style.Custom
}

func (def D) class() string {
	return style.CSSClass(style.Default, "table/row", def.CustomStyle)
}

func (def D) cells() []cell.D {
	switch cells := def.Cells.(type) {
	case []any:
		return def.cellsAny(cells)
	case []cell.D:
		return def.cellsCell(cells)
	case []string:
		return def.cellsString(cells)
	case []templ.Component:
		return def.cellsComponent(cells)
	}
	return nil
}

func (def D) cellsCell(cells []cell.D) []cell.D {
	for i := range cells {
		cells[i].Header = def.Header
		cells[i].CustomStyle = cells[i].CustomStyle.AddBefore(def.CustomStyle)
	}
	return cells
}

func (def D) cellsString(cells []string) []cell.D {
	res := make([]cell.D, len(cells))
	for i, c := range cells {
		res[i] = cell.D{
			Header:      def.Header,
			Content:     c,
			CustomStyle: def.CustomStyle,
		}
	}
	return res
}

func (def D) cellsComponent(cells []templ.Component) []cell.D {
	res := make([]cell.D, len(cells))
	for i, c := range cells {
		res[i] = cell.D{
			Header:      def.Header,
			Content:     c,
			CustomStyle: def.CustomStyle,
		}
	}
	return res
}

func (def D) cellsAny(cells []any) []cell.D {
	res := make([]cell.D, len(cells))
	for i, c := range cells {
		switch v := c.(type) {
		case cell.D:
			v.Header = def.Header
			v.CustomStyle = v.CustomStyle.AddBefore(def.CustomStyle)
			res[i] = v
		default:
			res[i] = cell.D{
				Header:      def.Header,
				Content:     c,
				CustomStyle: def.CustomStyle,
			}
		}
	}
	return res
}
