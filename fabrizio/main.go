package main

import (
	"fmt"

	"github.com/GFG/hackday_n2o/fabrizio/lib/csv"
	"github.com/GFG/hackday_n2o/fabrizio/lib/excel"
	"github.com/GFG/hackday_n2o/fabrizio/lib/validator"
	"github.com/GFG/hackday_n2o/fabrizio/srv"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

type (
	numberRule struct {
		originalValue string
		value float32
		mandatory bool
		min float32
		max float32
	}
)

func main() {

	num := numberRule{}
	fmt.Println(num)

	fieldValidator := validator.NewValidator()

	excelReader := excel.NewReader()
	csvReader := csv.NewReader()

	server := gin.Default()

	server.Use(static.Serve("/", static.LocalFile("./static", false)))
	srv.UploadHandler(server, fieldValidator, excelReader, csvReader)
	srv.DownloadHandler(server)

	fmt.Println()

	fmt.Println("Starting on port :1234")
	err := server.Run(":1234")

	if err != nil {
		fmt.Println(err.Error())
	}
}
