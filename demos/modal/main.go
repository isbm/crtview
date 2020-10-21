// Demo code for the Modal primitive.
package main

import (
	"github.com/isbm/crtview"
)

func main() {
	app := crtview.NewApplication()
	app.EnableMouse(true)

	modal := crtview.NewModal()
	modal.SetText("Do you want to quit the application?")
	modal.AddButtons([]string{"Quit", "Cancel"})
	modal.SetDoneFunc(func(buttonIndex int, buttonLabel string) {
		if buttonLabel == "Quit" {
			app.Stop()
		}
	})

	app.SetRoot(modal, false)
	if err := app.Run(); err != nil {
		panic(err)
	}
}
