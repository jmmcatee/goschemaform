package goschemaform

import (
	"testing"
)

func TestSimpleSchemaFormOneInputDefaults(t *testing.T) {
	sf := NewSchemaForm()

	text1 := NewTextInput("input1")

	sf.AddElement(text1)

	t.Logf("SchemaForm:\n%s\n", sf.SchemaForm())
}

func TestSimpleSchemaFormOneInputTitles(t *testing.T) {
	sf := NewSchemaForm()
	sf.SetTitle("New Form")

	text1 := NewTextInput("input1")
	text1.SetTitle("Input 1")

	sf.AddElement(text1)

	t.Logf("SchemaForm:\n%s\n", sf.SchemaForm())
}

func TestSchemaFormComplexTabs(t *testing.T) {
	sf := NewSchemaForm()
	sf.SetTitle("New Form")

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

	sf.AddElement(f)

	t.Logf("SchemaForm:\n%s\n", sf.SchemaForm())
}

func TestSchemaFormComplexTabsGlobal(t *testing.T) {
	sf := NewSchemaForm()
	sf.SetTitle("New Form")

	globalCheckbox := NewCheckBoxInput("globalcheckbox1")
	globalCheckbox.SetTitle("Global Check")
	sf.AddElement(globalCheckbox)

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
	text1.IsRequired(true)
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
	drop1.IsRequired(true)
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

	sf.AddElement(f)

	t.Logf("SchemaForm:\n%s\n", sf.SchemaForm())
}
