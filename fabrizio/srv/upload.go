package srv

import (
	"errors"
	"fmt"
	"log"
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

const (
	inputUploadedFile = "fileToUpload"
	fileTypeExcel = ".xslx"
	fileTypeCSV = ".csv"
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
	log.Println("Upload working")

	file, err := c.FormFile(inputUploadedFile)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	filename := filepath.Base(file.Filename)
	if err := c.SaveUploadedFile(file, filename); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	reader, err := u.getReaderByType(filename)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}


	rows, err := reader.Read(filename)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

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

	if ext == fileTypeExcel {
		return u.excelReader, nil
	}

	if ext == fileTypeCSV {
		return u.csvReader, nil
	}

	return nil, errors.New(fmt.Sprintf("unsupported extension: %s", ext))
}



