package main

import (
	"fmt"

	"github.com/hellskater/benchmark-cli/pkg/data"
	"github.com/hellskater/benchmark-cli/pkg/sorting"
)

func main() {
	stockData, err := data.LoadStockData("stock_data.csv")

	if err != nil {
		panic(err)
	}

	results := sorting.Benchmark(stockData)

	for _, result := range results {
		fmt.Println(result)
	}
}
