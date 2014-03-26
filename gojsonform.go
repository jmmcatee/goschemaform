package gojsonform

import ()

type Form struct {
	elements []FormItem
}

func (f *Form) Prepend(item FormItem) {
	newElem := []FormItem{}

	newElem = append(newElem, item)

	newElem = append(newElem, f.elements...)

	f.elements = newElem
}

func (f *Form) Append(item FormItem) {
	f.elements = append(f.elements, item)
}

func (f *Form) GetForm() string {
	schema := "\t\"schema\": {"
	form := "\t\"form\": ["

	for i, v := range f.elements {
		if i > 0 {
			schema += ","
			form += ","
		}

		schema += v.Schema()
		form += v.Form()
	}

	schema += "\n\t}"
	form += "\t\t\"*\": {"
	form += "\n\t\t\t\"type\":\"submit\","
	form += "\n\t\t\t\"title\":\"Submit\""
	form += "\n\t\t},"
	form += "\n\t]"

	return "{\n" + schema + ",\n" + form + "\n}\n"
}

type FormItem interface {
	Name() string
	Schema() string
	Form() string
}
