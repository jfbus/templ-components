package textarea_test

import (
	"context"
	"os"

	"github.com/jfbus/templui/components/icon"
	"github.com/jfbus/templui/components/size"
	"github.com/jfbus/templui/components/textarea"
)

func ExampleC() {
	c := textarea.C(textarea.D{
		Name:  "comment",
		Label: "Commentaire",
		Value: "Previous comment",
		Rows:  3,
		Size:  size.S,
		Icon:  icon.Text,
	})
	_ = c.Render(context.TODO(), os.Stdout)
}
