package mdterm

import "github.com/nsf/termbox-go"

type ThemeName string

type Theme interface {
	AttrsByElement(e Element) (termbox.Attribute, termbox.Attribute)
}
