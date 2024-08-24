package size

type Size string

const (
	Inherit Size = ""
	XS           = "xs"
	S            = "sm"
	Normal       = "base"
	L            = "lg"
	XL           = "xl"
	TwoXL        = "2xl"
	ThreeXL      = "3xl"
	FourXL       = "4xl"
	FiveXL       = "5xl"
	SixXL        = "6xl"
	SevenXL      = "7xl"
	EightXL      = "8xl"
	NineXL       = "9xl"
	Full         = "full"
)

var Default Size

var tailwindWidth = map[Size]string{
	XS:      "w-3",
	S:       "w-3.5",
	Normal:  "w-4",
	L:       "w-4",
	XL:      "w-5",
	TwoXL:   "w-6",
	ThreeXL: "w-7",
	FourXL:  "w-9",
	FiveXL:  "w-12",
	SixXL:   "w-14",
	SevenXL: "w-16",
	EightXL: "w-24",
	NineXL:  "w-32",
	Full:    "w-full",
}

var tailwindHeight = map[Size]string{
	XS:      "h-3",
	S:       "h-3.5",
	Normal:  "h-4",
	L:       "h-4",
	XL:      "h-5",
	TwoXL:   "h-6",
	ThreeXL: "h-7",
	FourXL:  "h-9",
	FiveXL:  "h-12",
	SixXL:   "h-14",
	SevenXL: "h-16",
	EightXL: "h-24",
	NineXL:  "h-32",
	Full:    "h-full",
}

func (s Size) Height() string {
	if s == Inherit {
		return ""
	}
	return tailwindHeight[s]
}

func (s Size) Width() string {
	if s == Inherit {
		return ""
	}
	return tailwindWidth[s]
}

func (s Size) Class() string {
	if s == Inherit {
		return ""
	}
	return s.Width() + " " + s.Height()
}
