package tooltip_test

import (
	"context"
	"os"

	"github.com/jfbus/templui/components/button"
	"github.com/jfbus/templui/components/tooltip"
)

func ExampleC() {
	c := button.C(button.D{
		Label: "Label",
		Tooltip: &tooltip.D{
			Style: tooltip.StylePositionBottom,
			Text:  "tooltip",
		},
	})
	_ = c.Render(context.TODO(), os.Stdout)
	// output: <button type="button" class="rounded-lg font-medium focus:ring-4 focus:outline-none p-2.5 text-sm relative before:opacity-0 hover:before:opacity-100 before:absolute before:content-[&#39;&#39;] before:z-100 before:border-solid before:border-t-4 before:border-x-[6px] before:border-b-0 after:opacity-0 hover:after:opacity-100 after:absolute after:content-[attr(tooltip)] after:z-100 after:px-3 after:py-2 after:text-sm after:font-medium after:rounded-lg after:shadow-sm before:top-full before:mt-2 before:left-1/2 before:-translate-x-1/2 before:-translate-y-full before:rotate-180 after:top-full after:mt-2 after:left-1/2 after:-translate-x-1/2" tooltip="tooltip"><span class="">Label</span> </button>
}
