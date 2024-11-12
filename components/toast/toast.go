package toast

import (
	"context"
	"net/http"
	"time"

	"github.com/jfbus/templui/components/icon"
	"github.com/jfbus/templui/components/style"
	"github.com/jfbus/templui/components/toast/container"
	"github.com/rs/xid"
)

const (
	StyleOK      style.Style = 1 << 8
	StyleWarning style.Style = 1 << 9
	StyleError   style.Style = 1 << 10
)

type Close int

const (
	CloseAuto Close = iota
	CloseButton
)

func init() {
	style.SetDefaults(style.Defaults{
		"toast": {
			style.Default: {
				style.Set("flex items-center w-full max-w-xs p-4 rounded-lg shadow"),
			},
		},
		"toast/icon": {
			style.Default: {
				style.Set("inline-flex items-center justify-center flex-shrink-0 w-8 h-8 rounded-lg"),
			},
		},
		"toast/content": {
			style.Default: {
				style.Set("ms-3 text-sm font-normal"),
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
	Close   Close
	// AutoCloseDelay defines the close delay (default 5s).
	AutoCloseDelay time.Duration
	// CustomStyle defines a custom style.
	// 	style.Custom{
	// 		"toast":      style.D{style.Add("...")},
	// 		"toast/icon": style.D{style.Add("...")},
	//	}
	CustomStyle style.Custom
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

func (def D) class(k string) string {
	return style.CSSClass(def.Style, k, def.CustomStyle)
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
