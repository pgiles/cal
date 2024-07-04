package main

import (
	cal "github.com/pgiles/cal/internal"
	"time"
)

func main() {
	c := cal.NewCalendar().AddMonth(time.February)
	c.AddMonth(time.March)
	c.Print()
}
