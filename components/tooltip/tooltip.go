package tooltip

import "github.com/jfbus/templui/components/style"

func init() {
	style.SetDefaults(style.Defaults{
		"tooltip": {
			style.Default: {
				style.Set("relative"),
				style.Add("before:opacity-0 hover:before:opacity-100 before:absolute before:content-[''] before:z-50 before:border-solid before:border-t-4 before:border-x-[6px] before:border-b-0"),
				style.Add("after:opacity-0 hover:after:opacity-100 after:absolute after:content-[attr(tooltip)] after:z-50 after:px-3 after:py-2 after:text-sm after:font-medium after:rounded-lg after:shadow-sm"),
			},
			StylePositionTop: {
				style.Add("before:-top-1.5 before:left-1/2 before:-translate-x-1/2"),
				style.Add("after:-top-1.5 after:left-1/2 after:-translate-x-1/2 after:-translate-y-full"),
			},
			StylePositionBottom: {
				style.Add("before:top-full before:mt-2 before:left-1/2 before:-translate-x-1/2 before:-translate-y-full before:rotate-180"),
				style.Add("after:top-full after:mt-2 after:left-1/2 after:-translate-x-1/2"),
			},
			StylePositionLeft: {
				style.Add("before:top-1/2 before:-ml-3 before:left-0 before:-translate-y-1/2 before:-rotate-90"),
				style.Add("after:top-1/2 after:-ml-2 after:left-0 after:-translate-x-full after:-translate-y-1/2"),
			},
			StylePositionRight: {
				style.Add("before:top-1/2 before:ml-px before:left-full before:-translate-y-full before:rotate-90"),
				style.Add("after:top-1/2 after:ml-2 after:left-full after:-translate-y-1/2"),
			},
		},
	})
}

const (
	StylePositionTop    style.Style = 1 << 8
	StylePositionBottom style.Style = 1 << 9
	StylePositionLeft   style.Style = 1 << 10
	StylePositionRight  style.Style = 1 << 11
)

type D struct {
	Style  style.Style
	Text   string
	Custom style.Custom
}

func (def D) style() style.Style {
	if def.Style > 0 {
		return def.Style
	}
	return StylePositionTop
}

func (def D) Class() string {
	return style.CSSClass(def.style(), "tooltip", def.Custom)
}
