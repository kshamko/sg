package database

import (
	"encoding/json"
	"io"
	"os"

	"github.com/kshamko/sg/internal/predict"
)

type JSON struct {
	data []LTVDataJSON
}

func NewJSON(filePath string) (*JSON, error) {

	data := []LTVDataJSON{}

	jsonFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	bytes, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, &data)

	if err != nil {
		return nil, err
	}

	return &JSON{data}, nil
}

func (jsn *JSON) GetData() []predict.LTVData {

	res := make([]predict.LTVData, len(jsn.data))

	for i, row := range jsn.data {
		res[i] = predict.LTVData{
			Country:    row.Country,
			CampaignID: row.CampaignID,
			UserCount:  row.Users,
			LTVs: []float64{
				row.LTV1, row.LTV2, row.LTV3, row.LTV4,
				row.LTV5, row.LTV6, row.LTV7,
			},
		}
	}

	return res
}
