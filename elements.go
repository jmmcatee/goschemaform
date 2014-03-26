package gojsonform

import (
	"strconv"
)

/*
	- Single Text
	- Enum (Select)
	- Checkbox
	- Array (1 -> many of an element or group)
	- Alternative Select

	JSON Form supports all JSON Schema simple types that make sense in the context of an HTML form:

		string for strings
		number for numbers, including floating numbers
		integer for integers
		boolean for booleans
		array for arrays
		object for objects


		The JSON Form library uses the information available in the JSON schema to complete each generated field. In particular:

		The title property serves as label for the input.
		The description property is displayed next to the input field to guide user input.
		The default property sets the initial value of a field.
		The maxLength property sets the maxlength attribute of the input field
		The minimum, maximum, exclusiveMinimum, exclusiveMaximum properties are used to set corresponding attributes for a range input field.
		The minItems and maxItems properties are used to restrict the number of items that may be added or removed to an array.
		The readOnly property sets the readonly attribute of the input field
		The required property marks the input field as required, adds a jsonform-required class to the container markup of the input field for styling purpose and adds a jsonform-hasrequired class to the <form> tag e.g. to allow one to add a legend that explains that fields marked with an asterisk are required, e.g. with CSS properties such as:
*/

type InputText struct {
	name        string
	title       string
	description string
	defValue    string
	maxLength   int
	required    bool
}

func NewInputText(v string) InputText {
	return InputText{name: v}
}

func (e *InputText) SetTitle(v string) {
	e.title = v
}

func (e *InputText) SetDesc(v string) {
	e.description = v
}

func (e *InputText) SetDefault(v string) {
	e.defValue = v
}

func (e *InputText) SetMaxLength(v int) {
	e.maxLength = v
}

func (e *InputText) SetRequired(v bool) {
	e.required = v
}

func (e *InputText) Name() string {
	return e.name
}

func (e *InputText) Schema() string {
	json := "\n\t\t\"" + e.name + "\": {\n\t\t\t" + "\"type\":\"string\""

	if e.title != "" {
		json += ",\n\t\t\t" + "\"title\":\"" + e.title + "\""
	}

	if e.description != "" {
		json += ",\n\t\t\t" + "\"description\":\"" + e.description + "\""
	}

	if e.defValue != "" {
		json += ",\n\t\t\t" + "\"default\":\"" + e.defValue + "\""
	}

	if e.maxLength != 0 {
		json += ",\n\t\t\t" + "\"maxLength\":" + strconv.Itoa(e.maxLength)
	}

	if e.required != false {
		json += ",\n\t\t\t" + "\"required\":true"
	}

	json += "\n\t\t}"

	return json
}

func (e *InputText) Form() string {
	json := "\n\t\t{\n"
	json += "\t\t\t\"key\":\"" + e.name + "\"\n"
	json += "\t\t}"

	return json
}

type SelectInput struct {
	name        string
	title       string
	description string
	defValue    string
	enum        []string
	required    bool
}

func NewSelectInput(v string) SelectInput {
	return SelectInput{name: v}
}

func (e *SelectInput) SetTitle(v string) {
	e.title = v
}

func (e *SelectInput) SetDesc(v string) {
	e.description = v
}

func (e *SelectInput) SetDefault(v string) {
	e.defValue = v
}

func (e *SelectInput) SetEnum(v []string) {
	e.enum = v
}

func (e *SelectInput) SetRequired(v bool) {
	e.required = v
}

func (e *SelectInput) Name() string {
	return e.name
}

func (e *SelectInput) Schema() string {
	json := "\n\t\t\"" + e.name + "\": {" + "\n\t\t\t\"type\":\"string\""

	if e.title != "" {
		json += ",\n\t\t\t" + "\"title\":\"" + e.title + "\""
	}

	if e.description != "" {
		json += ",\n\t\t\t" + "\"description\":\"" + e.description + "\""
	}

	for _, v := range e.enum {
		if v == e.defValue {
			json += ",\n\t\t\t" + "\"default\":\"" + e.defValue + "\""
		}
	}

	if len(e.enum) != 0 {
		json += ",\n\t\t\t" + "\"enum\": [\n"

		for i, v := range e.enum {
			if i > 0 {
				json += ",\n"
			}
			json += "\t\t\t\t\"" + v + "\""
		}

		json += "\n\t\t\t]"
	}

	if e.required != false {
		json += ",\n\t\t\t" + "\"required\":true\n"
	}

	json += "\t\t}"

	return json
}

func (e *SelectInput) Form() string {
	json := "\n\t\t{\n"
	json += "\t\t\t\"key\":\"" + e.name + "\"\n"
	json += "\t\t}"

	return json
}

type CheckboxInput struct {
	name        string
	title       string
	description string
	defValue    bool
	required    bool
}

func NewCheckboxInput(v string) CheckboxInput {
	return CheckboxInput{name: v}
}

func (e *CheckboxInput) SetTitle(v string) {
	e.title = v
}

func (e *CheckboxInput) SetDesc(v string) {
	e.description = v
}

func (e *CheckboxInput) SetDefault(v bool) {
	e.defValue = v
}

func (e *CheckboxInput) SetRequired(v bool) {
	e.required = v
}

func (e *CheckboxInput) Name() string {
	return e.name
}

func (e *CheckboxInput) Schema() string {
	json := "\n\t\t\"" + e.name + "\": {" + "\n\t\t\t\"type\":\"boolean\""

	if e.title != "" {
		json += ",\n\t\t\t" + "\"title\":\"" + e.title + "\""
	}

	if e.defValue {
		json += ",\n\t\t\t" + "\"default\": true"
	}

	if e.required != false {
		json += ",\n\t\t\t" + "\"required\":true"
	}

	json += "\n\t\t}"

	return json
}

func (e *CheckboxInput) Form() string {
	json := ""

	if e.description != "" {
		json += "\n\t\t{\n"
		json += "\t\t\t\"key\":\"" + e.name + "\",\n"
		json += "\t\t\t\"inlinetitle\":\"" + e.description + "\"\n"
		json += "\t\t}\n"
	}

	return json
}

type ArrayInput struct {
	name        string
	title       string
	description string
	elem        []FormItem
}

func NewArrayInput(v string) ArrayInput {
	return ArrayInput{name: v}
}

func (e *ArrayInput) SetTitle(v string) {
	e.title = v
}

func (e *ArrayInput) SetDesc(v string) {
	e.description = v
}

func (e *ArrayInput) AddItem(v FormItem) {
	e.elem = append(e.elem, v)
}

func (e *ArrayInput) Name() string {
	return e.name
}

func (e *ArrayInput) Schema() string {
	json := "\n\t\t\"" + e.name + "\": {" + "\n\t\t\t\"type\":\"array\""

	// if e.title != "" {
	// 	json += ",\n\t\t\t" + "\"title\":\"" + e.title + "\""
	// }

	if e.description != "" {
		json += ",\n\t\t\t" + "\"description\":\"" + e.description + "\""
	}

	if len(e.elem) > 0 {
		json += ",\n\t\t\t\"items\": {\n"
		json += "\t\t\t\t\"type\":\"object\",\n"
		json += "\t\t\t\t\"title\":\"" + e.title + "\",\n"
		json += "\t\t\t\t\"properties\": {\n"

		for _, v := range e.elem {
			json += v.Schema()
		}

		json += "\n\t\t\t\t}\n\t\t\t}\n"
	}

	json += "}"

	return json
}

func (e *ArrayInput) Form() string {
	//	json += "\t\t{\n"
	json := "\n\t\t{\n"
	json += "\t\t\t\"type\":\"array\",\n"
	json += "\t\t\t\"items\": {\n"
	json += "\t\t\t\t\"type\": \"section\",\n"
	json += "\t\t\t\t\"items\": ["

	for i, v := range e.elem {
		if i > 0 {
			json += ","
		}

		json += "\n\t\t\t\t\t\"" + e.name + "[]." + v.Name() + "\""
	}

	json += "\n\t\t\t\t]\n"
	json += "\t\t\t}\n"
	json += "\t\t}"

	return json
}

type SubFormInput struct {
	name        string
	title       string
	description string
	elem        []FormItem
}

func NewSubFormInput(v string) SubFormInput {
	return SubFormInput{name: v}
}

func (e *SubFormInput) SetTitle(v string) {
	e.title = v
}

func (e *SubFormInput) SetDesc(v string) {
	e.description = v
}

func (e *SubFormInput) AddItem(v FormItem) {
	e.elem = append(e.elem, v)
}

func (e *SubFormInput) Name() string {
	return e.name
}

func (e *SubFormInput) Schema() string {
	json := "\n\t\t\"" + e.name + "\": {" + "\n\t\t\t\"type\":\"string\""

	if e.title != "" {
		json += ",\n\t\t\t" + "\"title\":\"" + e.title + "\""
	}

	if e.description != "" {
		json += ",\n\t\t\t" + "\"description\":\"" + e.description + "\""
	}

	json += ",\n\t\t\t\"enum\": [\n"
	json += "\t\t\t\t\"\""
	for _, v := range e.elem {
		json += ",\n\t\t\t\t\"" + v.Name() + "\""
	}
	json += "\n\t\t\t]\n\t\t}"

	if len(e.elem) > 0 {
		for _, v := range e.elem {
			json += "," + v.Schema()
		}
	}

	return json
}

func (e *SubFormInput) Form() string {
	if len(e.elem) > 0 {
		json := "\n\t\t{"
		json += "\n\t\t\t\"type\":\"selectfieldset\""
		json += ",\n\t\t\t\"key\":\"" + e.name + "\""

		if e.title != "" {
			json += ",\n\t\t\t\"title\":\"" + e.title + "\""
		}

		json += ",\n\t\t\t\"items\": [\n"
		json += "\t\t\t\t\"\""
		for _, v := range e.elem {
			json += ",\n\t\t\t\t\"" + v.Name() + "\""
		}

		json += "\n\t\t\t]\n"

		json += "\t\t}\n"

		return json
	}

	return ""
}
