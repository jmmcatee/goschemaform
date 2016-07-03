package goschemaform

import (
	"bytes"
	"log"
	"strconv"
	"text/template"
)

// NewFileInput returns an initialized FileInput with the required key
// field set.
func NewFileInput(key string) *FileInput {
	return &FileInput{
		key:         key,
		title:       "",
		placeHolder: "",
		maxSize:     0,
		required:    false,
		condition:   "",
		condFlip:    false,
	}
}

// FileInput is a structure for a single line or multiline input field. The
// default FileInput structure is a single line input. You can use SetMultiline
// to change the input to a multiple line text input field (textarea).
type FileInput struct {
	key         string
	title       string
	placeHolder string
	maxSize     int
	required    bool
	condition   string
	condFlip    bool
}

// Form retuns the "form" section of the JSON Schema Form definition
func (t *FileInput) Form() string {
	// Compile the template for generating the Form section
	var tmplForm = template.Must(template.New("form").Parse(tmplFileInputForm))

	var cCheck = false
	if t.condition != "" {
		cCheck = true
	}

	tmplData := struct {
		Key            string
		PlaceHolder    string
		ConditionCheck bool
		Condition      string
		ConditionFlip  bool
	}{
		t.key,
		t.placeHolder,
		cCheck,
		t.condition,
		t.condFlip,
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
func (t *FileInput) Schema() string {
	// Compile the template for generating the Schema section
	var tmplSchema = template.Must(template.New("schema").Parse(tmplFileInputSchema))

	tmplData := struct {
		Key           string
		Title         string
		MaxSize       int
		MaxSizeString string
	}{
		t.key,
		t.title,
		t.maxSize,
		strconv.Itoa(t.maxSize),
	}

	schema := bytes.NewBuffer([]byte{})

	err := tmplSchema.Execute(schema, tmplData)
	if err != nil {
		log.Println(err.Error())
		return ""
	}

	return schema.String()
}

// Inputs returns the FileInput structure, which matches the Input interface.
// It only returns itself in the array as it holds no other Inputs.
func (t *FileInput) Inputs() []Input {
	return []Input{t}
}

// Key returns the name/id of the form input that will be used.
func (t *FileInput) Key() string {
	return t.key
}

// Required returns a boolean value to determine if the Input is required.
func (t *FileInput) Required() bool {
	return t.required
}

// MapDefinition returns the internal structure of the Input in a Go friendly
// format.
func (t *FileInput) MapDefinition() map[string]string {
	return map[string]string{
		"key":         t.key,
		"title":       t.title,
		"placeHolder": t.placeHolder,
		"maxSize":     strconv.Itoa(t.maxSize),
		"required":    strconv.FormatBool(t.required),
		"type":        t.getType(),
	}
}

// getType returns the type of the Input (text or textarea)
func (t *FileInput) getType() string {
	return "string"
}

// SetTitle sets the Input's title text
func (t *FileInput) SetTitle(title string) {
	t.title = title
}

// SetMaxSize sets the maximum length of the text input
func (t *FileInput) SetMaxSize(max int) {
	t.maxSize = max
}

// SetPlaceHolder configures the temporary text that is set in the text field
// before a use begins to type into it.
func (t *FileInput) SetPlaceHolder(text string) {
	t.placeHolder = text
}

// IsRequired configures the input to be required (true) or to not be required
// (false).
func (t *FileInput) IsRequired(req bool) {
	t.required = req
}

// SetCondition will set whether this item displays on the form based on if the provided
// key has a value or not. You can reverse the behaivor with the defaultHide switch. False
// for this option is the default and will make something only appear if the condition is
// is set, while true flips this and shows the control until the condition is met.
func (t *FileInput) SetCondition(text string, defaultHide bool) {
	t.condition = text
	t.condFlip = defaultHide
}
