package gojsonform

import (
	"testing"
)

func TestTextInput(t *testing.T) {
	form := Form{}

	textInput := NewInputText("username")
	textInput.SetTitle("Username")
	textInput.SetDesc("The username used to log into the site.")
	textInput.SetDefault("Enter Username Here")
	textInput.SetMaxLength(100)
	textInput.SetRequired(true)

	form.Append(&textInput)

	println(form.GetForm())
}

func TestInputAndSelect(t *testing.T) {
	form := Form{}

	textInput := NewInputText("username")
	textInput.SetTitle("Username")
	textInput.SetDesc("The username used to log into the site.")
	textInput.SetDefault("Enter Username Here")
	textInput.SetMaxLength(100)
	textInput.SetRequired(true)

	form.Append(&textInput)

	selectInput := NewSelectInput("group")
	selectInput.SetTitle("Group")
	selectInput.SetDesc("The group this user is in.")
	selectInput.SetDefault("Read-Only")
	selectInput.SetEnum([]string{"Read-Only", "Standard User", "Adminsitrator"})
	selectInput.SetRequired(true)

	form.Append(&selectInput)

	println(form.GetForm())
}

func TestInputSelectCheckbox(t *testing.T) {
	form := Form{}

	textInput := NewInputText("username")
	textInput.SetTitle("Username")
	textInput.SetDesc("The username used to log into the site.")
	textInput.SetDefault("Enter Username Here")
	textInput.SetMaxLength(100)
	textInput.SetRequired(true)

	form.Append(&textInput)

	selectInput := NewSelectInput("group")
	selectInput.SetTitle("Group")
	selectInput.SetDesc("The group this user is in.")
	selectInput.SetDefault("Read-Only")
	selectInput.SetEnum([]string{"Read-Only", "Standard User", "Adminsitrator"})
	selectInput.SetRequired(true)

	form.Append(&selectInput)

	checkbox := NewCheckboxInput("enable")
	checkbox.SetTitle("Enable User")
	checkbox.SetDesc("Enable this user immediately.")
	checkbox.SetDefault(true)
	checkbox.SetRequired(true)

	form.Append(&checkbox)

	println(form.GetForm())
}

func TestInputArray(t *testing.T) {
	form := Form{}

	textInput := NewInputText("username")
	textInput.SetTitle("Username")
	textInput.SetDesc("The username used to log into the site.")
	textInput.SetDefault("Enter Username Here")
	textInput.SetMaxLength(100)
	textInput.SetRequired(true)

	form.Append(&textInput)

	arrayInput := NewArrayInput("groups")
	arrayInput.SetTitle("Groups")
	arrayInput.SetDesc("Groups to add to this user account.")

	selectInput := NewSelectInput("group")
	selectInput.SetTitle("Group")
	selectInput.SetDesc("The group this user is in.")
	selectInput.SetDefault("Read-Only")
	selectInput.SetEnum([]string{"Read-Only", "Standard User", "Adminsitrator"})
	selectInput.SetRequired(true)

	arrayInput.AddItem(&selectInput)
	form.Append(&arrayInput)

	checkbox := NewCheckboxInput("enable")
	checkbox.SetTitle("Enable User")
	checkbox.SetDesc("Enable this user immediately.")
	checkbox.SetDefault(true)
	checkbox.SetRequired(true)

	form.Append(&checkbox)

	println(form.GetForm())
}

func TestSubForm(t *testing.T) {
	form := Form{}

	altInput := NewSubFormInput("uorg")
	altInput.SetTitle("Add Type")
	altInput.SetDesc("Add a User or Group.")

	textInput := NewInputText("username")
	textInput.SetTitle("Username")
	textInput.SetDesc("The username used to log into the site.")
	textInput.SetMaxLength(100)
	textInput.SetRequired(true)
	altInput.AddItem(&textInput)

	groupInput := NewInputText("groupname")
	groupInput.SetTitle("Group Name")
	groupInput.SetDesc("The group name used to log into the site.")
	groupInput.SetMaxLength(100)
	groupInput.SetRequired(false)
	altInput.AddItem(&groupInput)

	form.Append(&altInput)

	println(form.GetForm())
}
