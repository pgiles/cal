package internal

import (
	"fmt"
	"github.com/fatih/color"
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
	date    time.Time
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
			date:    Date(c.year, month, d+1),
			working: false,
			worked:  false,
		})
	}
	return c
}

func (c *Calendar) Print() {
	g := color.New(color.FgGreen)
	for m := range c.grid {
		t := Date(c.year, monthNameToMonth(m), 1)
		fmt.Print(t.Format("Jan"), ": ")
		for dayMeta := 0; dayMeta < len(c.grid[m]); dayMeta++ {
			if strings.HasPrefix(c.grid[m][dayMeta].name, "S") {
				continue
			}
			if c.grid[m][dayMeta].worked {
				_, _ = g.Printf("%s %d,", c.grid[m][dayMeta].name[0:2], c.grid[m][dayMeta].num)
			} else {
				fmt.Printf("%s %d,", c.grid[m][dayMeta].name[0:2], c.grid[m][dayMeta].num)
			}
		}
		fmt.Println()
	}
}

func (c *Calendar) AddWorkingDay(date time.Time) {
	for m := range c.grid {
		for dayMeta := 0; dayMeta < len(c.grid[m]); dayMeta++ {
			if EqualDate(date, c.grid[m][dayMeta].date) {
				fmt.Printf("added %v as a working day\n", date)
				c.grid[m][dayMeta].worked = true
			}
		}
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

// EqualDate compares two time.Time values based only on Year, Month, and Day
func EqualDate(t1, t2 time.Time) bool {
	y1, m1, d1 := t1.Date()
	y2, m2, d2 := t2.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}
