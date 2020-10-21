// Demo code for the CheckBox primitive.
package main

import (
	"github.com/isbm/crtview"
)

func main() {
	app := crtview.NewApplication()
	app.EnableMouse(true)

	checkbox := crtview.NewCheckBox()
	checkbox.SetLabel("Hit Enter to check box: ")

	app.SetRoot(checkbox, true)
	if err := app.Run(); err != nil {
		panic(err)
	}
}
