package main

import (
	"math/rand"
	"time"

	"github.com/nsf/termbox-go"
)

const animationSpeed = 10 * time.Millisecond

// Function main initializes termbox, renders the view, and starts
// handling events.
func main() {
	rand.Seed(time.Now().UnixNano())

	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	eventQueue := make(chan termbox.Event)
	go func() {
		for {
			eventQueue <- termbox.PollEvent()
		}
	}()

	g := NewGame()
	render(g)

	for {
		select {
		case ev := <-eventQueue:
			if ev.Type == termbox.EventKey {
				switch {
				case ev.Key == termbox.KeyArrowLeft:
					g.moveLeft()
				case ev.Key == termbox.KeyArrowRight:
					g.moveRight()
				case ev.Key == termbox.KeyArrowUp:
					g.rotate()
				case ev.Key == termbox.KeyArrowDown:
					g.moveDown()
				case ev.Key == termbox.KeySpace:
					g.fall()
				case ev.Ch == 's':
					g.start()
				case ev.Ch == 'p':
					g.pause()
				case ev.Ch == 'q' || ev.Key == termbox.KeyEsc || ev.Key == termbox.KeyCtrlC || ev.Key == termbox.KeyCtrlD:
					return
				}
			}
		case <-g.fallingTimer.C:
			g.play()
		default:
			render(g)
			time.Sleep(animationSpeed)
		}
	}
}
