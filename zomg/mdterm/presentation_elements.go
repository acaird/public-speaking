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

// Element is an element in a markdown file that is terminated by a \n at the end
type Element interface {
	Label() ElementType
	RawString() string
}

type ElementType string

const (
	// t1 is text heading 1 and is for the highest level titles
	t1 ElementType = "# "

	// t2 is text heading 2 for smaller subheadings
	t2 ElementType = "## "

	// t3 is text heading 3 for yet even smaller subheadings
	t3 ElementType = "### "

	// t4 is text 4 for general paragraph text
	t4 ElementType = "#### "

	// nl is for raw text to be parsed as a block
	nl ElementType = "//"

	// c is for comments or empty lines (they are parsed the same)
	c ElementType = "/!"

	// ex is for code execution
	ex ElementType = "$ "

	// tt is for Titles
	tt ElementType = "title"

	// ll is for Left
	ll ElementType = "left"

	// rr is for right
	rr ElementType = "right"
)

type T1Element struct {
	label   ElementType
	rawLine string
}

func (t *T1Element) Label() ElementType {
	return t.label
}

func (t *T1Element) RawString() string {
	return t.rawLine
}

type T2Element struct {
	label   ElementType
	rawLine string
}

func (t *T2Element) Label() ElementType {
	return t.label
}

func (t *T2Element) RawString() string {
	return t.rawLine
}

type T3Element struct {
	label   ElementType
	rawLine string
}

func (t *T3Element) Label() ElementType {
	return t.label
}
func (t *T3Element) RawString() string {
	return t.rawLine
}

type T4Element struct {
	label   ElementType
	rawLine string
}

func (t *T4Element) Label() ElementType {
	return t.label
}

func (t *T4Element) RawString() string {
	return t.rawLine
}

type NLElement struct {
	label   ElementType
	rawLine string
}

func (t *NLElement) Label() ElementType {
	return t.label
}

func (t *NLElement) RawString() string {
	return t.rawLine
}

type CElement struct {
	label   ElementType
	rawLine string
}

func (t *CElement) Label() ElementType {
	return t.label
}

func (t *CElement) RawString() string {
	return t.rawLine
}

type TElement struct {
	label   ElementType
	rawLine string
}

func (t *TElement) Label() ElementType {
	return t.label
}

func (t *TElement) RawString() string {
	return t.rawLine
}

type LElement struct {
	label   ElementType
	rawLine string
}

func (t *LElement) Label() ElementType {
	return t.label
}

func (t *LElement) RawString() string {
	return t.rawLine
}

type RElement struct {
	label   ElementType
	rawLine string
}

func (t *RElement) Label() ElementType {
	return t.label
}

func (t *RElement) RawString() string {
	return t.rawLine
}

type EXElement struct {
	label   ElementType
	rawLine string
}

func (t *EXElement) Label() ElementType {
	return t.label
}

func (t *EXElement) RawString() string {
	return t.rawLine
}
