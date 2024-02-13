package main

import (
	"encoding/xml"
)

type Form struct {
	XMLName    xml.Name  `xml:"Form"`
	Fields     []Field   `xml:"Field"`
	Sections   []Section `xml:"Section"`
	schemaFile string
}

type Field struct {
	Name      string  `xml:"Name,attr"`
	Type      string  `xml:"Type,attr"`
	Optional  string  `xml:"Optional,attr"`
	FieldType string  `xml:"FieldType,attr"`
	Caption   string  `xml:"Caption"`
	Labels    []Label `xml:"Labels>Label"`
}

type Label struct {
	Name    string `xml:"Name,attr"`
	Content string `xml:",chardata"`
}

type Section struct {
	Name     string   `xml:"Name,attr"`
	Optional string   `xml:"Optional,attr"`
	Title    string   `xml:"Title"`
	Contents Contents `xml:"Contents"`
}

type Contents struct {
	Fields []Field `xml:"Field"`
}
