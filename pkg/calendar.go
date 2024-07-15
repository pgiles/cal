package pkg

import (
	"fmt"
	"github.com/fatih/color"
	"sort"
	"strings"
	"time"
)

type Calendar struct {
	grid map[time.Month][]DayMeta
	year int
}

type DayMeta struct {
	num     int
	name    string
	date    time.Time
	working bool
	worked  bool
	dayOff  bool
}

func NewCalendar() *Calendar {
	return &Calendar{
		grid: map[time.Month][]DayMeta{},
		year: time.Now().Year(),
	}
}

func (c *Calendar) SetAndAddYear(year int) *Calendar {
	c.year = year
	for i := time.January; i <= time.December; i++ {
		c.AddMonth(i)
	}
	return c
}

func (c *Calendar) AddMonth(month time.Month) *Calendar {
	//build a grid
	days := daysInMonth(month, c.year)
	for d := 0; d < days; d++ {
		c.grid[month] = append(c.grid[month], DayMeta{
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
	g := color.New(color.FgGreen) // color for worked days
	r := color.New(color.FgRed)   // color for days off

	totalWorkedDays := 0
	for _, k := range sortedKeys(c.grid) {
		m := time.Month(k)
		fmt.Print(Date(c.year, m, 1).Format("Jan"), ": ") // Print month abbreviation to start a row
		workedDays := 0
		daysOff := 0
		for dayMeta := 0; dayMeta < len(c.grid[m]); dayMeta++ {
			if strings.HasPrefix(c.grid[m][dayMeta].name, "S") {
				// Work weeks; who works Saturday or Sunday?
				continue
			}
			c.grid[m][dayMeta].working = true
			if c.grid[m][dayMeta].worked {
				workedDays++
				_, _ = g.Printf("%s %d,", c.grid[m][dayMeta].name[0:2], c.grid[m][dayMeta].num)
			} else if c.grid[m][dayMeta].dayOff {
				daysOff++
				_, _ = r.Printf("%s %d,", c.grid[m][dayMeta].name[0:2], c.grid[m][dayMeta].num)
			} else {
				fmt.Printf("%s %d,", c.grid[m][dayMeta].name[0:2], c.grid[m][dayMeta].num)
			}
		}
		totalWorkedDays += workedDays
		workingDays := c.calculateWorkingDays(m)
		percentage := float64(workedDays) / float64(workingDays-daysOff) * 100
		fmt.Printf(" days worked: %.2f%%\n", percentage)
	}
	fmt.Println("total days worked:", totalWorkedDays)
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

func (c *Calendar) AddDayOff(date time.Time) {
	for m := range c.grid {
		for dayMeta := 0; dayMeta < len(c.grid[m]); dayMeta++ {
			if EqualDate(date, c.grid[m][dayMeta].date) {
				fmt.Printf("added %v as a day off\n", date)
				c.grid[m][dayMeta].dayOff = true
			}
		}
	}
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

// daysInMonth returns the number of days in a month for a given year.
func daysInMonth(m time.Month, year int) int {
	return time.Date(year, m+1, 0, 0, 0, 0, 0, time.UTC).Day()
}

func sortedKeys(m map[time.Month][]DayMeta) []int {
	keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)
	return keys
}

func (c *Calendar) calculateWorkingDays(m time.Month) int {
	result := 0
	for dayMeta := 0; dayMeta < len(c.grid[m]); dayMeta++ {
		if c.grid[m][dayMeta].working {
			result = result + 1
		}
	}
	return result
}
