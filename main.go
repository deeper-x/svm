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

const defaultOut = "action required: [show | major | minor | patch | undo | all]"

func main() {
	out, err := svm()
	if err != nil {
		panic(err)
	}

	fmt.Println(out)
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
		return "", err
	}

	minorInt, err := strconv.Atoi(minor)
	if err != nil {
		return "", err
	}

	patchInt, err := strconv.Atoi(patch)
	if err != nil {
		panic(err)
	}

	switch action := os.Args[1]; action {
	case "show":
		out := fmt.Sprintf("Current version: %s", getCurVer())
		return out, nil

	case "major":
		newVer := fmt.Sprintf("v%d.%d.%d", majorInt+1, minorInt, patchInt)
		err := setNewVer(newVer, "major")
		if err != nil {
			return newVer, err
		}

		return newVer, nil

	case "minor":
		newVer := fmt.Sprintf("v%d.%d.%d", majorInt, minorInt+1, patchInt)
		err := setNewVer(newVer, "minor")
		if err != nil {
			return newVer, err
		}
		return newVer, nil

	case "patch":
		newVer := fmt.Sprintf("v%d.%d.%d", majorInt, minorInt, patchInt+1)
		err := setNewVer(newVer, "patch")
		if err != nil {
			return newVer, err
		}

		return newVer, nil

	case "undo":
		cmd := exec.Command("bash", "-c", "git tag -d $( git tag -l --sort=creatordate | tail -n1 )")
		stdout, err := cmd.Output()
		if err != nil {
			return "", err
		}
		res := bytes.NewBuffer(stdout).String()

		return res, nil

	case "all":
		cmd := exec.Command("bash", "-c", "git tag -l --sort=creatordate")
		stdout, err := cmd.Output()
		if err != nil {
			return "", err
		}

		res := bytes.NewBuffer(stdout).String()
		return res, nil

	default:
		return "", errors.New(defaultOut)
	}
}

func getCurVer() string {
	cmd := exec.Command("bash", "-c", "git tag -l --sort=creatordate | tail -n1")
	stdout, err := cmd.Output()
	if err != nil {
		panic(err)
	}

	args := os.Args

	if len(args) != 2 {
		fmt.Println(defaultOut)
		return ""
	}

	var semVer = bytes.NewBuffer(stdout).String()
	return semVer
}

func setNewVer(newVer, what string) error {
	fmt.Printf("Setting new %s tag:\n", what)
	cmd := exec.Command("git", "tag", newVer)
	_, err := cmd.Output()
	if err != nil {
		return err
	}

	return nil
}
