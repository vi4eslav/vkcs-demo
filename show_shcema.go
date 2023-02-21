package main

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

type Schema struct {
	Type        schema.ValueType
	Elem        interface{}
	Description string

	ConfigMode       schema.SchemaConfigMode // pass by
	Required         bool
	Optional         bool
	Computed         bool
	MaxItems         int
	MinItems         int
	Set              schema.SchemaSetFunc
	ComputedWhen     []string // Deprecated
	ConflictsWith    []string
	ExactlyOneOf     []string
	AtLeastOneOf     []string
	RequiredWith     []string
	ValidateFunc     schema.SchemaValidateFunc
	ValidateDiagFunc schema.SchemaValidateDiagFunc // multi errors => diag

	ForceNew              bool
	DiffSuppressFunc      schema.SchemaDiffSuppressFunc
	DiffSuppressOnRefresh bool

	Default      interface{}
	DefaultFunc  schema.SchemaDefaultFunc
	InputDefault string

	Deprecated string
	Sensitive  bool
	StateFunc  schema.SchemaStateFunc // transform before save to state
}
