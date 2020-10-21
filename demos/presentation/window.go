package main

import (
	"github.com/isbm/crtview"
)

const loremIpsumText = "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."

// Window returns the window page.
func Window(nextSlide func()) (title string, content crtview.Primitive) {
	wm := crtview.NewWindowManager()

	list := crtview.NewList()
	list.ShowSecondaryText(false)
	list.AddItem(crtview.NewListItem("Item #1"))
	list.AddItem(crtview.NewListItem("Item #2"))
	list.AddItem(crtview.NewListItem("Item #3"))
	list.AddItem(crtview.NewListItem("Item #4"))
	list.AddItem(crtview.NewListItem("Item #5"))
	list.AddItem(crtview.NewListItem("Item #6"))
	list.AddItem(crtview.NewListItem("Item #7"))

	loremIpsum := crtview.NewTextView()
	loremIpsum.SetText(loremIpsumText)

	w1 := crtview.NewWindow(list)
	w1.SetPosition(2, 2)
	w1.SetSize(10, 7)

	w2 := crtview.NewWindow(loremIpsum)
	w2.SetPosition(7, 4)
	w2.SetSize(12, 12)

	w1.SetTitle("List")
	w2.SetTitle("Lorem Ipsum")

	wm.Add(w1, w2)

	return "Window", wm
}
