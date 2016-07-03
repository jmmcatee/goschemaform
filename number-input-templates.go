package goschemaform

const tmplNumberInputForm = `
{
    "key": "{{.Key}}",
    "type": "number"{{if .ConditionCheck}},
    "condition": "{{if .ConditionFlip}}!{{end}}model.{{.Condition}}"{{end}}
}
`

const tmplNumberInputSchema = `
"{{.Key}}": {
    "title": "{{.Title}}",
    "type": "number"{{if .Max}},
    "maximum": {{.MaxString}}{{end}}{{if .Min}},
    "minimum": {{.MinString}}{{end}}
}
`
