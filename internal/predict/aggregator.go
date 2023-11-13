package predict

type aggregatorEntry struct {
	valuesSum float64
	count     uint
}

type AggregatorResult struct {
	Key   string
	Value float64
}

type DefaultAggregator struct {
	data map[string]aggregatorEntry
}

func NewDefaultAggregator() *DefaultAggregator {
	return &DefaultAggregator{
		data: map[string]aggregatorEntry{},
	}
}

func (agg *DefaultAggregator) Collect(key string, value float64, count uint) {

	entry := agg.data[key]
	agg.data[key] = aggregatorEntry{
		valuesSum: entry.valuesSum + value,
		count:     entry.count + count,
	}
}

func (agg *DefaultAggregator) Result() []AggregatorResult {
	res := make([]AggregatorResult, len(agg.data))
	i := 0
	for k, v := range agg.data {
		res[i] = AggregatorResult{
			Key:   k,
			Value: v.valuesSum / float64(v.count),
		}
		i++
	}
	return res
}
