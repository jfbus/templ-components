package checkbox_test

import (
	"context"
	"os"

	"github.com/jfbus/templ-components/components/checkbox"
	"github.com/jfbus/templ-components/components/style"
	"github.com/jfbus/templ-components/skin"
)

func ExampleD() {
	style.SetSkin(skin.Default)
	c := checkbox.C(checkbox.D{
		Name:  "accept",
		Label: "Title",
	})
	_ = c.Render(context.TODO(), os.Stdout)
	// output: <div class="flex items-center"><input type="checkbox" id="accept" name="accept" class="ms-2 w-4 h-4 focus:ring-2 rounded text-blue-600 bg-gray-100 border-gray-300 focus:ring-blue-500 dark:focus:ring-blue-600 dark:ring-offset-gray-800 dark:bg-gray-700 dark:border-gray-600"><label for="accept" class="ms-2 text-sm font-medium" :class="hasError(&#39;&#39;) &amp;&amp; &#39;text-red-700 dark:text-red-500&#39;">Title</label></div>
}
