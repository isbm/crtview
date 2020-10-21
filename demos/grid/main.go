// Demo code for the Grid primitive.
package main

import (
	"github.com/isbm/crtview"
)

func main() {
	app := crtview.NewApplication()
	app.EnableMouse(true)

	newPrimitive := func(text string) crtview.Primitive {
		tv := crtview.NewTextView()
		tv.SetTextAlign(crtview.AlignCenter)
		tv.SetText(text)
		return tv
	}
	menu := newPrimitive("Menu")
	main := newPrimitive("Main content")
	sideBar := newPrimitive("Side Bar")

	grid := crtview.NewGrid()
	grid.SetRows(3, 0, 3)
	grid.SetColumns(30, 0, 30)
	grid.SetBorders(true)
	grid.AddItem(newPrimitive("Header"), 0, 0, 1, 3, 0, 0, false)
	grid.AddItem(newPrimitive("Footer"), 2, 0, 1, 3, 0, 0, false)

	// Layout for screens narrower than 100 cells (menu and side bar are hidden).
	grid.AddItem(menu, 0, 0, 0, 0, 0, 0, false)
	grid.AddItem(main, 1, 0, 1, 3, 0, 0, false)
	grid.AddItem(sideBar, 0, 0, 0, 0, 0, 0, false)

	// Layout for screens wider than 100 cells.
	grid.AddItem(menu, 1, 0, 1, 1, 0, 100, false)
	grid.AddItem(main, 1, 1, 1, 1, 0, 100, false)
	grid.AddItem(sideBar, 1, 2, 1, 1, 0, 100, false)

	app.SetRoot(grid, true)
	if err := app.Run(); err != nil {
		panic(err)
	}
}
