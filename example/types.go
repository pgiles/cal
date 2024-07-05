package example

import (
	"time"
)

type CommitList []CommitWrapper

type CommitWrapper struct {
	Commit Commit `json:"commit"`
}

type Commit struct {
	Name    string    `json:"name"`
	Date    time.Time `json:"date"`
	Message string    `json:"message"`
	URL     string    `json:"url"`
	SHA     string    `json:"sha"`
	CSTDate time.Time `json:"cst_date"`
}
