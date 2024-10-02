package social

import (
	"github.com/jfbus/templ-components/components/icon"
	"github.com/jfbus/templ-components/components/size"
	"github.com/jfbus/templ-components/components/style"
)

type Type int

const (
	TypeFacebook Type = iota
	TypeInstagram
	TypeGithub
	TypeTwitch
	TypeLinkedIn
	TypeYoutube
	TypeDribble
)

// D is a button.D preset that renders social links.
type D struct {
	Type  Type
	Size  size.Size
	Class style.D
}

func (def D) icon() string {
	switch def.Type {
	case TypeFacebook:
		return icon.Facebook
	case TypeInstagram:
		return icon.Instagram
	case TypeGithub:
		return icon.Github
	case TypeTwitch:
		return icon.Twitch
	case TypeLinkedIn:
		return icon.Linkedin
	case TypeYoutube:
		return icon.Youtube
	case TypeDribble:
		return icon.Dribbble
	}
	return ""
}
