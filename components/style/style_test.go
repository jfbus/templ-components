package style_test

import (
	"testing"

	"github.com/jfbus/templ-components/components/style"
	"github.com/stretchr/testify/assert"
)

const (
	StyleOverrideClass = 2
	StyleOverrideColor = 4
	StyleAdd           = 8
)

func TestDefaults(t *testing.T) {
	defaults := style.Defaults{
		style.StyleDefault: {
			"Test": {
				style.Set("class"), style.Color("color"),
			},
		},
		StyleOverrideClass: {
			"Test": {
				style.Set("override"),
			},
		},
		StyleOverrideColor: {
			"Test": {
				style.Color("override"),
			},
		},
		StyleAdd: {
			"Test": {
				style.Add("add"),
			},
		},
	}
	assert.Equal(t, "class color", style.D{}.CSSClass(defaults, style.StyleDefault, "Test"))
	assert.Equal(t, "override color", style.D{}.CSSClass(defaults, StyleOverrideClass, "Test"))
	assert.Equal(t, "class override", style.D{}.CSSClass(defaults, StyleOverrideColor, "Test"))
	assert.Equal(t, "class color add", style.D{}.CSSClass(defaults, StyleAdd, "Test"))
}

func TestOverrides(t *testing.T) {
	defaults := style.Defaults{
		style.StyleDefault: {
			"Test": {
				style.Set("class"), style.Color("color"),
			},
		},
	}
	{
		over := style.D{style.Set("override")}
		assert.Equal(t, "override color", over.CSSClass(defaults, style.StyleDefault, "Test"))
	}
	{
		over := style.D{style.Color("override")}
		assert.Equal(t, "class override", over.CSSClass(defaults, style.StyleDefault, "Test"))
	}
	{
		over := style.D{style.Add("add")}
		assert.Equal(t, "class color add", over.CSSClass(defaults, style.StyleDefault, "Test"))
	}
}

func TestReplace(t *testing.T) {
	defaults := style.Defaults{
		style.StyleDefault: {
			"Test": {
				style.Set("text-normal mb-4 hover:mb-2 w-2 dark:focus:mb-5 emb-2 mbt-2"), style.Color("text-normal mb-4 w-2"),
			},
		},
	}
	{
		over := style.D{style.ReplaceClass("mb", "replaced")}
		assert.Equal(t, "text-normal w-2 emb-2 mbt-2 replaced text-normal mb-4 w-2", over.CSSClass(defaults, style.StyleDefault, "Test"))
	}
	{
		over := style.D{style.ReplaceColor("mb", "replaced")}
		assert.Equal(t, "text-normal mb-4 hover:mb-2 w-2 dark:focus:mb-5 emb-2 mbt-2 text-normal w-2 replaced", over.CSSClass(defaults, style.StyleDefault, "Test"))
	}
	{
		over := style.D{style.Replace("mb", "replaced")}
		assert.Equal(t, "text-normal w-2 emb-2 mbt-2 replaced text-normal w-2 replaced", over.CSSClass(defaults, style.StyleDefault, "Test"))
	}
}
