package input_test

import (
	"context"
	"os"

	"github.com/jfbus/templui/components/icon"
	"github.com/jfbus/templui/components/input"
	"github.com/jfbus/templui/components/size"
)

func ExampleC() {
	c := input.C(input.D{
		Name:        "title",
		Type:        input.TypeText,
		Label:       "Title",
		Value:       "Previous title",
		Placeholder: "Enter a title",
		Size:        size.S,
		Icon:        icon.Bookmark,
	})
	_ = c.Render(context.TODO(), os.Stdout)
}
