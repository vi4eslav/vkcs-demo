package main

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

type Schema struct {
	Type        schema.ValueType
	Elem        interface{} // Elem represents the element type for a TypeList, TypeSet, or TypeMap
	Description string

	ConfigMode       schema.SchemaConfigMode // pass by
	Required         bool
	Optional         bool
	MaxItems         int
	MinItems         int
	Set              schema.SchemaSetFunc // Set defines custom hash algorithm for each TypeSet element.
	ExactlyOneOf     []string
	AtLeastOneOf     []string
	RequiredWith     []string
	ComputedWhen     []string // Deprecated
	ConflictsWith    []string
	ValidateFunc     schema.SchemaValidateFunc
	ValidateDiagFunc schema.SchemaValidateDiagFunc // multi errors => diag

	Computed              bool
	ForceNew              bool
	DiffSuppressFunc      schema.SchemaDiffSuppressFunc
	DiffSuppressOnRefresh bool

	Default      interface{}
	DefaultFunc  schema.SchemaDefaultFunc
	InputDefault string

	Deprecated string
	Sensitive  bool                   // backed is not scripted
	StateFunc  schema.SchemaStateFunc // transform before save to state
}
