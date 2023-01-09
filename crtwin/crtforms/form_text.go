package crtforms

import (
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/isbm/crtview"
	"github.com/isbm/textwrap"
)

type FormTextView struct {
	// Text is wrapped manually by the user.
	autofill bool

	*crtview.TextView
	*crtview.FormItemBaseMixin
}

func NewFormTextView() *FormTextView {
	return (&FormTextView{
		TextView: crtview.NewTextView(),
	}).init()
}

// Initialise the widget with default stuff (part of the constructor)
func (ftw *FormTextView) init() *FormTextView {
	return ftw
}

// GetFieldHeight returns current height of the field, if the item is not anyway in a Flex or Grid
func (ftw *FormTextView) GetFieldHeight() int {
	_, _, w, _ := ftw.GetRect()
	return len(textwrap.NewTextWrap().SetWidth(w).SetDropWhitespace(true).Wrap(strings.ReplaceAll(ftw.GetText(true), "\n", " ")))
}

// GetFieldWidth returns current width of the field, if the item is not anyway in a Flex or Grid
func (ftw *FormTextView) GetFieldWidth() int {
	_, _, w, _ := ftw.GetRect()
	return w
}

// GetLabel returns an empty string (and SetLabel sets nothing as well).
// This widget has no label, but is a label on its own.
func (ftw *FormTextView) GetLabel() string {
	return ""
}

// SetTextAutofill sets wrapping text manually or not. If set to true, all newlines
// will be discared and the text will be always reformatted according to the width of the field.
// The height will be as text goes. If set to false, then the text will be displayed as is.
// If there is not enough place for text anymore, then it will be not shown.
func (ftw *FormTextView) SetTextAutofill(autofill bool) *FormTextView {
	ftw.autofill = autofill
	return ftw
}

func (ftw *FormTextView) SetText(text string) *FormTextView {
	ftw.SetTextAutofill(!strings.Contains(text, "\n"))
	ftw.TextView.SetText(text)
	return ftw
}

// SetBackgroundColor sets the background color of the widget
func (ftw *FormTextView) SetBackgroundColor(color tcell.Color) {
	ftw.TextView.SetBackgroundColor(color)
}

func (ftw *FormTextView) SetTextAlign(align int) *FormTextView {
	ftw.TextView.SetTextAlign(align)
	return ftw
}

/*
These are no-op methods, but here are only for the interface requirements to satisfy.
*/
func (ftw *FormTextView) SetFieldBackgroundColor(color tcell.Color)        {}
func (ftw *FormTextView) SetFieldBackgroundColorFocused(color tcell.Color) {}
func (ftw *FormTextView) SetFieldTextColor(color tcell.Color)              {}
func (ftw *FormTextView) SetFieldTextColorFocused(color tcell.Color)       {}
func (ftw *FormTextView) SetFinishedFunc(action func(key tcell.Key))       {}
func (ftw *FormTextView) SetLabelColor(color tcell.Color)                  {}
func (ftw *FormTextView) SetLabelColorFocused(color tcell.Color)           {}
func (ftw *FormTextView) SetLabelWidth(width int)                          {}
func (ftw *FormTextView) SetLabel(label string)                            {}

// Draw the widget on its place. This also contains dynamic hooks,
// such as reformatting text on a fly, depending on the current dimensions.
func (ftw *FormTextView) Draw(screen tcell.Screen) {
	ftw.GetFieldHeight()
	x, y, w, _ := ftw.GetRect()

	// Place text on the widget where it belongs
	text := ftw.GetText(true)
	if ftw.autofill {
		text = textwrap.NewTextWrap().SetWidth(w).SetDropWhitespace(true).Fill(strings.ReplaceAll(text, "\n", " "))
	}

	ftw.SetText(text)
	ftw.SetRect(x, y, w, ftw.GetFieldHeight())
	ftw.TextView.Draw(screen)
}
