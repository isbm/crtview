package crtwin

import (
	"github.com/gdamore/tcell/v2"
	"github.com/isbm/crtview"
)

type DialogWindow struct {
	// Foreground (text stuff)
	fgcolor tcell.Color

	// Color of all the symbols on which shadow is falling. Default is RGB 0x666666.
	shFgColor tcell.Color

	// Color of the shadow. Default is black.
	shBgColor tcell.Color

	shadow   bool
	centered bool

	*crtview.Form
}

func NewDialogWindow() *DialogWindow {
	return (&DialogWindow{
		Form: crtview.NewForm(),
	}).init()
}

func (tmd *DialogWindow) init() *DialogWindow {
	tmd.SetBorder(true)
	tmd.SetShadow(true)
	tmd.SetTitle("")
	tmd.SetShadowColor(tcell.NewRGBColor(0x66, 0x66, 0x66), tcell.ColorBlack)

	return tmd
}

// SetShadowColor sets foreground and background colors of the shadow.
func (tmd *DialogWindow) SetShadowColor(foreground, background tcell.Color) *DialogWindow {
	tmd.shBgColor, tmd.shFgColor = background, foreground
	return tmd
}

// SetShadow enables or disables shadow painting of the dialog
func (tmd *DialogWindow) SetShadow(shadow bool) *DialogWindow {
	tmd.shadow = shadow
	return tmd
}

// SetSize of the window
func (tmw *DialogWindow) SetSize(w, h int) *DialogWindow {
	x, y, _, _ := tmw.GetRect()
	tmw.SetRect(x, y, w, h)
	return tmw
}

// SetPosition of the window
func (tmw *DialogWindow) SetPosition(x, y int) *DialogWindow {
	_, _, w, h := tmw.GetRect()
	tmw.SetRect(x, y, w, h)
	return tmw
}

// SetCentered window (or not)
func (tmw *DialogWindow) SetCentered(c bool) *DialogWindow {
	tmw.centered = c
	return tmw
}

// Draw the dialog window
func (tmw *DialogWindow) Draw(screen tcell.Screen) {
	// Bailout if we are hidden
	if !tmw.IsVisible() {
		return
	}

	// Draw content
	x, y, w, h := tmw.GetRect()
	if tmw.centered {
		sw, sh := screen.Size()
		x = (sw / 2) - (w / 2)
		y = (sh / 2) - (h / 2)
	}

	tmw.Form.SetRect(x+1, y+1, w-2, h-2)
	tmw.Form.Draw(screen)

	tmw.Lock()
	defer tmw.Unlock()

	if tmw.shadow {
		// Shadow
		shadowStyle := tcell.StyleDefault.Background(tmw.shBgColor).Foreground(tmw.shFgColor)
		x, y, w, h := tmw.GetRect()
		for i := 0; i < w; i++ {
			c, _, _, _ := screen.GetContent(x+i+1, y+h)
			screen.SetContent(x+i+1, y+h, c, nil, shadowStyle)
		}

		for i := 0; i < (h - 1); i++ {
			c, _, _, _ := screen.GetContent(x+w, y+i+1)
			screen.SetContent(x+w, y+i+1, c, nil, shadowStyle)
		}
	}

	tmw.SetRect(x, y, w, h)
}
