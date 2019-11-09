package mdterm

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/kris-nova/logger"
)

type Presentation struct {
	filepath       string
	slides         []*Slide
	theme          Theme
	renderedSlides []*RenderedSlide
}

type Slide struct {
	name         string
	content      []byte
	presentation *Presentation
}

func LoadPresentation(dir string) (*Presentation, error) {
	dir, err := filepath.Abs(dir)
	if err != nil {
		return nil, fmt.Errorf("unable to abs() dir: %v", dir)
	}
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("unable to parse dir (%s): %v", dir, err)
	}
	p := &Presentation{}
	logger.Info("number of files in (%s): %d", dir, len(files))
	if len(files) <= 1 {
		return nil, fmt.Errorf("unable to find enough files in (%s)", dir)
	}
	for _, f := range files {
		logger.Info("file: %s", f.Name())
		absf := filepath.Join(dir, f.Name())
		data, err := ioutil.ReadFile(absf)
		if err != nil {
			logger.Info("error parsing file: %v", err)
		}
		s := &Slide{
			name:         f.Name(),
			content:      data,
			presentation: p,
		}
		p.slides = append(p.slides, s)
	}
	return p, nil
}

func (p *Presentation) SetTheme(theme Theme) error {
	p.theme = theme
	return nil
}

func (p *Presentation) RenderPresentation() error {
	logger.Info("rendering %d slides", len(p.slides))
	for _, slide := range p.slides {
		pslide, err := RenderSlide(slide)
		if err != nil {
			return fmt.Errorf("unable to render slide %s: %v", slide.name, err)
		}
		p.renderedSlides = append(p.renderedSlides, pslide)
	}
	return nil
}

func (p *Presentation) DisplayPresentation() error {
	logger.Info("displaying slides")
	for _, rslide := range p.renderedSlides {
		err := rslide.DisplayWithOptions(&DisplayOptions{})
		if err != nil {
			return fmt.Errorf("unable to render slide %s: %v", rslide.slide.name, err)
		}
	}
	return nil
}
