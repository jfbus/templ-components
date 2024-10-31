package social

import (
	"github.com/jfbus/templ-components/components/icon"
	"github.com/jfbus/templ-components/components/size"
	"github.com/jfbus/templ-components/components/style"
)

type Type int

const (
	Facebook Type = iota
	Instagram
	Github
	Twitch
	LinkedIn
	Youtube
	Dribble
)

// D is a button.D preset that renders social links.
type D struct {
	Type  Type
	Link  string
	Size  size.Size
	Class style.D
}

func (def D) icon() string {
	switch def.Type {
	case Facebook:
		return icon.Facebook
	case Instagram:
		return icon.Instagram
	case Github:
		return icon.Github
	case Twitch:
		return icon.Twitch
	case LinkedIn:
		return icon.Linkedin
	case Youtube:
		return icon.Youtube
	case Dribble:
		return icon.Dribbble
	}
	return ""
}

func (def D) class() string {
	return def.Class.CSSClass(style.StyleDefault, "social")
}
