package main

import (
	"github.com/gdamore/tcell"
	"fmt"
	"os"
	"github.com/gdamore/tcell/encoding"
	"github.com/demouth/mario-go"
	"time"
)

var defStyle tcell.Style

func main() {
	encoding.Register()

	s, e := tcell.NewScreen()
	if e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}
	if e := s.Init(); e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}
	defStyle = tcell.StyleDefault.
		Background(tcell.ColorBlack).
		Foreground(tcell.ColorWhite)
	s.SetStyle(defStyle)
	s.Clear()

	st := tcell.StyleDefault.Background(tcell.ColorRed)
	w, h := s.Size()

	world := mario_go.NewWorld(s)
	world.Width = w
	world.Height = h
	world.Draw()
	sp := mario_go.NewMario()
	world.AddChild(sp)

	quit := make(chan struct{})

	go func() {
		for {
			ev := s.PollEvent()
			switch ev := ev.(type) {
			case *tcell.EventKey:
				if ev.Key() == tcell.KeyEscape {
					close(quit)
				} else if ev.Key() == tcell.KeyCtrlC {
					close(quit)
				} else if ev.Key() == tcell.KeyRight {
					sp.Right()
				} else if ev.Key() == tcell.KeyLeft {
					sp.Left()
				} else if ev.Key() == tcell.KeyUp {
					sp.Jump()
				} else if ev.Key() == tcell.KeyDown {
					sp.SetY(sp.Y()-1)
				}
			case *tcell.EventResize:
				w, h = s.Size()
				world.Width = w
				world.Height = h
			default:
				s.SetContent(w-1, h-1, 'X', nil, st)
			}
		}
	}()

	loop:
	for {
		select {
		case <-quit:
			break loop
		case <-time.After(time.Millisecond * 20):
		}
		sp.Move()
		world.Draw()
		s.Show()
		s.Clear()
	}

	s.Fini()
	os.Exit(0)
}
