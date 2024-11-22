# TemplUI - A library of templ UI components

<a href="https://pkg.go.dev/github.com/jfbus/templui"><img src="https://pkg.go.dev/badge/github.com/jfbus/templui.svg" alt="Go Reference" /></a>
<a href="https://goreportcard.com/report/github.com/jfbus/templui"><img src="https://goreportcard.com/badge/github.com/jfbus/templui" alt="Go Report Card" /></a>

A library of components to be used in a Go/templ/HTMX/Alpine.js project, loosely based on [Flowbite](https://flowbite.com/) components and the [Lucide](https://lucide.dev/) icon
library.

_Note: all Flowbite JS code has been rewritten using Alpine or pure CSS._

> This is a work in progress, breaking changes might happen.

Go support: 1.23+

## Setup

### Start a new project

Install this package :

```
go get -u github.com/jfbus/templui
```

Create a new `assets_src` directory & install Flowbite:

```
npm install flowbite
```

No need to configure Tailwind & Flowbite, it is handled by templui.

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

- [ ] Rating
- [ ] Only add used components in tailwind config
- [ ] Storybook-like viewer

## Sizes

Sizes vary from `size.XS` to `size.NineXL`.
Not all components support all sizes. Each component definition lists the allowed sizes.

## Overriding CSS classes

Default CSS classes are defined by a `style.Defaults` global variable.
It can be changed after init if you need to globally change the style.

Each component has one or multiple html tags, each having its CSS classes (e.g. "button" and "button/label").

Most components have a `CustomStyle` attribute that can be used to modify those default classes.

```
    CustomStyle: style.Custom{
      "button": style.D{style.Add("text-sm")},
	},
```

You may use :
* `style.Add` to add some CSS classes,
* `style.Set` to replace all CSS classes,
* `style.Replace` to replace some CSS classes by new ones,
* `style.ReplaceVariants` to replace all variants from a CSS class,
* `style.Remove` to remove some CSS classes.

As a shortcut to modify the CSS class of the top-level tag, you can use :

```
    CustomStyle: style.CustomAdd("text-sm"),
```

## Configuring a skin

You can configure a skin using `style.SetSkin`.

You may use the default skin by calling `style.SetSkin(skin.Default)` or copy `skin.Default` and build your own style.

## Using TemplUI

### Base layout

Replace your `<body>` tag with `body.C`, and add a navbar, a sidebar and a footer.

### Icons

Icons may be used by name:

```templ
import "github.com/jfbus/templui/components/icon"

@icon.C(icon.Flower)
```

or using a definition:

```templ
import "github.com/jfbus/templui/components/icon"

@icon.C(icon.D{
  Icon:icon.Flower,
})
```

Icon sizes are mapped to text sizes:

| xs | s   | normal | l    | xl | 2xl | 3xl  | 4xl | 5xl | 6xl  | 7xl  | 8xl | 9xl | full |
|----|-----|--------|------|----|-----|------|-----|-----|------|------|-----|-----|------|
| 3  | 3.5 | 4      | 18px | 5  | 6   | 30px | 9   | 12  | 60px | 72px | 24  | 32  | full |

`size.S` translates into a `w-3.5 h-3.5` class. `size.L` translates into a `w-[18px] h-[18px]` class.

### Forms

Add a `form.C` component and various input fields (input, textarea, select, radio/radiogroup, checkbox/checkboxgroup).

#### Validation errors

Add a `Message` container to your inputs: 

```templ
@input.C(input.D{
    Name:  "foo",
    Label: "Foo",
    Value: [your value],
    Message: &message.D{},
})
```

And return errors from your code:

```go
return validation.Retarget(ctx, validation.D{
		FormID: formID,
		Errors: map[string]string{
			"foo": "field is required",
		},
	}, w, http.StatusUnprocessableEntity)
```

Validation errors require Alpine.JS.

### Inline editing

```templ
import "github.com/jfbus/templui/components/inline"

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

Inline editing require Alpine.JS.

### Table

```templ
import (
    "github.com/jfbus/templui/components/table"
    "github.com/jfbus/templui/components/table/row"
    "github.com/jfbus/templui/components/table/cell"
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

### Toasts

You may return both your normal handler output and a toast:

```go
templ.Join(
	YourComponent(),
	&toast.C(toast.D{
		ContainerID: container.DefaultID,
		Style:       toast.StyleOK,
		Content:     "Success !",
    }),
)
```

If you wish to replace the normal handler output by an error toast:

```go
toast.Retarget(ctx, toast.D{
    Style:       toast.StyleError,
    Content:     "An error occured",
    Close:       toast.CloseButton,
}, r, http.StatusInternalServerError)
```

By default, toasts automatically close, use `Close: toast.CloseButton` to add a manually closing toast.

Toasts require Alpine.JS.

### Modals

A handler may display a modal:

```go
modal.Retarget(ctx, modal.D{}, w, http.StatusOK)
```

Display a component and close a modal:

```go
templ.Join(
	YourComponent(),
	modal.Close(modalID),
)
```

Modals require Alpine.JS.

### Tooltips

You may add tooltip to inputs, labels, buttons:

```templ
button.C(button.D{
  ToolTip: &tooltip.D{
    Text: "tooltip",
  },
})
```

## Helpers

```
import "github.com/jfbus/templui/components/helper"
```

### S

`helper.S` renders anything (numbers, booleans, ...). `{helper.S(1)}` is the equivalent of `{strconv.Itoa(1)}`

### IfEmpty

Renders the first non empty string from a list of string parameters.

```go
helper.IfEmpty(item.Value, "???")
```

## FAQ

### I use a new component and it looks broken !

Run `go generate` again, and update your Tailwind class (`npx tailwindcss -i [...] -o [...]`).

Check that your `tailwind.config.js` content section contains :
* your templates (something like `"../views/**/*.{templ,go}"`)
* templui (something like `"[your local path]/github.com/jfbus/templui/**/*.{templ,go}"`)