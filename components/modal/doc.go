// Package modal adds modals.
// Each modal embeds a form.C component, yous only have to add inputs and use Cancel/Confirm buttons.
//
// Usage:
//
//	// with a form, embedding the content in the modal
//	@modal.C(modal.D{
//		ID: "profile",
//		Title: "Profile",
//		Content: [...], // your form
//		Form: &form.D{},
//		Close: &button.Button(button.D{Style: button.StyleOutline, Label: "Cancel"}),
//		Confirm: &button.Button(button.D{Label: "Save"}),
//	})
//
//	// without form,
//	@modal.C(modal.D{
//		ID: "modal",
//		Title: "Modal example",
//		Content: [...], // your form
//		Close: &button.Button(button.D{Label: "OK"}),
//	})
//
//	// content as children
//	@modal.C(modal.D{[...]}) {
//		// your form
//	}
//
// The cancel button closes the modal, the confirm does not. To close a modal, just add
// modal.Close to your submit response:
//
//	components.Append(yourResponseComponent, modal.Close("profile"))
package modal
