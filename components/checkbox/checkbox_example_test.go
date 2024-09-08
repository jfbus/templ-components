package checkbox_test

import (
	"context"
	"os"

	"github.com/jfbus/templ-components/components/checkbox"
)

func ExampleD() {
	c := checkbox.C(checkbox.D{
		Name:  "accept",
		Label: "Title",
	})
	c.Render(context.TODO(), os.Stdout)
}
