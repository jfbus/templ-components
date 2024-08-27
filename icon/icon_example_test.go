package icon_test

import (
	"pkgs/templ-components/icon"
	"pkgs/templ-components/size"
)

func ExampleC() {
	_ = icon.C(icon.Flower)
}

func ExampleC_withSize() {
	_ = icon.C(icon.D{Icon: icon.Flower, Size: size.NineXL})
}
