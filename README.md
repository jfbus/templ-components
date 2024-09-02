# A library of templ components

<a href="https://pkg.go.dev/github.com/jfbus/templ-components"><img src="https://pkg.go.dev/badge/github.com/jfbus/templ-components.svg" alt="Go Reference" /></a>
<a href="https://goreportcard.com/report/github.com/jfbus/templ-components"><img src="https://goreportcard.com/badge/github.com/jfbus/templ-components" alt="Go Report Card" /></a>

A library of components to be used in a Go/templ/HTMX/AlpineJS project, based on [Flowbite](https://flowbite.com/) components and the [Lucide](https://lucide.dev/) icon
library.

_Note: all Flowbite JS code has been rewritten using Alpine._

> This is a work in progress, breaking changes might happen.

## Setup

### Start a new project

Install this package :

```
go get -u github.com/jfbus/templ-components
```

Create a new `assets_src` directory & install Flowbite:

```
npm install flowbite
```

No need to configure Tailwind & Flowbite, it is handled by templ-components.

Install [Alpine.js](https://alpinejs.dev/).

In the same `assets_src` directory, create a `tailwind.config.go` file :

```go
//go:generate tailwindconfig
package tailwind_config
```

Generate the tailwind config file:

```
go generate
```

### Add to an existing Tailwind project

Add the `tailwind.config.go` file beside your `tailwind.config.js` file.

Update the tailwind config file:

```
go generate
```

`go generate` can run with an existing config;
it renames the previous config file to `tailwind.config.js.saved`.

## Roadmap

- [ ] Checkbox
- [ ] Modal
- [ ] Toast
- [ ] Rating
- [ ] Tooltip
- [ ] Only add used components in tailwind config
- [ ] Storybook-like viewer

## Sizes

Sizes vary from `size.XS` to `size.NineXL`.
Not all components support all sizes. Each component definition lists the allowed sizes.

## Overriding CSS classes

There are usually `Class` attributes in each component, containing both `Color` and `Class` values.
`Color` contains all CSS classes defining colors (text, background, color).
`Class` contains all other CSS classes (border, padding, ...).

You change override default CSS classes either:
* by setting the relevant attribute in each component,
* by changing the default values used by every component (usually `packagename.Defaults`).

## Components

### Icon

Basic usage:

```templ
import "github.com/jfbus/templ-components/components/icon"

@icon.C(icon.Flower)
```

A size can be set:

```templ
import "github.com/jfbus/templ-components/components/size"

@icon.C(icon.D{Icon:icon.Flower, Size:size:S})
```

Icon sizes are mapped to text sizes:

| xs | s   | normal | l    | xl | 2xl | 3xl  | 4xl | 5xl | 6xl  | 7xl  | 8xl | 9xl | full |
|----|-----|--------|------|----|-----|------|-----|-----|------|------|-----|-----|------|
| 3  | 3.5 | 4      | 18px | 5  | 6   | 30px | 9   | 12  | 60px | 72px | 24  | 32  | full |

`size.S` translates into a `w-3.5 h-3.5` class. `size.L` translates into a `w-[18px] h-[18px]` class.

### Input Field

```templ
@input.C(input.D{
    Name:  "foo",
    Label: "Foo",
    Value: [your value],
    Size:  size.S,
    Icon:  icon.Flower,
})
```

With HTMX attributes and a spinning loader:

```templ
@input.C(input.D{
    Name:       "foo",
    Label:      "Foo",
    Value:      [your value],
    Loader:     true,
    Attributes: templ.Attributes{
        "hx-post":   "/add",
        "hx-target": "#list",
})
```

You can define a custom loader by changing `loader.DefaultLoader`. 

### Textarea

```templ
import "github.com/jfbus/templ-components/components/textarea"

@textarea.C(textarea.D{
    Name:  "foo",
    Label: "Foo",
    Value: [your value],
    Icon:  icon.Flower,
})
```

### Select

```templ
import (
	"github.com/jfbus/templ-components/components/selectfield"
	"github.com/jfbus/templ-components/components/selectfield/option"
)

@selectfield.C(selectfield.D{
    Name:  "country",
    Label: "Country",
    Options: []option.D{{
        Label: "Select a country",
    }, {
        Value: "FR",
        Label: "France",
    }, {
        Value: "DE",
        Label: "Germany",
    }, {
        Value: "GB",
        Label: "United Kingdom",
    }},
    Selected: "DE",
})
```

### Radio/Radiogroup

```templ
import (
  "github.com/jfbus/templ-components/components/radio"
  "github.com/jfbus/templ-components/components/radiogroup"
)

@radiogroup.C(radiogroup.D{
  Name: "choice",
  Style: radiogroup.StyleBordered,
  Inputs: []radio.D{{
    Value: "choice1",
    Label: "Choice 1",
  }, {
    Value: "choice2",
    Label: "Choice 2",
  }},
})
```

### Inline editing

```templ
import "github.com/jfbus/templ-components/components/inline"

@inline.C(inline.D{
    Value:    [your value],
    IconSize: size.S,
    Edit:     input.C(input.D{
        Name:         "title",
        Value:        [your value],
        Icon:         icon.CornerDownLeft,
        IconPosition: position.End,
        Size:         size.S,
        Attributes:   templ.Attributes{
            "hx-trigger": "keyup[key=='Enter']",
            "hx-post":    "/add",
            "hx-target":  "#item",
            "hx-swap":    "outerHTML",
      },
    }),
})
```

### Button

```templ
import "github.com/jfbus/templ-components/components/button"

@button.C(button.D{
    Name:  "foo",
    Label: "Foo",
    Value: [your value],
    Icon:  icon.Pencil,
})
```

### Button group

```templ
import "github.com/jfbus/templ-components/components/buttongroup"

@buttongroup.C(buttongroup.D{
    Size:    size.S,
    Buttons: []button.D{
        {
            Icon:  icon.ArrowDownNarrowWide,
            Label: label.D{
                  Label: "Sort",
                  Hide: true,
            },
        },
        {
            Icon:      icon.Heart,
            Label: label.D{
                  Label: "Rating",
                  Hide: true,
            },
        },
        {
            Icon:      icon.Banknote,
            Label: label.D{
                  Label: "Price",
                  Hide: true,
            },
        },
    },
})
```

### Table

```templ
import (
    "github.com/jfbus/templ-components/components/table"
    "github.com/jfbus/templ-components/components/table/row"
    "github.com/jfbus/templ-components/components/table/cell"
)

@table.C(table.D{
    Style:  table.StyleStripedRows,
    Header: &row.D{
        Cells: []string{"Email","Name","Status", ""},
    },
    Rows: []row.D{{
        Cells: []any{
            "John Doe",
            "john.doe@example.com",
            "active",
            cell.D{
                Class:   style.D{Class:"text-center"},
                Content: button.C(button.D{
                    Label: "disable",
                }),
            },
        },
    }},
})
```

Row contents can either be a slice of strings, a slice of `cell.D` definitions,
a slice of `templ.Component` components or a `[]any` slice containing any number of these.

## Helpers

```
import "github.com/jfbus/templ-components/components/helper"
```

### S

`helper.S` renders anything (numbers, booleans, ...). `{helper.S(1)}` is the equivalent of `{strconv.Itoa(1)}`

### IfEmpty

Renders the first non empty string from a list of string parameters.

```go
helper.IfEmpty(item.Value, "???")
```

### L

`L` returns a `templ.Component` based on a list of `templ.Component` values.

## FAQ

### I use a new component and it looks broken !

Run `go generate` again, and update your Tailwind class (`npx tailwindcss -i [...] -o [...]`).

Check that your `tailwind.config.js` content section contains :
* your templates (something like `"../views/**/*.{templ,go}"`)
* templ-components (something like `"[your local path]/github.com/jfbus/templ-components/**/*.{templ,go}"`)