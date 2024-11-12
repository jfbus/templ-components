package radio_test

import (
	"context"
	"os"

	"github.com/jfbus/templui/components/radio"
)

func ExampleD() {
	c := radio.C(radio.D{
		Name:  "accept",
		Label: "Title",
	})
	_ = c.Render(context.TODO(), os.Stdout)
}
