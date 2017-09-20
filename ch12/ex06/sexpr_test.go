package sexpr

import (
	"testing"
)

func Test(t *testing.T) {
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
		t.Fatalf("Marshal failed: %v", err)
	}
	t.Logf("\n%s\n", encodedData)
}
