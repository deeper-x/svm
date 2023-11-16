package main

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/deeper-x/svm/git"

	"github.com/deeper-x/svm/settings"
)

func main() {
	out, err := svm()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Print(out)
}

func svm() (string, error) {
	semVer, err := git.GetCurVer()
	if err != nil {
		return settings.DefaultErrMsg, err
	}

	digits := strings.Split(semVer, "v")[1]
	numParts := strings.Split(digits, ".")

	major := numParts[0]
	minor := numParts[1]
	patch := strings.TrimSuffix(numParts[2], "\n")

	majorInt, err := strconv.Atoi(major)
	if err != nil {
		return settings.DefaultErrMsg, err
	}

	minorInt, err := strconv.Atoi(minor)
	if err != nil {
		return settings.DefaultErrMsg, err
	}

	patchInt, err := strconv.Atoi(patch)
	if err != nil {
		panic(err)
	}

	err = git.CheckNumArgs(2)
	if err != nil {
		return settings.DefaultErrMsg, err
	}

	switch action := os.Args[1]; action {
	case "show":
		cuVer, err := git.GetCurVer()
		if err != nil {
			return settings.DefaultErrMsg, err
		}

		out := fmt.Sprintf("Current version: %s", cuVer)
		return out, nil

	case "major":
		newVer := fmt.Sprintf("v%d.%d.%d\n", majorInt+1, 0, 0)
		err := git.SetNewVer(newVer)
		if err != nil {
			return settings.DefaultErrMsg, err
		}

		return newVer, nil

	case "minor":
		newVer := fmt.Sprintf("v%d.%d.%d\n", majorInt, minorInt+1, 0)
		err := git.SetNewVer(newVer)
		if err != nil {
			return settings.DefaultErrMsg, err
		}

		return newVer, nil

	case "patch":
		newVer := fmt.Sprintf("v%d.%d.%d\n", majorInt, minorInt, patchInt+1)
		err := git.SetNewVer(newVer)
		if err != nil {
			return settings.DefaultErrMsg, err
		}

		return newVer, nil

	case "undo":
		res, err := undo()
		if err != nil {
			return settings.DefaultErrMsg, err
		}

		return res, nil

	case "all":
		res, err := git.ShowAll()
		if err != nil {
			return settings.DefaultErrMsg, err
		}

		return res, nil

	case "write":
		err = git.CheckNumArgs(3)
		if err != nil {
			return settings.DefaultErrMsg, err
		}

		fName := os.Args[2]

		err := writeFile(fName)
		if err != nil {
			return settings.DefaultErrMsg, err
		}

		return "Saved to file", nil

	default:
		return settings.DefaultErrMsg, errors.New(settings.DefaultOut)
	}
}

func writeFile(fName string) error {
	curVer, err := git.GetCurVer()
	if err != nil {
		return err
	}

	err = os.WriteFile(fName, []byte(curVer), 0644)

	if err != nil {
		return err
	}

	return nil
}

func undo() (string, error) {
	cmd := exec.Command("bash", "-c", "git tag -d $( git tag -l --sort=creatordate | tail -n1 )")
	stdout, err := cmd.Output()

	if err != nil {
		return settings.DefaultErrMsg, err
	}

	res := bytes.NewBuffer(stdout).String()

	return res, nil
}
