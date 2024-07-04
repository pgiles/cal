package internal

import (
	"fmt"
	"strings"
	"time"
)

type Calendar struct {
	grid map[string][]DayMeta
	year int
}

type DayMeta struct {
	num     int
	name    string
	working bool
	worked  bool
}

func NewCalendar() *Calendar {
	return &Calendar{
		grid: map[string][]DayMeta{},
		year: time.Now().Year(),
	}
}

func (c *Calendar) AddMonth(month time.Month) *Calendar {
	//build a grid
	days := daysIn(month, c.year)
	for d := 0; d < days; d++ {
		c.grid[month.String()] = append(c.grid[month.String()], DayMeta{
			num:     d + 1,
			name:    Date(c.year, month, d+1).Weekday().String(),
			working: false,
			worked:  false,
		})
	}
	return c
}

func (c *Calendar) Print() {
	for m := range c.grid {
		t := Date(c.year, monthNameToMonth(m), 1)
		fmt.Print(t.Format("Jan"), ": ")
		for dayMeta := 0; dayMeta < len(c.grid[m]); dayMeta++ {
			if !strings.HasPrefix(c.grid[m][dayMeta].name, "S") {
				fmt.Printf("%s %d,", c.grid[m][dayMeta].name[0:2], c.grid[m][dayMeta].num)
			}
			/*			if (dayMeta)%2 == 0 {
							fmt.Printf("%v ", c.grid[m][dayMeta].name)
						} else {
							fmt.Printf("%v ", c.grid[m][dayMeta].name)
						}*/
		}
		fmt.Println()
	}
}

// daysIn returns the number of days in a month for a given year.
func daysIn(m time.Month, year int) int {
	return time.Date(year, m+1, 0, 0, 0, 0, 0, time.UTC).Day()
}

func monthNameToMonth(name string) time.Month {
	for i := time.January; i <= time.December; i++ {
		if i.String() == name {
			return i
		}
	}
	return 0
}

func Date(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}
