package git

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/deeper-x/svm/settings"
)

// Tag object representing git tag
type Tag struct {
	Major int
	Minor int
	Patch int
}

// NewTag return tag
func NewTag() (Tag, error) {
	cmd := exec.Command("bash", "-c", "git tag -l --sort=creatordate | tail -n1")
	stdout, err := cmd.Output()
	if err != nil {
		return Tag{}, err
	}

	var semVer = bytes.NewBuffer(stdout).String()
	digits := strings.Split(semVer, "v")[1]
	numParts := strings.Split(digits, ".")

	major := numParts[0]
	minor := numParts[1]
	patch := strings.TrimSuffix(numParts[2], "\n")

	majorInt, err := strconv.Atoi(major)
	if err != nil {
		return Tag{}, err
	}

	minorInt, err := strconv.Atoi(minor)
	if err != nil {
		return Tag{}, err
	}

	patchInt, err := strconv.Atoi(patch)
	if err != nil {
		panic(err)
	}

	return Tag{
		Major: majorInt,
		Minor: minorInt,
		Patch: patchInt,
	}, nil
}

// SetNewVer set new git tag version
func (t *Tag) SetNewVer(newVer string) error {
	command := fmt.Sprintf("git tag %s", newVer)
	cmd := exec.Command("bash", "-c", command)
	_, err := cmd.Output()
	if err != nil {
		return err
	}

	return nil
}

// DelVer delete input version
func (t *Tag) DelVer(ver string) error {
	action := fmt.Sprintf("git tag -d %s", ver)
	cmd := exec.Command("bash", "-c", action)
	_, err := cmd.Output()
	if err != nil {
		return err
	}

	return nil
}

// ShowAll return all tags
func (t *Tag) ShowAll() (string, error) {
	cmd := exec.Command("bash", "-c", "git tag -l --sort=creatordate")
	stdout, err := cmd.Output()
	if err != nil {
		return settings.DefaultErrMsg, err
	}

	res := bytes.NewBuffer(stdout).String()

	return res, nil
}

// Undo last tag creation
func (t *Tag) Undo() (string, error) {
	cmd := exec.Command("bash", "-c", "git tag -d $( git tag -l --sort=creatordate | tail -n1 )")
	stdout, err := cmd.Output()

	if err != nil {
		return settings.DefaultErrMsg, err
	}

	res := bytes.NewBuffer(stdout).String()

	return res, nil
}

// WriteFile save semantic version to input target file
func (t *Tag) WriteFile(fName string) error {
	curVer := t.String()

	err := os.WriteFile(fName, []byte(curVer), 0644)

	if err != nil {
		return err
	}

	return nil
}

func (t *Tag) String() string {
	return fmt.Sprintf("v%d.%d.%d", t.Major, t.Minor, t.Patch)
}

// CheckNumArgs check input arguments numbers (minimum)
func CheckNumArgs(tot int) error {
	if len(os.Args) < tot {
		return errors.New(settings.DefaultOut)
	}

	return nil
}
