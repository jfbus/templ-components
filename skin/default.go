package skin

import (
	"github.com/jfbus/templ-components/components/button"
	"github.com/jfbus/templ-components/components/checkboxgroup"
	"github.com/jfbus/templ-components/components/icon"
	"github.com/jfbus/templ-components/components/radiogroup"
	"github.com/jfbus/templ-components/components/style"
	"github.com/jfbus/templ-components/components/table"
	"github.com/jfbus/templ-components/components/toast"
)

var Default = style.Defaults{
	"layout": {
		style.StyleDefault: {
			style.Set("text-gray-900 dark:text-white border-gray-300 dark:border-gray-600"),
		},
	},
	"accordion/title": {
		style.StyleDefault: {
			style.Set("text-gray-500 focus:ring-gray-200 dark:focus:ring-gray-800 dark:text-gray-400 hover:bg-gray-100 dark:hover:bg-gray-800"),
		},
	},
	"button": {
		style.StyleDefault: {
			style.Set("text-white bg-blue-700 hover:bg-blue-800 focus:ring-blue-300 dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"),
		},
		button.StyleAlternative: {
			style.Set("bg-white border-gray-200 hover:bg-gray-100 hover:text-blue-700 focus:ring-gray-100 dark:focus:ring-gray-700 dark:bg-gray-800 dark:text-gray-400 dark:border-gray-600 dark:hover:text-white dark:hover:bg-gray-700"),
		},
		button.StyleOutline: {
			style.Set("text-blue-700 hover:text-white border-blue-700 hover:bg-blue-800 focus:ring-blue-300 dark:border-blue-500 dark:text-blue-500 dark:hover:text-white dark:hover:bg-blue-500 dark:focus:ring-blue-800"),
		},
		button.StyleNoBorder: {
			style.Set("bg-white border-gray-200 hover:bg-gray-100 hover:text-blue-700 focus:ring-gray-100 dark:focus:ring-gray-700 dark:bg-gray-800 dark:text-gray-400 dark:border-gray-600 dark:hover:text-white dark:hover:bg-gray-700"),
		},
	},
	"checkbox/input": {
		style.StyleDefault: {
			style.Set("text-blue-600 bg-gray-100 border-gray-300 focus:ring-blue-500 dark:focus:ring-blue-600 dark:ring-offset-gray-800 dark:bg-gray-700 dark:border-gray-600"),
		},
	},
	"checkboxgroup/checkbox/label": {
		checkboxgroup.StyleLabelOnly: {
			style.Set("dark:peer-checked:text-blue-500 peer-checked:border-blue-600 peer-checked:text-blue-600 hover:text-gray-600 hover:bg-gray-100 dark:bg-gray-800 dark:hover:bg-gray-700"),
		},
	},
	"dropdown": {
		style.StyleDefault: {
			style.Set("bg-white divide-gray-100 dark:bg-gray-700 dark:divide-gray-600"),
		},
	},
	"dropdown/link": {
		style.StyleDefault: {
			style.Set("hover:bg-gray-100 dark:hover:bg-gray-600"),
		},
	},
	"form/validation/message": {
		style.StyleDefault: {
			style.Set("text-red-600 dark:text-red-500"),
		},
		style.StyleValid: {
			style.Set("text-green-600 dark:text-green-500"),
		},
		style.StyleInvalid: {
			style.Set("text-red-600 dark:text-red-500"),
		},
	},
	"icon": {
		icon.StyleBorder: {
			style.Set("dark:bg-gray-800 border-gray-300 dark:border-gray-600"),
		},
	},
	"input/input": {
		style.StyleDefault: {
			style.Set("border-gray-300 dark:border-gray-600 bg-gray-50 focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:placeholder-gray-400 dark:focus:ring-blue-500 dark:focus:border-blue-500"),
		},
		style.StyleDisabled: {
			style.ReplaceVariants("bg", "bg-gray-100 dark:bg-gray-700"),
		},
		style.StyleValid: {
			style.Set("bg-green-50 border-green-500 text-green-900 dark:text-green-400 placeholder-green-700 dark:placeholder-green-500 focus:ring-green-500 focus:border-green-500 dark:bg-gray-700 dark:border-green-500"),
		},
		style.StyleInvalid: {
			style.Set("bg-red-50 border-red-500 text-red-900 placeholder-red-700 focus:ring-red-500 dark:bg-gray-700 focus:border-red-500 dark:text-red-500 dark:placeholder-red-500 dark:border-red-500"),
		},
	},
	"input/icon": {
		style.StyleValid: {
			style.Set("text-green-700 dark:text-green-500"),
		},
		style.StyleInvalid: {
			style.Set("text-red-700 dark:text-red-500"),
		},
	},
	"label": {
		style.StyleValid: {
			style.Set("text-green-700 dark:text-green-500"),
		},
		style.StyleInvalid: {
			style.Set("text-red-700 dark:text-red-500"),
		},
	},
	"navbar": {
		style.StyleDefault: {
			style.Set("bg-white"),
		},
	},
	"radio/input": {
		style.StyleDefault: {
			style.Set("border-gray-300 dark:border-gray-600 text-blue-600 bg-gray-100 border-gray-300 focus:ring-blue-500 dark:focus:ring-blue-600 dark:ring-offset-gray-800 dark:bg-gray-700 dark:border-gray-600"),
		},
	},
	"radio/label": {
		style.StyleDisabled: {
			style.ReplaceVariants("text", "text-gray-400 dark:text-gray-500"),
		},
	},
	"radiogroup/radio/label": {
		radiogroup.StyleLabelOnly: {
			style.Set("dark:peer-checked:text-blue-500 peer-checked:border-blue-600 peer-checked:text-blue-600 hover:text-gray-600 hover:bg-gray-100 dark:bg-gray-800 dark:hover:bg-gray-700"),
		},
	},
	"select": {
		style.StyleDefault: {
			style.Set("border-gray-300 dark:border-gray-600 bg-gray-50 focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:placeholder-gray-400 dark:focus:ring-blue-500 dark:focus:border-blue-500"),
		},
		style.StyleDisabled: {
			style.ReplaceVariants("bg", "bg-gray-100 dark:bg-gray-700"),
		},
		style.StyleValid: {
			style.Set("bg-green-50 border-green-500 text-green-900 dark:text-green-400 placeholder-green-700 dark:placeholder-green-500 focus:ring-green-500 focus:border-green-500 dark:bg-gray-700 dark:border-green-500"),
		},
		style.StyleInvalid: {
			style.Set("bg-red-50 border-red-500 text-red-900 placeholder-red-700 focus:ring-red-500 dark:bg-gray-700 focus:border-red-500 dark:text-red-500 dark:placeholder-red-500 dark:border-red-500"),
		},
	},
	"sidebar": {
		style.StyleDefault: {
			style.Set("bg-gray-50 dark:bg-gray-800"),
		},
	},
	"table/header": {
		style.StyleDefault: {
			style.Set("text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400"),
		},
		table.StyleNoBorder: {
			style.Set("text-gray-900 uppercase dark:text-gray-400"),
		},
	},
	"table/row": {
		table.StyleStripedRows: {
			style.Set("even:bg-gray-50 even:dark:bg-gray-800 dark:border-gray-700"),
		},
		table.StyleNoBorder: {
			style.Set("whitespace-nowrap"),
		},
		table.StyleAddHighlightHover: {
			style.Add("hover:bg-gray-50 dark:hover:bg-gray-600"),
		},
	},
	"textarea/input": {
		style.StyleDefault: {
			style.Set("border-gray-300 dark:border-gray-600 bg-gray-50 focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:placeholder-gray-400 dark:focus:ring-blue-500 dark:focus:border-blue-500"),
		},
		style.StyleDisabled: {
			style.ReplaceVariants("bg", "bg-gray-100 dark:bg-gray-700"),
		},
		style.StyleValid: {
			style.Set("bg-green-50 border-green-500 text-green-900 dark:text-green-400 placeholder-green-700 dark:placeholder-green-500 focus:ring-green-500 focus:border-green-500 dark:bg-gray-700 dark:border-green-500"),
		},
		style.StyleInvalid: {
			style.Set("bg-red-50 border-red-500 text-red-900 placeholder-red-700 focus:ring-red-500 dark:bg-gray-700 focus:border-red-500 dark:text-red-500 dark:placeholder-red-500 dark:border-red-500"),
		},
	},
	"textarea/icon": {
		style.StyleValid: {
			style.Set("text-green-700 dark:text-green-500"),
		},
		style.StyleInvalid: {
			style.Set("text-red-700 dark:text-red-500"),
		},
	},
	"toast": {
		style.StyleDefault: {
			style.Set("bg-white"),
		},
	},
	"toast/icon": {
		toast.StyleOK: {
			style.Set("text-green-500 bg-green-100 dark:bg-green-800 dark:text-green-200"),
		},
		toast.StyleWarning: {
			style.Set("text-orange-500 bg-orange-100 dark:bg-orange-700 dark:text-orange-200"),
		},
		toast.StyleError: {
			style.Set("text-red-500 bg-red-100 dark:bg-red-800 dark:text-red-200"),
		},
	},
}
