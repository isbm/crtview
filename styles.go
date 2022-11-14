package crtview

import "github.com/gdamore/tcell/v2"

// Theme defines the colors used when primitives are initialized.
type Theme struct {
	// Title, border and other lines
	TitleColor    tcell.Color // Box titles.
	BorderColor   tcell.Color // Box borders.
	GraphicsColor tcell.Color // Graphics.

	// Text
	PrimaryTextColor           tcell.Color // Primary text.
	SecondaryTextColor         tcell.Color // Secondary text (e.g. labels).
	TertiaryTextColor          tcell.Color // Tertiary text (e.g. subtitles, notes).
	InverseTextColor           tcell.Color // Text on primary-colored backgrounds.
	ContrastPrimaryTextColor   tcell.Color // Primary text for contrasting elements.
	ContrastSecondaryTextColor tcell.Color // Secondary text on ContrastBackgroundColor-colored backgrounds.
	ButtonTextColor            tcell.Color // Button color for label
	ButtonTextFocusedColor     tcell.Color // Button color for focused label
	ShadowTextColor            tcell.Color // Shadow text color

	// Background
	PrimitiveBackgroundColor     tcell.Color // Main background color for primitives.
	ContrastBackgroundColor      tcell.Color // Background color for contrasting elements.
	MoreContrastBackgroundColor  tcell.Color // Background color for even more contrasting elements.
	ButtonBackgroundColor        tcell.Color // Background color for buttons
	ButtonFocusedBackgroundColor tcell.Color // Background color for focused/selected buttons
	ShadowColor                  tcell.Color // Shadow color

	// Check box
	CheckBoxCheckedRune rune

	// Context menu
	ContextMenuPaddingTop    int
	ContextMenuPaddingBottom int
	ContextMenuPaddingLeft   int
	ContextMenuPaddingRight  int

	// Drop down
	DropDownAbbreviationChars string // The chars to show when the option's text gets shortened.
	DropDownSymbol            rune   // The symbol to draw at the end of the field.

	// Scroll bar
	ScrollBarColor tcell.Color

	// Window
	WindowMinWidth  int
	WindowMinHeight int

	// Dialogs
	InfoDialogBackgroundColor    tcell.Color
	InfoDialogTextColor          tcell.Color
	InfoDialogBorderColor        tcell.Color
	AltInfoDialogBackgroundColor tcell.Color
	AltInfoDialogTextColor       tcell.Color
	AltInfoDialogBorderColor     tcell.Color
	WarningDialogBackgroundColor tcell.Color
	WarningDialogTextColor       tcell.Color
	WarningDialogBorderColor     tcell.Color
	AlertDialogBackgroundColor   tcell.Color
	AlertDialogTextColor         tcell.Color
	AlertDialogBorderColor       tcell.Color
}

// Styles defines the appearance of an application. The default is for a black
// background and some basic colors: black, white, yellow, green, cyan, and
// blue.
var Styles = Theme{
	TitleColor:    tcell.ColorWhite.TrueColor(),
	BorderColor:   tcell.ColorWhite.TrueColor(),
	GraphicsColor: tcell.ColorWhite.TrueColor(),

	PrimaryTextColor:           tcell.ColorWhite.TrueColor(),
	SecondaryTextColor:         tcell.ColorYellow.TrueColor(),
	TertiaryTextColor:          tcell.ColorGreen.TrueColor(),
	InverseTextColor:           tcell.ColorBlue.TrueColor(),
	ContrastPrimaryTextColor:   tcell.ColorBlack.TrueColor(),
	ContrastSecondaryTextColor: tcell.ColorDarkCyan.TrueColor(),

	PrimitiveBackgroundColor:    tcell.ColorBlack.TrueColor(),
	ContrastBackgroundColor:     tcell.ColorBlue.TrueColor(),
	MoreContrastBackgroundColor: tcell.ColorGreen.TrueColor(),

	ButtonBackgroundColor:        tcell.ColorWhite.TrueColor(),
	ButtonFocusedBackgroundColor: tcell.ColorYellow.TrueColor(),
	ButtonTextColor:              tcell.ColorBlack.TrueColor(),
	ButtonTextFocusedColor:       tcell.ColorBlack.TrueColor(),

	ShadowTextColor: tcell.NewRGBColor(0x66, 0x66, 0x66),
	ShadowColor:     tcell.ColorBlack,

	CheckBoxCheckedRune: 'X',

	ContextMenuPaddingTop:    0,
	ContextMenuPaddingBottom: 0,
	ContextMenuPaddingLeft:   1,
	ContextMenuPaddingRight:  1,

	DropDownAbbreviationChars: "...",
	DropDownSymbol:            '▼',

	ScrollBarColor: tcell.ColorWhite.TrueColor(),

	WindowMinWidth:  4,
	WindowMinHeight: 3,

	// Default dialogs theming, this usually doesn't need to be changed :)
	InfoDialogBackgroundColor:    tcell.ColorLightGrey.TrueColor(),
	InfoDialogTextColor:          tcell.ColorBlack.TrueColor(),
	InfoDialogBorderColor:        tcell.ColorBlack.TrueColor(),
	AltInfoDialogBackgroundColor: tcell.ColorTeal.TrueColor(),
	AltInfoDialogTextColor:       tcell.ColorWhite.TrueColor(),
	AltInfoDialogBorderColor:     tcell.ColorAqua.TrueColor(),
	WarningDialogBackgroundColor: tcell.ColorOlive.TrueColor(),
	WarningDialogTextColor:       tcell.ColorWhite.TrueColor(),
	WarningDialogBorderColor:     tcell.ColorYellow.TrueColor(),
	AlertDialogBackgroundColor:   tcell.ColorRed.TrueColor(),
	AlertDialogTextColor:         tcell.ColorWhite.TrueColor(),
	AlertDialogBorderColor:       tcell.ColorLightYellow.TrueColor(),
}
