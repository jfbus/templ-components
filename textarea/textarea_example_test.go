package textarea_test

import (
	"github.com/jfbus/templ-components/icon"
	"github.com/jfbus/templ-components/size"
	"github.com/jfbus/templ-components/textarea"
)

func ExampleC() {
	_ = textarea.C(textarea.D{
		Name:  "comment",
		Label: "Commentaire",
		Value: "Previous comment",
		Rows:  3,
		Size:  size.S,
		Icon:  icon.Text,
	})
}
