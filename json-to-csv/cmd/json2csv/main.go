package main

import (
	"flag"
	"fmt"

	"github.com/hellskater/json2csv/pkg/converter"
)

func main() {
	input := flag.String("input", "./pkg/converter/input.json", "The input file")
	output := flag.String("output", "output.csv", "The output file")

	flag.Parse()

	if *input == "" || *output == "" {
		fmt.Println("Input and output files must be specified")
		flag.Usage()
		return
	}

	err := converter.ConvertJsonToCsv(*input, *output)

	if err != nil {
		panic(err)
	}

	fmt.Println("Done!")
}
