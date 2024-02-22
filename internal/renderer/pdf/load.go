package pdfrenderer

import (
	"encoding/xml"
	"os"
	"path/filepath"
)

func (form *Form) Load(schemaFile string) error {

	fullPath := filepath.Join("schema", schemaFile)

	xmlData, err := os.ReadFile(fullPath)
	if err != nil {
		return err
	}

	err = xml.Unmarshal(xmlData, &form)
	if err != nil {
		return err
	}

	return nil
}
