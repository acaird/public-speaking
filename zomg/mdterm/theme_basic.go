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

const BasicTheme ThemeName = "basic"

type ThemeBasic struct {
}

func (t *ThemeBasic) AttrsByElement(e Element) (termbox.Attribute, termbox.Attribute) {
	return termbox.ColorDefault, termbox.ColorDefault
}
