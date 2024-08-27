package textarea_test

import (
	"pkgs/templ-components/icon"
	"pkgs/templ-components/size"
	"pkgs/templ-components/textarea"
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
