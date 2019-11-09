package mdterm

// Element is an element in a markdown file that is terminated by a \n at the end
type Element interface {
	Label() ElementType
	RawString() string
}

type ElementType string

const (
	t1 ElementType = "# "
	t2 ElementType = "## "
	t3 ElementType = "### "
	t4 ElementType = "#### "
	nl ElementType = "//"
	c  ElementType = "/!"
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
