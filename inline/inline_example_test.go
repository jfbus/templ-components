package inline_test

import (
	"github.com/a-h/templ"
	"github.com/jfbus/templ-components/inline"
	"github.com/jfbus/templ-components/input"
	"github.com/jfbus/templ-components/size"
)

func ExampleC() {
	_ = inline.C(inline.D{
		Value:    "Previous value",
		IconSize: size.S,
		Edit: input.C(input.D{
			Name:      "title",
			Label:     "Title",
			HideLabel: true,
			Value:     "Previous value",
			Size:      size.S,
			Loader:    true,
			Attributes: templ.Attributes{
				"hx-post":   "/items/update_title/24",
				"hx-target": "#item_24",
			},
		}),
	})
}
