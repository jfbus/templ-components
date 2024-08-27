package icon_test

import (
	"github.com/jfbus/templ-components/icon"
	"github.com/jfbus/templ-components/size"
)

func ExampleC() {
	_ = icon.C(icon.Flower)
}

func ExampleC_withSize() {
	_ = icon.C(icon.D{Icon: icon.Flower, Size: size.NineXL})
}
