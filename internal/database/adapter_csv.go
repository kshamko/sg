package database

import (
	"os"

	"github.com/gocarina/gocsv"
	"github.com/kshamko/sg/internal/predict"
)

type CSV struct {
	data []LTVDataCSV
}

func NewCSV(filePath string) (*CSV, error) {

	data := []LTVDataCSV{}

	csvFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer csvFile.Close()

	err = gocsv.UnmarshalFile(csvFile, &data)

	if err != nil {
		return nil, err
	}

	return &CSV{data}, nil
}

func (csv *CSV) GetData() []predict.LTVData {

	res := make([]predict.LTVData, len(csv.data))

	for i, row := range csv.data {
		res[i] = predict.LTVData{
			Country:    row.Country,
			CampaignID: row.CampaignID,
			UserCount:  1,
			LTVs: []float64{
				row.LTV1, row.LTV2, row.LTV3, row.LTV4,
				row.LTV5, row.LTV6, row.LTV7,
			},
		}
	}

	return res
}
