package excel

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/tealeg/xlsx"
)

type (
	Reader struct {}
	SlowerReader struct {}
)

func NewReader() *Reader {
	return &Reader{}
}

func (r *Reader) GetName() string {
	return "Excel reader"
}

func (r *Reader) Read(fileName string) ([][]string, error) {
	data, err := xlsx.FileToSlice(fileName)

	if err != nil {
		return nil, err
	}

	fmt.Println(fmt.Sprintf("Count: %d", len(data[0])))

	return data[0], nil
}

func NewSlowerReader() *SlowerReader {
	return &SlowerReader{}
}

func (fr *SlowerReader) GetName() string {
	return "Slower Excel reader"
}

func (fr *SlowerReader) Read(fileName string) ([][]string, error) {
	f, err := excelize.OpenFile(fileName)

	if err != nil {
		return nil, err
	}

	sheetName := f.GetSheetName(1)

	return f.GetRows(sheetName), nil
}
