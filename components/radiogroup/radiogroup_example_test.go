package radiogroup_test

import (
	"context"
	"os"

	"github.com/jfbus/templ-components/components/radio"
	"github.com/jfbus/templ-components/components/radiogroup"
)

func ExampleC() {
	c := radiogroup.C(radiogroup.D{
		Name: "group",
		Radios: []radio.D{
			{Label: "Option 1", Value: "1"},
			{Label: "Option 2", Value: "2"},
		},
	})
	c.Render(context.TODO(), os.Stdout)
}
