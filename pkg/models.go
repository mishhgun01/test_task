package pkg

type gasStatistics struct {
	wastePerMonth      map[string]float64
	averagePricePerDay map[string]float64
	pricePerHour       map[string]float64
	paid               float64
}

type Request struct {
	Data map[string]map[string][]IncomingData
}

type IncomingData struct {
	Time           string  `json:"time"`
	GasPrice       float64 `json:"gasPrice"`
	GasValue       float64 `json:"gasValue"`
	Average        float64 `json:"average"`
	MaxGasPrice    float64 `json:"maxGasPrice"`
	MedianGasPrice float64 `json:"medianGasPrice"`
}
