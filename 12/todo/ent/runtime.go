// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"
	"todo/ent/schema"
	"todo/ent/todo"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	todoFields := schema.Todo{}.Fields()
	_ = todoFields
	// todoDescText is the schema descriptor for text field.
	todoDescText := todoFields[0].Descriptor()
	// todo.TextValidator is a validator for the "text" field. It is called by the builders before save.
	todo.TextValidator = todoDescText.Validators[0].(func(string) error)
	// todoDescCreatedAt is the schema descriptor for created_at field.
	todoDescCreatedAt := todoFields[1].Descriptor()
	// todo.DefaultCreatedAt holds the default value on creation for the created_at field.
	todo.DefaultCreatedAt = todoDescCreatedAt.Default.(func() time.Time)
	// todoDescPriority is the schema descriptor for priority field.
	todoDescPriority := todoFields[3].Descriptor()
	// todo.DefaultPriority holds the default value on creation for the priority field.
	todo.DefaultPriority = todoDescPriority.Default.(int)
}
