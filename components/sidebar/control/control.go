package control

import (
	"github.com/a-h/templ"
	"github.com/jfbus/templui/components/button"
	"github.com/jfbus/templui/components/icon"
	"github.com/jfbus/templui/components/sidebar"
	"github.com/jfbus/templui/components/size"
	"github.com/jfbus/templui/components/style"
)

type D struct {
	SidebarID string
	Icon      any
}

func (def D) sidebarId() string {
	if def.SidebarID != "" {
		return def.SidebarID
	}
	return sidebar.DefaultID
}

func (def D) button() button.D {
	return button.D{
		Icon:        icon.Menu,
		Style:       button.StyleOutline,
		CustomStyle: style.Custom{"button": {style.Set("md:hidden")}},
		Size:        size.L,
		Attributes: templ.Attributes{
			"@click.stop": "sidebar = !sidebar",
		},
	}
}
