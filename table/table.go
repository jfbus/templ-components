package table

import (
	"github.com/jfbus/templ-components/table/row"
)

type Style int

const (
	StyleDefault           Style = 0
	StyleStripedRows             = 1
	StyleNoBorder                = 2
	StyleAddHighlightHover       = 16
)

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
		BodyClass: "hover:bg-gray-50 dark:hover:bg-gray-600",
	},
}

type D struct {
	Style       Style
	Color       string
	TableClass  string
	HeaderClass string
	BodyClass   string
	FooterClass string
	CellClass   string
	Header      *row.D
	Rows        []row.D
	Footer      *row.D
}

func (def D) cellClass() string {
	class := Defaults[StyleDefault].CellClass
	if def.CellClass != "" {
		class = def.CellClass
	}
	return class
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
	class := Defaults[StyleDefault].TableClass
	if def.Style&StyleStripedRows != 0 && Defaults[StyleStripedRows].TableClass != "" {
		class = Defaults[StyleStripedRows].TableClass
	}
	if def.Style&StyleNoBorder != 0 && Defaults[StyleNoBorder].TableClass != "" {
		class = Defaults[StyleNoBorder].TableClass
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
	class := Defaults[StyleDefault].HeaderClass
	if def.Style&StyleStripedRows != 0 && Defaults[StyleStripedRows].HeaderClass != "" {
		class = Defaults[StyleStripedRows].HeaderClass
	}
	if def.Style&StyleNoBorder != 0 && Defaults[StyleStripedRows].HeaderClass != "" {
		class = Defaults[StyleNoBorder].HeaderClass
	}
	if def.Style&StyleAddHighlightHover != 0 {
		class += " " + Defaults[StyleAddHighlightHover].HeaderClass
	}
	if def.HeaderClass != "" {
		class += " " + def.HeaderClass
	}
	return class
}

func (def D) trClass() string {
	class := Defaults[StyleDefault].BodyClass
	if def.Style&StyleStripedRows != 0 {
		class = Defaults[StyleStripedRows].BodyClass
	}
	if def.Style&StyleNoBorder != 0 {
		class = Defaults[StyleNoBorder].BodyClass
	}
	if def.Style&StyleAddHighlightHover != 0 {
		class += " " + Defaults[StyleAddHighlightHover].BodyClass
	}
	if def.BodyClass != "" {
		class += " " + def.BodyClass
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
	class := Defaults[StyleDefault].FooterClass
	if def.Style&StyleStripedRows != 0 {
		class = Defaults[StyleStripedRows].FooterClass
	}
	if def.Style&StyleNoBorder != 0 {
		class = Defaults[StyleNoBorder].FooterClass
	}
	if def.Style&StyleAddHighlightHover != 0 {
		class += " " + Defaults[StyleAddHighlightHover].FooterClass
	}
	if def.FooterClass != "" {
		class += " " + def.FooterClass
	}
	return class
}
