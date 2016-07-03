package goschemaform

const tmplFileInputForm = `
{
    "key": "{{.Key}}",
    "placeholder": "{{.PlaceHolder}}"{{if .ConditionCheck}},
    "condition": "{{if .ConditionFlip}}!{{end}}model.{{.Condition}}"{{end}}
}
`

const tmplFileInputSchema = `
"{{.Key}}": {
    "title": "{{.Title}}",
    "type": "string",
    "format": "base64"{{if .MaxSize}},
    "maxSize": "{{.MaxSizeString}}"{{end}}
}
`
