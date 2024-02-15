package pdfrenderer

import (
	"github.com/bartdeboer/api-exp/internal/form"
)

// Some overrides so that our local structs get loaded allowing us to extend logic here

type Form struct {
	form.Form
	Fields   []Field   `xml:"Field"`
	Sections []Section `xml:"Section"`
}

type Field struct {
	form.Field
}

type Section struct {
	form.Section
	Contents Contents `xml:"Contents"`
}

type Contents struct {
	form.Contents
	Fields []Field `xml:"Field"`
}
