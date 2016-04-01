package goschemaform

const tmplTextInputForm = `
{
    "key": "{{.Key}}",
    "type": "{{.Type}}",
    "placeholder": "{{.PlaceHolder}}"
}
`

const tmplTextInputSchema = `
"{{.Key}}": {
    "title": "{{.Title}}",
    "type": "string"{{if .MaxLength}},
    "maxLength": "{{.MaxLengthString}}"{{end}}
}
`
