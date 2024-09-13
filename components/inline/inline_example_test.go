package inline_test

import (
	"context"
	"os"

	"github.com/a-h/templ"
	"github.com/jfbus/templ-components/components/icon"
	"github.com/jfbus/templ-components/components/inline"
	"github.com/jfbus/templ-components/components/input"
	"github.com/jfbus/templ-components/components/label"
	"github.com/jfbus/templ-components/components/position"
	"github.com/jfbus/templ-components/components/size"
)

func ExampleC() {
	c := inline.C(inline.D{
		Value:    "Previous value",
		IconSize: size.S,
		Focus:    "itemtitle",
		Edit: input.C(input.D{
			Name:         "title",
			Label:        label.D{Label: "Title", Hide: true},
			Value:        "Previous value",
			Icon:         icon.CornerDownLeft,
			IconPosition: position.End,
			Size:         size.S,
			NoValidation: true,
			Loader:       true,
			Attributes: templ.Attributes{
				"hx-post":   "/items/update_title/24",
				"hx-target": "#item_24",
				"x-ref":     "itemtitle",
			},
		}),
	})
	c.Render(context.TODO(), os.Stdout)
}
