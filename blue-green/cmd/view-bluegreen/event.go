package main

import (
	"context"
	"log"

	"github.com/gdamore/tcell"
)

const (
	done   = "done"
	resize = "resize"
)

// Event represents a user or system event. It contains a Type and optionally
// the X and Y coordinates at which the event occurred.
type Event struct {
	Type string
	X    int
	Y    int
}

func eventLoop(ctx context.Context, s tcell.Screen, event chan<- Event) {
	for {
		select {
		case <-ctx.Done():
			return

		default:

			ev := s.PollEvent()
			switch ev := ev.(type) {

			case *tcell.EventResize:
				event <- Event{Type: resize}

			case *tcell.EventKey:
				if ev.Key() == tcell.KeyEsc || ev.Key() == tcell.KeyCtrlC || ev.Rune() == 'q' {
					log.Print("event: exit")
					event <- Event{Type: done}
					return
				}

			default:
				continue
			}

		}
	}
}
