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
			style.Default: {
				style.Set("w-full text-sm text-left"),
			},
		},
		"table/header": {
			style.Default: {
				style.Set("text-xs"),
			},
		},
		"table/row": {
			style.Default: {
				style.Set("border-b"),
			},
			StyleNoBorder: {
				style.Set(""),
			},
		},
		"table/cell": {
			style.Default: {
				style.Set("px-6 py-4"),
			},
		},
	})
}

// D is the table definition.
type D struct {
	// Style defines the table style.
	Style style.Style
	// Header defines an optional header row (thead).
	//playground:import:github.com/jfbus/templ-components/components/table/row
	//playground:default:&row.D{Cells: []string{"Name", "Description", ""}}
	Header *row.D
	// Rows defines the body rows (tbody).
	//playground:default:[]row.D{{Cells: []string{"Lorem", "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed non risus.", ""}},{Cells: []string{"Ipsum", "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed non risus.", ""}},{Cells: []string{"Dolor", "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed non risus.", ""}}}
	Rows []row.D
	// Footer defines an optional footer fow (tfoot).
	Footer *row.D
	// CustomStyle defines a custom style.
	// 	style.Custom{
	// 		"table":             style.D{style.Add("...")},
	// 		"table/header":       style.D{style.Add("...")},
	// 		"table/row": style.D{style.Add("...")},
	// 		"table/cell":     style.D{style.Add("...")},
	// 		"table/footer":       style.D{style.Add("...")},
	//	}
	CustomStyle style.Custom
}

func (def D) class(k string) string {
	return style.CSSClass(def.Style, k, def.CustomStyle)
}

func (def D) rows() []row.D {
	cc := style.Custom{
		"table/row":  style.Compute(def.Style, "table/row", def.CustomStyle),
		"table/cell": style.Compute(def.Style, "table/cell", def.CustomStyle),
	}
	for i := range def.Rows {
		def.Rows[i].CustomStyle = def.Rows[i].CustomStyle.AddBefore(cc)
	}
	return def.Rows
}

func (def D) header() row.D {
	if def.Header == nil {
		return row.D{}
	}
	cc := style.Custom{
		"table/header": style.Compute(def.Style, "table/header", def.CustomStyle),
		"table/cell":   style.Compute(def.Style, "table/cell", def.CustomStyle),
	}
	fd := *def.Header
	fd.Header = true
	fd.CustomStyle = fd.CustomStyle.AddBefore(cc)
	return fd
}

func (def D) headerClass() string {
	if def.Header == nil {
		return ""
	}
	return style.CSSClass(def.Style, "table/header", def.CustomStyle)
}

func (def D) footer() row.D {
	if def.Footer == nil {
		return row.D{}
	}
	cc := style.Custom{
		"table/footer": style.Compute(def.Style, "table/footer", def.CustomStyle),
		"table/cell":   style.Compute(def.Style, "table/cell", def.CustomStyle),
	}
	fd := *def.Footer
	fd.CustomStyle = fd.CustomStyle.AddBefore(cc)
	return fd
}

func (def D) footerClass() string {
	if def.Footer == nil {
		return ""
	}
	return style.CSSClass(def.Style, "table/footer", def.CustomStyle)
}
