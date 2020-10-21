package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/isbm/crtview"
)

// End shows the final slide.
func End(nextSlide func()) (title string, content crtview.Primitive) {
	textView := crtview.NewTextView()
	textView.SetDoneFunc(func(key tcell.Key) {
		nextSlide()
	})
	url := "https://github.com/isbm/crtview"
	fmt.Fprint(textView, url)
	return "End", Center(len(url), 1, textView)
}
