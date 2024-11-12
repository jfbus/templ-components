package button_test

import (
	"context"
	"os"

	"github.com/a-h/templ"
	"github.com/jfbus/templui/components/button"
	"github.com/jfbus/templui/components/icon"
	"github.com/jfbus/templui/components/size"
)

func ExampleC() {
	c := button.C(button.D{
		ID:     "item_submit",
		Type:   button.Submit,
		Style:  button.StyleOutline | button.StylePill | button.StyleHideLabelAlways | button.StylePill,
		Label:  "Save",
		Size:   size.L,
		Icon:   icon.Save,
		Loader: true,
		Attributes: templ.Attributes{
			"hx-post":   "/items/save",
			"hx-target": "#item_list",
		},
	})
	_ = c.Render(context.TODO(), os.Stdout)
	// output: <button type="submit" id="item_submit" class="rounded-full font-medium focus:ring-4 focus:outline-none border p-2.5 text-base inline-flex items-center justify-center" hx-post="/items/save" hx-target="#item_list" hx-indicator="#item_submit-indicator"><svg class="h-5 w-5" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M15.2 3a2 2 0 0 1 1.4.6l3.8 3.8a2 2 0 0 1 .6 1.4V19a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2z"/><path d="M17 21v-7a1 1 0 0 0-1-1H8a1 1 0 0 0-1 1v7"/><path d="M7 3v4a1 1 0 0 0 1 1h7"/></svg><span class="sr-only">Save</span> <span id="item_submit-indicator" class="htmx-indicator"><svg class="animate-spin h-5 w-5" aria-hidden="true" viewBox="0 0 100 101" fill="none" xmlns="http://www.w3.org/2000/svg"><path d="M100 50.5908C100 78.2051 77.6142 100.591 50 100.591C22.3858 100.591 0 78.2051 0 50.5908C0 22.9766 22.3858 0.59082 50 0.59082C77.6142 0.59082 100 22.9766 100 50.5908ZM9.08144 50.5908C9.08144 73.1895 27.4013 91.5094 50 91.5094C72.5987 91.5094 90.9186 73.1895 90.9186 50.5908C90.9186 27.9921 72.5987 9.67226 50 9.67226C27.4013 9.67226 9.08144 27.9921 9.08144 50.5908Z" fill="currentColor"></path><path d="M93.9676 39.0409C96.393 38.4038 97.8624 35.9116 97.0079 33.5539C95.2932 28.8227 92.871 24.3692 89.8167 20.348C85.8452 15.1192 80.8826 10.7238 75.2124 7.41289C69.5422 4.10194 63.2754 1.94025 56.7698 1.05124C51.7666 0.367541 46.6976 0.446843 41.7345 1.27873C39.2613 1.69328 37.813 4.19778 38.4501 6.62326C39.0873 9.04874 41.5694 10.4717 44.0505 10.1071C47.8511 9.54855 51.7191 9.52689 55.5402 10.0491C60.8642 10.7766 65.9928 12.5457 70.6331 15.2552C75.2735 17.9648 79.3347 21.5619 82.5849 25.841C84.9175 28.9121 86.7997 32.2913 88.1811 35.8758C89.083 38.2158 91.5421 39.6781 93.9676 39.0409Z" fill="currentFill"></path></svg></span></button>
}
