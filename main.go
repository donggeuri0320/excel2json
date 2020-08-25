package main

import (
	e2j "donggeuri0320/excel2json/excel2json"
	"github.com/360EntSecGroup-Skylar/excelize"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"strings"
	"sync"
)

func main()  {
	runtime.GOMAXPROCS(runtime.NumCPU())

	files, err := ioutil.ReadDir("./xlsx")
	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup

	for _, f := range files {
		if !strings.HasSuffix(f.Name(), ".xlsx") {
			continue
		}

		wg.Add(1)
		go func(xlsxFileName string) {
			defer wg.Done()

			xf, err := excelize.OpenFile("./xlsx/" + xlsxFileName)
			if err != nil {
				log.Fatal("`",xlsxFileName,"` open error!")
			}

			excel2json := e2j.Excel2json{}
			excel2json.Init()

			for _, name := range xf.GetSheetMap() {
				data := xf.GetRows(name)
				tables := e2j.Parse(data)
				excel2json.AddTables(tables)
			}

			of, err := os.Create("./json/" + strings.ReplaceAll(xlsxFileName, ".xlsx", ".json"))
			of.WriteString(excel2json.ToString())
		}(f.Name())
	}
	wg.Wait()
	log.Println("Done!")
}