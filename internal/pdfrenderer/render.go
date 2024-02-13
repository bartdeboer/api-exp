package pdfrenderer

import (
	"net/http"
	"strings"

	"github.com/jung-kurt/gofpdf"

	"github.com/bartdeboer/api-exp/internal/form"
)

func generateField(field *form.Field, pdf *gofpdf.Fpdf, value string) {
	pdf.SetFont("Arial", "", 12)
	pdf.CellFormat(0, 10, field.Caption+": "+value, "0", 1, "", false, 0, "")
	pdf.Ln(10)
}

func generateForm(form *form.Form, r *http.Request) *gofpdf.Fpdf {
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
			generateField(&field, pdf, strings.Join(value, ", "))
		}
	}

	for _, section := range form.Sections {
		pdf.SetFont("Arial", "B", 14)
		pdf.CellFormat(0, 10, section.Title, "0", 1, "", false, 0, "")
		pdf.SetFont("Arial", "", 12)
		for _, field := range section.Contents.Fields {
			if value, ok := r.Form[field.Name]; ok {
				generateField(&field, pdf, strings.Join(value, ", "))
			}
		}
	}

	return pdf
}
