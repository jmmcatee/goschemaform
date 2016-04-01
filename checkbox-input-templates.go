package goschemaform

const tmplCheckBoxForm = `
{
    "key": "{{.Key}}",
    "type": "checkbox"
}
`

const tmplCheckBoxSchema = `
        "{{.Key}}": {
            "type": "boolean",
            "title": "{{.Title}}"
        }
`
