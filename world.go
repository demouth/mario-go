package mario_go

import (
	"github.com/gdamore/tcell"
)

type World struct {
	screen  tcell.Screen
	objects []Drawable
	Width   int
	Height  int
	CameraX int
	CameraY int
	Mario   *Mario
}

func NewWorld(s tcell.Screen) *World {
	w := &World{}
	w.screen = s
	w.objects = []Drawable{}
	w.CameraX = 0
	w.CameraY = 0
	return w
}

func (world *World) SetMario(mario *Mario) {
	mario.SetScreen(world.screen)
	world.Mario = mario
}

func (world *World) AddObject(object Drawable) {
	object.SetScreen(world.screen)
	world.objects = append(world.objects, object)
}

func (w *World) HitTest() {
	for _, child := range w.objects {
		if child.X() < w.Mario.X()+w.Mario.Width() && w.Mario.X() < child.X()+child.Width() && child.Y() < w.Mario.Y()+w.Mario.Height() && w.Mario.Y() < child.Y()+child.Height() {
			leftOverlap := w.Mario.X() - (child.X() + child.Width())
			rightOverlap := child.X() - (w.Mario.X() + w.Mario.Width())
			topOverlap := w.Mario.Y() - (child.Y() + child.Height())
			bottomOverlap := child.Y() - (w.Mario.Y() + w.Mario.Height())
			if rightOverlap < leftOverlap && topOverlap < leftOverlap && bottomOverlap < leftOverlap {
				w.Mario.SetX(child.X() + child.Width())
				w.Mario.StopX()
			} else if leftOverlap < rightOverlap && topOverlap < rightOverlap && bottomOverlap < rightOverlap {
				w.Mario.SetX(child.X() - w.Mario.Width())
				w.Mario.StopX()
			} else if leftOverlap < topOverlap && rightOverlap < topOverlap && bottomOverlap < topOverlap {
				w.Mario.SetY(child.Y() + child.Height())
				if !w.Mario.Rising() {
					w.Mario.Land()
				}
			} else if leftOverlap < bottomOverlap && rightOverlap < bottomOverlap && topOverlap < bottomOverlap {
				w.Mario.SetY(child.Y() - w.Mario.Height())
				w.Mario.Fall()
			}
		}
	}
}

func (w *World) Draw() {
	for _, object := range w.objects {
		w.DrawDots(object)
	}
	w.DrawDots(w.Mario)
}

func (w World) DrawDots(child Drawable) {
	dots := child.Dots()
	for _, dot := range dots {
		st := tcell.StyleDefault.Background(tcell.NewHexColor(dot.Color))
		w.screen.SetContent(w.CameraX*2+child.X()*2+dot.X*2, w.Height-w.CameraY-child.Y()-dot.Y-1, ' ', nil, st)
		w.screen.SetContent(w.CameraX*2+child.X()*2+dot.X*2+1, w.Height-w.CameraY-child.Y()-dot.Y-1, ' ', nil, st)
	}
}
