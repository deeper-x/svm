package git

import (
	"os"
	"testing"
)

var testSemVer = "v.199.199.199"

func TestNewTag(t *testing.T) {
	_, err := NewTag()

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
	tag, err := NewTag()

	if err != nil {
		t.Error(err)
	}

	err = tag.SetNewVer(testSemVer)
	if err != nil {
		t.Error(err)
	}
}

func TestDelVer(t *testing.T) {
	tag, err := NewTag()
	if err != nil {
		t.Error(err)
	}

	err = tag.DelVer(testSemVer)
	if err != nil {
		t.Error(err)
	}
}

func TestShowAll(t *testing.T) {
	tag, err := NewTag()
	if err != nil {
		t.Error(err)
	}

	_, err = tag.ShowAll()
	if err != nil {
		t.Error(err)
	}
}
