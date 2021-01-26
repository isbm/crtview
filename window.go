package crtview

import (
	"sync"

	"github.com/gdamore/tcell/v2"
)

// Window is a draggable, resizable frame around a primitive. Windows must be
// added to a WindowManager.
type Window struct {
	*Box

	primitive      Primitive
	x, y           int
	width, height  int
	fullscreen     bool
	centered       bool
	statusbar      string
	statusbarAlign int
	statusbarColor tcell.Color

	dragX, dragY   int
	dragWX, dragWY int
	marginTop      int
	marginRight    int
	marginBottom   int
	marginLeft     int

	sync.RWMutex
}

// NewWindow returns a new window around the given primitive.
func NewWindow(primitive Primitive) *Window {
	w := new(Window)
	w.Box = NewBox()
	w.primitive = primitive
	w.dragWX, w.dragWY = -1, -1
	w.Box.focus = w
	w.statusbarAlign = AlignRight

	return w
}

// SetBackgroundColor of the window
func (w *Window) SetBackgroundColor(color tcell.Color) *Window {
	w.Box.SetBackgroundColor(color)
	return w
}

// SetStatus widget to be displayed on the window
func (w *Window) SetStatus(text string) *Window {
	w.Lock()
	defer w.Unlock()

	w.statusbar = text
	return w
}

// GetStatus widget
func (w *Window) GetStatus() string {
	return w.statusbar
}

// SetStatusBarAlign on the window's footer (left/centered/right). Default: right
func (w *Window) SetStatusBarAlign(align int) *Window {
	w.Lock()
	defer w.Unlock()

	w.statusbarAlign = align
	return w
}

// SetStatusBarColor for the text. However, in case the string is dynamic, this is overridden.
func (w *Window) SetStatusBarColor(color tcell.Color) *Window {
	w.Lock()
	defer w.Unlock()

	w.statusbarColor = color
	return w
}

// GetStatusBarColor for the text.
func (w *Window) GetStatusBarColor() tcell.Color {
	return w.statusbarColor
}

// GetStatusBarAlign current alignment
func (w *Window) GetStatusBarAlign() int {
	return w.statusbarAlign
}

// SetPosition sets the position of the window.
func (w *Window) SetPosition(x, y int) *Window {
	w.Lock()
	defer w.Unlock()

	w.x, w.y = x, y
	w.centered = false
	return w
}

// SetSize sets the size of the window.
func (w *Window) SetSize(width, height int) *Window {
	w.Lock()
	defer w.Unlock()

	w.width, w.height = width, height
	return w
}

// SetFullscreen sets the flag indicating whether or not the the window should
// be drawn fullscreen.
func (w *Window) SetFullscreen(fullscreen bool) *Window {
	w.Lock()
	defer w.Unlock()

	w.fullscreen = fullscreen
	return w
}

// SetMarginBorder creates a space around window, outside of any defined borders.
// This works in full-screen mode too.
func (w *Window) SetMarginBorder(top int, right int, bottom int, left int) *Window {
	w.marginTop, w.marginRight, w.marginBottom, w.marginLeft = top, right, bottom, left
	return w
}

// GetMarginBorder returns top/right/bottom/left space.
// This method is mainly used by the WindowManager to properly resize the window on final render.
func (w *Window) GetMarginBorder() (int, int, int, int) {
	return w.marginTop, w.marginRight, w.marginBottom, w.marginLeft
}

// SetPositionCenter sets the flag to the Window Manager that the current window should be displayed centered.
// If SetPosition is called, this flag is reset to false.
func (w *Window) SetPositionCenter() *Window {
	w.centered = true
	return w
}

// Show window
func (w *Window) Show() *Window {
	w.Box.Show()
	return w
}

// Hide window
func (w *Window) Hide() *Window {
	w.Box.Hide()
	return w
}

// GetSize gets the size of the window
func (w *Window) GetSize() (int, int) {
	return w.width, w.height
}

// IsCentered returns true if window is meant to be displayed center
func (w *Window) IsCentered() bool {
	return w.centered
}

// Focus is called when this primitive receives focus.
func (w *Window) Focus(delegate func(p Primitive)) {
	w.Lock()
	defer w.Unlock()

	w.Box.Focus(delegate)
	w.primitive.Focus(delegate)
}

// Blur is called when this primitive loses focus.
func (w *Window) Blur() {
	w.Lock()
	defer w.Unlock()

	w.Box.Blur()
	w.primitive.Blur()
}

// HasFocus returns whether or not this primitive has focus.
func (w *Window) HasFocus() bool {
	w.RLock()
	defer w.RUnlock()

	focusable := w.primitive.GetFocusable()
	if focusable != nil {
		return focusable.HasFocus()
	}

	return w.Box.HasFocus()
}

// Draw draws this primitive onto the screen.
func (w *Window) Draw(screen tcell.Screen) {
	if !w.IsVisible() {
		return
	}

	w.RLock()
	defer w.RUnlock()

	w.Box.Draw(screen)

	x, y, width, height := w.GetInnerRect()
	w.primitive.SetRect(x, y, width, height)
	w.primitive.Draw(screen)
}

// InputHandler returns the handler for this primitive.
func (w *Window) InputHandler() func(event *tcell.EventKey, setFocus func(p Primitive)) {
	return w.primitive.InputHandler()
}

// MouseHandler returns the mouse handler for this primitive.
func (w *Window) MouseHandler() func(action MouseAction, event *tcell.EventMouse, setFocus func(p Primitive)) (consumed bool, capture Primitive) {
	return w.WrapMouseHandler(func(action MouseAction, event *tcell.EventMouse, setFocus func(p Primitive)) (consumed bool, capture Primitive) {
		if !w.InRect(event.Position()) {
			return false, nil
		}

		if action == MouseLeftDown || action == MouseMiddleDown || action == MouseRightDown {
			setFocus(w)
		}

		if action == MouseLeftDown {
			x, y, width, height := w.GetRect()
			mouseX, mouseY := event.Position()

			leftEdge := mouseX == x
			rightEdge := mouseX == x+width-1
			bottomEdge := mouseY == y+height-1
			topEdge := mouseY == y

			if mouseY >= y && mouseY <= y+height-1 {
				if leftEdge {
					w.dragX = -1
				} else if rightEdge {
					w.dragX = 1
				}
			}

			if mouseX >= x && mouseX <= x+width-1 {
				if bottomEdge {
					w.dragY = -1
				} else if topEdge {
					if leftEdge || rightEdge {
						w.dragY = 1
					} else {
						w.dragWX = mouseX - x
						w.dragWY = mouseY - y
					}
				}
			}
		}

		_, capture = w.primitive.MouseHandler()(action, event, setFocus)
		return true, capture
	})
}
