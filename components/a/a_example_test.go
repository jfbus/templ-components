package a_test

import (
	"context"
	"os"

	"github.com/jfbus/templ-components/components/a"
	"github.com/jfbus/templ-components/components/style"
)

func ExampleC() {
	c := a.C(a.D{
		Href: "https://www.example.com",
		Text: "Example",
		CustomStyle: style.Custom{
			"a": style.D{style.Add("text-sm")},
		},
	})
	_ = c.Render(context.TODO(), os.Stdout)
	// output: <a href="https://www.example.com" class="hover:underline text-sm">Example</a>
}
