package cal_test

import (
	"testing"
	"time"

	"github.com/juergen-holtz/cal"
)

func TestNewCalendarCurrentMonthYear(t *testing.T) {
	c := cal.NewCalendar()
	want := cal.Calendar{Year: time.Now().Year(), Month: int(time.Now().Month())}
	if c.Month != want.Month || c.Year != want.Year {
		t.Errorf("Not current year/month calendar. Want %v, got %v\n", want, c)
	}
}

func TestNewCalendarForMonth(t *testing.T) {
	month := 11
	c := cal.NewCalendar(cal.ForMonth(month))
	want := cal.Calendar{Year: time.Now().Year(), Month: month}
	if c.Month != want.Month {
		t.Errorf("Invalid month. Want %d, got %d\n", want.Month, c.Month)
	}
	if c.Year != want.Year {
		t.Errorf("Invalid default year. Want %d, got %d\n", want.Year, c.Year)
	}
}

func TestNewCalendarForYear(t *testing.T) {
	year := 2006
	c := cal.NewCalendar(cal.ForYear(year))
	want := cal.Calendar{Year: year, Month: int(time.Now().Month())}
	if c.Year != want.Year {
		t.Errorf("Invalid year. Want %d, got %d\n", want.Year, c.Year)
	}
	if c.Month != want.Month {
		t.Errorf("Invalid default month. Want %d, got %d\n", want.Month, c.Month)
	}
}

func TestNewCalendarForMonthYear(t *testing.T) {
	year := 2006
	month := 1
	c := cal.NewCalendar(cal.ForMonth(month), cal.ForYear(year))
	want := cal.Calendar{Year: year, Month: month}
	if c.Month != want.Month || c.Year != want.Year {
		t.Errorf("Not expected year/month calendar. Want %v, got %v\n", want, c)
	}
}

func TestLeapYear(t *testing.T) {
	type testYears struct {
		year int
		leap bool
	}
	var testData []testYears = []testYears{
		{1900, false},
		{1901, false},
		{1902, false},
		{1903, false},
		{1904, true},
		{1905, false},
		{1999, false},
		{2000, true},
		{2001, false},
		{2004, true},
		{2100, false},
	}
	for _, td := range testData {
		got := cal.IsLeapYear(td.year)
		want := td.leap
		if want != got {
			t.Errorf("Leap year calculation failed for year %d. Want %v, got %v\n", td.year, want, got)
		}
	}
}
