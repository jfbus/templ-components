// Package dropzone implements drap & drop upload drop zones.
package dropzone

import (
	"github.com/a-h/templ"
	"github.com/jfbus/templui/components/form/validation/message"
	"github.com/jfbus/templui/components/icon"
	"github.com/jfbus/templui/components/size"
	"github.com/jfbus/templui/components/style"
)

func init() {
	style.SetDefaults(style.Defaults{
		"dropzone": {
			style.Default: {
				style.Set("relative flex items-center justify-center w-full"),
			},
		},
		"dropzone/input": {
			style.Default: {
				style.Set("absolute top-0 left-0 w-full h-full opacity-0 cursor-pointer"),
			},
			style.Disabled: {
				style.Add("cursor-not-allowed"),
			},
		},
		"dropzone/label": {
			style.Default: {
				style.Set("w-full flex flex-col items-center justify-center text-center py-6 gap-2 bg-gray-50 text-gray-500 dark:text-gray-400 dark:bg-gray-700 dark:border-gray-600 border-2 border-gray-300 border-dashed rounded-lg"),
			},
		},
		"dropzone/label/drag": {
			style.Default: {
				style.Set("!bg-gray-100 !border-solid"),
			},
		},
		"dropzone/label/icon": {
			style.Default: {
				style.Set("w-8 h-8"),
			},
		},
		"dropzone/label/label": {
			style.Default: {
				style.Set("text-sm"),
			},
		},
		"dropzone/label/allowed": {
			style.Default: {
				style.Set("text-xs"),
			},
		},
	})
}

const DefaultDragMessage = `<span class="font-semibold">Click to upload</span> or drag and drop`
const DefaultDropMessage = `Drop to upload`

// D is the definition for input fields.
type D struct {
	// ID is the input id (Name if not set).
	ID string
	// Name is the input name.
	Name     string
	Multiple bool
	// AllowedTypes is the list of allowed mime types
	AllowedTypes []string
	// Style defines the style (style.Default, Valid or Invalid).
	Style style.Style
	// Icon allows you to define a custom icon.
	Icon *icon.D
	// DropIcon sets the icon displayed when dragging a file over the drop zone.
	DropIcon *icon.D
	// IconSize defines the default icon size (ignored when using a custom icon).
	IconSize size.Size
	// DragMessage is the default label message (default: "Click to upload or drag and drop").
	DragMessage string
	// DropMessage is the label message displayed when dragging a file over the drop zone.
	DropMessage string
	// AllowedTypesMessage is the text that specifies allowed types, max size (e.g. "PNG, JPG or GIF (Max. 600x300px)")
	AllowedTypesMessage string
	// Message adds a validation message below the field.
	// Just add &message.D{} to add automatic validation.
	//playground:import:github.com/jfbus/templui/components/form/validation/message
	//playground:default:&message.D{Message: "Validation message"}
	Message *message.D
	// Disabled disables the input.
	Disabled bool
	// Invalid marks the input as invalid.
	Invalid bool
	// Loader displays a spinning loader when an HTMX action is triggered by the input.
	Loader bool
	// CustomStyle defines a custom style.
	CustomStyle style.Custom
	// Attributes stores additional attributes (e.g. HTMX attributes).
	Attributes templ.Attributes
}

func (def D) dragMessage() string {
	if def.DragMessage == "" {
		return DefaultDragMessage
	}
	return def.DragMessage
}

func (def D) dropMessage() string {
	if def.DropMessage == "" {
		return DefaultDropMessage
	}
	return def.DropMessage
}

func (def D) style() style.Style {
	if def.Invalid {
		return def.Style | style.Invalid
	}
	if def.Disabled {
		return def.Style | style.Disabled
	}
	return def.Style
}

func (def D) id() string {
	if def.ID != "" {
		return def.ID
	}
	return def.Name
}

func (def D) icon() icon.D {
	if def.Icon != nil {
		if def.Icon.Attributes == nil {
			def.Icon.Attributes = make(templ.Attributes)
		}
		def.Icon.Attributes[":class"] = "drag && 'hidden'"
		return *def.Icon
	}
	sz := def.IconSize
	if sz == size.Inherit {
		sz = size.FourXL
	}
	cc := style.Custom{
		"icon": style.Compute(def.style(), "dropzone/label/icon", def.CustomStyle),
	}
	return icon.D{
		Icon:        icon.CloudUpload,
		CustomStyle: cc,
		Size:        def.IconSize,
		Attributes:  templ.Attributes{":class": "drag && 'hidden'"},
	}
}

func (def D) dropIcon() icon.D {
	if def.Icon != nil {
		if def.Icon.Attributes == nil {
			def.Icon.Attributes = make(templ.Attributes)
		}
		def.Icon.Attributes[":class"] = "drag || 'hidden'"
		return *def.Icon
	}
	sz := def.IconSize
	if sz == size.Inherit {
		sz = size.FourXL
	}
	cc := style.Custom{
		"icon": style.Compute(def.style(), "dropzone/label/icon", def.CustomStyle),
	}
	return icon.D{
		Icon:        icon.ArrowBigDown,
		CustomStyle: cc,
		Size:        def.IconSize,
		Attributes:  templ.Attributes{":class": "drag || 'hidden'"},
	}
}

func (def D) class(k string) string {
	return style.CSSClass(def.style(), k, def.CustomStyle)
}

func (def D) inputClassInvalid() string {
	return style.Delta(def.Style, style.Invalid, "dropzone/input", def.CustomStyle)
}

func (def D) message() message.D {
	m := *def.Message
	m.InputName = def.Name
	m.Style = def.style()
	return m
}
