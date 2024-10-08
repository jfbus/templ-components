package toast

import (
	"context"
	"net/http"
	"time"

	"github.com/jfbus/templ-components/components/icon"
	"github.com/jfbus/templ-components/components/style"
	"github.com/jfbus/templ-components/components/toast/container"
	"github.com/rs/xid"
)

const (
	StyleOK      style.Style = 1 << 8
	StyleWarning style.Style = 1 << 9
	StyleError   style.Style = 1 << 10
)

type Close int

const (
	CloseButton Close = iota
	CloseAuto
)

func init() {
	style.SetDefaults(style.Defaults{
		"toast": {
			style.StyleDefault: {
				style.Set("flex items-center w-full max-w-xs p-4 rounded-lg shadow"),
			},
		},
		"toast/icon": {
			style.StyleDefault: {
				style.Set("inline-flex items-center justify-center flex-shrink-0 w-8 h-8 rounded-lg"),
			},
		},
	})
}

type D struct {
	_id string
	// ContainerID is the ID of the container.C where the toast will be added.
	//playground:default:"toasts"
	ContainerID string
	// Style defines the toast style.
	Style style.Style
	// Icon defines an optional icon. StyleOK/StyleWarning/StyleError include a default icon
	// that can be overriden.
	Icon string
	// Content defines the text content of the toast.
	// You may also define a custom child content :
	//  @toast.C(toast.D{}){
	//    // your custom content
	//  }
	Content string
	// ContainerClass stores overrides to the container CSS classes.
	ContainerClass style.D
	// IconClass stores overrides to the icon CSS classes.
	IconClass style.D
	// Close defines if a close button will be added or if the toast will close after AutoCloseDelay.
	Close Close
	// AutoCloseDelay defines the close delay (default 5s).
	AutoCloseDelay time.Duration
}

func (def *D) id() string {
	if def._id == "" {
		def._id = xid.New().String()
	}
	return def._id
}

func (def D) containerID() string {
	if def.ContainerID != "" {
		return def.ContainerID
	}
	return container.DefaultID
}

func (def D) icon() string {
	switch {
	case def.Icon != "":
		return def.Icon
	case def.Style&StyleOK != 0:
		return icon.CircleCheck
	case def.Style&StyleWarning != 0:
		return icon.CircleAlert
	case def.Style&StyleError != 0:
		return icon.CircleX
	default:
		return ""
	}
}

func (def D) containerClass() string {
	return def.ContainerClass.CSSClass(def.Style, "toast")
}

func (def D) iconClass() string {
	return def.IconClass.CSSClass(def.Style, "toast/icon")
}

func (def D) closeDelayMS() int {
	d := def.AutoCloseDelay
	if d == 0 {
		d = 5 * time.Second
	}
	return int(d.Milliseconds())
}

// Retarget returns a HTMX response that retargets the response to display the response
// to the toast container, ignoring the initial target.
// When using error status codes, do not forget to add your error codes to the list of
// codes for which HTMX swaps the content - https://htmx.org/docs/#requests
func Retarget(ctx context.Context, def D, w http.ResponseWriter, statusCode int) error {
	w.Header().Set("Content-Type", "text/html, charset=UTF-8")
	w.Header().Set("HX-Retarget", "#"+def.containerID())
	w.Header().Set("HX-Reswap", "beforeend")
	w.WriteHeader(statusCode)
	return C(def).Render(ctx, w)
}
