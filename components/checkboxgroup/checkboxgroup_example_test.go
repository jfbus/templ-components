package checkboxgroup_test

import (
	"context"
	"os"

	"github.com/jfbus/templui/components/checkbox"
	"github.com/jfbus/templui/components/checkboxgroup"
)

func ExampleC() {
	c := checkboxgroup.C(checkboxgroup.D{
		Name: "group",
		Checkboxes: []checkbox.D{
			{Label: "Option 1", Value: "1"},
			{Label: "Option 2", Value: "2"},
		},
	})
	_ = c.Render(context.TODO(), os.Stdout)
}
