package toast_test

import (
	"context"
	"os"

	"github.com/jfbus/templ-components/components/toast"
)

func ExampleC() {
	c := toast.C(toast.D{
		Style:   toast.StyleError,
		Content: "An error occurred !",
	})
	c.Render(context.TODO(), os.Stdout)
}
