package row

import (
	"github.com/a-h/templ"
	"github.com/jfbus/templ-components/table/cell"
)

// D defines a table row (tr).
type D struct {
	// Header defines if it is a Header row or not.
	Header bool
	// Class is the CSS class, inherited from the table if not set.
	Class string
	// DefaultCellClass is the class for all cells of the row.
	DefaultCellClass string
	// Cells is the slice of cell contents
	// (either []string, []cell.D or []templ.Component or []any mixing the previous types).
	Cells any
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
		if cells[i].Class == "" {
			cells[i].Class = def.DefaultCellClass
		}
	}
	return cells
}

func (def D) cellsString(cells []string) []cell.D {
	res := make([]cell.D, len(cells))
	for i, c := range cells {
		res[i] = cell.D{
			Header:  def.Header,
			Content: c,
			Class:   def.DefaultCellClass,
		}
	}
	return res
}

func (def D) cellsComponent(cells []templ.Component) []cell.D {
	res := make([]cell.D, len(cells))
	for i, c := range cells {
		res[i] = cell.D{
			Header:  def.Header,
			Content: c,
			Class:   def.DefaultCellClass,
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
			if v.Class == "" {
				v.Class = def.DefaultCellClass
			}
			res[i] = v
		default:
			res[i] = cell.D{
				Header:  def.Header,
				Content: c,
				Class:   def.DefaultCellClass,
			}
		}
	}
	return res
}
