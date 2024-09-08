// Package table implements tables.
package table

import (
	"github.com/jfbus/templ-components/components/style"
	"github.com/jfbus/templ-components/components/table/row"
)

const (
	StyleStripedRows       style.Style = 1
	StyleNoBorder          style.Style = 2
	StyleAddHighlightHover style.Style = 32
)

// Defaults are default CSS classes for each style. The are overriden by D.TableClass/D.HeaderClass/D.BodyClass/D.CellClass.
var Defaults = style.Defaults{
	style.StyleDefault: {
		"Class":       {Class: "w-full text-sm text-left"},
		"HeaderClass": {Class: "text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400"},
		"BodyClass":   {Class: "bg-white border-b dark:bg-gray-800 dark:border-gray-700"},
		"CellClass":   {Class: "px-6 py-4"},
	},
	StyleStripedRows: {
		"BodyClass": {Class: "odd:bg-white odd:dark:bg-gray-900 even:bg-gray-50 even:dark:bg-gray-800 border-b dark:border-gray-700"},
	},
	StyleNoBorder: {
		"HeaderClass": {Class: "text-gray-900 uppercase dark:text-gray-400"},
		"BodyClass":   {Class: "text-gray-900 whitespace-nowrap dark:text-white"},
	},
	StyleAddHighlightHover: {
		// This style is added to the previous class or to D.Class
		"BodyClass": {Class: "hover:bg-gray-50 dark:hover:bg-gray-600"},
	},
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
	return def.CellClass.CSSClass(style.Default(Defaults, def.Style, "CellClass"))
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
	return def.Class.CSSClass(style.Default(Defaults, def.Style, "Class"))
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
	return def.HeaderClass.CSSClass(style.Default(Defaults, def.Style, "HeaderClass"))
}

func (def D) trClass() string {
	return def.BodyClass.CSSClass(style.Default(Defaults, def.Style, "BodyClass"))
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
	return def.FooterClass.CSSClass(style.Default(Defaults, def.Style, "FooterClass"))
}
