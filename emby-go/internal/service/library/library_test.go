package library

import (
	"testing"
)

func TestGenerateItemID(t *testing.T) {
	id1 := generateItemID("/path/to/file.mp4")
	id2 := generateItemID("/path/to/file.mp4")

	if id1 == "" {
		t.Error("generateItemID returned empty string")
	}
	if id1 == id2 {
		t.Error("generateItemID should generate unique IDs")
	}
}

func TestScanResult(t *testing.T) {
	result := &ScanResult{
		TotalItemsFound: 10,
		NewItems:        5,
		UpdatedItems:    3,
		RemovedItems:    2,
		Errors:          0,
	}

	if result.TotalItemsFound != 10 {
		t.Errorf("expected TotalItemsFound 10, got %d", result.TotalItemsFound)
	}
	if result.NewItems != 5 {
		t.Errorf("expected NewItems 5, got %d", result.NewItems)
	}
}