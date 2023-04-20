package main

import "fmt"

func parseCSV(csv string) [][]string {
	var res [][]string
	var fields []string
	var field string
	var inQuotes bool

	for i := 0; i < len(csv); i++ {
		c := csv[i]

		if c == '"' {
			if inQuotes && i+1 < len(csv) && csv[i+1] == '"' {
				// 处理双引号的转义问题
				field += "\""
				i++
			} else {
				inQuotes = !inQuotes
			}
		} else if c == ',' && !inQuotes {
			fields = append(fields, field)
			field = ""
		} else if c == '\n' && !inQuotes {
			fields = append(fields, field)
			res = append(res, fields)
			fields = []string{}
			field = ""
		} else {
			field += string(c)
		}
	}

	if field != "" {
		fields = append(fields, field)
	}

	if len(fields) > 0 {
		res = append(res, fields)
	}

	return res
}

func main() {
	csv := `"Name","Age" "Tom",25 "Jerry","20"`

	res := parseCSV(csv)
	fmt.Println(res)
}
