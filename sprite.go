package mario_go

import "github.com/gdamore/tcell"

type Sprite struct {
	screen tcell.Screen
	x      float32
	y      float32
}

func NewSprite() *Sprite {
	s := &Sprite{}
	s.x = 0
	s.y = 0
	return s
}

func (s Sprite) X() int {
	return int(s.x)
}

func (s Sprite) Y() int {
	return int(s.y)
}

func (s *Sprite) SetX(v int) {
	s.x = float32(v)
}

func (s *Sprite) SetY(v int) {
	s.y = float32(v)
}

func (s *Sprite) SetScreen(screen tcell.Screen) {
	s.screen = screen
}

func (s *Sprite) Draw() {
	st := tcell.StyleDefault.Background(tcell.Color21)
	s.screen.SetContent(4, 4, ' ', nil, st)
}

func (s Sprite) Dots() Dots {
	var dots Dots
	d := NewDot(0, 0, 0x8A7301)
	d2 := NewDot(2, 1, 0xFF0000)
	dots = Dots{*d, *d2}
	return dots
}
