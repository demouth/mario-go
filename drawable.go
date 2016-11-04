package mario_go

import "github.com/gdamore/tcell"

type Drawable interface {
	SetScreen(screen tcell.Screen)
	Dots() Dots
	X() int
	Y() int
	SetX(v int)
	SetY(v int)
	Width() int
	Height() int
}
