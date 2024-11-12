package selectfield_test

import (
	"context"
	"os"

	"github.com/jfbus/templui/components/selectfield"
	"github.com/jfbus/templui/components/selectfield/option"
)

func ExampleC() {
	c := selectfield.C(selectfield.D{
		Name:  "country",
		Label: "Country",
		Options: []option.D{{
			Label: "Select a country",
		}, {
			Value: "FR",
			Label: "France",
		}, {
			Value: "DE",
			Label: "Germany",
		}, {
			Value: "UK",
			Label: "United Kingdom",
		}},
		Selected: "DE",
	})
	c.Render(context.TODO(), os.Stdout)
}
