package pkg

import (
	"testing"
	"time"
)

func TestAddWorkingDay(t *testing.T) {
	c := NewCalendar().AddMonth(time.January)
	workedDay := 3
	c.AddWorkingDay(time.Date(2024, time.January, workedDay, 0, 0, 0, 0, time.UTC))

	for k := range c.grid[time.Month(1)] {
		if c.grid[time.Month(1)][k].num == workedDay {
			if c.grid[time.Month(1)][k].worked != true {
				t.Errorf("Jan 3 is declared a worked day, but was not added as one")
			}
		}
	}
}

func TestAddDayOff(t *testing.T) {
	c := NewCalendar().AddMonth(time.February)
	dayOff := 7
	c.AddDayOff(time.Date(2024, time.February, dayOff, 0, 0, 0, 0, time.UTC))

	for k := range c.grid[time.Month(1)] {
		if c.grid[time.Month(1)][k].num == dayOff {
			if c.grid[time.Month(1)][k].dayOff != true {
				t.Errorf("Jan 3 is declared a day off, but was not added as one")
			}
		}
	}
}

func TestSetAndAddYear(t *testing.T) {
	c := NewCalendar().SetAndAddYear(2024)
	c.Print()
}
