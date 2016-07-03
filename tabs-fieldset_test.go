package goschemaform

import (
	"testing"
)

func TestOneTabOneInputDefaults(t *testing.T) {
	f := NewTabFieldset()
	t1 := NewTab()

	text1 := NewTextInput("text1")

	t1.AddElement(text1)

	f.AddTab(t1)

	t.Logf("Form:\n%s\n\nSchema:\n%s\n", f.Form(), f.Schema())
}

func TestOneTabOneInputTitle(t *testing.T) {
	f := NewTabFieldset()
	f.SetTitle("Tabbed Form")
	t1 := NewTab()
	t1.SetTitle("Tab 1")

	text1 := NewTextInput("text1")
	text1.SetTitle("Input 1")

	t1.AddElement(text1)

	f.AddTab(t1)

	t.Logf("Form:\n%s\n\nSchema:\n%s\n", f.Form(), f.Schema())
}

func TestMultiTabMulitInput(t *testing.T) {
	f := NewTabFieldset()
	f.SetTitle("Tabbed Form")
	t1 := NewTab()
	t1.SetTitle("Tab 1")
	t2 := NewTab()
	t2.SetTitle("Tab 2")
	t3 := NewTab()
	t3.SetTitle("Tab 3")
	t4 := NewTab()
	t4.SetTitle("Tab 4")

	text1 := NewTextInput("text1")
	text1.SetTitle("Input 1")
	text2 := NewTextInput("text2")
	text2.SetTitle("Input 2")
	text3 := NewTextInput("textarea1")
	text3.SetTitle("Textarea 1")
	text3.SetMultiline(true)

	check1 := NewCheckBoxInput("checkbox1")
	check1.SetTitle("Checkbox 1")
	check2 := NewCheckBoxInput("checkbox2")
	check2.SetTitle("Checkbox 2")

	drop1 := NewDropDownInput("drop1")
	drop1.SetTitle("Drop 1")
	dopt1 := NewDropDownInputOption("option1")
	dopt2 := NewDropDownInputOption("option2")
	dopt3 := NewDropDownInputOption("option3")
	dopt4 := NewDropDownInputOption("option4")
	drop1.AddOption(dopt1)
	drop1.AddOption(dopt2)
	drop1.AddOption(dopt3)
	drop1.AddOption(dopt4)

	t1.AddElement(text1)
	t2.AddElement(text2)
	t3.AddElement(text3)
	t4.AddElement(drop1)
	t3.AddElement(check2)
	t3.AddElement(check1)

	f.AddTab(t1)
	f.AddTab(t2)
	f.AddTab(t3)
	f.AddTab(t4)

	t.Logf("Form:\n%s\n\nSchema:\n%s\n", f.Form(), f.Schema())
}

func TestMultiTabMulitInputConditional(t *testing.T) {
	f := NewTabFieldset()
	f.SetTitle("Tabbed Form")
	f.SetConidition("check1")
	t1 := NewTab()
	t1.SetTitle("Tab 1")
	t2 := NewTab()
	t2.SetTitle("Tab 2")
	t3 := NewTab()
	t3.SetTitle("Tab 3")
	t4 := NewTab()
	t4.SetTitle("Tab 4")

	text1 := NewTextInput("text1")
	text1.SetTitle("Input 1")
	text2 := NewTextInput("text2")
	text2.SetTitle("Input 2")
	text3 := NewTextInput("textarea1")
	text3.SetTitle("Textarea 1")
	text3.SetMultiline(true)

	check1 := NewCheckBoxInput("checkbox1")
	check1.SetTitle("Checkbox 1")
	check2 := NewCheckBoxInput("checkbox2")
	check2.SetTitle("Checkbox 2")

	drop1 := NewDropDownInput("drop1")
	drop1.SetTitle("Drop 1")
	dopt1 := NewDropDownInputOption("option1")
	dopt2 := NewDropDownInputOption("option2")
	dopt3 := NewDropDownInputOption("option3")
	dopt4 := NewDropDownInputOption("option4")
	drop1.AddOption(dopt1)
	drop1.AddOption(dopt2)
	drop1.AddOption(dopt3)
	drop1.AddOption(dopt4)

	t1.AddElement(text1)
	t2.AddElement(text2)
	t3.AddElement(text3)
	t4.AddElement(drop1)
	t3.AddElement(check2)
	t3.AddElement(check1)

	f.AddTab(t1)
	f.AddTab(t2)
	f.AddTab(t3)
	f.AddTab(t4)

	t.Logf("Form:\n%s\n\nSchema:\n%s\n", f.Form(), f.Schema())
}
