# A library of templ components

A library of components to be used in a Go/templ/HTMX/AlpineJS project, based on Flowbite components and the Lucide icon
library.

Note: all Flowbite JS code has been rewritten using Alpine.

> This is a work in progress, breaking changes will probably happen.

## Setup

Install this package :

```
go get -u github.com/jfbus/templ-components
```

Install flowbite :

```
npm install flowbite
```

No need to configure Flowbite, it is handled by templ-components.

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

- [ ] Button
- [ ] Radio
- [ ] Checkbox
- [ ] Modal
- [ ] Only add used components in tailwind config
- [ ] Storybook-like viewer

## Sizes

Sizes vary from `size.XS` to `size.NineXL`.

## Components

### Icons

```
import "github.com/jfbus/templ-components/icon"
```

Basic usage:

```
@icon.Icon(icon.Flower)
```

A size can be set:

```
import "github.com/jfbus/templ-components/size"

@icon.Icon(icon.IconDefinition{Icon:icon.Flower, Size:size:S})
```

Icon sizes are mapped to text sizes:

| xs | s   | normal | l    | xl | 2xl | 3xl  | 4xl | 5xl | 6xl  | 7xl  | 8xl | 9xl | full |
|----|-----|--------|------|----|-----|------|-----|-----|------|------|-----|-----|------|
| 3  | 3.5 | 4      | 18px | 5  | 6   | 30px | 9   | 12  | 60px | 72px | 24  | 32  | full |

`size.S` translates into a `w-3.5 h-3.5` class. `size.L` translates into a `w-[18px] h-[18px]` class.

### Form

```
import "github.com/jfbus/templ-components/form"
```

#### Input Field

Basic usage:

```
@form.InputField(form.InputFieldDefinition{
    Name:  "foo",
    Label: "Foo",
    Value: [your value],    
})
```

With an icon (don't forget the icon import):

```
import "github.com/jfbus/templ-components/icon"

@form.InputField(form.InputFieldDefinition{
    Name:  "foo",
    Label: "Foo",
    Value: [your value],
    Icon:  icon.Flower,
})
```

Icon can be placed on the right side:

```
import (
    "github.com/jfbus/templ-components/icon"
    "github.com/jfbus/templ-components/position"
)

@form.InputField(form.InputFieldDefinition{
    Name:  "foo",
    Label: "Foo",
    Value: [your value],
    Icon:  icon.Flower,
    IconPosition: position.End,
})
```

With HTMX attributes and a spinning loader:

```
@form.InputField(form.InputFieldDefinition{
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

```
import (
    "github.com/jfbus/templ-components/size"
)

@form.InputField(form.InputFieldDefinition{
    Name:  "foo",
    Label: "Foo",
    Value: [your value],
    Size:  size.S,
})
```

### Textarea

```
@form.Textarea(form.TextareaDefinition{
    Name: "foo",
    Label: "Foo",
    Value: [your value],
    Icon: icon.Flower,
})
```

### Inline editing

```
@form.InlineEdit(form.InlineEditDefinition{
    Value: [your value],
    IconSize: size.S,
    Edit: form.InputField(form.InputFieldDefinition{
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