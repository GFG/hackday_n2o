package csv

import (
	//"bufio"
	"encoding/csv"
	"io"
	"os"
)

type (
	Reader struct {}
)

func NewReader() *Reader {
	return &Reader{}
}

func (r *Reader) GetName() string {
	return "CSV reader"
}

func (r *Reader) Read(fileName string) ([][]string, error) {

	var rows [][]string

	// Open the file
	csvFile, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	// Parse the file
	reader := csv.NewReader(csvFile)
	//reader := csv.NewReader(bufio.NewReader(csvFile))

	reader.LazyQuotes = true
	reader.Comma = ';'
	reader.Comment = '#'

	// Iterate through the records
	for {
		// Read each record from csv
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		rows = append(rows, record)
	}

	return rows, nil
}
