package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/isbm/crtview"
)

// Grid demonstrates the grid layout.
func Grid(nextSlide func()) (title string, content crtview.Primitive) {
	modalShown := false
	panels := crtview.NewPanels()

	newPrimitive := func(text string) crtview.Primitive {
		tv := crtview.NewTextView()
		tv.SetTextAlign(crtview.AlignCenter)
		tv.SetText(text)
		tv.SetDoneFunc(func(key tcell.Key) {
			if modalShown {
				nextSlide()
				modalShown = false
			} else {
				panels.ShowPanel("modal")
				modalShown = true
			}
		})
		return tv
	}

	menu := newPrimitive("Menu")
	main := newPrimitive("Main content")
	sideBar := newPrimitive("Side Bar")

	grid := crtview.NewGrid()
	grid.SetRows(3, 0, 3)
	grid.SetColumns(0, -4, 0)
	grid.SetBorders(true)
	grid.AddItem(newPrimitive("Header"), 0, 0, 1, 3, 0, 0, true)
	grid.AddItem(newPrimitive("Footer"), 2, 0, 1, 3, 0, 0, false)

	// Layout for screens narrower than 100 cells (menu and side bar are hidden).
	grid.AddItem(menu, 0, 0, 0, 0, 0, 0, false)
	grid.AddItem(main, 1, 0, 1, 3, 0, 0, false)
	grid.AddItem(sideBar, 0, 0, 0, 0, 0, 0, false)

	// Layout for screens wider than 100 cells.
	grid.AddItem(menu, 1, 0, 1, 1, 0, 100, false)
	grid.AddItem(main, 1, 1, 1, 1, 0, 100, false)
	grid.AddItem(sideBar, 1, 2, 1, 1, 0, 100, false)

	modal := crtview.NewModal()
	modal.SetText("Resize the window to see how the grid layout adapts")
	modal.AddButtons([]string{"Ok"})
	modal.SetDoneFunc(func(buttonIndex int, buttonLabel string) {
		panels.HidePanel("modal")
	})

	panels.AddPanel("grid", grid, true, true)
	panels.AddPanel("modal", modal, false, false)

	return "Grid", panels
}
