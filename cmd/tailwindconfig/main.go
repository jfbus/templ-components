package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
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
	cur, err := os.Getwd()
	if err == nil {
		p, err = filepath.Rel(cur, filepath.Dir(p))
		if err == nil {
			err = updateConfig(p)
			if errors.Is(err, os.ErrNotExist) {
				err = createConfig(p, defaultConfig)
			}
		}
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err) //nolint:errcheck
		os.Exit(1)
	}
}

func updateConfig(path string) error {
	content, err := os.ReadFile("tailwind.config.js")
	if err != nil {
		return fmt.Errorf("read tailwind.config.js: %w", err)
	}
	err = os.Rename("tailwind.config.js", "tailwind.config.js.saved")
	if err != nil {
		return fmt.Errorf("rename tailwind.config.js: %w", err)
	}
	return createConfig(path, string(content))
}

func createConfig(path string, content string) (rerr error) {
	fd, err := os.Create("tailwind.config.js")
	if err != nil {
		return fmt.Errorf("create tailwind.config.js: %w", err)
	}
	defer func() {
		err := fd.Close()
		if rerr == nil {
			rerr = err
		}
	}()
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
				rerr = fmt.Errorf("write tailwind.config.js: %w", err)
				return
			}
			incontent = false
		case incontent && strings.Contains(line, path):
			continue
		case incontent && !strings.HasSuffix(strings.TrimSpace(line), ","):
			line += ","
		}
		_, err := fmt.Fprintln(fd, line)
		if err != nil {
			rerr = fmt.Errorf("write tailwind.config.js: %w", err)
			return
		}
	}
	if err := scanner.Err(); err != nil {
		rerr = fmt.Errorf("rewriting: %w", err)
		return
	}
	return nil
}
