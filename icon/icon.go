//go:generate go run .build/generate.go
package icon

// IconDefinition is the definition of icons.
// Usage: @icon.Icon(icon.IconDefinition{})
type IconDefinition struct {
	// Icon is the SVG string (e.g. icon.Flower).
	Icon string
	// Class is an option CSS class to apply to the SVG tag.
	Class string
}
