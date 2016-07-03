package goschemaform

import (
	"bytes"
	"log"
	"strconv"
	"text/template"
)

// NewDropDownInput returns an initialized DropDownInput
func NewDropDownInput(key string) *DropDownInput {
	return &DropDownInput{
		key:       key,
		title:     "",
		required:  false,
		enum:      []DropDownInputOption{},
		condition: "",
		condFlip:  false,
	}
}

// DropDownInput represents a form dropdown or single select
type DropDownInput struct {
	key       string
	title     string
	required  bool
	enum      []DropDownInputOption
	condition string
	condFlip  bool
}

// options is a structure used in processing the templates
type options struct {
	Value string
	Name  string
	Group string
}

// Form returns the "form" section of the JSON Schema Form definition
func (d *DropDownInput) Form() string {
	// Compile the template for generating the form section
	var tmplForm = template.Must(template.New("form").Parse(tmplDropDownForm))

	var cCheck = false
	if d.condition != "" {
		cCheck = true
	}

	tmplData := struct {
		Key            string
		Enum           []options
		ConditionCheck bool
		Condition      string
		ConditionFlip  bool
	}{}

	tmplData.Key = d.key
	tmplData.ConditionCheck = cCheck
	tmplData.Condition = d.condition
	tmplData.ConditionFlip = d.condFlip
	for i := range d.enum {
		o := options{
			Value: d.enum[i].value,
			Name:  d.enum[i].name,
			Group: d.enum[i].group,
		}

		tmplData.Enum = append(tmplData.Enum, o)
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
func (d *DropDownInput) Schema() string {
	// Compile the template for generating the Schema section
	var tmplSchema = template.Must(template.New("schema").Parse(tmplDropDownSchema))

	tmplData := struct {
		Key   string
		Title string
		Enum  []options
	}{
		d.key,
		d.title,
		[]options{},
	}

	for i := range d.enum {
		o := options{
			Value: d.enum[i].value,
			Name:  d.enum[i].name,
			Group: d.enum[i].group,
		}

		tmplData.Enum = append(tmplData.Enum, o)
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
func (d *DropDownInput) Inputs() []Input {
	return []Input{d}
}

// Key returns the name/id of the form input that will be used
func (d *DropDownInput) Key() string {
	return d.key
}

// Required returns a boolean value to determine if the Input is required
func (d *DropDownInput) Required() bool {
	return d.required
}

// MapDefinition returns the internal structure of the Input in a Go friendly
// format.
func (d *DropDownInput) MapDefinition() map[string]string {
	return map[string]string{
		"key":      d.key,
		"title":    d.title,
		"required": strconv.FormatBool(d.required),
		"type":     "select",
		"enum":     d.serialEnum(),
	}
}

// serialEnum returns a : seperated list of the value portion of the enum array
func (d *DropDownInput) serialEnum() string {
	var serial string
	for i := range d.enum {
		serial += d.enum[i].value

		if i < len(d.enum)-1 {
			serial += ":"
		}
	}

	return serial
}

// SetTitle sets the Input's title text
func (d *DropDownInput) SetTitle(title string) {
	d.title = title
}

// IsRequired configures the input to be required (true) or to not be required
// (false).
func (d *DropDownInput) IsRequired(req bool) {
	d.required = req
}

// AddOption adds a DropDownInputOption to the DropDownInput
func (d *DropDownInput) AddOption(option *DropDownInputOption) {
	d.enum = append(d.enum, *option)
}

// SetConidition will set whether this item displays on the form based on if the provided
// key has a value or not. You can reverse the behaivor with the defaultHide switch. False
// for this option is the default and will make something only appear if the condition is
// is set, while true flips this and shows the control until the condition is met.
func (d *DropDownInput) SetConidition(text string, defaultHide bool) {
	d.condition = text
	d.condFlip = defaultHide
}

// NewDropDownInputOption returns an initialized DropDownInputOption
func NewDropDownInputOption(value string) *DropDownInputOption {
	return &DropDownInputOption{
		value: value,
		name:  "",
		group: "",
	}
}

// DropDownInputOption is an option added to a DropDownInput
type DropDownInputOption struct {
	value string
	name  string
	group string
}

// SetName sets the display name of the option
func (d *DropDownInputOption) SetName(name string) {
	d.name = name
}

// SetGroup sets the group for the options to be put in
func (d *DropDownInputOption) SetGroup(group string) {
	d.group = group
}
