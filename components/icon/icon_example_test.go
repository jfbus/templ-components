package icon_test

import (
	"context"
	"os"

	"github.com/jfbus/templ-components/components/icon"
	"github.com/jfbus/templ-components/components/size"
)

func ExampleC() {
	c := icon.C(icon.Flower)
	c.Render(context.TODO(), os.Stdout)
}

func ExampleC_withSize() {
	c := icon.C(icon.D{Icon: icon.Flower, Size: size.NineXL})
	c.Render(context.TODO(), os.Stdout)
}
