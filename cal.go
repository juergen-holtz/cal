// Package cal is a slightly different calendar command line tool.
package cal

import (
	"errors"
	"fmt"
	"time"
)

type Calendar struct {
	Year     int
	Month    int
	StartDay int
}

type CalendarOpts func(c *Calendar) error

// currentYear and currentMonth are defined here, so they can be set once when the
// package is initialized.
var currentYear int
var currentMonth int

// monthDays initializes the days in a month for non leap years.
var monthDays [13]int = [13]int{
	-1, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31,
}

func init() {
	now := time.Now()
	currentYear = now.Year()
	currentMonth = int(now.Month())
}

// IsLeapYear returns true, if the calendar's year is a leap year.
func IsLeapYear(year int) bool {
	if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
		return true
	}
	return false
}

// ForYear is an initializer function to set the calendar's year.
func ForYear(year int) CalendarOpts {
	return func(c *Calendar) error {
		if year < 1 {
			return errors.New("invalid year")
		}
		c.Year = year
		return nil
	}
}

// ForMonth is an initializer function to set the calendar's month.
func ForMonth(month int) CalendarOpts {
	return func(c *Calendar) error {
		if month < 1 {
			return errors.New("invalid month")
		}
		c.Month = month
		return nil
	}
}

// Constructs a Calendar object with the specified options.c
func NewCalendar(options ...CalendarOpts) (*Calendar, error) {
	// If nothing else, construct a calendar for current month and year.
	c := Calendar{Year: currentYear, Month: currentMonth}
	for _, opt := range options {
		err := opt(&c)
		if err != nil {
			return nil, err
		}
	}
	// Set location
	loc, err := time.LoadLocation("Local")
	if err != nil {
		return nil, err
	}
	// Set weekday of the 1st (e.g. Tue)
	c.StartDay = int(time.Date(c.Year, time.Month(c.Month), 1, 0, 0, 0, 0, loc).Weekday())
	return &c, nil
}

// Print calendar by implementing the Stringer interface.
func (c *Calendar) String() string {
	out := ""
	// Days in Month
	daysInMonth := monthDays[c.Month]
	if c.Month == 2 && IsLeapYear(c.Year) {
		daysInMonth++
	}
	// Print header
	out = out + fmt.Sprintln("Mo Di Mi Do Fr Sa So")
	// Print calendar
	var p int
	for p = 1; p < c.StartDay; p++ {
		out = out + "   "
	}
	for d := 1; d <= daysInMonth; d++ {
		out = out + fmt.Sprintf("%2d ", d)
		if p%7 == 0 {
			out = out + "\n"
			p = 0
		}
		p++
	}
	return out
}
