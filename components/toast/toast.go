package toast

import (
	"time"

	"github.com/jfbus/templ-components/components/icon"
	"github.com/jfbus/templ-components/components/style"
)

const (
	StyleOK      style.Style = 2
	StyleWarning style.Style = 4
	StyleError   style.Style = 8
)

type Close int

const (
	CloseButton Close = iota
	CloseAuto
)

var Defaults = style.Defaults{
	style.StyleDefault: {
		"ContainerClass": {
			style.Color("text-gray-500 bg-white dark:text-gray-400 dark:bg-gray-800"),
			style.Class("flex items-center w-full max-w-xs p-4 rounded-lg shadow"),
		},
		"IconClass": {
			style.Class("inline-flex items-center justify-center flex-shrink-0 w-8 h-8 rounded-lg"),
		},
	},
	StyleOK: {
		"IconClass": {
			style.Color("text-green-500 bg-green-100 dark:bg-green-800 dark:text-green-200"),
		},
	},
	StyleWarning: {
		"IconClass": {
			style.Color("text-orange-500 bg-orange-100 dark:bg-orange-700 dark:text-orange-200"),
		},
	},
	StyleError: {
		"IconClass": {
			style.Color("text-red-500 bg-red-100 dark:bg-red-800 dark:text-red-200"),
		},
	},
}

type D struct {
	ID             string
	Style          style.Style
	Icon           string
	Content        string
	ContainerClass style.D
	IconClass      style.D
	Close          Close
	AutoCloseDelay time.Duration
}

func (def D) id() string {
	return def.ID
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
	return def.ContainerClass.CSSClass(Defaults, def.Style, "ContainerClass")
}

func (def D) iconClass() string {
	return def.IconClass.CSSClass(Defaults, def.Style, "IconClass")
}

func (def D) closeDelayMS() int {
	d := def.AutoCloseDelay
	if d == 0 {
		d = 5 * time.Second
	}
	return int(d.Milliseconds())
}
