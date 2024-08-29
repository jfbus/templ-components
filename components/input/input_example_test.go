package input_test

import (
	"github.com/jfbus/templ-components/components/icon"
	"github.com/jfbus/templ-components/components/input"
	"github.com/jfbus/templ-components/components/size"
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
