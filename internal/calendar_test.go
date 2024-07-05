package internal

import (
	"testing"
	"time"
)

func TestRenewVaultToken(t *testing.T) {
	c := NewCalendar().AddMonth(time.January)
	c.AddWorkingDay(time.Date(2024, time.January, 3, 0, 0, 0, 0, time.UTC))
	c.Print()

	//	if err != nil {
	//		t.Errorf("%v", err)
	//	}
}
