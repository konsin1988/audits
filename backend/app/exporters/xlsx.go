package exporters

import (
    "fmt"

    "github.com/xuri/excelize/v2"

    "konsin1988/go-server/models"
)

type XlsxExporter struct{}

func (e *XlsxExporter) Export(data models.TableData) ([]byte, string, error){

    f := excelize.NewFile()
    sheet := data.Title 
    f.SetSheetName("Sheet1", sheet) 

    titleStyle, _ := f.NewStyle(&excelize.Style{
        Font: &excelize.Font{Bold: true, Size: 14},
        Alignment: &excelize.Alignment{
            Horizontal: "center",
            Vertical: "center",
        },
    })

    headerStyle, _ := f.NewStyle(&excelize.Style{
        Font: &excelize.Font{Bold: true},
        Alignment: &excelize.Alignment{
            Horizontal: "center",
            Vertical: "center",
            WrapText: true,
        },
        Border: []excelize.Border{
            {Type: "left", Color: "000000", Style: 1},
            {Type: "right", Color: "000000", Style: 1},
            {Type: "top", Color: "000000", Style: 1},
            {Type: "bottom", Color: "000000", Style: 1},
        },
    })
    
    cellStyle, _ := f.NewStyle(&excelize.Style{
        Alignment: &excelize.Alignment{
            Vertical: "top",
            WrapText: true,
        },
        Border: []excelize.Border{
            {Type: "left", Color: "000000", Style: 1},
            {Type: "right", Color: "000000", Style: 1},
            {Type: "top", Color: "000000", Style: 1},
            {Type: "bottom", Color: "000000", Style: 1},
        },
    })

    lastCol, _ := excelize.ColumnNumberToName(len(data.Headers))
    f.MergeCell(sheet, "A1", lastCol+"1")
    f.SetCellValue(sheet, "A1", data.Title)
    f.SetCellStyle(sheet, "A1", lastCol+"1", titleStyle)

    headerRow := 2
    for i, h := range data.Headers {
        col, _ := excelize.ColumnNumberToName(i+1)
        cell := fmt.Sprintf("%s%d", col, headerRow)
        f.SetCellValue(sheet, cell, h)
        f.SetCellStyle(sheet, cell, cell, headerStyle)
    }

    startRow := headerRow + 1
    for r, row := range data.Rows {
        for c, value := range row {
            col, _ := excelize.ColumnNumberToName(c + 1)
            cell := fmt.Sprintf("%s%d", col, startRow+r)
            f.SetCellValue(sheet, cell, value)
            f.SetCellStyle(sheet, cell, cell, cellStyle)
        }
    }

    for i := range data.Headers {
        col, _ := excelize.ColumnNumberToName(i + 1)
        f.SetColWidth(sheet, col, col, 25)
    }

    buf, _ := f.WriteToBuffer()
    //if err != nil {
    //    return []byte{}, "", err
    //}
	
    return buf.Bytes(), "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", nil
}

