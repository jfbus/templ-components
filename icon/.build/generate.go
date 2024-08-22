package main

import (
	"os"
	"path"
	"regexp"
	"strings"
	"unicode"
)

const assetsPath = ".build/node_modules/lucide-static/icons"

var re = regexp.MustCompile("[\n\r\t ]+")

func main() {
	des, err := os.ReadDir(assetsPath)
	if err != nil {
		panic(err)
	}

	var names []string
	svgs := map[string]string{}
	max := 0
	for _, de := range des {
		if !strings.HasSuffix(de.Name(), ".svg") {
			continue
		}
		svg, err := os.ReadFile(path.Join(assetsPath, de.Name()))
		if err != nil {
			panic(err)
		}
		var name []rune
		upper := true
		for _, r := range strings.TrimSuffix(de.Name(), ".svg") {
			if r == '-' {
				upper = true
				continue
			}
			if upper {
				name = append(name, unicode.ToUpper(r))
				upper = false
			} else {
				name = append(name, r)
			}
		}
		ssvg := strings.TrimSpace(string(svg))
		if i := strings.Index(ssvg, "<svg"); i >= 0 {
			ssvg = ssvg[i:]
		} else {
			continue
		}
		if len(name) > max {
			max = len(name)
		}
		names = append(names, string(name))
		svgs[string(name)] = re.ReplaceAllString(ssvg, " ")
	}
	file := `// Auto-generated code, DO NOT EDIT.
package icon

const (`
	for _, name := range names {
		file += `
	` + name + strings.Repeat(" ", max-len(name)) + " = `" + svgs[name] + "`"
	}
	file += `
)`
	err = os.WriteFile("icon_lucide.go", []byte(file), 0644)
	if err != nil {
		panic(err)
	}
}
