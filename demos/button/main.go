// Demo code for the Button primitive.
package main

import "github.com/isbm/crtview"

func main() {
	app := crtview.NewApplication()
	app.EnableMouse(true)

	button := crtview.NewButton("Hit Enter to close")
	button.SetBorder(true)
	button.SetRect(0, 0, 22, 3)
	button.SetSelectedFunc(func() {
		app.Stop()
	})

	app.SetRoot(button, false)
	if err := app.Run(); err != nil {
		panic(err)
	}
}
