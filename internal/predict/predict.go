package predict

import "fmt"

type LTVData struct {
	Country    string
	CampaignID string
	UserCount  uint
	LTVs       []float64
}

type Database interface {
	GetData() []LTVData
}

type Inferencer interface {
	PredictLTV(LTVs []float64, day int) float64
}

// Agregator has
type Aggregator interface {
	Collect(key string, value float64, count uint)
	Result() []AggregatorResult
}

type aggregateKeyFunc func(row LTVData) string

var aggregateRules = map[string]aggregateKeyFunc{
	"country": func(row LTVData) string {
		return row.Country
	},
	"campaign": func(row LTVData) string {
		return row.CampaignID
	},
}

type Predict struct {
	ds         Database
	model      Inferencer
	agg        Aggregator
	aggKeyFunc aggregateKeyFunc
}

func New(ds Database, mod Inferencer, agg Aggregator, aggRule string) (*Predict, error) {

	aggKeyFunc, ok := aggregateRules[aggRule]

	if !ok {
		return nil, fmt.Errorf("aggregation rule for %s not defined", aggRule)
	}

	return &Predict{
		ds:         ds,
		model:      mod,
		agg:        agg,
		aggKeyFunc: aggKeyFunc,
	}, nil
}

func (p *Predict) Predict(day int) []AggregatorResult {

	data := p.ds.GetData()

	for _, row := range data {

		revenue := []float64{}
		for _, ltv := range row.LTVs {
			if ltv == 0 {
				break
			}
			revenue = append(revenue, ltv*float64(row.UserCount))
		}

		predictedRev := p.model.PredictLTV(revenue, 60)

		p.agg.Collect(p.aggKeyFunc(row), predictedRev, row.UserCount)
	}

	return p.agg.Result()
}
