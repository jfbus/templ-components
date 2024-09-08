package buttongroup_test

import (
	"context"
	"os"

	"github.com/a-h/templ"
	"github.com/jfbus/templ-components/components/button"
	"github.com/jfbus/templ-components/components/buttongroup"
	"github.com/jfbus/templ-components/components/icon"
	"github.com/jfbus/templ-components/components/size"
	"github.com/jfbus/templ-components/components/style"
)

func ExampleC() {
	c := buttongroup.C(buttongroup.D{
		Size: size.S,
		Buttons: []button.D{
			{
				Icon:      icon.ArrowDownNarrowWide,
				Label:     "Sort",
				HideLabel: true,
				Class: style.D{
					Color: "text-gray-800 bg-gray-100 border-gray-300 hover:bg-gray-200 focus:ring-gray-100 dark:bg-gray-700 dark:text-white dark:border-gray-700 dark:hover:bg-gray-600 dark:hover:border-gray-600 dark:focus:ring-gray-700",
				},
				Attributes: templ.Attributes{
					"@click": "$dispatch('input', '')",
				},
			},
			{
				Icon:      icon.Heart,
				Label:     "Rating",
				HideLabel: true,
				Attributes: templ.Attributes{
					"@click": "$dispatch('input', 'rating')",
				},
			},
			{
				Icon:      icon.Banknote,
				Label:     "Price",
				HideLabel: true,
				Attributes: templ.Attributes{
					"@click": "$dispatch('input', 'price')",
				},
			},
		},
		Attributes: templ.Attributes{
			"x-model": "sort",
		},
	})
	c.Render(context.TODO(), os.Stdout)
}
