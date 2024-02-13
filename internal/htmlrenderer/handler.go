package htmlrenderer

import (
	"encoding/xml"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/bartdeboer/api-exp/internal/form"
)

func HandleSchemaForm(w http.ResponseWriter, r *http.Request) {
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

	var form form.Form
	err = xml.Unmarshal(xmlData, &form)
	if err != nil {
		http.Error(w, "Failed to parse XML file", http.StatusInternalServerError)
		return
	}

	form.SchemaFile = filePath
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(generateHtmlDoc(&form)))
}
