package accordion_test

import (
	"context"
	"os"

	"github.com/a-h/templ"
	"github.com/jfbus/templ-components/components/accordion"
	"github.com/jfbus/templ-components/components/accordion/element"
)

var yourcomponent templ.Component

func ExampleC() {
	c := accordion.C(accordion.D{
		ID: "accordion",
		Children: []element.D{{
			Open:    true,
			Title:   "First",
			Content: yourcomponent,
		}, {
			Title:   "Second",
			Content: yourcomponent,
		}},
	})
	c.Render(context.TODO(), os.Stdout)
}
