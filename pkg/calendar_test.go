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
				t.Errorf("Jan 3 is declared a worked day, but was no added as one")
			}
		}
	}
}

func TestSetAndAddYear(t *testing.T) {
	c := NewCalendar().SetAndAddYear(2024)
	c.Print()
}
