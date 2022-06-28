package excel2json

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"math"
	"strconv"
	"strings"
)

func Parse(datas [][]string) []*Table {
	lenRows := len(datas)
	if lenRows == 0 {
		return nil
	}
	readline := -1
	lenCols := len(datas[0])
	titleStr := ""
	objs := make([]*Table, 0)
	lenDefineCols := 0
	lenSubTitlecols := 0
	var obj *Table
	for rIdx:=0; rIdx<lenRows; rIdx++ {
		if readline == -1 && datas[rIdx][0] == "" {
			continue
		}

		if readline == -1 {
			readline = 0
			obj = &Table{}
			objs = append(objs, obj)
			for i:=1; i<lenCols-1; i++ {

			}
		} else {
			readline++
		}

		if readline == 0 {
			lenDefineCols = 0
			lenSubTitlecols = 0
			for i:=1; i<lenCols; i++ {
				if datas[rIdx-1][i] != "" {
					lenDefineCols++
				}
				if datas[rIdx][i] != "" {
					lenSubTitlecols++
				}
			}

			if lenDefineCols != lenSubTitlecols {
				log.Fatal("`", titleStr, "` you do define data type!")
			}

			titleStr = datas[rIdx][0]

			// define data type check
			for i:=1; i<lenDefineCols+1; i++ {
				if datas[rIdx-1][i] == "" {
					log.Fatal("`", titleStr, "` you do define data type!")
				} else {
					switch datas[rIdx-1][i] {
					case Type_N:
					case Type_NARR:
					case Type_S:
					case Type_SARR:
					default:
						log.Fatal(titleStr, " wrong define data type!: ", datas[rIdx-1][i])
					}
				}
			}
			
			obj.SetFormat(titleStr, datas[rIdx][1:lenDefineCols+1], datas[rIdx-1][1:lenDefineCols+1])
			continue
		}

		isExist := false
		// data type check
		for i:=1; i<lenDefineCols+1; i++ {
			if !isExist && datas[rIdx][i] != "" {
				isExist = true
			}

			if *obj.ColTypes[i-1] == Type_N && datas[rIdx][i] != "" {
				v, err := strconv.ParseFloat(datas[rIdx][i], 64)
				if err != nil {
					log.Fatal("`", titleStr, "`json data type error! row:", readline, ", col:", i, ", wrong data: ", datas[rIdx][i])
				}
				if strings.Contains(datas[rIdx][i], ".") {
					var temp float64 = math.Round(v * 1000.0)
					cv := fmt.Sprintf("%f", temp / 1000)
					dIdx := strings.Index(cv, ".")
					if dIdx + 4 < len(cv)-1 {
						datas[rIdx][i] = cv[:dIdx+5]
					} else {
						datas[rIdx][i] = cv
					}
				}
			} else if *obj.ColTypes[i-1] == Type_NARR && datas[rIdx][i] != "" {
				// temp := strings.ReplaceAll(, "~", "")
				temp := strings.ReplaceAll(datas[rIdx][i], " ", "")
				temp = strings.ReplaceAll(temp, "\r\n", "")
				tempArr := strings.Split(temp, ",")
				for _, d := range tempArr {
					if strings.Contains(d, "~") && !strings.Contains(d, ".") {
						tArr := strings.Split(d, "~")
						for _, tv := range tArr {
							_, err := strconv.ParseInt(tv, 10,64)
							if err != nil {
								log.Fatal("`", titleStr, "`json data type error! row:", readline, ", col:", i, ", wrong data: ", d)
							}
						}
					} else if strings.Contains(d, "."){
						_, err := strconv.ParseFloat(d, 64)
						if err != nil {
							log.Fatal("`", titleStr, "`json data type error! row:", readline, ", col:", i, ", wrong data: ", d)
						}
					} else {
						_, err := strconv.ParseInt(d, 10,64)
						if err != nil {
							log.Fatal("`", titleStr, "`json data type error! row:", readline, ", col:", i, ", wrong data: ", d)
						}
					}
				}
			}
		}

		if !isExist {
			readline = -1
			continue
		}

		colDatas := make([]string,lenDefineCols,lenDefineCols)
		for i:=0; i<lenDefineCols; i++ {
			colDatas[i] = datas[rIdx][i+1]
		}
		obj.AddRow(colDatas)

	}
	return objs
}