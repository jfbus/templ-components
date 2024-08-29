package selectfield_test

import (
	"github.com/jfbus/templ-components/selectfield"
	"github.com/jfbus/templ-components/selectfield/option"
)

func ExampleC() {
	_ = selectfield.C(selectfield.D{
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
}
