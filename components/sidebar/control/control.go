package control

import (
	"github.com/a-h/templ"
	"github.com/jfbus/templ-components/components/button"
	"github.com/jfbus/templ-components/components/icon"
	"github.com/jfbus/templ-components/components/sidebar"
	"github.com/jfbus/templ-components/components/size"
	"github.com/jfbus/templ-components/components/style"
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
		Icon:  icon.Menu,
		Style: button.StyleOutline,
		Class: style.D{style.Class("md:hidden")},
		Size:  size.L,
		Attributes: templ.Attributes{
			"@click.stop": "sidebar = !sidebar",
		},
	}
}
