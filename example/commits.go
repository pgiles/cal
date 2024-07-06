package example

import (
	"encoding/json"
	cal "github.com/pgiles/cal/pkg"
	"os"
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
