package validation

import (
	"context"
	"errors"
	"net/http"

	"github.com/a-h/templ"
)

// D is the validation retarget response.
type D struct {
	// FormID is the ID of the form to target.
	FormID string
	// Errors is the list of errors to display.
	Errors map[string]string
}

func (def D) setErrors() string {
	s, err := templ.JSONString(def.Errors)
	if err != nil {
		return ""
	}
	return "errors=" + s
}

// Retarget retargets the response to display validation errors in a form.
func Retarget(ctx context.Context, def D, w http.ResponseWriter, statusCode int) error {
	if def.FormID == "" {
		return errors.New("form id is required")
	}
	w.Header().Set("Content-Type", "text/html, charset=UTF-8")
	w.Header().Set("HX-Retarget", "#"+def.FormID)
	w.Header().Set("HX-Reswap", "beforeend")
	w.WriteHeader(statusCode)
	return C(def).Render(ctx, w)
}
