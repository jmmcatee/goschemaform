package goschemaform

import (
	"bytes"
	"log"
	"text/template"
)

// NewTabFieldset returns an initialized TabFieldset
func NewTabFieldset() *TabFieldset {
	return &TabFieldset{
		title: "",
		elem:  []Tab{},
	}
}

// TabFieldset represents the fieldset that holds multiple tabs
type TabFieldset struct {
	title string
	elem  []Tab
}

// Form returns the "form" section of the JSON Schema Form definition
func (f *TabFieldset) Form() string {
	// Compile the template for generating the form section
	var tmplForm = template.Must(template.New("form").Parse(tmplTabFieldsetForm))

	tmplData := struct {
		Title string
		Elem  []element
	}{
		Title: f.title,
		Elem:  []element{},
	}

	for i := range f.elem {
		e := element{
			Form: f.elem[i].Form(),
		}

		tmplData.Elem = append(tmplData.Elem, e)
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
func (f *TabFieldset) Schema() string {
	// Compile the template for generating the Schema section
	var tmplSchema = template.Must(template.New("schema").Parse(tmplTabFieldsetSchema))

	tmplData := struct {
		Title string
		Elem  []element
	}{
		Title: f.title,
		Elem:  []element{},
	}

	for i := range f.elem {
		e := element{
			Schema: f.elem[i].Schema(),
		}

		tmplData.Elem = append(tmplData.Elem, e)
	}

	schema := bytes.NewBuffer([]byte{})

	err := tmplSchema.Execute(schema, tmplData)
	if err != nil {
		log.Println(err.Error())
		return ""
	}

	return schema.String()
}

// Inputs returns all Inputs from the underlying tabs of the fieldset.
func (f *TabFieldset) Inputs() []Input {
	var inputs []Input
	for _, v := range f.elem {
		inputs = append(inputs, v.Inputs()...)
	}

	return inputs
}

// SetTitle sets the fieldset title
func (f *TabFieldset) SetTitle(title string) {
	f.title = title
}

// AddTab adds a tab to the fieldset
func (f *TabFieldset) AddTab(tab *Tab) {
	f.elem = append(f.elem, *tab)
}

// NewTab returns an initialized Tab
func NewTab() *Tab {
	return &Tab{
		title: "",
		elem:  []Element{},
	}
}

// Tab represents an individual Tab in a Fieldset
type Tab struct {
	title string
	elem  []Element
}

// Form returns the "form" section of the JSON Schema Form definition
func (t *Tab) Form() string {
	// Compile the template for generating the form section
	var tmplForm = template.Must(template.New("form").Parse(tmplTabForm))

	tmplData := struct {
		Title string
		Elem  []element
	}{
		Title: t.title,
		Elem:  []element{},
	}

	for i := range t.elem {
		e := element{
			Form: t.elem[i].Form(),
		}

		tmplData.Elem = append(tmplData.Elem, e)
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
func (t *Tab) Schema() string {
	// Compile the template for generating the Schema section
	var tmplSchema = template.Must(template.New("schema").Parse(tmplTabSchema))

	tmplData := struct {
		Title string
		Elem  []element
	}{
		Title: t.title,
		Elem:  []element{},
	}

	for i := range t.elem {
		e := element{
			Schema: t.elem[i].Schema(),
		}

		tmplData.Elem = append(tmplData.Elem, e)
	}

	schema := bytes.NewBuffer([]byte{})

	err := tmplSchema.Execute(schema, tmplData)
	if err != nil {
		log.Println(err.Error())
		return ""
	}

	return schema.String()
}

// Inputs returns all Inputs added to this Tab
func (t *Tab) Inputs() []Input {
	var inputs []Input
	for _, v := range t.elem {
		inputs = append(inputs, v.Inputs()...)
	}

	return inputs
}

// SetTitle sets the Tab title text
func (t *Tab) SetTitle(title string) {
	t.title = title
}

// AddElement adds a form Input to the Tab
func (t *Tab) AddElement(el Element) {
	t.elem = append(t.elem, el)
}
