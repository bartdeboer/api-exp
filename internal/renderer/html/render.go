package htmlrenderer

import (
	"fmt"
	"net/http"
)

func (form *Form) Output(r *http.Request, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(form.GenerateHTMLDoc()))
}

func (form *Form) GenerateHTMLDoc() string {
	return fmt.Sprintf(`<!DOCTYPE html>
<html lang="en">
<head>
<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/@picocss/pico@2/css/pico.min.css" />
<meta charset="UTF-8">
<title>Simple Form</title>
</head>
<body><main class="container">%s</main></body>
</html>
`, form.GenerateHTML())
}

func (field *Field) GenerateHTML() string {
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

func (form *Form) GenerateHTML() string {
	html := fmt.Sprintf(`<form action="/handle-form?schema=%s&renderer=pdf" method="post">`, form.SchemaFile)
	for _, field := range form.Fields {
		html += field.GenerateHTML()
	}
	for _, section := range form.Sections {
		html += fmt.Sprintf("<fieldset><legend>%s</legend>", section.Title)
		for _, field := range section.Contents.Fields {
			html += field.GenerateHTML()
		}
		html += "</fieldset>"
	}
	html += `<input type="submit" value="Submit" /></form>`
	return html
}
