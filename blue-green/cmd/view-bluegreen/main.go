package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/gdamore/tcell"
)

func main() {
	u := flag.String("url", "", "the URL to query (required)")
	flag.Parse()

	if *u == "" {
		flag.Usage()
		os.Exit(1)
	}

	p, err := url.Parse(*u)
	if err != nil {
		fmt.Fprintln(os.Stderr, "unable to parse url:", err)
		flag.Usage()
		os.Exit(1)
	}

	s, err := tcell.NewScreen()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	if err := s.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	defer s.Fini()

	width, height := s.Size()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	events := make(chan Event)
	go eventLoop(ctx, s, events)

	ticker := time.NewTicker(50 * time.Millisecond)
	defer ticker.Stop()

	i := 0
	width--
	maxI := width * height

	for {
		select {
		case ev := <-events:
			switch ev.Type {
			case resize:
				width, height = s.Size()
				width--
				maxI = width * height
			case done:
				return
			}
		case <-ticker.C:
			if i > maxI {
				i = 0
			}

			x := i % width
			y := i / width

			s.SetContent(x, y, '◼', nil, tcell.StyleDefault.Foreground(tcell.ColorWhite))
			s.Show()

			go ping(ctx, p, s, x, y)

			i++
		}
	}
}

func ping(ctx context.Context, u *url.URL, s tcell.Screen, x, y int) {
	defer s.Show()

	client := &http.Client{Timeout: 10 * time.Second}
	r, err := client.Get(u.String())
	if err != nil {
		s.SetContent(x, y, '◼', nil, tcell.StyleDefault.Foreground(tcell.ColorRed))
		return
	}
	defer r.Body.Close()

	switch r.Header.Get("Deployment-Colour") {
	case "blue":
		s.SetContent(x, y, '◼', nil, tcell.StyleDefault.Foreground(tcell.ColorBlue))
	case "green":
		s.SetContent(x, y, '◼', nil, tcell.StyleDefault.Foreground(tcell.ColorGreen))
	default:
		fmt.Fprintf(os.Stderr, "unable to determine deployment colour")
		s.SetContent(x, y, '◼', nil, tcell.StyleDefault.Foreground(tcell.ColorRed))
	}
}
