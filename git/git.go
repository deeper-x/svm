package git

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"

	"github.com/deeper-x/svm/settings"
)

// GetCurVer return current tag version
func GetCurVer() (string, error) {
	cmd := exec.Command("bash", "-c", "git tag -l --sort=creatordate | tail -n1")
	stdout, err := cmd.Output()
	if err != nil {
		return settings.DefaultErrMsg, err
	}

	var semVer = bytes.NewBuffer(stdout).String()
	return semVer, nil
}

// CheckNumArgs check input arguments numbers (minimum)
func CheckNumArgs(tot int) error {
	if len(os.Args) < tot {
		return errors.New(settings.DefaultOut)
	}

	return nil
}

// SetNewVer set new git tag version
func SetNewVer(newVer string) error {
	command := fmt.Sprintf("git tag %s", newVer)
	cmd := exec.Command("bash", "-c", command)
	_, err := cmd.Output()
	if err != nil {
		return err
	}

	return nil
}

// DelVer delete input version
func DelVer(ver string) error {
	action := fmt.Sprintf("git tag -d %s", ver)
	cmd := exec.Command("bash", "-c", action)
	_, err := cmd.Output()
	if err != nil {
		return err
	}

	return nil
}

// ShowAll return all tags
func ShowAll() (string, error) {
	cmd := exec.Command("bash", "-c", "git tag -l --sort=creatordate")
	stdout, err := cmd.Output()
	if err != nil {
		return settings.DefaultErrMsg, err
	}

	res := bytes.NewBuffer(stdout).String()

	return res, nil
}
