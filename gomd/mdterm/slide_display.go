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

	// Calculate the longest message
	highestLen := 0
	for _, ec := range r.elements {
		l := len(ec.RawString())
		if len(ec.RawString()) > highestLen {
			highestLen = l
		}
	}

	for i, e := range r.elements {
		// T's are center aligned
		// nl's are centered but aligned
		if e.Label() == nl {
			err := r.DisplayAlignedMessage(e, i, highestLen)
			if err != nil {
				return fmt.Errorf("failed to display slide: %v", err)
			}
		} else {
			err := r.DisplayCenteredMessage(e, i)
			if err != nil {
				return fmt.Errorf("failed to display slide: %v", err)
			}
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

func (r *RenderedSlide) DisplayAlignedMessage(e Element, lnum, longestMsg int) error {
	msg := e.RawString()
	// Time for some arithmetic!
	// Horizontal Axis
	yEmpty := r.maxy - longestMsg
	if yEmpty < 0 {
		return fmt.Errorf("slide %s too many horizontal chars (%d) max chars (%d)", r.slide.name, (yEmpty * -1), longestMsg)
	}
	yStart := yEmpty/2 + ((yEmpty % 2) / 2)
	// Vertical Axis
	totalVertI := len(r.elements)
	xEmpty := r.maxx - totalVertI
	if xEmpty < 0 {
		return fmt.Errorf("slide %s too many vertical lines (%d) max lines (%d)", r.slide.name, (xEmpty * -1), totalVertI)
	}
	xStart := xEmpty/2 + (xEmpty % 2)
	theme := r.slide.presentation.theme
	a1, a2 := theme.AttrsByElement(e)
	for i := 0; i < longestMsg; i++ {
		for i, r := range msg {
			termbox.SetCell(i+yStart, lnum+xStart, r, a1, a2)
			i++
		}
		termbox.SetCell(i+yStart, lnum+xStart, 32, a1, a2)
	}
	return nil
}

func (r *RenderedSlide) DisplayCenteredMessage(e Element, lnum int) error {
	// Time for some arithmetic!
	// Horizontal Axis
	msg := e.RawString()
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
	theme := r.slide.presentation.theme
	a1, a2 := theme.AttrsByElement(e)
	for i, r := range msg {
		termbox.SetCell(i+yStart, lnum+xStart, r, a1, a2)
	}
	return nil
}
