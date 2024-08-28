package helper

import (
	"fmt"
	"strconv"

	"github.com/a-h/templ"
)

func List(l ...any) []templ.Component {
	res := make([]templ.Component, len(l))
	for _, c := range l {
		if c, ok := c.(templ.Component); ok {
			res = append(res, c)
			continue
		}
		res = append(res, S(c))
	}
	return res
}

func tostring(v any) string {
	switch v := v.(type) {
	case string:
		return v
	case int:
		return strconv.Itoa(v)
	case uint:
		return strconv.FormatUint(uint64(v), 10)
	case uint64:
		return strconv.FormatUint(v, 10)
	case float64:
		return strconv.FormatFloat(v, 'g', -1, 64)
	case float32:
		return strconv.FormatFloat(float64(v), 'g', -1, 32)
	default:
		return fmt.Sprintf("%v", v)
	}
}
