package input_test

import (
	"pkgs/templ-components/icon"
	"pkgs/templ-components/input"
	"pkgs/templ-components/size"
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
