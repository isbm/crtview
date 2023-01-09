package crtforms

import (
	"reflect"
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

	// Adds another column in front that shows selected items.
	showSelectedColumn bool

	hiddenColumns []int

	header []string
	rows   [][]string

	// Index of newly selectedRows and deselectedRows items, if multi-select mode
	selectedRows, deselectedRows []int

	*crtview.Table
	*crtview.FormItemBaseMixin
}

func NewFormTabularChoice(label string, header []string, rows [][]string, showSelected bool, hidden ...int) *FormTabularChoice {
	return (&FormTabularChoice{
		Table:              crtview.NewTable(),
		selectedRows:       []int{},
		deselectedRows:     []int{},
		hiddenColumns:      hidden,
		expandedColumn:     -1,
		isMultiSelect:      false,
		hasSearchFilter:    false,
		showSelectedColumn: showSelected,
		fieldHeight:        5,
		valueColumn:        0,
		label:              label,
		header:             header,
		rows:               rows,
	}).init()
}

func (tbc *FormTabularChoice) skip(v int) bool {
	for _, c := range tbc.hiddenColumns {
		if !tbc.showSelectedColumn {
			c--
		}
		if v == c {
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

	// Alter data, if there is selected column to display
	if tbc.showSelectedColumn {
		insert := func(row []string, data string, pos int) []string {
			return append(row[:pos], append([]string{data}, row[pos:]...)...)
		}

		tbc.header = insert(tbc.header, "...", 0)
		for i, r := range tbc.rows {
			tbc.rows[i] = insert(r, "   ", 0)
		}
	}

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
	col--
	if col < 0 {
		return tbc
	}

	tbc.expandedColumn = col
	if tbc.showSelectedColumn {
		tbc.expandedColumn++
	}

	for idx := range tbc.rows {
		tbc.GetCell(idx, tbc.expandedColumn).SetExpansion(0xB)
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
		offset := tbc.valueColumn
		if tbc.showSelectedColumn {
			offset++
		}
		return tbc.rows[row][offset]
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

func (tbc *FormTabularChoice) SetSelectedFunc(handler func(row, column int)) *FormTabularChoice {
	tbc.Table.SetSelectedFunc(func(row, column int) {
		tbc.setSelectedMarker(row)
		handler(row, column)
	})
	return tbc
}

func (tbc *FormTabularChoice) Select(row, col int) {
	tbc.Table.Select(row, col)
	tbc.setSelectedMarker(row)
}

func (tbc *FormTabularChoice) setSelectedMarker(row int) {
	if tbc.showSelectedColumn {
		for i := 1; i < tbc.GetRowCount(); i++ { // skip the header
			label := "   "
			if i == row {
				label = " ◆ "
			}
			tbc.GetCell(i, 0).SetText(label)
		}
	}
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

func (tbc *FormTabularChoice) GetWidgetType() string {
	return reflect.TypeOf(tbc).Elem().Name()
}
