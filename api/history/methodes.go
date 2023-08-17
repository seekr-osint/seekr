package history

import (
	"reflect"
	"time"
)

func (h *History[T]) Initialized() bool {
	return h.Latest != nil
}
func (h *History[T]) AddOrUpdateLatestItem(data T) {
	if h.Initialized() && reflect.DeepEqual(h.Latest.Data, data) { // any is used to support structs
		h.Latest.EndDate = time.Now()
		h.Latest.KnownDates = append(h.Latest.KnownDates, time.Now())
	} else {
		newLatest := HistoryItem[T]{
			Data:       data,
			StartDate:  time.Now(),
			EndDate:    time.Now(),
			KnownDates: []time.Time{time.Now()},
		}
		h.Latest = &newLatest
		h.History = append([]HistoryItem[T]{newLatest}, h.History...)
	}
}

func (h *History[T]) GetLatest() *T {
	if !h.Initialized() {
		return nil
	}
	return &h.Latest.Data
}
