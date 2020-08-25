package excel2json

import "fmt"

type RowData struct {
	Cols []*ColData
}

func (d *RowData) ToString(format *string) string {
	colStrArr := make([]interface{},0, len(d.Cols))
	for _, c := range d.Cols {
		colStrArr = append(colStrArr, c.ToString())
	}
	return fmt.Sprintf(*format, colStrArr...)
}