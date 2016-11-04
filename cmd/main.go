package main

import (
	"fmt"
	"github.com/demouth/mario-go"
	"github.com/gdamore/tcell"
	"github.com/gdamore/tcell/encoding"
	"os"
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
	sp := mario_go.NewMario()
	sp.SetX(10)
	sp.SetY(50)
	world.SetMario(sp)
	ground := mario_go.NewGround()
	ground.SetX(0)
	ground.SetY(0)
	ground2 := mario_go.NewGround()
	ground2.SetX(16)
	ground2.SetY(0)
	ground3 := mario_go.NewGround()
	ground3.SetX(32)
	ground3.SetY(0)
	ground4 := mario_go.NewGround()
	ground4.SetX(60)
	ground4.SetY(10)
	ground5 := mario_go.NewGround()
	ground5.SetX(100)
	ground5.SetY(24)
	world.AddObject(ground)
	world.AddObject(ground2)
	world.AddObject(ground3)
	world.AddObject(ground4)
	world.AddObject(ground5)
	world.Draw()

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
					sp.SetY(sp.Y() - 1)
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
		case <-time.After(time.Millisecond * 25):
		}
		st := tcell.StyleDefault.Background(tcell.NewHexColor(0x6AADFD))
		s.Fill(' ', st)
		//s.Clear()
		sp.Move()
		world.HitTest()
		world.CameraX = -sp.X() + 30
		if world.CameraX > 0 {
			world.CameraX = 0
		}
		world.Draw()
		s.Show()
	}

	s.Fini()
	os.Exit(0)
}
