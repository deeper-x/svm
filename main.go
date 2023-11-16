package main

import (
	"errors"
	"fmt"
	"os"

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
	tag, err := git.NewTag()
	if err != nil {
		return settings.DefaultErrMsg, err
	}

	err = git.CheckNumArgs(2)
	if err != nil {
		return settings.DefaultErrMsg, err
	}

	switch action := os.Args[1]; action {
	case "show":
		tag, err := git.NewTag()
		if err != nil {
			return settings.DefaultErrMsg, err
		}

		out := fmt.Sprintf("Current version: %s\n", tag.String())
		return out, nil

	case "major":
		newVer := fmt.Sprintf("v%d.%d.%d\n", tag.Major+1, 0, 0)
		err := tag.SetNewVer(newVer)
		if err != nil {
			return settings.DefaultErrMsg, err
		}

		return newVer, nil

	case "minor":
		newVer := fmt.Sprintf("v%d.%d.%d\n", tag.Major, tag.Minor+1, 0)
		err := tag.SetNewVer(newVer)
		if err != nil {
			return settings.DefaultErrMsg, err
		}

		return newVer, nil

	case "patch":
		newVer := fmt.Sprintf("v%d.%d.%d\n", tag.Major, tag.Minor, tag.Patch+1)
		err := tag.SetNewVer(newVer)
		if err != nil {
			return settings.DefaultErrMsg, err
		}

		return newVer, nil

	case "undo":
		res, err := tag.Undo()
		if err != nil {
			return settings.DefaultErrMsg, err
		}

		return res, nil

	case "all":
		res, err := tag.ShowAll()
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

		err := tag.WriteFile(fName)
		if err != nil {
			return settings.DefaultErrMsg, err
		}

		return "Saved to file", nil

	default:
		return settings.DefaultErrMsg, errors.New(settings.DefaultOut)
	}
}
