package backend

import (
	"testing"
)

func TestSearchGoogleImage(t *testing.T) {
	// Skip online test in CI or if network is not reliable, 
	// but for our test bed, we want to see it works.
	query := "Taylor Swift Speak Now"
	results, err := SearchGoogleImage(query)
	if err != nil {
		t.Fatalf("Failed to search google images: %v", err)
	}

	if len(results) == 0 {
		t.Errorf("Expected some results, got 0")
	}

	for i, res := range results {
		if res.URL == "" {
			t.Errorf("Result %d has empty URL", i)
		}
		t.Logf("Result %d: %s", i+1, res.URL)
	}
}

func TestDecodeUnicode(t *testing.T) {
	input := "http://example.com/a\\u003db&amp;c"
	expected := "http://example.com/a=b&c"
	output := decodeUnicode(input)
	if output != expected {
		t.Errorf("Expected %s, got %s", expected, output)
	}
}
