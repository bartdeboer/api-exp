package pdfrenderer

import (
	"encoding/xml"
	"net/http"
	"os"
	"path/filepath"

	"github.com/bartdeboer/api-exp/internal/form"
)

func HandlePDFFormSubmission(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method is not supported.", http.StatusMethodNotAllowed)
		return
	}

	schema := r.URL.Query().Get("schema")
	if schema == "" {
		http.Error(w, "Schema parameter is missing", http.StatusBadRequest)
		return
	}

	fullPath := filepath.Join("schema", schema)

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

	pdf := generateForm(&form, r)

	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", "attachment; filename=\"form_submission.pdf\"")

	pdf.Output(w)
}
