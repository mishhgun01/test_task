package pkg

import (
	"strconv"
	"time"
)

// helpers - вспомогательные функции, в частности, для подсчета статистики

func statistics(data []IncomingData) (gasStatistics, error) {
	var currentMonth string
	var currentDay string
	var currentHour string
	var output gasStatistics
	var currGasVal float64
	var currGasPrice float64
	var aGasVal float64
	var aGasPrice float64
	var pricePerHour float64
	var mapOfHourPrices = make(map[string][]float64)
	output.wastePerMonth = make(map[string]float64)
	output.pricePerHour = make(map[string]float64)
	output.averagePricePerDay = make(map[string]float64)
	for _, item := range data {
		date, err := time.Parse("06-01-02 15:04", item.Time)
		if err != nil {
			return gasStatistics{}, err
		}
		// поторачено за месяц
		if date.Month().String() != currentMonth {
			currentMonth = date.Month().String()
			currGasVal = item.GasValue
			output.wastePerMonth[currentMonth] = currGasVal
		} else {
			output.wastePerMonth[currentMonth] = currGasVal - item.GasValue
		}
		// средняя цена за день
		if date.Month().String()+strconv.Itoa(date.Day()) != currentDay {
			currentDay = date.Month().String() + strconv.Itoa(date.Day())
			currGasPrice = item.GasPrice
			output.averagePricePerDay[currentDay] = currGasPrice
		} else {
			currGasPrice += item.GasPrice
			sum := currGasPrice / float64(daysOfMonth(date.Month().String()))
			output.averagePricePerDay[currentDay] = sum
		}
		if strconv.Itoa(date.Hour()) != currentHour {
			currentHour = strconv.Itoa(date.Hour())
			pricePerHour = item.GasPrice
		} else {
			mapOfHourPrices[currentHour] = append(mapOfHourPrices[currentHour], pricePerHour)
		}
		aGasPrice += item.GasPrice
		aGasVal += item.GasValue
	}
	for k, v := range mapOfHourPrices {
		output.pricePerHour[k] = averageOfFloats(v)
	}
	output.paid = aGasVal * aGasVal
	return output, nil
}

func daysOfMonth(month string) int {
	var days int
	switch month {
	case "January":
		days = 31
	case "February":
		days = 28
	case "March":
		days = 31
	case "April":
		days = 30
	case "May":
		days = 31
	case "June":
		days = 30
	case "July":
		days = 31
	case "August":
		days = 31
	case "September":
		days = 30
	case "October":
		days = 31
	case "November":
		days = 30
	case "December":
		days = 31
	}
	return days
}

func averageOfFloats(slice []float64) float64 {
	var sum float64
	for _, v := range slice {
		sum += v
	}
	return sum / float64(len(slice))
}
