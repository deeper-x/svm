package main

import (
	"os"
	"testing"
)

func TestShow(t *testing.T) {
	os.Args = []string{"svm", "show"}

	_, err := svm()
	if err != nil {
		t.Error(err)
	}
}

func TestPatch(t *testing.T) {
	os.Args = []string{"svm", "patch"}

	_, err := svm()
	if err != nil {
		t.Error(err)
	}
}

func TestMinor(t *testing.T) {
	os.Args = []string{"svm", "minor"}

	_, err := svm()
	if err != nil {
		t.Error(err)
	}
}

func TestMajor(t *testing.T) {
	os.Args = []string{"svm", "major"}

	_, err := svm()
	if err != nil {
		t.Error(err)
	}
}

func TestUndo(t *testing.T) {
	os.Args = []string{"svm", "undo"}

	_, err := svm()
	if err != nil {
		t.Error(err)
	}
}

func TestAll(t *testing.T) {
	os.Args = []string{"svm", "all"}

	_, err := svm()
	if err != nil {
		t.Error(err)
	}
}

func TestWritre(t *testing.T) {
	os.Args = []string{"svm", "write", "test/out_file.txt"}

	_, err := svm()
	if err != nil {
		t.Error(err)
	}
}

func TestSetNewVer(t *testing.T) {
	err := setNewVer("v1.100.0")
	if err != nil {
		t.Error(err)
	}
}

func TestDelVer(t *testing.T) {
	err := delVer("v1.100.0")
	if err != nil {
		t.Error(err)
	}
}
