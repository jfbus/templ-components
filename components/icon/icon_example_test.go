package icon_test

import (
	"context"
	"os"

	"github.com/jfbus/templ-components/components/icon"
)

func ExampleC() {
	c := icon.C(icon.Flower)
	c.Render(context.TODO(), os.Stdout)
}
