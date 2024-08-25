package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
)

const defaultConfig = `
/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
  ],
  theme: {
    extend: {},
  },
  plugins: [
    require('flowbite/plugin'),
  ],
}`

func main() {
	_, p, _, _ := runtime.Caller(0)
	p = path.Dir(p)
	err := updateConfig(p)
	if errors.Is(err, os.ErrNotExist) {
		err = createConfig(p, defaultConfig)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func updateConfig(path string) error {
	content, err := os.ReadFile("tailwind.config.js")
	if err != nil {
		return fmt.Errorf("read tailwind.config.js: %w", path, err)
	}
	err = os.Rename("tailwind.config.js", "tailwind.config.js.saved")
	if err != nil {
		return fmt.Errorf("rename tailwind.config.js: %w", path, err)
	}
	return createConfig(path, string(content))
}

func createConfig(path string, content string) error {
	fd, err := os.Create("tailwind.config.js")
	if err != nil {
		return fmt.Errorf("create tailwind.config.js: %w", err)
	}
	defer fd.Close()
	var tabs string
	incontent := false
	scanner := bufio.NewScanner(strings.NewReader(content))
	for scanner.Scan() {
		line := scanner.Text()
		switch {
		case strings.Contains(line, "content:"):
			incontent = true
			tabs = line[:strings.Index(line, "content:")]
		case incontent && strings.Contains(line, "]"):
			_, err := fmt.Fprintln(fd, tabs+tabs+`"`+path+`/**/*.{templ,go}",`)
			if err != nil {
				return fmt.Errorf("write tailwind.config.js: %w", err)
			}
			incontent = false
		case incontent && strings.Contains(line, path):
			continue
		case incontent && !strings.HasSuffix(strings.TrimSpace(line), ","):
			line += ","
		}
		_, err := fmt.Fprintln(fd, line)
		if err != nil {
			return fmt.Errorf("write tailwind.config.js: %w", err)
		}
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("rewriting: %w", err)
	}
	return nil
}
