package mario_go

import (
	"github.com/gdamore/tcell"
)

type World struct {
	screen tcell.Screen
	children []Drawable
	Width int
	Height int
	CameraX int
	CameraY int
}


func NewWorld(s tcell.Screen) *World {
	w := &World{}
	w.screen = s
	w.children = []Drawable{}
	w.CameraX=0
	w.CameraY=0
	return w
}

func (world *World) AddChild(child Drawable) {
	child.SetScreen(world.screen)
	world.children = append(world.children,child)
}

func (w *World) Draw() {
	for _, child := range w.children {
		w.DrawDots(child)
	}
}


func (w World) DrawDots(child Drawable) {
	dots := child.Dots()
	for _,dot := range dots {
		st := tcell.StyleDefault.Background(tcell.NewHexColor(dot.Color))
		w.screen.SetContent(w.CameraX+child.X()*2+dot.X*2,  w.Height-w.CameraY-child.Y()-dot.Y-1,' ',nil,st)
		w.screen.SetContent(w.CameraX+child.X()*2+dot.X*2+1,w.Height-w.CameraY-child.Y()-dot.Y-1,' ',nil,st)
	}
}