package accordion

import (
	"strconv"

	"github.com/jfbus/templui/components/accordion/element"
	"github.com/jfbus/templui/components/style"
)

// D defines an accordion.
type D struct {
	// ID is the accordion ID (mandatory).
	ID string
	// Children defines the list of elements in the accordion.
	//playground:import:github.com/jfbus/templui/components/text
	//playground:import:github.com/jfbus/templui/components/accordion/element
	//playground:default:[]element.D{{Title: "Section 1", Content:text.C("Content 1")},{Title: "Section 2", Content:text.C("Content 2")}}
	Children []element.D
	// CustomStyle defines a custom style for children elements.
	// 	style.Custom{
	// 		"accordion/element/title":   style.D{style.Add("text-xl")},
	// 		"accordion/element/content": style.D{style.Add("text-sm")},
	//	}
	CustomStyle style.Custom
}

func (def D) id() string {
	return def.ID
}

func (def D) defaultState() string {
	st := "{ _open:'"
	for i := range def.Children {
		if def.Children[i].Open {
			st += def.ID + strconv.Itoa(i+1)
			break
		}
	}
	st += `',
	isOpen(value) {
		return this._open == value
	},
	toggle(value) {
		if (this._open == value) {
			this._open = ''
		} else {
			this._open = value
		}
	},
}`
	return st
}

func (def D) children() []element.D {
	for i := range def.Children {
		def.Children[i].ID = def.ID + strconv.Itoa(i+1)
		def.Children[i].CustomStyle.AddBefore(def.CustomStyle)
	}
	return def.Children
}
