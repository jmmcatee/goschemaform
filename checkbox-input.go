package goschemaform

import (
	"bytes"
	"log"
	"strconv"
	"text/template"
)

// NewCheckBoxInput returns an initialized CheckBoxInput
func NewCheckBoxInput(key string) *CheckBoxInput {
	return &CheckBoxInput{
		key:       key,
		title:     "",
		required:  false,
		condition: "",
	}
}

// CheckBoxInput represents a form checkxbox
type CheckBoxInput struct {
	key       string
	title     string
	required  bool
	condition string
}

// Form returns the "form" section of the JSON Schema Form definition
func (c *CheckBoxInput) Form() string {
	// Compile the template for generating the form section
	var tmplForm = template.Must(template.New("form").Parse(tmplCheckBoxForm))

	var cCheck = false
	if c.condition != "" {
		cCheck = true
	}

	tmplData := struct {
		Key            string
		Title          string
		ConditionCheck bool
		Condition      string
	}{
		Key:            c.key,
		Title:          c.title,
		ConditionCheck: cCheck,
		Condition:      c.condition,
	}

	form := bytes.NewBuffer([]byte{})

	err := tmplForm.Execute(form, tmplData)
	if err != nil {
		log.Println(err.Error())
		return ""
	}

	return form.String()
}

// Schema returns the "schema" section of the JSON Schema Form definition
func (c *CheckBoxInput) Schema() string {
	// Compile the template for generating the Schema section
	var tmplSchema = template.Must(template.New("schema").Parse(tmplCheckBoxSchema))

	tmplData := struct {
		Key   string
		Title string
	}{
		Key:   c.key,
		Title: c.title,
	}

	schema := bytes.NewBuffer([]byte{})

	err := tmplSchema.Execute(schema, tmplData)
	if err != nil {
		log.Println(err.Error())
		return ""
	}

	return schema.String()
}

// Inputs returns the DropDownInput structure, which matches the Input interface.
// It only returns itself in the array as it holds no other Inputs.
func (c *CheckBoxInput) Inputs() []Input {
	return []Input{c}
}

// Key returns the name/id of the form input that will be used
func (c *CheckBoxInput) Key() string {
	return c.key
}

// Required returns a boolean value to determine if the Input is required
func (c *CheckBoxInput) Required() bool {
	return c.required
}

// MapDefinition returns the internal structure of the Input in a Go friendly
// format.
func (c *CheckBoxInput) MapDefinition() map[string]string {
	return map[string]string{
		"key":      c.key,
		"title":    c.title,
		"required": strconv.FormatBool(c.required),
		"type":     "checkbox",
	}
}

// SetTitle sets the Input's title text
func (c *CheckBoxInput) SetTitle(title string) {
	c.title = title
}

// IsRequired configures the input to be required (true) or to not be required
// (false).
func (c *CheckBoxInput) IsRequired(req bool) {
	c.required = req
}

// SetConidition will set whether this item displays on the form based on if the provided
// key has a value or not.
func (c *CheckBoxInput) SetConidition(text string) {
	c.condition = text
}
