package example

import (
	"encoding/json"
	cal "github.com/pgiles/cal/pkg"
	"os"
	"strconv"
	"strings"
	"time"
)

func ApplyCommits() {
	c := cal.NewCalendar().AddMonth(time.January).AddMonth(time.February)
	commitList := jsonToCommitObject("example/commits.json")
	for _, cw := range *commitList {
		c.AddWorkingDay(cw.Commit.CSTDate)
	}
	c.Print()
}

func ApplyCommitsAndPTO() {
	c := cal.NewCalendar().SetAndAddYear(2024)
	commitList := jsonToCommitObject("example/commits.json")
	for _, cw := range *commitList {
		c.AddWorkingDay(cw.Commit.CSTDate)
	}

	ptoList := jsonToObject("example/pto.json")
	//daysOff := ptoList["employee"].(map[string]interface{})["daysOff"].([]interface{})
	daysOff := ptoList["daysOff"].([]interface{})
	for _, itm := range daysOff {
		ymd := strings.Split(itm.(string), "-")
		y, _ := strconv.Atoi(ymd[0])
		m, _ := strconv.Atoi(ymd[1])
		d, _ := strconv.Atoi(ymd[2])
		c.AddDayOff(cal.Date(y, time.Month(m), d))
	}
	c.Print()
}

func jsonToObject(filepath string) map[string]interface{} {
	jsonData, err := os.ReadFile(filepath)
	handleError(err)

	var daysOff map[string]interface{}
	err = json.Unmarshal(jsonData, &daysOff)
	handleError(err)

	return daysOff
}

func jsonToCommitObject(filepath string) *CommitList {
	jsonData, err := os.ReadFile(filepath)
	handleError(err)

	var commits CommitList
	err = json.Unmarshal(jsonData, &commits)
	handleError(err)

	return &commits
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
