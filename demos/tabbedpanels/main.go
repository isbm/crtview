// Demo code for the TabbedPanels primitive.
package main

import (
	"fmt"

	"github.com/isbm/crtview"
)

const panelCount = 5

func main() {
	app := crtview.NewApplication()
	app.EnableMouse(true)

	panels := crtview.NewTabbedPanels()
	for panel := 0; panel < panelCount; panel++ {
		func(panel int) {
			form := crtview.NewForm()
			form.SetBorder(true)
			form.SetTitle(fmt.Sprintf("This is tab %d. Choose another tab.", panel+1))
			form.AddButton("Next", func() {
				panels.SetCurrentTab(fmt.Sprintf("panel-%d", (panel+1)%panelCount))
			})
			form.AddButton("Quit", func() {
				app.Stop()
			})
			form.SetCancelFunc(func() {
				app.Stop()
			})

			panels.AddTab(fmt.Sprintf("panel-%d", panel), fmt.Sprintf("Panel #%d", panel), form)
		}(panel)
	}

	app.SetRoot(panels, true)
	if err := app.Run(); err != nil {
		panic(err)
	}
}
