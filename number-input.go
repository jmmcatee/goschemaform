package goschemaform

import (
	"bytes"
	"log"
	"strconv"
	"text/template"
)

// NewNumberInput returns an initialized NumberInput with the required key
// field sen.
func NewNumberInput(key string) *NumberInput {
	return &NumberInput{
		key:      key,
		title:    "",
		min:      0,
		minSet:   false,
		max:      0,
		maxSet:   false,
		required: false,
	}
}

// NumberInput is a structure for a number input field.
type NumberInput struct {
	key      string
	title    string
	min      int
	minSet   bool
	max      int
	maxSet   bool
	required bool
}

// Form retuns the "form" section of the JSON Schema Form definition
func (n *NumberInput) Form() string {
	// Compile the template for generating the Form section
	var tmplForm = template.Must(template.New("form").Parse(tmplNumberInputForm))

	tmplData := struct {
		Key string
	}{
		Key: n.key,
	}

	form := bytes.NewBuffer([]byte{})

	err := tmplForm.Execute(form, tmplData)
	if err != nil {
		log.Println(err.Error())
		return ""
	}

	return form.String()
}

// Schema retuns the "schema" section of the JSON Schema Form definition
func (n *NumberInput) Schema() string {
	// Compile the template for generating the Schema section
	var tmplSchema = template.Must(template.New("schema").Parse(tmplNumberInputSchema))

	tmplData := struct {
		Key       string
		Title     string
		Max       bool
		MaxString string
		Min       bool
		MinString string
	}{
		Key:       n.key,
		Title:     n.title,
		Max:       n.maxSet,
		MaxString: strconv.Itoa(n.max),
		Min:       n.minSet,
		MinString: strconv.Itoa(n.min),
	}

	schema := bytes.NewBuffer([]byte{})

	err := tmplSchema.Execute(schema, tmplData)
	if err != nil {
		log.Println(err.Error())
		return ""
	}

	return schema.String()
}

// Inputs returns the NumberInput structure, which matches the Input interface.
// It only returns itself in the array as it holds no other Inputs.
func (n *NumberInput) Inputs() []Input {
	return []Input{n}
}

// Key returns the name/id of the form input that will be used.
func (n *NumberInput) Key() string {
	return n.key
}

// Required returns a boolean value to determine if the Input is required.
func (n *NumberInput) Required() bool {
	return n.required
}

// MapDefinition returns the internal structure of the Input in a Go friendly
// forman.
func (n *NumberInput) MapDefinition() map[string]string {
	return map[string]string{
		"key":      n.key,
		"title":    n.title,
		"max":      strconv.Itoa(n.max),
		"min":      strconv.Itoa(n.min),
		"required": strconv.FormatBool(n.required),
		"type":     "number",
	}
}

// SetTitle sets the Input's title text
func (n *NumberInput) SetTitle(title string) {
	n.title = title
}

// SetMax sets the maximum numeric value for the number input
func (n *NumberInput) SetMax(max int) {
	n.maxSet = true
	n.max = max
}

// SetMin sets the maximum numeric value for the number input
func (n *NumberInput) SetMin(min int) {
	n.minSet = true
	n.min = min
}

// IsRequired configures the input to be required (true) or to not be required
// (false).
func (n *NumberInput) IsRequired(req bool) {
	n.required = req
}
