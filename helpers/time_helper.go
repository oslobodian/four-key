package helpers

import (
	"fmt"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"sort"
	"strconv"
	"time"
)

type TagData struct {
	tagDate time.Time
	tag     *plumbing.Reference
}

type timeSlice []TagData

func (p timeSlice) Len() int {
	return len(p)
}

func (p timeSlice) Less(i, j int) bool {
	return p[i].tagDate.Before(p[j].tagDate)
}

func (p timeSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func GetAscendingOrderByTagDate(r *git.Repository) (tags []TagData) {

	var tagDataList = make(map[string]TagData)
	rTags, err := r.Tags()
	if err != nil {
		println(err)
	}

	var i = 0
	err = rTags.ForEach(func(t *plumbing.Reference) error {
		cm, err := GetCommitFromTagHash(r, t.Hash())
		if err != nil {
			fmt.Println(err)
		}

		if cm != nil {
			tagDataList[strconv.Itoa(i)] = TagData{cm.Committer.When, t}
			i++
		}

		return nil
	})

	sortedTagDataList := make(timeSlice, 0, len(tagDataList))
	for _, tag := range tagDataList {
		sortedTagDataList = append(sortedTagDataList, tag)
	}
	sort.Sort(sortedTagDataList)

	return sortedTagDataList
}

type CommitData struct {
	CommitDate time.Time
	Commit     object.Commit
}
type timeCommitSlice []CommitData

func (p timeCommitSlice) Len() int {
	return len(p)
}

func (p timeCommitSlice) Less(i, j int) bool {
	return p[i].CommitDate.After(p[j].CommitDate)
}

func (p timeCommitSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func GetDescendingCommits(r *git.Repository) (tags []CommitData) {

	var commitDataList = make(map[string]CommitData)
	repoCommits, err := r.CommitObjects()
	if err != nil {
		println(err)
	}

	var i = 0
	err = repoCommits.ForEach(func(commit *object.Commit) error {
		if err != nil {
			//fmt.Println(err)
		}
		commitDataList[strconv.Itoa(i)] = CommitData{commit.Committer.When, *commit}
		i++

		return nil
	})

	sortedCommitDataList := make(timeCommitSlice, 0, len(commitDataList))
	for _, tempCommit := range commitDataList {
		sortedCommitDataList = append(sortedCommitDataList, tempCommit)
	}
	sort.Sort(sortedCommitDataList)

	return sortedCommitDataList
}

func IsDateWithinRange(dateStamp, startDate, finishDate time.Time) bool {

	if dateStamp.Before(startDate) && dateStamp.After(finishDate) {
		return true
	}

	return false
}
func inTimeSpan(start, end, check time.Time) bool {
	return check.After(start) && check.Before(end)
}

func GetSecondsToHour(seconds float64) float64 {
	return seconds / 3600
}

func GetSecondsToDays(seconds float64) float64 {
	return seconds / 86400
}
