package history

import (
	"reflect"
	"testing"
)
func TestMerge(t *testing.T) {
	h := History[int]{}
	h.AddOrUpdateLatestItem(9)

	h2 := History[int]{}

	h.AddOrUpdateLatestItem(12)

	h.Merge(h2)

	if h.Latest.Data != 12 {
		t.Errorf("Merge failed: Expected Latest Data to be %d, but got %d", 12, h.Latest.Data)
	}

	if len(h.History) != 2 {
		t.Errorf("Merge failed: Expected History length to be 2, but got %d", len(h.Latest.KnownDates))
	}
}


func TestHistoryGetLatestAndInitialized(t *testing.T) {
	var h History[int]

	if h.Initialized() {
		t.Errorf("Expected Initialized to retrun false, but got %t", h.Initialized())
	}
	if h.GetLatest() != nil {
		t.Errorf("Expected GetLatest to retrun nil, but got %d", h.GetLatest())
	}
	h.AddOrUpdateLatestItem(9)

	if *h.GetLatest() != 9 {
		t.Errorf("Expected GetLatest to retrun 9, but got %d", h.GetLatest())
	}
	
}

func TestHistory_AddOrUpdateLatestItem(t *testing.T) {
	var h History[int]

	h.AddOrUpdateLatestItem(42)
	if len(h.History) != 1 {
		t.Errorf("Expected history length to be 1, but got %d", len(h.History))
	}
	if h.Latest == nil {
		t.Error("Expected Latest item to be not nil")
	}
	if h.Latest.Data != 42 {
		t.Errorf("Expected Latest data to be 42, but got %d", h.Latest.Data)
	}

	originalLatest := h.Latest
	h.AddOrUpdateLatestItem(42)
	if len(h.History) != 1 {
		t.Errorf("Expected history length to still be 1, but got %d", len(h.History))
	}
	if h.Latest != originalLatest {
		t.Error("Expected Latest item to be the same after update")
	}
	if !reflect.DeepEqual(h.Latest.KnownDates, originalLatest.KnownDates) {
		t.Error("Expected KnownDates to be equal after update")
	}

	h.AddOrUpdateLatestItem(24)
	if len(h.History) != 2 {
		t.Errorf("Expected history length to be 2, but got %d", len(h.History))
	}
	if h.Latest.Data != 24 {
		t.Errorf("Expected Latest data to be 24, but got %d", h.Latest.Data)
	}
	h.AddOrUpdateLatestItem(42)
	if len(h.History) != 3 {
		t.Errorf("Expected history length to be 2, but got %d", len(h.History))
	}
	if h.Latest == nil {
		t.Error("Expected Latest item to be not nil")
	}
	if h.Latest.Data != 42 {
		t.Errorf("Expected Latest data to be 42, but got %d", h.Latest.Data)
	}
	originalLatest = h.Latest
	h.AddOrUpdateLatestItem(42)
	if len(h.History) != 3 {
		t.Errorf("Expected history length to still be 3, but got %d", len(h.History))
	}
	if h.Latest != originalLatest {
		t.Error("Expected Latest item to be the same after update")
	}
	if !reflect.DeepEqual(h.Latest.KnownDates, originalLatest.KnownDates) {
		t.Error("Expected KnownDates to be equal after update")
	}
}
