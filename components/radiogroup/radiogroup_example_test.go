package radiogroup_test

import (
	"context"
	"os"

	"github.com/jfbus/templ-components/components/radio"
	"github.com/jfbus/templ-components/components/radiogroup"
	"github.com/jfbus/templ-components/components/style"
	"github.com/jfbus/templ-components/skin"
)

func ExampleC() {
	style.SetSkin(skin.Default)
	c := radiogroup.C(radiogroup.D{
		Name:  "group",
		Style: radiogroup.StyleBordered,
		Radios: []radio.D{
			{Label: "Option 1", Value: "1"},
			{Label: "Option 2", Value: "2"},
		},
	})
	_ = c.Render(context.TODO(), os.Stdout)
	// output: <div class="flex flex-col sm:flex-row gap-4"><div class="flex items-center px-4 border rounded w-full"><input type="radio" id="group-1" name="group" value="1" class="w-4 h-4 focus:ring-2 border-gray-300 dark:border-gray-600 text-blue-600 bg-gray-100 focus:ring-blue-500 dark:focus:ring-blue-600 dark:ring-offset-gray-800 dark:bg-gray-700"><label for="group-1" class="ms-2 text-sm font-medium py-3" :class="hasError(&#39;&#39;) &amp;&amp; &#39;text-red-700 dark:text-red-500&#39;">Option 1</label></div><div class="flex items-center px-4 border rounded w-full"><input type="radio" id="group-2" name="group" value="2" class="w-4 h-4 focus:ring-2 border-gray-300 dark:border-gray-600 text-blue-600 bg-gray-100 focus:ring-blue-500 dark:focus:ring-blue-600 dark:ring-offset-gray-800 dark:bg-gray-700"><label for="group-2" class="ms-2 text-sm font-medium py-3" :class="hasError(&#39;&#39;) &amp;&amp; &#39;text-red-700 dark:text-red-500&#39;">Option 2</label></div></div>
}
