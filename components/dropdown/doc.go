// Package dropdown implements dropdowns.
// When clicked, the button will display a dropdown menu containing:
// - a header
// - link sections, each separated by an horizontal separator.
//
// Usage:
//
//	@dropdown.C(dropdown.D{
//		Button: button.D{Label: "Menu"},
//		Header: usernameComponent,
//		Links: [][]a.D{{
//			{Href: "/profile", Text: "Profile"},
//		}, {
//			{Href: "/logout", Text: "Logout"},
//		}},
//	})
package dropdown
