package row

import (
	"github.com/a-h/templ"
	"github.com/jfbus/templ-components/table/cell"
)

type D struct {
	Header           bool
	Class            string
	DefaultCellClass string
	Cells            any
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
