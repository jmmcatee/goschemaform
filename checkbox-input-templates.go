package goschemaform

const tmplCheckBoxForm = `
{
    "key": "{{.Key}}",
    "type": "checkbox"{{if .ConditionCheck}},
    "condition": "{{if .ConditionFlip}}!{{end}}model.{{.Condition}}"{{end}}
}`

const tmplCheckBoxSchema = `
        "{{.Key}}": {
            "type": "boolean",
            "title": "{{.Title}}"
        }`
