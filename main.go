package main

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

const defaultOut = "action required: [show | major | minor | patch | undo | all | write <file>]"
const defaultErr = "no action executed"

func main() {
	out, err := svm()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Print(out)
}

func svm() (string, error) {
	semVer := getCurVer()
	digits := strings.Split(semVer, "v")[1]
	parts := strings.Split(digits, ".")

	major := parts[0]
	minor := parts[1]
	patch := strings.TrimSuffix(parts[2], "\n")

	majorInt, err := strconv.Atoi(major)
	if err != nil {
		return defaultErr, err
	}

	minorInt, err := strconv.Atoi(minor)
	if err != nil {
		return defaultErr, err
	}

	patchInt, err := strconv.Atoi(patch)
	if err != nil {
		panic(err)
	}

	err = checkNumArgs(2)
	if err != nil {
		return defaultErr, err
	}

	switch action := os.Args[1]; action {
	case "show":
		out := fmt.Sprintf("Current version: %s", getCurVer())
		return out, nil

	case "major":
		newVer := fmt.Sprintf("v%d.%d.%d\n", majorInt+1, minorInt, patchInt)
		err := setNewVer(newVer)
		if err != nil {
			return defaultErr, err
		}

		return newVer, nil

	case "minor":
		newVer := fmt.Sprintf("v%d.%d.%d\n", majorInt, minorInt+1, patchInt)
		err := setNewVer(newVer)
		if err != nil {
			return defaultErr, err
		}

		return newVer, nil

	case "patch":
		newVer := fmt.Sprintf("v%d.%d.%d\n", majorInt, minorInt, patchInt+1)
		err := setNewVer(newVer)
		if err != nil {
			return defaultErr, err
		}

		return newVer, nil

	case "undo":
		cmd := exec.Command("bash", "-c", "git tag -d $( git tag -l --sort=creatordate | tail -n1 )")
		stdout, err := cmd.Output()
		if err != nil {
			return defaultErr, err
		}
		res := bytes.NewBuffer(stdout).String()

		return res, nil

	case "all":
		cmd := exec.Command("bash", "-c", "git tag -l --sort=creatordate")
		stdout, err := cmd.Output()
		if err != nil {
			return defaultErr, err
		}

		res := bytes.NewBuffer(stdout).String()
		return res, nil

	case "write":
		err = checkNumArgs(3)
		if err != nil {
			return defaultErr, err
		}

		fName := os.Args[2]
		curVer := getCurVer()
		err := os.WriteFile(fName, []byte(curVer), 0644)
		if err != nil {
			return defaultErr, err
		}

		return "Saved to file", nil

	default:
		return defaultErr, errors.New(defaultOut)
	}
}

func getCurVer() string {
	cmd := exec.Command("bash", "-c", "git tag -l --sort=creatordate | tail -n1")
	stdout, err := cmd.Output()
	if err != nil {
		panic(err)
	}

	var semVer = bytes.NewBuffer(stdout).String()
	return semVer
}

func checkNumArgs(tot int) error {
	if len(os.Args) < tot {
		return errors.New(defaultOut)
	}

	return nil
}

func setNewVer(newVer string) error {
	command := fmt.Sprintf("git tag %s", newVer)
	cmd := exec.Command("bash", "-c", command)
	_, err := cmd.Output()
	if err != nil {
		return err
	}

	return nil
}

func delVer(ver string) error {
	action := fmt.Sprintf("git tag -d %s", ver)
	cmd := exec.Command("bash", "-c", action)
	_, err := cmd.Output()
	if err != nil {
		return err
	}

	return nil
}
