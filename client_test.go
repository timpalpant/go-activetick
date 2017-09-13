package activetick

import (
	"encoding/csv"
	"os"
	"path/filepath"
	"testing"
)

func loadCSVData(filename string) ([][]string, error) {
	f, err := os.Open(filepath.Join("testdata", filename))
	if err != nil {
		return nil, err
	}
	defer f.Close()

	reader := csv.NewReader(f)
	return reader.ReadAll()
}

func TestParseBarData(t *testing.T) {
	rows, err := loadCSVData("barDataResponse.csv")
	if err != nil {
		t.Error(err)
	}

	for _, row := range rows {
		_, err := parseBarData(row)
		if err != nil {
			t.Error(err)
		}
	}
}

func TestParseTickData(t *testing.T) {
	rows, err := loadCSVData("tickDataResponse.csv")
	if err != nil {
		t.Error(err)
	}

	for _, row := range rows {
		_, err := parseTickData(row)
		if err != nil {
			t.Error(err)
		}
	}
}
