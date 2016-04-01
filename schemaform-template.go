package goschemaform

const tmplSchemaForm = `
{
    "form": [{{range $index, $element := .Form}}{{if $index}},{{end}}{{$element}}{{end}}
    ],
    "schema": {
        "type": "object",
        "title": "{{.Title}}",
        "properties": { {{range $index, $element := .Schema}}{{if $index}},{{end}}{{$element}}{{end}}
        },
        "required": [
            {{range $index, $element := .Required}}{{if $index}},
            {{end}}"{{$element}}"{{end}}
        ]
    }
}
`
