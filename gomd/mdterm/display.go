package mdterm

import (
	"fmt"

	"github.com/kris-nova/logger"

	"github.com/nsf/termbox-go"
)

type DisplayOptions struct {
}

func (r *RenderedSlide) DisplayWithOptions(opt *DisplayOptions) error {
	err := termbox.Init()
	if err != nil {
		return fmt.Errorf("unable to init termbox: %v", err)
	}
	defer termbox.Close()

	err = termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	if err != nil {
		return fmt.Errorf("unable to clear screen: %v", err)
	}

	for i, e := range r.elements {
		err := r.DisplayCenteredMessage(e.RawString(), i)
		if err != nil {
			return fmt.Errorf("failed to display slide: %v", err)
		}
	}
	termbox.Flush()
	event := termbox.PollEvent()
	switch event.Type {
	case termbox.EventKey:
		logger.Info("event [%v] (%v)", event.Key, event.Type)
		break
	case termbox.EventNone:
		break
	}
	return nil
}

func (r *RenderedSlide) DisplayCenteredMessage(msg string, lnum int) error {
	// Time for some arithmetic!
	// Horizontal Axis
	totalLineI := len(msg)
	yEmpty := r.maxy - totalLineI
	if yEmpty < 0 {
		return fmt.Errorf("slide %s too many horizontal chars (%d) max chars (%d)", r.slide.name, (yEmpty * -1), totalLineI)
	}
	yStart := yEmpty/2 + ((yEmpty % 2) / 2)
	// Vertical Axis
	totalVertI := len(r.elements)
	xEmpty := r.maxx - totalVertI
	if xEmpty < 0 {
		return fmt.Errorf("slide %s too many vertical lines (%d) max lines (%d)", r.slide.name, (xEmpty * -1), totalVertI)
	}
	xStart := xEmpty/2 + (xEmpty % 2)

	for i, r := range msg {
		termbox.SetCell(i+yStart, lnum+xStart, r, termbox.ColorDefault, termbox.ColorDefault)
	}
	return nil
}
