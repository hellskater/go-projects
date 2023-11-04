package converter

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
)

func ConvertJsonToCsv(input string, output string) error {
	// read the input file
	data, err := os.ReadFile(input)

	if err != nil {
		return err
	}

	// parse into a map
	var parsedData []map[string]interface{}
	if err := json.Unmarshal(data, &parsedData); err != nil {
		return err
	}

	// open the output file
	file, err := os.Create(output)

	if err != nil {
		return err
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// write the header
	var header []string
	for key := range parsedData[0] {
		header = append(header, key)
	}

	if err := writer.Write(header); err != nil {
		return err
	}

	// write the data
	for _, item := range parsedData {
		var row []string
		for _, key := range header {
			value := item[key]
			row = append(row, fmt.Sprintf("%v", value))
		}
		if err := writer.Write(row); err != nil {
			return err
		}
	}

	return nil
}
