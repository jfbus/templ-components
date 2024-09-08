package main

import (
	"os"
	"path"
	"regexp"
	"strings"
	"unicode"
)

const assetsPath = "internal/build/node_modules/lucide-static/icons"

var re = regexp.MustCompile("[\n\r\t ]+")

var reSVG = regexp.MustCompile("<svg [^>]+> *(.*) *</svg>")

func main() {
	des, err := os.ReadDir(assetsPath)
	if err != nil {
		panic(err)
	}

	var icons []string
	svgs := map[string]string{}
	names := map[string]string{}
	maxLenIcon, maxLenName := 0, 0
	for _, de := range des {
		if !strings.HasSuffix(de.Name(), ".svg") {
			continue
		}
		svg, err := os.ReadFile(path.Join(assetsPath, de.Name()))
		if err != nil {
			panic(err)
		}
		var icon []rune
		upper := true
		name := strings.TrimSuffix(de.Name(), ".svg")
		for _, r := range name {
			if r == '-' {
				upper = true
				continue
			}
			if upper {
				icon = append(icon, unicode.ToUpper(r))
				upper = false
			} else {
				icon = append(icon, r)
			}
		}
		ssvg := re.ReplaceAllString(string(svg), " ")
		if sm := reSVG.FindStringSubmatch(ssvg); len(sm) == 2 {
			ssvg = sm[1]
		} else {
			continue
		}
		if len(icon) > maxLenIcon {
			maxLenIcon = len(icon)
		}
		if len(name) > maxLenName {
			maxLenName = len(name)
		}
		icons = append(icons, string(icon))
		svgs[string(icon)] = strings.ReplaceAll(ssvg, " /> ", "/>")
		names[string(icon)] = name
	}
	file := `// Code generated - DO NOT EDIT\.$
package icon

const (`
	for _, icon := range icons {
		file += `
	` + icon + strings.Repeat(" ", maxLenIcon-len(icon)) + " = `" + svgs[icon] + "`"
	}
	file += `
)

var NameToIcon = map[string]string{`
	for _, icon := range icons {
		file += `
	"` + names[icon] + `":` + strings.Repeat(" ", maxLenName-len(names[icon])) + " `" + svgs[icon] + "`,"
	}
	file += `
}`
	err = os.WriteFile("icon_lucide.go", []byte(file), 0644)
	if err != nil {
		panic(err)
	}
}
