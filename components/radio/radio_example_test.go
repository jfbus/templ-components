package radio_test

import (
	"context"
	"os"

	"github.com/jfbus/templ-components/components/radio"
)

func ExampleD() {
	c := radio.C(radio.D{
		Name:  "accept",
		Label: "Title",
	})
	c.Render(context.TODO(), os.Stdout)
}
