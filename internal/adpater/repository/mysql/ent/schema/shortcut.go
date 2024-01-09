package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Shortcut holds the schema definition for the Shortcut entity.
type Shortcut struct {
	ent.Schema
}

// Fields of the Shortcut.
func (Shortcut) Fields() []ent.Field {
	return []ent.Field{
		field.String("url").
			NotEmpty(),
		field.String("shortcut").
			NotEmpty(),
		field.Time("exp"),
	}
}

// Edges of the Shortcut.
func (Shortcut) Edges() []ent.Edge {
	return nil
}
