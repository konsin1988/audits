package exporters

import (
    "bytes"
    "fmt"
    "time"

    "github.com/phpdave11/gofpdf"

    "konsin1988/go-server/models"
)

type PdfExporter struct{}

func (e *PdfExporter) Export(data models.TableData) ([]byte, string, error){

    var buf bytes.Buffer

	pdf := gofpdf.New("L", "mm", "A4", "")
	pdf.SetAutoPageBreak(true, 20)
	pdf.AddPage()
    pdf.AddUTF8Font("Times", "", "fonts/Times_New_Roman.ttf")
    pdf.AddUTF8Font("TimesBold", "", "fonts/Times_New_Roman_Bold.ttf")

	pdf.SetFont("TimesBold", "", 18)
	pdf.CellFormat(0, 10, data.Title, "", 1, "C", false, 0, "")
    pdf.Ln(8)
    
	pdf.SetFont("Times", "", 14)
	creationDate := time.Now().Format("02.01.2006 15:04")
	dateText := fmt.Sprintf("Дата создания: %s", creationDate)
	pdf.CellFormat(0, 10, dateText, "", 1, "L", false, 0, "")

	pdf.Ln(5)

	if len(data.Rows) > 0 {
		e.createTable(pdf, data)
	}
    err := pdf.Output(&buf)
	if err != nil {
		return buf.Bytes(), "", err
	}
	
    return buf.Bytes(), "application/pdf", nil
}

func (e *PdfExporter) createTable(pdf *gofpdf.Fpdf, data models.TableData) {
	margin := 10.0

	// Calculate column widths
	colCount := len(data.Headers) - 1
	pageWidth, _ := pdf.GetPageSize()
	availableWidth := pageWidth - 2*margin - 17
	colWidth := availableWidth / float64(colCount)

	// Draw table header
	pdf.SetLineWidth(0.3)

	// Header cells
    drawRow(pdf, data.Headers, colWidth, 5, "C", true)

	// Draw table rows
	for i := 0; i < len(data.Rows); i++ {
        drawRow(pdf, data.Rows[i], colWidth, 5, "C", false)
	}
}

func drawRow(
    pdf *gofpdf.Fpdf,
    cells []string,
    colWidth float64,
    lineHeight float64,
    align string,
    bold bool,
) {
    x, y := pdf.GetXY()
	pdf.SetTextColor(0, 0, 0)
    if bold {
        pdf.SetFont("TimesBold", "", 13)
    } else {
        pdf.SetFont("Times", "", 12)
    }
	evenRowColor := []int{250, 240, 232} 
	oddRowColor := []int{255, 255, 255} 

    maxHeight := 0.0
    for _, text := range cells {
        h := cellHeight(pdf, text, colWidth, lineHeight)
        if h > maxHeight {
            maxHeight = h
        }
    }
    var c_width float64
    for idx, text := range cells {
		if idx%2 == 0 {
			pdf.SetFillColor(evenRowColor[0], evenRowColor[1], evenRowColor[2])
		} else {
			pdf.SetFillColor(oddRowColor[0], oddRowColor[1], oddRowColor[2])
		}
        pdf.SetDrawColor(150,150,150)

        if idx == 0 {
            c_width = 17.0
        } else {
            c_width = colWidth
        }
        pdf.Rect(x, y, c_width, maxHeight, "DF")
        pdf.SetXY(x, y)
        pdf.MultiCell(c_width, lineHeight, text, "", align, false)

        x += c_width 
    }
    pdf.SetXY(pdf.GetX(), y + maxHeight)
    checkPageBreak(pdf, 6)
}

func cellHeight(pdf *gofpdf.Fpdf, text string, width, lineHeight float64) float64 {
    lines := pdf.SplitLines([]byte(text), width)
    return float64(len(lines)) * lineHeight
}

func checkPageBreak(pdf *gofpdf.Fpdf, height float64) {
    _, pageH := pdf.GetPageSize()
    _, _, _, bottom := pdf.GetMargins()

    if pdf.GetY() + height > pageH-bottom {
        pdf.AddPage()
    }
}
