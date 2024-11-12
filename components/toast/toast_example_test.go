package toast_test

import (
	"context"
	"os"

	"github.com/jfbus/templui/components/toast"
)

func ExampleC() {
	c := toast.C(toast.D{
		ContainerID: "toasts",
		Style:       toast.StyleError,
		Content:     "An error occurred !",
	})
	_ = c.Render(context.TODO(), os.Stdout)
}
