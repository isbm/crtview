package crtview

import (
	"strings"

	"github.com/gdamore/tcell/v2"
)

type InfoText struct {
	// Text color
	defaultColor tcell.Color

	// Height of the field
	fieldheight int

	// A callback function set by the Form class and called when the user leaves
	// this form item.
	finished func(tcell.Key)

	// An optional function which is called when the user indicated that they
	// are done entering text. The key which was pressed is provided (tab,
	// shift-tab, or escape).
	done func(tcell.Key)

	*Box
	TextView
	*FormItemBaseMixin
	FormItem
}

func NewInfoText(content string) *InfoText {
	tb := *NewTextView()
	tb.SetText(content)

	nfo := &InfoText{
		Box:      NewBox(),
		TextView: tb,
	}
	nfo.fieldheight = len(strings.Split(content, "\n"))
	return nfo
}

// Draw draws this primitive onto the screen. Implementers can call the
// screen's ShowCursor() function but should only do so when they have focus.
// (They will need to keep track of this themselves.)
func (nt *InfoText) Draw(screen tcell.Screen) {
	x, y, w, _ := nt.GetRect()
	nt.SetRect(x, y, w, nt.GetFieldHeight())
	nt.TextView.SetTextColor(nt.defaultColor)
	nt.TextView.Draw(screen)
}

// GetRect returns the current position of the primitive, x, y, width, and
// height.
func (nt *InfoText) GetRect() (int, int, int, int) {
	return nt.TextView.GetRect()
}

// SetRect sets a new position of the primitive.
func (nt *InfoText) SetRect(x, y, width, height int) {
	nt.TextView.SetRect(x, y, width, height)
}

// IsVisible returns whether or not the primitive is visible.
func (nt *InfoText) IsVisible() bool {
	return nt.TextView.IsVisible()
}

// SetVisible sets whether or not the primitive is visible.
func (nt *InfoText) setVisible(v bool) {
	nt.TextView.setVisible(v)
}

// InputHandler returns a handler which receives key events when it has focus.
// It is called by the Application class.
//
// A value of nil may also be returned, in which case this primitive cannot
// receive focus and will not process any key events.
//
// The handler will receive the key event and a function that allows it to
// set the focus to a different primitive, so that future key events are sent
// to that primitive.
//
// The Application's Draw() function will be called automatically after the
// handler returns.
//
// The Box class provides functionality to intercept keyboard input. If you
// subclass from Box, it is recommended that you wrap your handler using
// Box.WrapInputHandler() so you inherit that functionality.
func (nt *InfoText) InputHandler() func(event *tcell.EventKey, setFocus func(p Primitive)) {
	return nt.TextView.InputHandler()
}

// Focus is called by the application when the primitive receives focus.
// Implementers may call delegate() to pass the focus on to another primitive.
func (nt *InfoText) Focus(delegate func(p Primitive)) {
	nt.TextView.Focus(delegate)

	// Immediately skip to the next field
	if nt.finished != nil {
		nt.finished(tcell.KeyTab) // XXX: This does not know from where previous focus happened
	}
}

// Blur is called by the application when the primitive loses focus.
func (nt *InfoText) Blur() {
	nt.TextView.Blur()
}

// GetFocusable returns the item's Focusable.
func (nt *InfoText) GetFocusable() Focusable {
	return nt.TextView.GetFocusable()
}

// MouseHandler returns a handler which receives mouse events.
// It is called by the Application class.
//
// A value of nil may also be returned to stop the downward propagation of
// mouse events.
//
// The Box class provides functionality to intercept mouse events. If you
// subclass from Box, it is recommended that you wrap your handler using
// Box.WrapMouseHandler() so you inherit that functionality.
func (nt *InfoText) MouseHandler() func(action MouseAction, event *tcell.EventMouse, setFocus func(p Primitive)) (consumed bool, capture Primitive) {
	return nt.TextView.MouseHandler()
}

// GetLabel returns the item's label text.
func (nt *InfoText) GetLabel() string {
	return ""
}

// SetLabelWidth sets the screen width of the label. A value of 0 will cause the
// primitive to use the width of the label string.
func (nt *InfoText) SetLabelWidth(int) {}

// SetLabelColor sets the color of the label.
func (nt *InfoText) SetLabelColor(tcell.Color) {}

// SetLabelColor sets the color of the label when focused.
func (nt *InfoText) SetLabelColorFocused(tcell.Color) {}

// GetFieldWidth returns the width of the form item's field (the area which
// is manipulated by the user) in number of screen cells. A value of 0
// indicates the the field width is flexible and may use as much space as
// required.
func (nt *InfoText) GetFieldWidth() int { return 0 }

// GetFieldHeight returns the height of the form item.
func (nt *InfoText) GetFieldHeight() int {
	return nt.fieldheight
}

// SetFieldTextColor sets the text color of the input area.
func (nt *InfoText) SetFieldTextColor(c tcell.Color) {
	nt.TextView.SetTextColor(c)
}

// SetFieldTextColorFocused sets the text color of the input area when focused.
func (nt *InfoText) SetFieldTextColorFocused(c tcell.Color) {
	//nt.TextView.SetTextColor(tcell.ColorGreenYellow)
}

func (nt *InfoText) SetFieldBackgroundColor(c tcell.Color)        {}
func (nt *InfoText) SetFieldBackgroundColorFocused(c tcell.Color) {}

// SetBackgroundColor sets the background color of the form item.
func (nt *InfoText) SetBackgroundColor(c tcell.Color) {
	nt.TextView.SetBackgroundColor(c)
}

// SetFinishedFunc sets a callback invoked when the user leaves the form item.
func (nt *InfoText) SetFinishedFunc(handler func(key tcell.Key)) {
	nt.Lock()
	defer nt.Unlock()

	nt.finished = handler
}

func (nt *InfoText) SetDoneFunc(handler func(key tcell.Key)) {
	nt.Lock()
	defer nt.Unlock()

	nt.done = handler
}

// IsMaximised returns if a widget can be vertically maximised.
func (nt *InfoText) IsMaximised() bool {
	return false
}

// SetMaximised sets widget to be maximised vertically, as long as it is the last one
// and is maximise-able (like tabular view or text entry). One-unit high fields won't
// be affected, such as field text entry or password or dropdown etc.
func (nt *InfoText) SetMaximised(maximised bool) {}

// GetWidgetType returns a class of the widget.
func (nt *InfoText) GetWidgetType() string {
	return ""
}
