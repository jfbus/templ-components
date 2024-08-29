// Package table implements tables.
package table

import (
	"github.com/jfbus/templ-components/helper"
	"github.com/jfbus/templ-components/table/row"
)

type Style int

const (
	StyleDefault           Style = 0
	StyleStripedRows             = 1
	StyleNoBorder                = 2
	StyleAddHighlightHover       = 16
)

// Defaults are default CSS classes for each style. The are overriden by D.TableClass/D.HeaderClass/D.BodyClass/D.CellClass.
var Defaults = map[Style]D{
	StyleDefault: {
		TableClass:  "w-full text-sm text-left",
		HeaderClass: "text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400",
		BodyClass:   "bg-white border-b dark:bg-gray-800 dark:border-gray-700",
		CellClass:   "px-6 py-4",
	},
	StyleStripedRows: {
		BodyClass: "odd:bg-white odd:dark:bg-gray-900 even:bg-gray-50 even:dark:bg-gray-800 border-b dark:border-gray-700",
	},
	StyleNoBorder: {
		HeaderClass: "text-gray-900 uppercase dark:text-gray-400",
		BodyClass:   "text-gray-900 whitespace-nowrap dark:text-white",
	},
	StyleAddHighlightHover: {
		// This style is added to the previous class or to D.TableClass
		BodyClass: "hover:bg-gray-50 dark:hover:bg-gray-600",
	},
}

// D is the table definition.
type D struct {
	// Style defines the table style.
	Style Style
	// Color defines the table color CSS classes. If not set, Defaults is used.
	Color string
	// TableClass defines the CSS classes for the table tag.
	TableClass string
	// HeaderClass defines the CSS classes for the thead tag.
	HeaderClass string
	// BodyClass defines the CSS classes for the tr tags within tbody.
	BodyClass string
	// FooterClass defines the CSS classes for the tfoot tag.
	FooterClass string
	// CellClass defines the CSS classes tor all td tags.
	CellClass string
	// Header defines an optional header row (thead).
	Header *row.D
	// Rows defines the body rows (tbody).
	Rows []row.D
	// Footer defines an optional footer fow (tfoot).
	Footer *row.D
}

func (def D) cellClass() string {
	return helper.IfEmpty(def.CellClass, Defaults[def.Style].CellClass, Defaults[StyleDefault].CellClass)
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
	var class string
	switch {
	case def.Style&StyleStripedRows != 0:
		class = helper.IfEmpty(def.TableClass, Defaults[StyleStripedRows].TableClass, Defaults[StyleDefault].TableClass)
	case def.Style&StyleNoBorder != 0:
		class = helper.IfEmpty(def.TableClass, Defaults[StyleNoBorder].TableClass, Defaults[StyleDefault].TableClass)
	default:
		class = helper.IfEmpty(def.TableClass, Defaults[StyleDefault].TableClass)
	}
	if def.Style&StyleAddHighlightHover != 0 {
		class += " " + Defaults[StyleAddHighlightHover].TableClass
	}
	if def.TableClass != "" {
		class += " " + def.TableClass
	}
	return class
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
	var class string
	switch {
	case def.Style&StyleStripedRows != 0:
		class = helper.IfEmpty(def.HeaderClass, Defaults[StyleStripedRows].HeaderClass, Defaults[StyleDefault].HeaderClass)
	case def.Style&StyleNoBorder != 0:
		class = helper.IfEmpty(def.HeaderClass, Defaults[StyleNoBorder].HeaderClass, Defaults[StyleDefault].HeaderClass)
	default:
		class = helper.IfEmpty(def.HeaderClass, Defaults[StyleDefault].HeaderClass)
	}
	if def.Style&StyleAddHighlightHover != 0 {
		class += " " + Defaults[StyleAddHighlightHover].HeaderClass
	}
	return class
}

func (def D) trClass() string {
	var class string
	switch {
	case def.Style&StyleStripedRows != 0:
		class = helper.IfEmpty(def.BodyClass, Defaults[StyleStripedRows].BodyClass, Defaults[StyleDefault].BodyClass)
	case def.Style&StyleNoBorder != 0:
		class = helper.IfEmpty(def.BodyClass, Defaults[StyleNoBorder].BodyClass, Defaults[StyleDefault].BodyClass)
	default:
		class = helper.IfEmpty(def.BodyClass, Defaults[StyleDefault].BodyClass)
	}
	if def.Style&StyleAddHighlightHover != 0 {
		class += " " + Defaults[StyleAddHighlightHover].BodyClass
	}
	return class
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
	var class string
	switch {
	case def.Style&StyleStripedRows != 0:
		class = helper.IfEmpty(def.FooterClass, Defaults[StyleStripedRows].FooterClass, Defaults[StyleDefault].FooterClass)
	case def.Style&StyleNoBorder != 0:
		class = helper.IfEmpty(def.FooterClass, Defaults[StyleNoBorder].FooterClass, Defaults[StyleDefault].FooterClass)
	default:
		class = helper.IfEmpty(def.FooterClass, Defaults[StyleDefault].FooterClass)
	}
	if def.Style&StyleAddHighlightHover != 0 {
		class += " " + Defaults[StyleAddHighlightHover].FooterClass
	}
	return class
}
