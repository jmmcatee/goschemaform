package goschemaform

const tmplDropDownForm = `
{
    "key": "{{.Key}}",
    "type": "select"{{if len .Enum}},
    "titleMap": [
        {{range $index, $element := .Enum}}{{if $index}},{{end}}
        {
            "value": "{{.Value}}",
            "name": {{if $element.Name}}"{{$element.Name}}"{{else}}"{{$element.Value}}"{{end}}{{if $element.Group}},
            "group": "{{$element.Group}}"{{end}}
        }{{end}}
    ]{{end}}
}`

const tmplDropDownSchema = `
        "{{.Key}}": {
            "title": "{{.Title}}",
            "type": "string",
            "enum": [
                {{range $index, $element := .Enum}}{{if $index}},
                {{end}}"{{$element.Value}}"{{end}}
            ]
        }`
