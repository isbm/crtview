package crtview

const (
	BorderDouble = iota
	BorderSingle
	BorderHeavy
)

// Borders defines various borders used when primitives are drawn.
// These may be changed to accommodate a different look and feel.
var Borders = struct {
	Horizontal  rune
	Vertical    rune
	TopLeft     rune
	TopRight    rune
	BottomLeft  rune
	BottomRight rune

	LeftT   rune
	RightT  rune
	TopT    rune
	BottomT rune
	Cross   rune

	HorizontalFocus  rune
	VerticalFocus    rune
	TopLeftFocus     rune
	TopRightFocus    rune
	BottomLeftFocus  rune
	BottomRightFocus rune
}{
	Horizontal:  BoxDrawingsLightHorizontal,
	Vertical:    BoxDrawingsLightVertical,
	TopLeft:     BoxDrawingsLightDownAndRight,
	TopRight:    BoxDrawingsLightDownAndLeft,
	BottomLeft:  BoxDrawingsLightUpAndRight,
	BottomRight: BoxDrawingsLightUpAndLeft,

	LeftT:   BoxDrawingsLightVerticalAndRight,
	RightT:  BoxDrawingsLightVerticalAndLeft,
	TopT:    BoxDrawingsLightDownAndHorizontal,
	BottomT: BoxDrawingsLightUpAndHorizontal,
	Cross:   BoxDrawingsLightVerticalAndHorizontal,

	/*
		HeavyHorizontalFocus:  BoxDrawingsHeavyHorizontal,
		HeavyVerticalFocus:    BoxDrawingsHeavyVertical,
		HeavyTopLeftFocus:     BoxDrawingsHeavyDownAndRight,
		HeavyTopRightFocus:    BoxDrawingsHeavyDownAndLeft,
		HeavyBottomLeftFocus:  BoxDrawingsHeavyUpAndRight,
		HeavyBottomRightFocus: BoxDrawingsHeavyUpAndLeft,
	*/

	HorizontalFocus:  BoxDrawingsDoubleHorizontal,
	VerticalFocus:    BoxDrawingsDoubleVertical,
	TopLeftFocus:     BoxDrawingsDoubleDownAndRight,
	TopRightFocus:    BoxDrawingsDoubleDownAndLeft,
	BottomLeftFocus:  BoxDrawingsDoubleUpAndRight,
	BottomRightFocus: BoxDrawingsDoubleUpAndLeft,
}
