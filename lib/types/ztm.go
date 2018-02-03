package types

import "time"

// BusStop result of GetBusStop request
type BusStop struct {
	BusID string
	Name  string
}

type Line string

type TimeTableRecord struct {
	Brigade   string
	Direction string
	Time      time.Time
}

type TimeTable struct {
	Record []TimeTableRecord
}
