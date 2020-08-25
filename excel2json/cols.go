package excel2json

import (
	"fmt"
	"strconv"
	"strings"
)

var (
	Type_N = "n"
	Type_NARR = "narr"
	Type_S = "s"
	Type_SARR = "sarr"
)

type ColData struct {
	Type *string
	Value string
}

func (d *ColData) ToString() string {
	switch *d.Type {
	case Type_N:
		if d.Value == "" {
			return "0"
		}
		return d.Value
	case Type_NARR:
		if d.Value == "" {
			return "[]"
		}

		if strings.Contains(d.Value, "~") {
			tempArr := strings.Split(strings.ReplaceAll(d.Value, " ", ""), ",")
			lenTempArr := len(tempArr)
			for i:=0; i<lenTempArr; i++ {
				if strings.Contains(tempArr[i], "~") && !strings.Contains(tempArr[i],".") {
					tempStr := strings.Split(tempArr[i], "~")
					lenTempStr := len(tempStr)
					if lenTempStr != 2 {
						continue
					}
					beginNum, _ := strconv.ParseInt(tempStr[0], 10, 64)
					endNum, _ := strconv.ParseInt(tempStr[1], 10, 64)
					sb := strings.Builder{}
					for j:=beginNum; j<=endNum; j++ {
						sb.WriteString(fmt.Sprintf("%v", j))
						if j < endNum {
							sb.WriteString(",")
						}
					}
					tempArr[i] = sb.String()
				}
			}
			return fmt.Sprintf("[%s]", strings.Join(tempArr, ","))
		} else {
			return fmt.Sprintf("[%s]", strings.ReplaceAll(d.Value, " ", ""))
		}
	case Type_S:
		if d.Value == "" {
			return "\"\""
		}
		return fmt.Sprintf("\"%s\"", strings.ReplaceAll(d.Value, "\"", "\\\"") )
	case Type_SARR:
		if d.Value == "" {
			return "[]"
		}

		sArr := strings.Split(strings.ReplaceAll(d.Value, " ", ""), ",")
		lenSArr := len(sArr)
		for i:=0; i<lenSArr; i++ {
			sArr[i] = strings.ReplaceAll(sArr[i], "\"", "\\\"")
		}

		return fmt.Sprintf("[\"%s\"]", strings.Join(sArr, "\",\""))
	}
	return ""
}