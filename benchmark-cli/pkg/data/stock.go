package data

import (
	"encoding/csv"
	"os"
	"strconv"
)

func LoadStockData(fileName string) ([]int, error) {
	file, err := os.Open(fileName)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()

	if err != nil {
		return nil, err
	}

	var prices []int
	for _, record := range records {
		price, err := strconv.Atoi(record[1])
		if err != nil {
			continue
		}
		prices = append(prices, price)
	}

	return prices, nil
}
