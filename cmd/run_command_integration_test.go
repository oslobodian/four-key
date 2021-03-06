package cmd

import (
	"os/exec"
	"strings"
	"testing"
)

func TestRunCommand_WhenRunsWithEmptyArgs_ReturnsOutput(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	runCmd := exec.Command("go", "run", "../main.go", "run",
		"--repository", "",
		"--startDate", "",
		"--endDate", "")

	out, err := runCmd.Output()
	if err != nil {
		t.Errorf(err.Error())
	}

	want := "date will be like -s YYYY-MM-DD"
	got := string(out)
	if !strings.Contains(got, want) {
		t.Errorf("Unexpected data.\nGot: %s\nExpected: %s", got, want)
	}
}

func TestRunCommand_WhenRunsWithCorrectCloneAddressButWrongArgs_ReturnsOutput(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	addCmd := exec.Command("go", "run", "../main.go", "add",
		"--cloneAddress", "https://github.com/Trendyol/medusa.git",
		"--team", "trendyol-team",
		"--releaseTagPattern", "release-v",
		"--fixCommitPatterns", "fix", "-f", "hotfix")

	_, err := addCmd.Output()
	if err != nil {
		t.Errorf(err.Error())
	}

	runCmd := exec.Command("go", "run", "../main.go", "run",
		"--repository", "medusa",
		"--startDate", "",
		"--endDate", "")

	out, err := runCmd.Output()
	if err != nil {
		t.Errorf(err.Error())
	}

	removeCmd := exec.Command("go", "run", "../main.go", "remove",
		"--repository", "medusa")

	_, err = removeCmd.Output()
	if err != nil {
		t.Errorf(err.Error())
	}

	want := "date will be like -s YYYY-MM-DD"
	got := string(out)
	if !strings.Contains(got, want) {
		t.Errorf("Unexpected data.\nGot: %s\nExpected: %s", got, want)
	}
}

func TestRunCommand_WhenRunsWithCorrectArgs_ReturnsOutput(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	addCmd := exec.Command("go", "run", "../main.go", "add",
		"--cloneAddress", "https://github.com/Trendyol/android-ui-components.git",
		"--team", "trendyol-team",
		"--releaseTagPattern", "toolbar-",
		"--fixCommitPatterns", "fix", "-f", "hotfix")

	_, err := addCmd.Output()
	if err != nil {
		t.Errorf(err.Error())
	}

	runCmd := exec.Command("go", "run", "../main.go", "run",
		"--repository", "android-ui-components",
		"--startDate", "2019-01-01",
		"--endDate", "2020-12-31")

	out, err := runCmd.Output()
	if err != nil {
		t.Errorf(err.Error())
	}

	removeCmd := exec.Command("go", "run", "../main.go", "remove",
		"--repository", "medusa")

	_, err = removeCmd.Output()
	if err != nil {
		t.Errorf(err.Error())
	}

	want := "metrics file generated"
	got := string(out)
	if !strings.Contains(got, want) {
		t.Errorf("Unexpected data.\nGot: %s\nExpected: %s", got, want)
	}
}