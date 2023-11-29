package csv

import (
	"encoding/csv"
	"os"
)

type reader struct {
	filePath string
	file     *os.File
}

func (r *reader) Read() ([][]string, map[string]int, error) {

	csvReader := csv.NewReader(r.file)
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, nil, err
	}

	titleMap := make(map[string]int)
	for idx, title := range records[0] {
		titleMap[title] = idx
	}
	return records, titleMap, err
}

func Init(filePath string) (Reader, error) {

	f, err := os.Open(filePath)
	defer f.Close()

	return &reader{
		filePath: filePath,
		file:     f,
	}, err
}
