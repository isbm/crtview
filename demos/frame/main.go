// Demo code for the Frame primitive.
package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/isbm/crtview"
)

func main() {
	app := crtview.NewApplication()
	app.EnableMouse(true)

	box := crtview.NewBox()
	box.SetBackgroundColor(tcell.ColorBlue.TrueColor())

	frame := crtview.NewFrame(box)
	frame.SetBorders(2, 2, 2, 2, 4, 4)
	frame.AddText("Header left", true, crtview.AlignLeft, tcell.ColorWhite.TrueColor())
	frame.AddText("Header middle", true, crtview.AlignCenter, tcell.ColorWhite.TrueColor())
	frame.AddText("Header right", true, crtview.AlignRight, tcell.ColorWhite.TrueColor())
	frame.AddText("Header second middle", true, crtview.AlignCenter, tcell.ColorRed.TrueColor())
	frame.AddText("Footer middle", false, crtview.AlignCenter, tcell.ColorGreen.TrueColor())
	frame.AddText("Footer second middle", false, crtview.AlignCenter, tcell.ColorGreen.TrueColor())

	app.SetRoot(frame, true)
	if err := app.Run(); err != nil {
		panic(err)
	}
}
