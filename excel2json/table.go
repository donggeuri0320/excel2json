package excel2json

import (
	"fmt"
	"log"
	"strings"
)

type Table struct {
	TitleFormat string
	RowFormat string
	ColTypes []*string
	Rows []*RowData
}

func (d *Table) Init(titleStr string, subTitleArr []string, colTypes []string) {
	d.ColTypes = make([]*string,0,len(colTypes))
	for i:=0; i<len(colTypes); i++ {
		d.ColTypes = append(d.ColTypes, &colTypes[i])
	}
	d.Rows = make([]*RowData,0)
}

func (d *Table) ToString() string {
	lenRows := len(d.Rows)
	sb := strings.Builder{}
	for i, v := range d.Rows {
		sb.WriteString(v.ToString(&d.RowFormat))
		if i < lenRows-1 {
			sb.WriteString(",\r\n\t")
		}
	}
	return fmt.Sprintf(d.TitleFormat, sb.String())
}

// 행 저장
func (d *Table) AddRow(cols []string) {
	rd := &RowData{}
	rd.Cols = make([]*ColData,0, len(cols))

	for i, col := range cols {
		rd.Cols = append(rd.Cols, &ColData{Type: d.ColTypes[i], Value: col})
	}
	d.Rows = append(d.Rows, rd)
}

func (d *Table) SetFormat(titleStr string, subTitleArr []string, colTypes []string) {
	bKey, bValue := false, false
	lenSubTitleArr := len(subTitleArr)
	if lenSubTitleArr > 1 {
		if subTitleArr[0] == "$key" {
			bKey = true
			if colTypes[0] != "s" {
				log.Fatal("$key data define type is must `s`")
			}
		}
		if subTitleArr[1] == "$value" {
			bValue = true
		}
	}

	if lenSubTitleArr > 2 {
		for i:=2; i<lenSubTitleArr; i++ {
			if bKey && bValue {
				log.Fatal("`", titleStr, "` data type define error! $key is first, $value is second!")
			}
			if subTitleArr[i] == "$key" || subTitleArr[i] == "$value" {
				log.Fatal("`", titleStr, "` data type define error! $key is first, $value is second!")
			}
		}
	}

	if bKey && bValue {
		d.TitleFormat = fmt.Sprintf("\"%s\"", titleStr) + ":{\r\n\t%s\r\n\t}"
		d.RowFormat = "\t%s: %s"
	} else if bKey {
		d.TitleFormat = fmt.Sprintf("\"%s\"", titleStr) + ":{\r\n\t%s\r\n\t}"
		sb := strings.Builder{}
		sb.WriteString("\t%s:{")
		for i, v := range subTitleArr {
			if i == 0 {
				continue
			}
			sb.WriteString(fmt.Sprintf("\"%s\":", v) + "%s")
			if i < len(subTitleArr)-1 {
				sb.WriteString(", ")
			}
		}
		sb.WriteString("}")
		d.RowFormat = sb.String()
	} else {
		d.TitleFormat = fmt.Sprintf("\"%s\"", titleStr) + ":[\r\n\t%s\r\n\t]"
		sb := strings.Builder{}
		sb.WriteString("\t{")
		for i, v := range subTitleArr {
			sb.WriteString(fmt.Sprintf("\"%s\":", v) + "%s")
			if i < len(subTitleArr)-1 {
				sb.WriteString(", ")
			}
		}
		sb.WriteString("}")
		d.RowFormat = sb.String()
	}

	d.ColTypes = make([]*string,0,len(colTypes))
	for i:=0; i<len(colTypes); i++ {
		d.ColTypes = append(d.ColTypes, &colTypes[i])
	}
	d.Rows = make([]*RowData,0)
}