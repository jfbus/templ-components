# A library of templ components

A library of components to be used in a Go/templ/HTMX/AlpineJS project, based on Flowbite components and the Lucide icon
library.

Note: all Flowbite JS code has been rewritten using Alpine.

> This is a work in progress, breaking changes might happen.

## Setup

Install this package :

```
go get -u github.com/jfbus/templ-components
```

Install flowbite :

```
npm install flowbite
```

No need to configure Tailwind & Flowbite, it is handled by templ-components.

In the same directory as your source js & css files, create a `tailwind.config.go` file :

```go
//go:generate tailwindconfig
package tailwind_config
```

Generate the tailwind config file:

```
go generate
```

## Update Tailwind config

`go generate` can run with an existing config;
it renames the previous config file to `tailwind.config.js.saved`.

## Roadmap

- [ ] Radio
- [ ] Checkbox
- [ ] Modal
- [ ] Only add used components in tailwind config
- [ ] Storybook-like viewer

## Sizes

Sizes vary from `size.XS` to `size.NineXL`.
Not all components support all sizes. Each component definition lists the allowed sizes.

## Components

### Icon

Basic usage:

```templ
import "github.com/jfbus/templ-components/icon"

@icon.C(icon.Flower)
```

A size can be set:

```templ
import "github.com/jfbus/templ-components/size"

@icon.C(icon.D{Icon:icon.Flower, Size:size:S})
```

Icon sizes are mapped to text sizes:

| xs | s   | normal | l    | xl | 2xl | 3xl  | 4xl | 5xl | 6xl  | 7xl  | 8xl | 9xl | full |
|----|-----|--------|------|----|-----|------|-----|-----|------|------|-----|-----|------|
| 3  | 3.5 | 4      | 18px | 5  | 6   | 30px | 9   | 12  | 60px | 72px | 24  | 32  | full |

`size.S` translates into a `w-3.5 h-3.5` class. `size.L` translates into a `w-[18px] h-[18px]` class.

### Input Field

Basic usage:

```templ
import "github.com/jfbus/templ-components/input"

@input.C(input.D{
    Name:  "foo",
    Label: "Foo",
    Value: [your value],    
})
```

With an icon:

```templ
@input.C(input.D{
    Name:  "foo",
    Label: "Foo",
    Value: [your value],
    Icon:  icon.Flower,
})
```

Icon can be placed on the right side:

```templ
@input.C(input.D{
    Name:  "foo",
    Label: "Foo",
    Value: [your value],
    Icon:  icon.Flower,
    IconPosition: position.End,
})
```

With HTMX attributes and a spinning loader:

```templ
@input.C(input.D{
    Name:   "foo",
    Label:  "Foo",
    Value:  [your value],
    Loader: true,
    Attributes:templ.Attributes{
        "hx-post":"/add",
        "hx-target":"#list",
})
```

Sizes (only `size.S`, `size.Normal` and `size.L` are available)

```templ
@input.C(input.D{
    Name:  "foo",
    Label: "Foo",
    Value: [your value],
    Size:  size.S,
})
```

### Textarea

```templ
import "github.com/jfbus/templ-components/textarea"

@textarea.C(textarea.D{
    Name: "foo",
    Label: "Foo",
    Value: [your value],
    Icon: icon.Flower,
})
```

### Inline editing

```templ
import "github.com/jfbus/templ-components/inline"

@inline.C(inline.D{
    Value: [your value],
    IconSize: size.S,
    Edit: input.C(input.D{
        Name: "title",
        Value: [your value],
        Icon: icon.CornerDownLeft,
        IconPosition: position.End,
        Size: size.S,
        Attributes: templ.Attributes{
            "hx-trigger":"keyup[key=='Enter']",
            "hx-post":"/add",
            "hx-target":"#item",
            "hx-swap":"outerHTML",
      },
    }),
})
```

### Button

Basic usage:

```templ
import "github.com/jfbus/templ-components/button"

@button.C(button.D{
    ID:  "foo",
    Label: "Foo",
})
```

With an icon:

```templ
@button.C(button.D{
    Name:  "foo",
    Label: "Foo",
    Value: [your value],
    Icon:  icon.Pencil,
})
```
