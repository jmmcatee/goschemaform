package goschemaform

import (
	"testing"
)

func TestCheckBoxDefaults(t *testing.T) {
	check := NewCheckBoxInput("checkbox1")

	t.Logf("Form:\n%s\n\nSchema:\n%s\n", check.Form(), check.Schema())
}

func TestCheckBoxTitle(t *testing.T) {
	check := NewCheckBoxInput("checkbox1")
	check.SetTitle("Checkbox 1")

	t.Logf("Form:\n%s\n\nSchema:\n%s\n", check.Form(), check.Schema())
}

func TestCheckBoxTitleCondition(t *testing.T) {
	check := NewCheckBoxInput("checkbox1")
	check.SetTitle("Checkbox 1")
	check.SetConidition("check1", false)

	t.Logf("Form:\n%s\n\nSchema:\n%s\n", check.Form(), check.Schema())
}

func TestCheckBoxTitleConditionFlip(t *testing.T) {
	check := NewCheckBoxInput("checkbox1")
	check.SetTitle("Checkbox 1")
	check.SetConidition("check1", true)

	t.Logf("Form:\n%s\n\nSchema:\n%s\n", check.Form(), check.Schema())
}
