package table_test

import (
	"github.com/a-h/templ"
	"github.com/jfbus/templ-components/components/button"
	"github.com/jfbus/templ-components/components/table"
	"github.com/jfbus/templ-components/components/table/cell"
	"github.com/jfbus/templ-components/components/table/row"
)

func ExampleC() {
	_ = table.C(table.D{
		Style: table.StyleStripedRows,
		Header: &row.D{
			Cells: []string{"Email", "Name", "Status", ""},
		},
		Rows: []row.D{{
			Cells: []any{
				"John Doe",
				"john.doe@example.com",
				"active",
				cell.D{
					Class: "text-center",
					Content: button.C(button.D{
						Label: "disable",
						Attributes: templ.Attributes{
							"hx-delete": "/users/1",
						},
					}),
				},
			},
		}},
	})

}
