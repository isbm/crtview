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
		flags:        flags,
		DialogWindow: NewDialogWindow(),
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
	tmd.SetTitleColor(brc)

	//tmd.SetButtonsToBottom(true) /XXX: Buggy!

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
	x, y, w, _ := tmd.GetRect()
	_, _, _, mh := tmd.msg.GetRect()
	tmd.msg.SetRect(x, y, w, mh)

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
