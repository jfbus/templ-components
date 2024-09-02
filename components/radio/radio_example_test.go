package radio_test

import (
	"github.com/jfbus/templ-components/components/radio"
)

func ExampleD() {
	_ = radio.C(radio.D{
		Name:  "accept",
		Label: "Title",
	})
}
