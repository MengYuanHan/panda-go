package util

import (
	"github.com/extrame/xls"
	"github.com/tealeg/xlsx"
	"reflect"
	"strings"
)

// ReadXlsx xlsx解析
func ReadXlsx(sheet *xlsx.Sheet) (res [][]string) {
	temp := make([][]string, len(sheet.Rows))
	for k, row := range sheet.Rows {
		var data []string
		for _, cell := range row.Cells {
			if strings.Index(cell.NumFmt, "yy") > -1 {
				rv := reflect.ValueOf(*cell)
				date1904 := rv.FieldByName("date1904").Bool()
				dateTime, timeErr := cell.GetTime(date1904)
				if timeErr != nil {
					data = append(data, cell.Value)
				} else {
					data = append(data, dateTime.Format("2006-01-02 15:04:05"))
				}
			} else {
				data = append(data, cell.Value)
			}
		}
		temp[k] = data
	}
	res = append(res, temp...)
	return res
}

// ReadXls xls解析
func ReadXls(sheet *xls.WorkSheet) (res [][]string) {
	if sheet.MaxRow == 0 {
		return res
	}
	temp := make([][]string, sheet.MaxRow+1)
	for i := 0; i <= int(sheet.MaxRow); i++ {
		row := sheet.Row(i)
		data := make([]string, 0)
		if row.LastCol() > 0 {
			for j := 0; j < row.LastCol(); j++ {
				col := row.Col(j)
				data = append(data, col)
			}
			temp[i] = data
		}
	}
	res = append(res, temp...)
	return res
}

// ExecList2map list转map
func ExecList2map(res [][]string) (result []map[string]string) {
	if len(res) < 2 {
		return result
	}
	//第一行为标题
	title := res[0]
	titleMap := make(map[int]string)
	for index, val := range title {
		if val != "" {
			titleMap[index] = val
		}
	}
	for i := 1; i < len(res); i++ {
		row := res[i]
		rowMap := make(map[string]string)
		for index, val := range row {
			if titleMap[index] != "" {
				sVal := strings.ToLower(val)
				// 处理Excel网址超链接
				if strings.Index(sVal, "(http") > -1 {
					sCol := strings.Split(sVal, "(http")
					val = sCol[0]
				}
				// 处理Excel邮箱超链接
				if strings.Index(sVal, "(mailto") > -1 {
					sCol := strings.Split(sVal, "(mailto")
					val = sCol[0]
				}
				rowMap[titleMap[index]] = val
			}
		}
		result = append(result, rowMap)
	}
	return result
}
