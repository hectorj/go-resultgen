package main

import (
	"io/ioutil"
	"os/exec"
	"reflect"
	"testing"
)

func TestOutputHasNotChanged(t *testing.T) {
	nonStrictFilePath := "./tests/dummy_result.go"
	strictFilePath := "./tests/dummy_result_strict.go"

	originalNonStrictContent, err := ioutil.ReadFile(nonStrictFilePath)
	if err != nil {
		t.Fatal(err)
	}

	originalStrictContent, err := ioutil.ReadFile(strictFilePath)
	if err != nil {
		t.Fatal(err)
	}

	cmd := exec.Command("go", "generate", "./tests")
	if err := cmd.Run(); err != nil {
		t.Fatal(err)
	}

	newNonStrictContent, err := ioutil.ReadFile(nonStrictFilePath)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(originalNonStrictContent, newNonStrictContent) {
		t.Error(nonStrictFilePath + " changed. You should commit that.")
	}

	newStrictContent, err := ioutil.ReadFile(strictFilePath)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(originalStrictContent, newStrictContent) {
		t.Error(strictFilePath + " changed. You should commit that.")
	}
}
