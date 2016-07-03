package goschemaform

import (
	"bytes"
	"log"
	"strconv"
	"text/template"
)

// NewTextInput returns an initialized TextInput with the required key
// field set.
func NewTextInput(key string) *TextInput {
	return &TextInput{
		key:           key,
		multiline:     false,
		title:         "",
		placeHolder:   "",
		maxLength:     0,
		required:      false,
		condition:     "",
		conditionFlip: false,
	}
}

// TextInput is a structure for a single line or multiline input field. The
// default TextInput structure is a single line input. You can use SetMultiline
// to change the input to a multiple line text input field (textarea).
type TextInput struct {
	key           string
	multiline     bool
	title         string
	placeHolder   string
	maxLength     int
	required      bool
	condition     string
	conditionFlip bool
}

// Form retuns the "form" section of the JSON Schema Form definition
func (t *TextInput) Form() string {
	// Compile the template for generating the Form section
	var tmplForm = template.Must(template.New("form").Parse(tmplTextInputForm))

	var cCheck = false
	if t.condition != "" {
		cCheck = true
	}

	tmplData := struct {
		Key            string
		Type           string
		PlaceHolder    string
		ConditionCheck bool
		Condition      string
		ConditionFlip  bool
	}{
		t.key,
		t.getType(),
		t.placeHolder,
		cCheck,
		t.condition,
		t.conditionFlip,
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
func (t *TextInput) Schema() string {
	// Compile the template for generating the Schema section
	var tmplSchema = template.Must(template.New("schema").Parse(tmplTextInputSchema))

	tmplData := struct {
		Key             string
		Title           string
		MaxLength       int
		MaxLengthString string
	}{
		t.key,
		t.title,
		t.maxLength,
		strconv.Itoa(t.maxLength),
	}

	schema := bytes.NewBuffer([]byte{})

	err := tmplSchema.Execute(schema, tmplData)
	if err != nil {
		log.Println(err.Error())
		return ""
	}

	return schema.String()
}

// Inputs returns the TextInput structure, which matches the Input interface.
// It only returns itself in the array as it holds no other Inputs.
func (t *TextInput) Inputs() []Input {
	return []Input{t}
}

// Key returns the name/id of the form input that will be used.
func (t *TextInput) Key() string {
	return t.key
}

// Required returns a boolean value to determine if the Input is required.
func (t *TextInput) Required() bool {
	return t.required
}

// MapDefinition returns the internal structure of the Input in a Go friendly
// format.
func (t *TextInput) MapDefinition() map[string]string {
	return map[string]string{
		"key":         t.key,
		"multiline":   strconv.FormatBool(t.multiline),
		"title":       t.title,
		"placeHolder": t.placeHolder,
		"maxLength":   strconv.Itoa(t.maxLength),
		"required":    strconv.FormatBool(t.required),
		"type":        t.getType(),
	}
}

// getType returns the type of the Input (text or textarea)
func (t *TextInput) getType() string {
	var inputType string
	if t.multiline {
		inputType = "textarea"
	} else {
		inputType = "text"
	}

	return inputType
}

// SetTitle sets the Input's title text
func (t *TextInput) SetTitle(title string) {
	t.title = title
}

// SetMaxLength sets the maximum length of the text input
func (t *TextInput) SetMaxLength(max int) {
	t.maxLength = max
}

// SetMultiline configures the TextInput as a multiline (true) or a single line
// text input (false).
func (t *TextInput) SetMultiline(multi bool) {
	t.multiline = multi
}

// SetPlaceHolder configures the temporary text that is set in the text field
// before a use begins to type into it.
func (t *TextInput) SetPlaceHolder(text string) {
	t.placeHolder = text
}

// IsRequired configures the input to be required (true) or to not be required
// (false).
func (t *TextInput) IsRequired(req bool) {
	t.required = req
}

// SetConidition will set whether this item displays on the form based on if the provided
// key has a value or not. You can reverse the behaivor with the defaultHide switch. False
// for this option is the default and will make something only appear if the condition is
// is set, while true flips this and shows the control until the condition is met.
func (t *TextInput) SetConidition(text string, defaultHide bool) {
	t.condition = text
	t.conditionFlip = defaultHide
}
