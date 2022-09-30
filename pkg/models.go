package pkg

type GasStatistics struct {
	WastePerMonth      map[string]float64
	AveragePricePerDay map[string]float64
	PricePerHour       map[string]float64
	Paid               float64
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
