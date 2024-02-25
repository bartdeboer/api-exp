package form

import (
	"encoding/xml"
	"os"
	"path/filepath"

	"github.com/bartdeboer/api-exp/internal/renderer/renderer"
)

func Load(form renderer.FormRenderer, schemaFile string) error {

	fullPath := filepath.Join("schema", schemaFile)

	xmlData, err := os.ReadFile(fullPath)
	if err != nil {
		return err
	}

	err = xml.Unmarshal(xmlData, &form)
	if err != nil {
		return err
	}

	form.AfterLoad(schemaFile)

	return nil
}
