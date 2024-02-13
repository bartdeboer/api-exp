package main

import (
	"encoding/xml"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/jung-kurt/gofpdf"
)

func handlePDFFormSubmission(w http.ResponseWriter, r *http.Request) {
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

	var form Form
	err = xml.Unmarshal(xmlData, &form)
	if err != nil {
		http.Error(w, "Failed to parse XML file", http.StatusInternalServerError)
		return
	}

	pdf := form.generatePDF(r)

	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", "attachment; filename=\"form_submission.pdf\"")

	pdf.Output(w)
}

func (field *Field) generatePDF(pdf *gofpdf.Fpdf, value string) {
	pdf.SetFont("Arial", "", 12)
	pdf.CellFormat(0, 10, field.Caption+": "+value, "0", 1, "", false, 0, "")
	pdf.Ln(10)
}

func (form *Form) generatePDF(r *http.Request) *gofpdf.Fpdf {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.CellFormat(0, 10, "Simple Form", "0", 1, "C", false, 0, "")
	pdf.Ln(20)

	if err := r.ParseForm(); err != nil {
		// Handle error
	}

	for _, field := range form.Fields {
		if value, ok := r.Form[field.Name]; ok {
			field.generatePDF(pdf, strings.Join(value, ", "))
		}
	}

	for _, section := range form.Sections {
		pdf.SetFont("Arial", "B", 14)
		pdf.CellFormat(0, 10, section.Title, "0", 1, "", false, 0, "")
		pdf.SetFont("Arial", "", 12)
		for _, field := range section.Contents.Fields {
			if value, ok := r.Form[field.Name]; ok {
				field.generatePDF(pdf, strings.Join(value, ", "))
			}
		}
	}

	return pdf
}
