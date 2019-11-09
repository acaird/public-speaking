package mdterm

import "github.com/nsf/termbox-go"

const NovaTheme ThemeName = "nova"

type ThemeNova struct {
}

func (t *ThemeNova) AttrsByElement(e Element) (termbox.Attribute, termbox.Attribute) {
	switch e.Label() {
	case t1:
		return termbox.ColorBlue | termbox.AttrBold | termbox.AttrUnderline, termbox.ColorWhite
	case t2:
		return termbox.ColorGreen | termbox.AttrBold, termbox.ColorDefault
	case t3:
		return termbox.ColorBlue, termbox.ColorDefault
	default:
		return termbox.ColorGreen, termbox.ColorDefault
	}
}
