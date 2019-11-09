package mdterm

import "github.com/nsf/termbox-go"

const NovaTheme ThemeName = "nova"

type ThemeNova struct {
}

func (t *ThemeNova) AttrsByElement(e Element) (termbox.Attribute, termbox.Attribute) {
	switch e.Label() {
	case t1:
		return termbox.ColorGreen | termbox.AttrBold | termbox.AttrUnderline, termbox.ColorDefault
	case t2:
		return termbox.ColorGreen | termbox.AttrBold, termbox.ColorDefault
	case t3:
		return termbox.ColorGreen, termbox.ColorDefault
	case nl:
		return termbox.ColorGreen | termbox.AttrBold, termbox.ColorWhite
	default:
		return termbox.ColorDefault, termbox.ColorDefault
	}
}
