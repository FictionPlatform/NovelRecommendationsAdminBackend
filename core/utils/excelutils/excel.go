package excelutils

import (
	"strings"
	"unicode"
)

// SanitizeCell 清洗单元格值，防止 Excel/CSV 公式注入：
// 对以 = + - @ 开头（允许前导空白/控制字符）的字符串内容前置单引号，非字符串原样返回
func SanitizeCell(v interface{}) interface{} {
	s, ok := v.(string)
	if !ok {
		return v
	}
	trimmed := strings.TrimLeftFunc(s, func(r rune) bool {
		return r == ' ' || r == '\t' || r == '\r' || r == '\n' || unicode.IsControl(r)
	})
	if trimmed != "" && strings.ContainsRune("=+-@", rune(trimmed[0])) {
		return "'" + s
	}
	return s
}

// SafeRow 构造经过公式注入清洗的表格行，用于 SetSheetRow
func SafeRow(vals ...interface{}) *[]interface{} {
	row := make([]interface{}, len(vals))
	for i, v := range vals {
		row[i] = SanitizeCell(v)
	}
	return &row
}
