package main

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func handleSchema(w http.ResponseWriter, r *http.Request) {
	filePath := r.URL.Path

	if !strings.HasSuffix(filePath, ".xml") {
		http.NotFound(w, r)
		return
	}

	fullPath := filepath.Join("schema", filePath)

	xmlData, err := os.ReadFile(fullPath)
	if err != nil {
		http.Error(w, "Failed to read XML file", http.StatusInternalServerError)
		return
	}

	var form Form
	err = xml.Unmarshal(xmlData, &form)
	if err != nil {
		http.Error(w, "Failed to parse XML file", http.StatusInternalServerError)
		return
	}

	form.schemaFile = filePath
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(form.generateHtmlDoc()))
}

func (form *Form) generateHtmlDoc() string {
	return fmt.Sprintf(`<!DOCTYPE html>
<html lang="en">
<head>
<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/@picocss/pico@2/css/pico.min.css" />
<meta charset="UTF-8">
<title>Simple Form</title>
</head>
<body><main class="container">%s</main></body>
</html>
`, form.generateHTML())
}

func (field *Field) generateHTML() string {
	html := fmt.Sprintf(`<label for="%s">%s</label>`, field.Name, field.Caption)
	switch field.FieldType {
	case "Select":
		html += fmt.Sprintf(`<select name="%s">`, field.Name)
		for _, label := range field.Labels {
			html += fmt.Sprintf(`<option value="%s">%s</option>`, label.Name, label.Content)
		}
		html += "</select>"
	case "TextBox":
		html += fmt.Sprintf(`<textarea name="%s"></textarea>`, field.Name)
	case "File":
		html += fmt.Sprintf(`<input type="file" name="%s" />`, field.Name)
	}
	return html
}

func (form *Form) generateHTML() string {
	html := fmt.Sprintf(`<form action="/submit-pdf-form?schema=%s" method="post">`, form.schemaFile)
	for _, field := range form.Fields {
		html += field.generateHTML()
	}
	for _, section := range form.Sections {
		html += fmt.Sprintf("<fieldset><legend>%s</legend>", section.Title)
		for _, field := range section.Contents.Fields {
			html += field.generateHTML()
		}
		html += "</fieldset>"
	}
	html += `<input type="submit" value="Submit" /></form>`
	return html
}
