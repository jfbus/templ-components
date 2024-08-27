package input_test

import (
	"github.com/jfbus/templ-components/icon"
	"github.com/jfbus/templ-components/input"
	"github.com/jfbus/templ-components/size"
)

func ExampleD() {
	_ = input.C(input.D{
		Name:        "title",
		Type:        input.TypeText,
		Label:       "Title",
		Value:       "Previous title",
		Placeholder: "Enter a title",
		Size:        size.S,
		Icon:        icon.Bookmark,
	})
}
