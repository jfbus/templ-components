package size

type Size int

const (
	Inherit Size = iota
	XS
	S
	Normal
	L
	XL
	TwoXL
	ThreeXL
	FourXL
	FiveXL
	SixXL
	SevenXL
	EightXL
	NineXL
	Full
)

var Default Size

var tailwindWidth = map[Size]string{
	XS:      "w-3",
	S:       "w-3.5",
	Normal:  "w-4",
	L:       "w-[18px]",
	XL:      "w-5",
	TwoXL:   "w-6",
	ThreeXL: "w-[30px]",
	FourXL:  "w-9",
	FiveXL:  "w-12",
	SixXL:   "w-[60px]",
	SevenXL: "w-[72px]",
	EightXL: "w-24",
	NineXL:  "w-32",
	Full:    "w-full",
}

var tailwindHeight = map[Size]string{
	XS:      "h-3",
	S:       "h-3.5",
	Normal:  "h-4",
	L:       "h-[18px]",
	XL:      "h-5",
	TwoXL:   "h-6",
	ThreeXL: "h-[30px]",
	FourXL:  "h-9",
	FiveXL:  "h-12",
	SixXL:   "h-[60px]",
	SevenXL: "h-[72px]",
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
