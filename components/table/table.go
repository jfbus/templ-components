// Package table implements tables.
package table

import (
	"github.com/jfbus/templ-components/components/style"
	"github.com/jfbus/templ-components/components/table/row"
)

const (
	StyleStripedRows       style.Style = 1 << 8
	StyleNoBorder          style.Style = 1 << 9
	StyleAddHighlightHover style.Style = 1 << 10
)

func init() {
	style.SetDefaults(style.Defaults{
		"table": {
			style.StyleDefault: {
				style.Set("w-full text-sm text-left"),
			},
		},
		"table/header": {
			style.StyleDefault: {
				style.Set("text-xs"),
			},
		},
		"table/row": {
			style.StyleDefault: {
				style.Set("border-b"),
			},
			StyleNoBorder: {
				style.Set(""),
			},
		},
		"table/cell": {
			style.StyleDefault: {
				style.Set("px-6 py-4"),
			},
		},
	})
}

// D is the table definition.
type D struct {
	// Style defines the table style.
	Style style.Style
	// Class defines the CSS classes for the table tag.
	Class style.D
	// HeaderClass defines the CSS classes for the thead tag.
	HeaderClass style.D
	// BodyClass defines the CSS classes for the tr tags within tbody.
	BodyClass style.D
	// FooterClass defines the CSS classes for the tfoot tag.
	FooterClass style.D
	// CellClass defines the CSS classes tor all td tags.
	CellClass style.D
	// Header defines an optional header row (thead).
	//playground:import:github.com/jfbus/templ-components/components/table/row
	//playground:default:&row.D{Cells: []string{"Name", "Description", ""}}
	Header *row.D
	// Rows defines the body rows (tbody).
	//playground:default:[]row.D{{Cells: []string{"Lorem", "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed non risus.", ""}},{Cells: []string{"Ipsum", "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed non risus.", ""}},{Cells: []string{"Dolor", "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed non risus.", ""}}}
	Rows []row.D
	// Footer defines an optional footer fow (tfoot).
	Footer *row.D
}

func (def D) cellClass() string {
	return def.CellClass.CSSClass(def.Style, "table/cell")
}

func (def D) rows() []row.D {
	class := def.cellClass()
	for i := range def.Rows {
		if def.Rows[i].Class == "" {
			def.Rows[i].Class = def.trClass()
			def.Rows[i].DefaultCellClass = class
		}
	}
	return def.Rows
}

func (def D) tableClass() string {
	return def.Class.CSSClass(def.Style, "table")
}

func (def D) header() row.D {
	if def.Header == nil {
		return row.D{}
	}
	fd := *def.Header
	fd.Header = true
	if fd.DefaultCellClass == "" {
		fd.DefaultCellClass = def.cellClass()
	}
	return fd
}

func (def D) headerClass() string {
	if def.Header == nil {
		return ""
	}
	return def.HeaderClass.CSSClass(def.Style, "table/header")
}

func (def D) trClass() string {
	return def.BodyClass.CSSClass(def.Style, "table/row")
}

func (def D) footer() row.D {
	if def.Footer == nil {
		return row.D{}
	}
	fd := *def.Footer
	if fd.DefaultCellClass == "" {
		fd.DefaultCellClass = def.cellClass()
	}
	return fd
}

func (def D) footerClass() string {
	if def.Footer == nil {
		return ""
	}
	return def.FooterClass.CSSClass(def.Style, "table/footer")
}
