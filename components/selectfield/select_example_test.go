package selectfield_test

import (
	selectfield2 "github.com/jfbus/templ-components/components/selectfield"
	"github.com/jfbus/templ-components/components/selectfield/option"
)

func ExampleC() {
	_ = selectfield2.C(selectfield2.D{
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
