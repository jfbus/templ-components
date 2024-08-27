package button_test

import (
	"github.com/jfbus/templ-components/button"
	"github.com/jfbus/templ-components/icon"
	"github.com/jfbus/templ-components/size"
)

func ExampleC() {
	_ = button.C(button.D{
		ID:     "item_submit",
		Type:   button.Submit,
		Label:  "Save",
		Size:   size.L,
		Icon:   icon.Save,
		Loader: true,
		Attributes: templ.Attributes{
			"hx-post":   "/items/save",
			"hx-target": "#item_list",
		},
	})
}
