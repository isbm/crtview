package crtforms

import (
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/isbm/crtview"
)

type FormTabularChoice struct {
	label             string
	labelColor        tcell.Color
	labelColorFocused tcell.Color

	// Height of the field on the form. Default 5.
	fieldHeight int

	// This adds another column at the
	// beginning with ticks which fields are selected.
	isMultiSelect bool

	// Displays a search field (limiter of the displayed data)
	// to narrow the visibility
	hasSearchFilter bool

	// Which column has value.
	// Default is 0 (i.e. the first one)
	valueColumn int

	// Which column is expanded. Default is -1
	expandedColumn int

	hiddenColumns []int

	header []string
	rows   [][]string

	// Index of newly selected and deselected items, if multi-select mode
	selected, deselected []int

	*crtview.Table
}

func NewFormTabularChoice(label string, header []string, rows [][]string, hidden ...int) *FormTabularChoice {
	return (&FormTabularChoice{
		Table:           crtview.NewTable(),
		selected:        []int{},
		deselected:      []int{},
		hiddenColumns:   hidden,
		expandedColumn:  -1,
		isMultiSelect:   false,
		hasSearchFilter: false,
		fieldHeight:     5,
		valueColumn:     0,
		label:           label,
		header:          header,
		rows:            rows,
	}).init()
}

func (tbc *FormTabularChoice) skip(v int) bool {
	for _, c := range tbc.hiddenColumns {
		if v == c-1 {
			return true
		}
	}
	return false
}

func (tbc *FormTabularChoice) init() *FormTabularChoice {
	tbc.SetBorder(true)
	tbc.SetTitle(tbc.label)
	tbc.SetBorders(false)
	tbc.SetSeparator(crtview.Borders.Vertical)
	tbc.SetSelectable(true, false)
	tbc.Select(1, 1)
	tbc.SetFixed(1, 1)

	// Make header
	col := 0
	for idx, label := range tbc.header {
		if tbc.skip(idx) {
			continue
		}
		cell := crtview.NewTableCell(label)
		cell.SetSelectable(false)
		cell.SetTextColor(tcell.ColorYellow.TrueColor())
		tbc.SetCell(0, col, cell)
		col++
	}

	// Add rows
	for ridx, row := range tbc.rows {
		col = 0
		for cidx, label := range row {
			if tbc.skip(cidx) {
				continue
			}
			cell := crtview.NewTableCell(label)
			tbc.SetCell(ridx+1, col, cell)
			col++
		}
	}

	tbc.fillEmpty()

	return tbc
}

func (tbc *FormTabularChoice) Focus(delegate func(p crtview.Primitive)) {
	tbc.Table.Focus(delegate)
}

func (tbc *FormTabularChoice) removeEmpty() {
	visible := len(tbc.rows) + 1
	filler := tbc.fieldHeight - visible - 2

	if filler < 1 {
		return
	}

	for r := 0; r < filler+1; r++ {
		tbc.RemoveRow(visible)
	}
}

func (tbc *FormTabularChoice) fillEmpty() {
	visible := len(tbc.rows) + 1
	filler := tbc.fieldHeight - visible - 2

	if filler < 1 {
		return
	}

	for r := 0; r < filler; r++ {
		col := 0
		for c := range tbc.header {
			if tbc.skip(c) {
				continue
			}
			cell := crtview.NewTableCell(" ")
			cell.SetSelectable(false)
			if c == tbc.expandedColumn {
				cell.SetExpansion(0xB)
			}
			tbc.SetCell(r+visible, col, cell)
			col++
		}
	}
}

// SetFieldHeight sets the height of the field. It will never go less than 5, because
// there needs to be at least two options visible, including header and the border around.
func (tbc *FormTabularChoice) SetFieldHeight(height int) *FormTabularChoice {
	if height < 5 {
		height = 5
	}
	tbc.fieldHeight = height

	tbc.removeEmpty()
	tbc.fillEmpty()

	return tbc
}

// SetMultiselect sets the ability to select multiple choices in the table.
// NOTE: this will display an additional column at the beginning.
func (tbc *FormTabularChoice) SetMultiselect(multiselect bool) *FormTabularChoice {
	tbc.isMultiSelect = multiselect
	return tbc
}

// SetHasSearch allows to display an additional form element that adds a search entry
// to filter-out data, narrowing the selection.
func (tbc *FormTabularChoice) SetHasSearch(search bool) *FormTabularChoice {
	tbc.hasSearchFilter = search
	return tbc
}

// SetExpandingColumn sets the column that is expanding.
// NOTE: 1st column start from 1, not from 0.
func (tbc *FormTabularChoice) SetExpandingColumn(col int) *FormTabularChoice {
	if col >= 0 {
		tbc.expandedColumn = col
	}

	for idx := range tbc.rows {
		tbc.GetCell(idx, tbc.expandedColumn-1).SetExpansion(0xB)
	}

	return tbc
}

// SetValueColumn points to the value column as a result of selected choice.
// NOTE: Index starts from 1, not from 0.
func (tbc *FormTabularChoice) SetValueColumn(col int) *FormTabularChoice {
	tbc.valueColumn = col - 1
	return tbc
}

// GetFieldHeight returns current height of the field, if the item is not anyway in a Flex or Grid
func (tbc *FormTabularChoice) GetFieldHeight() int {
	return tbc.fieldHeight
}

// GetFieldWidth returns current width of the field, if the item is not anyway in a Flex or Grid
func (tbc *FormTabularChoice) GetFieldWidth() int {
	return 10
}

// GetValueAt (row). Note: rows here are counted from 0, not from 1.
func (tbc *FormTabularChoice) GetValueAt(row int) string {
	if tbc.valueColumn > -1 {
		return tbc.rows[row][tbc.valueColumn]
	}

	return strings.Join(tbc.rows[row], ",")
}

// GetLabel returns the label of the list
func (tbc *FormTabularChoice) GetLabel() string {
	return tbc.label
}

// SetLabel sets the label of the list
func (tbc *FormTabularChoice) SetLabel(label string) {
	tbc.label = label
}

// SetLabelColor
func (tbc *FormTabularChoice) SetLabelColor(color tcell.Color) {
	tbc.labelColor = color
}

// SetLabelColorFocused
func (tbc *FormTabularChoice) SetLabelColorFocused(color tcell.Color) {
	tbc.labelColorFocused = color
}

func (tbc *FormTabularChoice) SetBackgroundColor(color tcell.Color) {
	tbc.Table.SetBackgroundColor(color)
}

func (tbc *FormTabularChoice) SetFieldBackgroundColor(color tcell.Color)        {}
func (tbc *FormTabularChoice) SetFieldBackgroundColorFocused(color tcell.Color) {}
func (tbc *FormTabularChoice) SetFieldTextColor(color tcell.Color)              {}
func (tbc *FormTabularChoice) SetFieldTextColorFocused(color tcell.Color)       {}
func (tbc *FormTabularChoice) SetFinishedFunc(action func(key tcell.Key)) {
	tbc.Table.SetDoneFunc(action)
}
func (tbc *FormTabularChoice) SetLabelWidth(width int) {}

// Draw table-list pick
func (tbc *FormTabularChoice) Draw(screen tcell.Screen) {
	tbc.Table.Draw(screen)
}
