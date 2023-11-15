package main

import (
	"fmt"
	"os"
	"testing"
)

func TestShow(t *testing.T) {
	os.Args = []string{"svm", "show"}

	out, err := svm()
	fmt.Println(out)
	if err != nil {
		t.Error(err)
	}
}

func TestPatch(t *testing.T) {
	os.Args = []string{"svm", "patch"}

	out, err := svm()
	fmt.Println(out)
	if err != nil {
		t.Error(err)
	}
}

func TestMinor(t *testing.T) {
	os.Args = []string{"svm", "minor"}

	out, err := svm()
	fmt.Println(out)
	if err != nil {
		t.Error(err)
	}
}

func TestMajor(t *testing.T) {
	os.Args = []string{"svm", "major"}

	out, err := svm()
	fmt.Println(out)
	if err != nil {
		t.Error(err)
	}
}

func TestUndo(t *testing.T) {
	os.Args = []string{"svm", "undo"}

	out, err := svm()
	fmt.Println(out)
	if err != nil {
		t.Error(err)
	}
}

func TestAll(t *testing.T) {
	os.Args = []string{"svm", "all"}

	out, err := svm()
	fmt.Println(out)
	if err != nil {
		t.Error(err)
	}
}
