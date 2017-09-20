package sexpr

import (
	"fmt"
)

func Example_nil() {
	type sample struct {
		sampleMap   map[string][]string
		sampleSlice []string
	}
	data := sample{
		sampleMap: map[string][]string{
			"key1": nil,
			"key2": []string{
				"value2",
			},
		},
		sampleSlice: nil,
	}

	// Encode it
	encodedData, err := Marshal(data)
	if err != nil {
		return
	}
	fmt.Printf("%s", encodedData)

	// Output:
	// {
	//	"sampleMap": {
	//		"key2": [
	//			"value2"
	//		]
	//	}
	// }
}
