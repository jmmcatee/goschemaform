package goschemaform

import (
	"testing"
)

func TestFileInputDefaults(t *testing.T) {
	input := NewFileInput("test1")

	input.SetTitle("Test 1")

	t.Logf("Form:\n%s\n\nSchema:\n%s\n", input.Form(), input.Schema())
}

func TestFileInputPlaceHolder(t *testing.T) {
	input := NewFileInput("test1")

	input.SetTitle("Test 1")
	input.SetPlaceHolder("Please type here...")

	t.Logf("Form:\n%s\n\nSchema:\n%s\n", input.Form(), input.Schema())
}

func TestFileInputMaxSize(t *testing.T) {
	input := NewFileInput("test1")

	input.SetTitle("Test 1")
	input.SetPlaceHolder("Please type here...")
	input.SetMaxSize(10)

	t.Logf("Form:\n%s\n\nSchema:\n%s\n", input.Form(), input.Schema())
}

func TestFileInputConditional(t *testing.T) {
	input := NewFileInput("test1")

	input.SetTitle("Test 1")
	input.SetConidition("check1")

	t.Logf("Form:\n%s\n\nSchema:\n%s\n", input.Form(), input.Schema())
}
