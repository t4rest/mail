package main

import (
	"time"
)

type Calendar struct {
	time time.Time
}

// returns Current Quarter
func (c *Calendar) CurrentQuarter() int {
	month := c.time.UTC().Month()

	if month < 4 {
		return 1
	} else if month < 7 {
		return 2
	} else if month < 10 {
		return 3
	} else {
		return 4
	}
}

// create new Calendar
func NewCalendar(time time.Time) *Calendar {
	return &Calendar{time}
}