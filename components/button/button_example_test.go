package button_test

import (
	"context"
	"os"

	"github.com/a-h/templ"
	"github.com/jfbus/templ-components/components/button"
	"github.com/jfbus/templ-components/components/icon"
	"github.com/jfbus/templ-components/components/size"
)

func ExampleC() {
	c := button.C(button.D{
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
	c.Render(context.TODO(), os.Stdout)
}
