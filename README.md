# templ-components

A library of components to be used in a Go/templ/HTMX/AlpineJS project, based on Flowbite components and the Lucide icon
library.

Note: all JS code has been rewritten to use Alpine instead of the Flowbite code.

> This is a work in progress, breaking changes will probably occur.

## Roadmap

[ ] Basic components
[ ] Styles (based on a tailwind plugin)
[ ] Storybook-like viewer

# Components

## Icons

CSS classes can be applied to the SVG icons:

```
import "github.com/jfbus/templ-components/icon"

[...]

@icon.Icon(icon.IconDefinition{Icon:icon.Flower, Class:"p-1 w-8 h-8"})
```

## Form

```
import "github.com/jfbus/templ-components/form"
```

### Input Field

Basic usage:

```
@form.InputField(form.InputFieldDefinition{
    Name: "foo",
    Label: "Foo",
    Value: [your value],    
})
```

With an icon (don't forget the import):

```
@form.InputField(form.InputFieldDefinition{
    Name: "foo",
    Label: "Foo",
    Value: [your value],
    Icon: icon.Flower,    
})
```

With HTMX attributes and a spinning loader:

```
@form.InputField(form.InputFieldDefinition{
    Name: "foo",
    Label: "Foo",
    Value: [your value],
    Loader: true,
    Attributes:templ.Attributes{
      "hx-post":"/send",
      "hx-target":"#list",
})
```
