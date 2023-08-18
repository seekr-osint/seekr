package history

import "time"

type History[T any] struct {
	Latest  *HistoryItem[T]  `json:"latest"`
	History []HistoryItem[T] `json:"history"`
}
type HistoryItem[T any] struct {
	Data       T           `json:"data"`
	StartDate  time.Time   `json:"start_date"`
	EndDate    time.Time   `json:"end_date"`
	KnownDates []time.Time `json:"known_dates"`
}
