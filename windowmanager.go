package crtview

import (
	"sync"

	"github.com/gdamore/tcell/v2"
)

// WindowManager provides an area which windows may be added to.
type WindowManager struct {
	windows    []*Window
	fullScreen bool

	sync.RWMutex
	*Box
}

// NewWindowManager returns a new window manager.
func NewWindowManager() *WindowManager {
	wm := new(WindowManager)
	wm.Box = NewBox()

	return wm
}

// Add adds a window to the manager.
func (wm *WindowManager) Add(w ...*Window) *WindowManager {
	wm.Lock()
	defer wm.Unlock()

	for _, window := range w {
		window.SetBorder(true)
	}

	wm.windows = append(wm.windows, w...)
	return wm
}

// Clear removes all windows from the manager.
func (wm *WindowManager) Clear() *WindowManager {
	wm.Lock()
	defer wm.Unlock()

	wm.windows = nil
	return wm
}

// SetFullScreen sets the window manager to a full screen.
// The size then is inherited from the screen directly.
func (wm *WindowManager) SetFullScreen(state bool) *WindowManager {
	wm.Lock()
	defer wm.Unlock()

	wm.fullScreen = state
	return wm
}

// IsFullScreen returns true if the window manager is set to be full screen
func (wm *WindowManager) IsFullScreen() bool {
	return wm.fullScreen
}

// Focus is called when this primitive receives focus.
func (wm *WindowManager) Focus(delegate func(p Primitive)) {
	wm.Lock()
	defer wm.Unlock()

	if len(wm.windows) == 0 {
		return
	}

	wm.windows[len(wm.windows)-1].Focus(delegate)
}

// HasFocus returns whether or not this primitive has focus.
func (wm *WindowManager) HasFocus() bool {
	wm.RLock()
	defer wm.RUnlock()

	for _, w := range wm.windows {
		if w.HasFocus() {
			return true
		}
	}

	return false
}

// Draw draws this primitive onto the screen.
func (wm *WindowManager) Draw(screen tcell.Screen) {
	if !wm.IsVisible() {
		return
	}

	wm.RLock()
	defer wm.RUnlock()

	var width, height, x, y int
	if wm.IsFullScreen() {
		width, height = screen.Size()
		x, y = 1, 0
		width--
		height--
	} else {
		x, y, width, height = wm.GetInnerRect()
	}

	var hasFullScreen bool
	for _, w := range wm.windows {
		if !w.fullscreen || !w.IsVisible() {
			continue
		}

		marginTop, marginRight, marginBottom, marginLeft := w.GetMarginBorder()
		hasFullScreen = true
		w.SetRect(x-1+marginLeft, y+marginTop, width+2-marginRight-marginLeft, height+1-marginBottom-marginTop)

		w.Draw(screen)
	}
	if hasFullScreen {
		return
	}

	for _, w := range wm.windows {
		if !w.IsVisible() {
			continue
		}

		if w.IsCentered() {
			sw, sh := screen.Size()
			ww, wh := w.GetSize()
			w.x, w.y = sw/2-ww/2, sh/2-wh/2
		}

		marginTop, marginRight, marginBottom, marginLeft := w.GetMarginBorder()
		w.SetBorder(true)
		w.SetRect(x+w.x+marginLeft, y+w.y+marginTop, w.width-marginRight, w.height-marginBottom)

		w.Draw(screen)

		if w.GetStatus() != "" {
			Print(screen, []byte(w.GetStatus()), w.x+1, w.y+w.height-1, w.width-2, w.GetStatusBarAlign(), w.GetStatusBarColor())
		}
	}
}

// MouseHandler returns the mouse handler for this primitive.
func (wm *WindowManager) MouseHandler() func(action MouseAction, event *tcell.EventMouse, setFocus func(p Primitive)) (consumed bool, capture Primitive) {
	return wm.WrapMouseHandler(func(action MouseAction, event *tcell.EventMouse, setFocus func(p Primitive)) (consumed bool, capture Primitive) {
		if !wm.InRect(event.Position()) {
			return false, nil
		}

		if action == MouseMove {
			x, y, _, _ := wm.GetInnerRect()
			mouseX, mouseY := event.Position()

			for _, w := range wm.windows {
				if w.dragWX != -1 || w.dragWY != -1 {
					offsetX := w.x - (mouseX - x)
					offsetY := w.y - (mouseY - y)

					w.x -= offsetX + w.dragWX
					w.y -= offsetY + w.dragWY

					consumed = true
				}

				if w.dragX != 0 {
					if w.dragX == -1 {
						offsetX := w.x - (mouseX - x)

						if w.width+offsetX >= Styles.WindowMinWidth {
							w.x -= offsetX
							w.width += offsetX
						}
					} else {
						offsetX := mouseX - (x + w.x + w.width)

						if w.width+offsetX >= Styles.WindowMinWidth {
							w.width += offsetX
						}
					}

					consumed = true
				}

				if w.dragY != 0 {
					if w.dragY == -1 {
						offsetY := mouseY - (y + w.y + w.height)

						if w.height+offsetY >= Styles.WindowMinHeight {
							w.height += offsetY
						}
					} else {
						offsetY := w.y - (mouseY - y)

						if w.height+offsetY >= Styles.WindowMinHeight {
							w.y -= offsetY
							w.height += offsetY
						}
					}

					consumed = true
				}
			}
		} else if action == MouseLeftUp {
			for _, w := range wm.windows {
				w.dragX, w.dragY = 0, 0
				w.dragWX, w.dragWY = -1, -1
			}
		}

		// Focus window on mousedown
		var (
			focusWindow      *Window
			focusWindowIndex int
		)
		for i := len(wm.windows) - 1; i >= 0; i-- {
			if wm.windows[i].InRect(event.Position()) {
				focusWindow = wm.windows[i]
				focusWindowIndex = i
				break
			}
		}
		if focusWindow != nil {
			if action == MouseLeftDown || action == MouseMiddleDown || action == MouseRightDown {
				for _, w := range wm.windows {
					if w != focusWindow {
						w.Blur()
					}
				}

				wm.windows = append(append(wm.windows[:focusWindowIndex], wm.windows[focusWindowIndex+1:]...), focusWindow)
			}

			return focusWindow.MouseHandler()(action, event, setFocus)
		}

		return consumed, nil
	})
}
