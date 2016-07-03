package goschemaform

// element is a structure for executing the tab fieldset templates
type element struct {
	Form   string
	Schema string
}

const tmplTabFieldsetForm = `
{
    "type": "fieldset",
    "title": "{{.Title}}",
    "items": [
        {
            "type": "tabs",
            "tabs": [{{range $index, $element := .Elem}}{{if $index}},{{end}}{{$element.Form}}{{end}}]
        }
    ]{{if .ConditionCheck}},
    "condition": "{{if .ConditionFlip}}!{{end}}model.{{.Condition}}"{{end}}
}
`

const tmplTabFieldsetSchema = `
    {{range $index, $element := .Elem}}{{if $index}},{{end}}{{$element.Schema}}{{end}}
`

const tmplTabForm = `
                {
                    "title": "{{.Title}}",
                    "items": [{{range $index, $element := .Elem}}{{if $index}},{{end}}{{$element.Form}}{{end}}
                    ]
                }`

const tmplTabSchema = `
        {{range $index, $element := .Elem}}{{if $index}},{{end}}{{$element.Schema}}{{end}}
`
