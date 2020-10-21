// Demo code for the DropDown primitive.
package main

import "github.com/isbm/crtview"

func main() {
	app := crtview.NewApplication()
	app.EnableMouse(true)

	dropdown := crtview.NewDropDown()
	dropdown.SetLabel("Select an option (hit Enter): ")
	dropdown.SetOptions(nil,
		crtview.NewDropDownOption("First"),
		crtview.NewDropDownOption("Second"),
		crtview.NewDropDownOption("Third"),
		crtview.NewDropDownOption("Fourth"),
		crtview.NewDropDownOption("Fifth"))

	app.SetRoot(dropdown, true)
	if err := app.Run(); err != nil {
		panic(err)
	}
}
