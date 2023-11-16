package git

import (
	"os"
	"testing"
)

var testSemVer = "v.199.199.199"

func TestGetCurVer(t *testing.T) {
	_, err := GetCurVer()

	if err != nil {
		t.Error(err)
	}
}

func TestCheckNumArgs(t *testing.T) {
	os.Args = []string{"svm", "test"}

	err := CheckNumArgs(2)
	if err != nil {
		t.Error(err)
	}
}

func TestCheckNumArgsMinimum(t *testing.T) {
	os.Args = []string{"svm", "test1", "test2"}
	err := CheckNumArgs(5)
	if err == nil {
		t.Error("this should fail because arguments should be 5 at least")
	}
}

func TestSetNewVer(t *testing.T) {
	err := SetNewVer(testSemVer)
	if err != nil {
		t.Error(err)
	}
}

func TestDelVer(t *testing.T) {
	err := DelVer(testSemVer)
	if err != nil {
		t.Error(err)
	}
}

func TestShowAll(t *testing.T) {
	_, err := ShowAll()
	if err != nil {
		t.Error(err)
	}
}
