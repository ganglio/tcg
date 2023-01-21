package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/gdamore/tcell/v2"
	"github.com/msoap/tcg"
)

func main() {
	mode := tcg.Mode2x3
	flag.Var(&mode, "mode", "screen mode, one of 1x1, 1x2, 2x2, 2x3, 2x4Braille")
	flag.Parse()

	tg, err := tcg.New(mode)
	if err != nil {
		log.Fatal(err)
	}

	maxX, maxY := tg.Width, tg.Height
	x, y := maxX/2, maxY/2
	drawNext := func() {
		tg.Buf.Set(x, y, tcg.Black)
		tg.Show()
	}

	width, height := tg.ScreenSize()
	if err := tg.SetClip(0, 0, width, height-1); err != nil {
		tg.Finish()
		log.Fatalf("SetClip: %s", err)
	}

	if tg.TCellScreen.HasMouse() {
		tg.TCellScreen.EnableMouse(tcell.MouseMotionEvents)
	}

	var (
		mx, my int
	)

	tg.PrintStrStyle(18, height-1, " <q> ", tcell.StyleDefault.Background(tcell.ColorGray))
	tg.PrintStr(23, height-1, " Quit ")
	tg.PrintStrStyle(29, height-1, " <c> ", tcell.StyleDefault.Background(tcell.ColorGray))
	tg.PrintStr(34, height-1, " Clear ")

LOOP:
	for {
		ev := tg.TCellScreen.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			if ev.Rune() == 'q' {
				break LOOP
			}
			if ev.Rune() == 'c' {
				tg.Buf.Clear()
				tg.Show()
				continue LOOP
			}
			switch ev.Key() {
			case tcell.KeyDown:
				y++
				drawNext()
			case tcell.KeyUp:
				y--
				drawNext()
			case tcell.KeyLeft:
				x--
				drawNext()
			case tcell.KeyRight:
				x++
				drawNext()
			case tcell.KeyEscape:
				break LOOP
			}
		case *tcell.EventMouse:
			cx, cy := ev.Position()
			if cx > mx {
				x++
			}
			if cx < mx {
				x--
			}
			if cy > my {
				y++
			}
			if cy < my {
				y--
			}
			mx, my = cx, cy
			drawNext()
		}
		tg.PrintStrStyle(0, height-1, " Coord: ", tcell.StyleDefault.Background(tcell.ColorGray))
		tg.PrintStr(9, height-1, fmt.Sprintf("%3d x%3d", x, y))
	}

	tg.Finish()
}
