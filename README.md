# crtview - Terminal-based user interface toolkit

This project is a fork of [cview](https://gitlab.com/tslocum/cview),
which is also fork of original
[tview](https://github.com/rivo/tview). :-) See
``docs/about-crtview.md`` for more details.

## Features

Available widgets:

- __Input forms__ (including __input/password fields__, __drop-down selections__, __checkboxes__, and __buttons__)
- Navigable multi-color __text views__
- Selectable __lists__ with __context menus__
- Modal __dialogs__
- Horizontal and vertical __progress bars__
- __Grid__, __Flexbox__ and __tabbed panel layouts__
- Sophisticated navigable __table views__
- Flexible __tree views__
- Draggable and resizable __windows__
- An __application__ wrapper

Widgets may be customized and extended to suit any application.

## Installation

```bash
go get github.com/isbm/crtview
```

## Hello World

This basic example creates a TextView titled "Hello, World!" and displays it in your terminal:

```go
package main

import (
	"github.com/isbm/crtview"
)

func main() {
	app := crtview.NewApplication()
	
	box := crtview.NewTextView()
		.SetBorder(true)
		.SetTitle("Hello, world!")
		.SetText("Here is some meaning-less text for your app.")
	
	app.SetRoot(box, true)
	if err := app.Run(); err != nil {
		panic(err)
	}
}
```
