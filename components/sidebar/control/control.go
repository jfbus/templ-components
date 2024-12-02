package control

import (
	"github.com/a-h/templ"
	"github.com/jfbus/templui/components/button"
	"github.com/jfbus/templui/components/icon"
	"github.com/jfbus/templui/components/size"
	"github.com/jfbus/templui/components/style"
)

type D struct {
	Icon string
}

func (def D) icon() string {
	if def.Icon == "" {
		return icon.Menu
	}
	return def.Icon
}

func (def D) button() button.D {
	return button.D{
		Icon:        def.icon(),
		Style:       button.StyleOutline,
		CustomStyle: style.Custom{"button": {style.Set("md:hidden")}},
		Size:        size.L,
		Attributes: templ.Attributes{
			"@click.stop": "sidebar = !sidebar",
		},
	}
}
