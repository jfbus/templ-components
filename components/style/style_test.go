package style_test

import (
	"testing"

	"github.com/jfbus/templ-components/components/style"
	"github.com/stretchr/testify/assert"
)

const (
	StyleOverride = 2
	StyleAdd      = 8
)

func TestDefaults(t *testing.T) {
	style.SetDefaults(style.Defaults{
		"TestDefaults": {
			style.Default: {
				style.Set("class color"),
			},
			StyleOverride: {
				style.Set("override"),
			},
			StyleAdd: {
				style.Add("add"),
			},
		},
	})
	assert.Equal(t, "class color", style.CSSClass(style.Default, "TestDefaults", nil))
	assert.Equal(t, "override", style.CSSClass(StyleOverride, "TestDefaults", nil))
	assert.Equal(t, "class color add", style.CSSClass(StyleAdd, "TestDefaults", nil))
}

func TestSkin(t *testing.T) {
	style.SetDefaults(style.Defaults{
		"TestSkin": {
			style.Default: {
				style.Set("class"),
			},
		},
	})
	style.SetSkin(style.Defaults{
		"TestSkin": {
			style.Default: {
				style.Set("color"),
			},
		},
	})
	assert.Equal(t, "class color", style.CSSClass(style.Default, "TestSkin", nil))
}

func TestCustom(t *testing.T) {
	style.SetDefaults(style.Defaults{
		"TestCustom": {
			style.Default: {
				style.Set("class color"),
			},
		},
	})
	{
		custom := style.Custom{"TestCustom": {style.Set("override")}}
		assert.Equal(t, "override", style.CSSClass(style.Default, "TestCustom", custom))
	}
	{
		custom := style.Custom{"TestCustom": {style.Add("add")}}
		assert.Equal(t, "class color add", style.CSSClass(style.Default, "TestCustom", custom))
	}
	{
		custom := style.Custom{"TestCustom": {style.Replace("class", "override")}}
		assert.Equal(t, "override color", style.CSSClass(style.Default, "TestCustom", custom))
	}
}

func TestReplaceVariants(t *testing.T) {
	style.SetDefaults(style.Defaults{
		"TestReplaceVariants": {
			style.Default: {
				style.Set("text-normal mb-4 hover:mb-2 w-2 mb-2/50 dark:focus:mb-5 emb-2 mbt-2"),
			},
		},
	})
	custom := style.Custom{"TestReplaceVariants": {style.ReplaceVariants("mb", "replaced")}}
	assert.Equal(t, "text-normal w-2 emb-2 mbt-2 replaced", style.CSSClass(style.Default, "TestReplaceVariants", custom))
}
