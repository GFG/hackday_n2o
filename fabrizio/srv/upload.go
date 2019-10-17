package srv

import (
	"errors"
	"fmt"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

type (
	Reader interface {
		GetName() string
		Read(fileName string) ([][]string, error)
	}

	Validator interface {
		ValidateFields(fields [][]string) (bool, error)
		ValidateField(name string, value interface{}) (bool, error)
	}

	uploadHandler struct {
		validator Validator
		excelReader Reader
		csvReader Reader
	}
)

func UploadHandler(server *gin.Engine, validator Validator, excelReader Reader, csvReader Reader) {
	u := &uploadHandler{
		validator: validator,
		excelReader: excelReader,
		csvReader: csvReader,
	}

	server.POST("/upload", u.Upload)
}

func (u *uploadHandler) Upload(c *gin.Context) {
	t1 := time.Now()
	fmt.Println("Upload working")

	// Source
	file, err := c.FormFile("fileToUpload")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}

	filename := filepath.Base(file.Filename)
	if err := c.SaveUploadedFile(file, filename); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
		return
	}

	reader, err := u.getReaderByType(filename)

	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}


	rows, err := reader.Read(filename)

	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	//printRows(rows)

	//isValid, err := u.validator.ValidateFields(rows)
	//
	//if err != nil {
	//	c.String(http.StatusBadRequest, err.Error())
	//	return
	//}


	t2 := time.Now()

	diff := t2.Sub(t1)

	c.JSON(http.StatusOK, gin.H{
		"filename": file.Filename,
		"extension": filepath.Ext(file.Filename),
		"rows": len(rows),
		"elapsed": diff.String(),
		"reader": reader.GetName(),
		"valid": true,
	})
}

func (u *uploadHandler) getReaderByType(filename string) (Reader, error) {
	ext := filepath.Ext(filename)

	if ext == ".xlsx" {
		return u.excelReader, nil
	}

	if ext == ".csv" {
		return u.csvReader, nil
	}

	return nil, errors.New(fmt.Sprintf("unsupported extension: %s", ext))
}

func printRows(rows [][]string) {
	var product map[string]string
	headers := rows[0]

	for k, row := range rows {
		if k == 0 {
			continue
		}

		product = map[string]string{}

		for pos, value := range row {
			product[headers[pos]] = value
		}

		fmt.Println(product)
	}
}


