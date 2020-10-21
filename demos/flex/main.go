// Demo code for the Flex primitive.
package main

import (
	"github.com/isbm/crtview"
)

func demoBox(title string) *crtview.Box {
	b := crtview.NewBox()
	b.SetBorder(true)
	b.SetTitle(title)
	return b
}

func main() {
	app := crtview.NewApplication()
	app.EnableMouse(true)

	subFlex := crtview.NewFlex()
	subFlex.SetDirection(crtview.FlexRow)
	subFlex.AddItem(demoBox("Top"), 0, 1, false)
	subFlex.AddItem(demoBox("Middle (3 x height of Top)"), 0, 3, false)
	subFlex.AddItem(demoBox("Bottom (5 rows)"), 5, 1, false)

	flex := crtview.NewFlex()
	flex.AddItem(demoBox("Left (1/2 x width of Top)"), 0, 1, false)
	flex.AddItem(subFlex, 0, 2, false)
	flex.AddItem(demoBox("Right (20 cols)"), 20, 1, false)

	app.SetRoot(flex, true)
	if err := app.Run(); err != nil {
		panic(err)
	}
}
