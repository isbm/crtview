// Demo code for the InputField primitive.
package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/isbm/crtview"
)

func main() {
	app := crtview.NewApplication()
	app.EnableMouse(true)

	inputField := crtview.NewInputField()
	inputField.SetLabel("Enter a number: ")
	inputField.SetPlaceholder("E.g. 1234")
	inputField.SetFieldWidth(10)
	inputField.SetAcceptanceFunc(crtview.InputFieldInteger)
	inputField.SetDoneFunc(func(key tcell.Key) {
		app.Stop()
	})

	app.SetRoot(inputField, true)
	if err := app.Run(); err != nil {
		panic(err)
	}
}
