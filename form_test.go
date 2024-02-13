package main

import (
	"encoding/xml"
	"os"
	"path/filepath"
	"testing"
)

func TestHandleSchema_Unmarshal(t *testing.T) {

	filePath := filepath.Join("schema", "form.xml")

	xmlData, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("Failed to read XML file: %v", err)
	}

	var form Form
	if err := xml.Unmarshal(xmlData, &form); err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}
}
