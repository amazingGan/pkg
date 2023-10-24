package excel

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
)

// 通用excel导出方法
func ExportToExcel(save bool, autoAdjustWidth bool, data [][]interface{}, sheetName string, heading []string) (res []byte, err error) {
	file := excelize.NewFile()
	headingStyle, err := file.NewStyle(`{"font":{"bold":true},"alignment":{"horizontal":"center"}}`)
	if err != nil {
		return
	}
	// 将中文表头写入第一行
	for columnIndex, headingValue := range heading {
		cell := excelize.ToAlphaString(columnIndex) + "1"
		file.SetCellValue(sheetName, cell, headingValue)
		file.SetCellStyle(sheetName, cell, cell, headingStyle)
	}
	// 写入数据
	for rowIndex, rowData := range data {
		for columnIndex, cellData := range rowData {
			cell := excelize.ToAlphaString(columnIndex) + fmt.Sprintf("%d", rowIndex+2)
			file.SetCellValue(sheetName, cell, cellData)
		}
	}

	// 自动调整列宽度
	if autoAdjustWidth {
		for columnIndex := range heading {
			column := excelize.ToAlphaString(columnIndex)
			maxColumnWidth := 0
			for _, rowData := range data {
				cellData := fmt.Sprintf("%v", rowData[columnIndex])
				cellWidth := len(cellData)
				if cellWidth > maxColumnWidth {
					maxColumnWidth = cellWidth
				}
				// 如果数据超过50个字符，将列宽度限制为50，您可以根据需要进行调整
				if maxColumnWidth > 50 {
					maxColumnWidth = 50
				}
			}
			file.SetColWidth(sheetName, column, column, float64(maxColumnWidth+2))
		}
	}

	if save {
		if err = file.SaveAs("output.xlsx"); err != nil {
			return
		}
	}
	buffer, err := file.WriteToBuffer()
	if err != nil {
		return
	}
	return buffer.Bytes(), nil
}
