package goschemaform

import (
	"testing"
)

func TestTextInputDefaults(t *testing.T) {
	input := NewTextInput("test1")

	input.SetTitle("Test 1")

	t.Logf("Form:\n%s\n\nSchema:\n%s\n", input.Form(), input.Schema())
}

func TestTextInputPlaceHolder(t *testing.T) {
	input := NewTextInput("test1")

	input.SetTitle("Test 1")
	input.SetPlaceHolder("Please type here...")

	t.Logf("Form:\n%s\n\nSchema:\n%s\n", input.Form(), input.Schema())
}

func TestTextInputMaxLength(t *testing.T) {
	input := NewTextInput("test1")

	input.SetTitle("Test 1")
	input.SetPlaceHolder("Please type here...")
	input.SetMaxLength(10)

	t.Logf("Form:\n%s\n\nSchema:\n%s\n", input.Form(), input.Schema())
}

func TestTextInputMultiline(t *testing.T) {
	input := NewTextInput("test1")

	input.SetTitle("Test 1")
	input.SetMultiline(true)

	t.Logf("Form:\n%s\n\nSchema:\n%s\n", input.Form(), input.Schema())
}

func TestTextInputConditional(t *testing.T) {
	input := NewTextInput("test1")

	input.SetTitle("Test 1")
	input.SetCondition("check1", false)

	t.Logf("Form:\n%s\n\nSchema:\n%s\n", input.Form(), input.Schema())
}

func TestTextInputConditionalFlip(t *testing.T) {
	input := NewTextInput("test1")

	input.SetTitle("Test 1")
	input.SetCondition("check1", true)

	t.Logf("Form:\n%s\n\nSchema:\n%s\n", input.Form(), input.Schema())
}
