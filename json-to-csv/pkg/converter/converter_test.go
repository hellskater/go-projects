package converter

import (
	"encoding/csv"
	"encoding/json"
	"os"
	"slices"
	"testing"
)

func TestConvertJsonToCsv(t *testing.T) {
	input := "input.json"
	output := "output.csv"

	err := ConvertJsonToCsv(input, output)

	if err != nil {
		t.Fatalf("Error converting json to csv: %v", err)
	}

	// verify the header
	data, err := os.ReadFile(input)

	if err != nil {
		t.Errorf("Error reading input file: %v", err)
	}

	var parsedData []map[string]interface{}
	if err := json.Unmarshal(data, &parsedData); err != nil {
		t.Errorf("Error parsing json: %v", err)
	}

	file, err := os.Open(output)

	if err != nil {
		t.Errorf("Error opening output file: %v", err)
	}

	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()

	if err != nil {
		t.Errorf("Error reading csv: %v", err)
	}

	header := records[0]

	// for every key of parsedData[0], there should be a corresponding headeritem in the array
	for key := range parsedData[0] {
		if isPresent := slices.Contains(header, key); !isPresent {
			t.Errorf("Expected header to contain %v", key)
		}
	}

	t.Log("Test successful!")
}
