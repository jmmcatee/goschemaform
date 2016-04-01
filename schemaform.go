package goschemaform

import (
	"bytes"
	"log"
	"text/template"
)

// NewSchemaForm returns an initialized SchemaForm
func NewSchemaForm() *SchemaForm {
	return &SchemaForm{
		elem: []Element{},
	}
}

// SchemaForm is a struct that represents the whole Schema Form structure.
type SchemaForm struct {
	title string
	elem  []Element
}

// SchemaForm returns the JSON Schema definition
func (s *SchemaForm) SchemaForm() string {
	var tmpl = template.Must(template.New("schemaform").Parse(tmplSchemaForm))

	tmplData := struct {
		Title    string
		Form     []string
		Schema   []string
		Required []string
	}{
		Title:    s.title,
		Form:     []string{},
		Schema:   []string{},
		Required: []string{},
	}

	for ei := range s.elem {
		tmplData.Form = append(tmplData.Form, s.elem[ei].Form())
		tmplData.Schema = append(tmplData.Schema, s.elem[ei].Schema())

		inputs := s.elem[ei].Inputs()
		for ii := range inputs {
			if inputs[ii].Required() {
				tmplData.Required = append(tmplData.Required, inputs[ii].Key())
			}
		}
	}

	schemaForm := bytes.NewBuffer([]byte{})

	err := tmpl.Execute(schemaForm, tmplData)
	if err != nil {
		log.Println(err.Error())
		return ""
	}

	return schemaForm.String()
}

// MapDefinitions returns a slice of the input definitions
func (s *SchemaForm) MapDefinitions() map[string]map[string]string {
	var defs = make(map[string]map[string]string)

	for ei := range s.elem {
		inputs := s.elem[ei].Inputs()
		for ii := range inputs {
			defs[inputs[ii].Key()] = inputs[ii].MapDefinition()
		}
	}

	return defs
}

// SetTitle sets the form title
func (s *SchemaForm) SetTitle(title string) {
	s.title = title
}

// AddElement adds a form element to the schema form
func (s *SchemaForm) AddElement(elem Element) {
	s.elem = append(s.elem, elem)
}

// Element is an interface representing some form element such as
// an input or fieldset. The Form() and Schema() functions return the JSON
// string defining the element. The Inputs() function is used to
// return all the Inputs of the given form. This would exclude
// structure items such as fieldsets.
type Element interface {
	Form() string
	Schema() string
	Inputs() []Input
}

// Input is an interface for form items that are input fields. This
// would exclude things such as fieldsets and tabs. The Key() method
// returns the form name/id/key that is used to identify the input.
// The Required() function returns whether an element is required,
// which is needed when writing the final form structure. Finally,
// MapDefinition returns a key value set of parameters related to
// the input such as the type, key, enum fields, etc. This can be used
// to get information in a more Go friendly fashion than the JSON Form
// Schema.
type Input interface {
	Key() string
	Required() bool
	MapDefinition() map[string]string
}
