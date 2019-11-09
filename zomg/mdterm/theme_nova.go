/*
Copyright © 2019 Kris Nova <kris@nivenly.com> 2019

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
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
	case ll, tt, rr:
		return termbox.ColorBlue | termbox.AttrBold, termbox.ColorDefault
	default:
		return termbox.ColorDefault, termbox.ColorDefault
	}
}
