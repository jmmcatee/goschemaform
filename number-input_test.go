package goschemaform

import (
	"testing"
)

func TestNumberInputDefaults(t *testing.T) {
	input := NewNumberInput("test1")

	input.SetTitle("Test 1")

	t.Logf("Form:\n%s\n\nSchema:\n%s\n", input.Form(), input.Schema())
}

func TestNumberInputMinimum(t *testing.T) {
	input := NewNumberInput("test1")

	input.SetTitle("Test 1")
	input.SetMin(0)

	t.Logf("Form:\n%s\n\nSchema:\n%s\n", input.Form(), input.Schema())
}

func TestNumberInputMaximum(t *testing.T) {
	input := NewNumberInput("test1")

	input.SetTitle("Test 1")
	input.SetMax(10)

	t.Logf("Form:\n%s\n\nSchema:\n%s\n", input.Form(), input.Schema())
}

func TestNumberInputMinMax(t *testing.T) {
	input := NewNumberInput("test1")

	input.SetTitle("Test 1")
	input.SetMin(-10)
	input.SetMax(5)

	t.Logf("Form:\n%s\n\nSchema:\n%s\n", input.Form(), input.Schema())
}
