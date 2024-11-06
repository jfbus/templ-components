package dropdown_test

import (
	"context"
	"os"

	"github.com/a-h/templ"
	"github.com/jfbus/templ-components/components/a"
	"github.com/jfbus/templ-components/components/button"
	"github.com/jfbus/templ-components/components/dropdown"
)

func usernameComponent() templ.Component {
	return templ.Raw("User Name")
}

func ExampleC() {
	c := dropdown.C(dropdown.D{
		Button: button.D{Label: "Menu"},
		Header: usernameComponent(),
		Links: [][]a.D{{
			{Href: "/profile", Text: "Profile"},
		}, {
			{Href: "/logout", Text: "Logout"},
		}},
	})

	_ = c.Render(context.TODO(), os.Stdout)
}
