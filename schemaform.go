package goschemaform

// SchemaForm is a struct that represents the whole Schema Form structure.
type SchemaForm struct {
	elem []Element
}

// Element is an interface representing some form element such as
// an input or fieldset. The Form() and Schema() functions return the JSON
// string defining the element. The Inputs() function is used to
// return all the Inputs of the given form. This would exclude
// structure items such as fieldsets.
type Element interface {
	Form() string
	Schema() string
	Inputs() []Input
}

// Input is an interface for form items that are input fields. This
// would exclude things such as fieldsets and tabs. The Key() method
// returns the form name/id/key that is used to identify the input.
// The Required() function returns whether an element is required,
// which is needed when writing the final form structure. Finally,
// MapDefinition returns a key value set of parameters related to
// the input such as the type, key, enum fields, etc. This can be used
// to get information in a more Go friendly fashion than the JSON Form
// Schema.
type Input interface {
	Key() string
	Required() bool
	MapDefinition() map[string]string
}
