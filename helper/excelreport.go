package helper

import (
	_ "Skripsi/gislus/webapp/models"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/tealeg/xlsx"
)

func GenerateExcel() {
	//mdl := make([]models.GenerationModel, 0)
	excelFileName, _ := filepath.Abs("assets/docs/template.xlsx")
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		fmt.Println(err.Error())
	}

	t := time.Now()

	sheet := xlFile.Sheets[0]
	sheet.Cell(1, 1).SetValue("Report Generated on " + t.Format("02 Jan 2006 15:04:05"))

	pathToSave, _ := filepath.Abs("assets/docs/attachments")
	os.MkdirAll(pathToSave, 0777)

	fileAttach := pathToSave + "/" + "Report_" + t.Format("20060102_1504") + ".xlsx"

	err = xlFile.Save(fileAttach)
	if err != nil {
		fmt.Println(err.Error())
	}
}
