package layout

import (
	"strconv"

	"github.com/jfbus/templ-components/components/navbar"
	"github.com/jfbus/templ-components/components/sidebar"
	"github.com/jfbus/templ-components/components/size"
	"github.com/jfbus/templ-components/components/style"
	"github.com/jfbus/templ-components/components/toast/container"
)

type NavbarHeight size.Size

func (h NavbarHeight) paddingTop() string {
	return "pt-14"
}

type D struct {
	Navbar       *navbar.D
	Sidebar      *sidebar.D
	Toasts       *container.D
	State        map[string]string
	NavbarHeight NavbarHeight
	Class        style.D
}

func (def D) state() string {
	st := `{
  sidebar: false`
	if def.State != nil {
		for k, v := range def.State {
			st += `,
  ` + k + ": " + strconv.Quote(v)
		}
	}
	return st + "}"
}

func (def D) sidebar() sidebar.D {
	d := *def.Sidebar
	d.ContainerClass = append(d.ContainerClass, style.Add(def.NavbarHeight.paddingTop()))
	return d
}

func (def D) toasts() container.D {
	if def.Toasts == nil {
		return container.D{
			Class: style.D{style.Add(def.NavbarHeight.paddingTop())},
		}
	}
	d := *def.Toasts
	d.Class = append(d.Class, style.Add(def.NavbarHeight.paddingTop()))
	return d
}

func (def D) class() string {
	return def.Class.CSSClass(style.StyleDefault, "layout")
}
