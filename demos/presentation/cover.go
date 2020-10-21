package main

import (
	"fmt"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/isbm/crtview"
)

const logo = `
 ======= ===  === === ======== ===  ===  ===
===      ===  === === ===      ===  ===  ===
===      ===  === === ======   ===  ===  ===
===       ======  === ===       ===========
 =======    ==    === ========   ==== ====
`

const (
	subtitle   = `Terminal-based user interface toolkit`
	mouse      = `Navigate with your keyboard or mouse.`
	navigation = `Next slide: Ctrl-N   Previous: Ctrl-P   Exit: Ctrl-C`
)

// Cover returns the cover page.
func Cover(nextSlide func()) (title string, content crtview.Primitive) {
	// What's the size of the logo?
	lines := strings.Split(logo, "\n")
	logoWidth := 0
	logoHeight := len(lines)
	for _, line := range lines {
		if len(line) > logoWidth {
			logoWidth = len(line)
		}
	}
	logoBox := crtview.NewTextView()
	logoBox.SetTextColor(tcell.ColorGreen.TrueColor())
	logoBox.SetDoneFunc(func(key tcell.Key) {
		nextSlide()
	})
	fmt.Fprint(logoBox, logo)

	// Create a frame for the subtitle and navigation infos.
	frame := crtview.NewFrame(crtview.NewBox())
	frame.SetBorders(0, 0, 0, 0, 0, 0)
	frame.AddText(subtitle, true, crtview.AlignCenter, tcell.ColorWhite.TrueColor())
	frame.AddText("", true, crtview.AlignCenter, tcell.ColorWhite.TrueColor())
	frame.AddText(mouse, true, crtview.AlignCenter, tcell.ColorDarkMagenta.TrueColor())
	frame.AddText(navigation, true, crtview.AlignCenter, tcell.ColorDarkMagenta.TrueColor())

	// Create a Flex layout that centers the logo and subtitle.
	subFlex := crtview.NewFlex()
	subFlex.AddItem(crtview.NewBox(), 0, 1, false)
	subFlex.AddItem(logoBox, logoWidth, 1, true)
	subFlex.AddItem(crtview.NewBox(), 0, 1, false)

	flex := crtview.NewFlex()
	flex.SetDirection(crtview.FlexRow)
	flex.AddItem(crtview.NewBox(), 0, 7, false)
	flex.AddItem(subFlex, logoHeight, 1, true)
	flex.AddItem(frame, 0, 10, false)

	return "Start", flex
}
