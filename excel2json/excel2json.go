package excel2json

import (
	"strings"
	"sync"
)

type Excel2json struct {
	tables []*Table
	mtx *sync.Mutex
}

func (d *Excel2json) Init() {
	d.tables = make([]*Table,0)
	d.mtx = new(sync.Mutex)
}

func (d *Excel2json) AddTables(tables []*Table) {
	d.mtx.Lock()
	defer d.mtx.Unlock()
	d.tables = append(d.tables, tables...)
}

func (d *Excel2json) ToString() string {
	sb := strings.Builder{}
	sb.WriteString("{\r\n\t")
	lenTables := len(d.tables)

	for i:=0; i<lenTables; i++ {
		sb.WriteString(d.tables[i].ToString())
		if i < lenTables-1 {
			sb.WriteString(",\r\n\t")
		}
	}
	sb.WriteString("\r\n}")
	return sb.String()
}