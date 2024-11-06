package dropzone_test

import (
	"context"
	"os"

	"github.com/jfbus/templ-components/components/dropzone"
)

func ExampleD() {
	c := dropzone.C(dropzone.D{
		Name: "upload",
	})
	c.Render(context.TODO(), os.Stdout)
}
