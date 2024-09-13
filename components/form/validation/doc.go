// Package validation enables your handler to return validation errors.
// Each validation error is displayed by the corresponding input.Input field:
// - the input switches to style.StyleInvalid
// - the error message is displayed below the field
//
// Handler code:
//
//	_ = validation.Retarget(r.Context(), validation.D{
//		FormID: "profile",
//		Errors: map[string]string{"Email": "Email is mandatory"},
//	}, w, http.StatusUnprocessableEntity)
//
// Template code:
//
//	@form.C(form.D{ID:"profile"}) {
//		@input.C(input.D{Name: "Email", Type: input.TypeEmail})
//	}
package validation
