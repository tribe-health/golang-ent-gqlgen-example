// Code generated by entc, DO NOT EDIT.

package ent

import (
	"golang-ent-gqlgen-example/ent/article"
	"golang-ent-gqlgen-example/ent/schema"
	"golang-ent-gqlgen-example/ent/user"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	articleFields := schema.Article{}.Fields()
	_ = articleFields
	// articleDescTitle is the schema descriptor for title field.
	articleDescTitle := articleFields[0].Descriptor()
	// article.DefaultTitle holds the default value on creation for the title field.
	article.DefaultTitle = articleDescTitle.Default.(string)
	// articleDescDescription is the schema descriptor for description field.
	articleDescDescription := articleFields[1].Descriptor()
	// article.DefaultDescription holds the default value on creation for the description field.
	article.DefaultDescription = articleDescDescription.Default.(string)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[0].Descriptor()
	// user.DefaultName holds the default value on creation for the name field.
	user.DefaultName = userDescName.Default.(string)
	// userDescAge is the schema descriptor for age field.
	userDescAge := userFields[1].Descriptor()
	// user.AgeValidator is a validator for the "age" field. It is called by the builders before save.
	user.AgeValidator = userDescAge.Validators[0].(func(int) error)
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[2].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
}
