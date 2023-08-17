package history

import (
	"log"
	"reflect"
	"time"
)

func (h *History[T]) Initialized() bool {
	return h.Latest != nil
}

func (h *History[T]) Merge(h2 History[T]) { // no error handeling for merging h2 with more then one history item
	log.Printf("merging")
	if !h2.Initialized() {
		return
	}
	data := *h2.GetLatest()
	h.AddOrUpdateLatestItem(data)
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
