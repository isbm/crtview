package main

import "github.com/isbm/crtview"

// Center returns a new primitive which shows the provided primitive in its
// center, given the provided primitive's size.
func Center(width, height int, p crtview.Primitive) crtview.Primitive {
	subFlex := crtview.NewFlex()
	subFlex.SetDirection(crtview.FlexRow)
	subFlex.AddItem(crtview.NewBox(), 0, 1, false)
	subFlex.AddItem(p, height, 1, true)
	subFlex.AddItem(crtview.NewBox(), 0, 1, false)

	flex := crtview.NewFlex()
	flex.AddItem(crtview.NewBox(), 0, 1, false)
	flex.AddItem(subFlex, width, 1, true)
	flex.AddItem(crtview.NewBox(), 0, 1, false)

	return flex
}
