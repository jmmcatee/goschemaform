package goschemaform

import (
	"testing"
)

func TestDropDownDefaults(t *testing.T) {
	dd := NewDropDownInput("dd1")
	o1 := NewDropDownInputOption("value1")
	o2 := NewDropDownInputOption("value2")
	o3 := NewDropDownInputOption("value3")

	dd.AddOption(o1)
	dd.AddOption(o2)
	dd.AddOption(o3)

	t.Logf("Form:\n%s\n\nSchema:\n%s\n", dd.Form(), dd.Schema())
}

func TestDropDownTitle(t *testing.T) {
	dd := NewDropDownInput("dd1")
	dd.SetTitle("Drop Down 1")

	o1 := NewDropDownInputOption("value1")
	o2 := NewDropDownInputOption("value2")
	o3 := NewDropDownInputOption("value3")

	dd.AddOption(o1)
	dd.AddOption(o2)
	dd.AddOption(o3)

	t.Logf("Form:\n%s\n\nSchema:\n%s\n", dd.Form(), dd.Schema())
}

func TestDropDownOptionTitles(t *testing.T) {
	dd := NewDropDownInput("dd1")
	o1 := NewDropDownInputOption("value1")
	o1.SetName("Value One")
	o2 := NewDropDownInputOption("value2")
	o2.SetName("Value Two")
	o3 := NewDropDownInputOption("value3")
	o3.SetName("Value Three")

	dd.AddOption(o1)
	dd.AddOption(o2)
	dd.AddOption(o3)

	t.Logf("Form:\n%s\n\nSchema:\n%s\n", dd.Form(), dd.Schema())
}

func TestDropDownTitleGroups(t *testing.T) {
	dd := NewDropDownInput("dd1")
	o1 := NewDropDownInputOption("value1")
	o1.SetName("Value One")
	o1.SetGroup("Group 1")
	o2 := NewDropDownInputOption("value2")
	o2.SetName("Value Two")
	o3 := NewDropDownInputOption("value3")
	o3.SetGroup("Group 1")
	o3.SetName("Value Three")

	dd.AddOption(o1)
	dd.AddOption(o2)
	dd.AddOption(o3)

	t.Logf("Form:\n%s\n\nSchema:\n%s\n", dd.Form(), dd.Schema())
}
