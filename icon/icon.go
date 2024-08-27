// Package icon implements SVG icons.
//
//go:generate go run .build/generate.go
package icon

import "github.com/jfbus/templ-components/size"

// D is the definition of icons.
// Usage: @icon.C(icon.D{Icon:icon.Flower}) or @icon.C(icon.Flower) or &icon.C("flower")
type D struct {
	// Icon is the SVG string (e.g. icon.Flower).
	Icon string
	// Class is an option CSS class to apply to the SVG tag.
	Class string
	// Size is the icon size.
	Size size.Size
}
