package crtwin

import (
	"github.com/gdamore/tcell/v2"
	"github.com/isbm/crtview"
	"github.com/isbm/crtview/crtwin/crtforms"
)

const (
	DIALOG_OK_CANCEL = 1 << iota
	DIALOG_YES_NO
	DIALOG_OK
	DIALOG_TIMED
	DIALOG_TYPE_INFO
	DIALOG_TYPE_ALT_INFO // Alternative, darker coloring
	DIALOG_TYPE_WARNING
	DIALOG_TYPE_ALERT
)

type ModalDialog struct {
	msg           *crtforms.FormTextView
	flags         int
	confirmAction func()
	cancelAction  func()

	*DialogWindow
}

func NewModalDialog(flags int) *ModalDialog {
	return (&ModalDialog{
		flags:         flags,
		DialogWindow:  NewDialogWindow(),
		confirmAction: func() {},
		cancelAction:  func() {},
	}).init()
}

func (tmd *ModalDialog) init() *ModalDialog {
	var bgc, brc, txc tcell.Color
	if tmd.flags&DIALOG_TYPE_ALT_INFO != 0 {
		bgc = crtview.Styles.AltInfoDialogBackgroundColor
		brc = crtview.Styles.AltInfoDialogBorderColor
		txc = crtview.Styles.AltInfoDialogTextColor
	} else if tmd.flags&DIALOG_TYPE_WARNING != 0 {
		bgc = crtview.Styles.WarningDialogBackgroundColor
		brc = crtview.Styles.WarningDialogBorderColor
		txc = crtview.Styles.WarningDialogTextColor
	} else if tmd.flags&DIALOG_TYPE_ALERT != 0 {
		bgc = crtview.Styles.AlertDialogBackgroundColor
		brc = crtview.Styles.AlertDialogBorderColor
		txc = crtview.Styles.AlertDialogTextColor
	} else {
		// DIALOG_TYPE_INFO
		bgc = crtview.Styles.InfoDialogBackgroundColor
		brc = crtview.Styles.InfoDialogBorderColor
		txc = crtview.Styles.InfoDialogTextColor
	}

	tmd.SetBorder(true)
	tmd.SetBackgroundColor(bgc)
	tmd.SetBorderColor(brc)
	tmd.SetBorderColorFocused(brc)
	tmd.SetTitleColor(brc)

	tmd.msg = crtforms.NewFormTextView()
	tmd.msg.SetBackgroundColor(bgc)
	tmd.msg.SetTextColor(txc)

	tmd.AddFormItem(tmd.msg)

	if tmd.flags&DIALOG_OK_CANCEL != 0 {
		tmd.AddButton("OK", func() { tmd.confirmAction() })
		tmd.AddButton("Cancel", func() { tmd.cancelAction() })
	} else if tmd.flags&DIALOG_YES_NO != 0 {
		tmd.AddButton("Yes", func() { tmd.confirmAction() })
		tmd.AddButton("No", func() { tmd.cancelAction() })
	} else if tmd.flags&DIALOG_OK != 0 {
		tmd.AddButton("OK", func() { tmd.confirmAction() })
	} else if tmd.flags&DIALOG_TIMED != 0 {

	} else {
		panic("Unknown dialog type")
	}

	// Default alert settings
	tmd.SetTextAlign(crtview.AlignCenter)
	tmd.SetCentered(true)
	tmd.SetTextAutofill(true)
	tmd.SetButtonsToBottom(true)

	return tmd
}

func (tmd *ModalDialog) SetMessage(msg string) {
	tmd.msg.SetText(msg)
}

// Draw the alert window. This is a bit tricky, because alert window is dynamic on the height,
// depending on the amount of the text passed as a message.
func (tmd *ModalDialog) Draw(screen tcell.Screen) {
	// Resize alert according to the text content size
	w := tmd.msg.GetFieldWidth()
	if len(tmd.GetTitle()) > tmd.msg.GetFieldWidth() {
		w = len(tmd.GetTitle())
	}
	tmd.SetSize(w+6, tmd.msg.GetFieldHeight()+8) // 6 & 8 is a padding for shadows, borders etc

	// Pass parent width to the text field
	fx, fy, _, fh := tmd.msg.GetRect()
	tmd.msg.SetRect(fx, fy, w+6, fh)

	// Draw the rest
	tmd.DialogWindow.Draw(screen)
}

func (tmd *ModalDialog) SetTextAutofill(autofill bool) *ModalDialog {
	tmd.msg.SetTextAutofill(autofill)
	return tmd
}

func (tmd *ModalDialog) SetTextAlign(align int) *ModalDialog {
	tmd.msg.SetTextAlign(align)
	return tmd
}

func (tmd *ModalDialog) SetOnConfirmAction(action func()) {
	tmd.confirmAction = action
}

func (tmd *ModalDialog) SetOnCancelAction(action func()) {
	tmd.cancelAction = action
}
