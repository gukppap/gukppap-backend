package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Url holds the schema definition for the Url entity.
type Url struct {
	ent.Schema
}

// Fields of the Url.
func (Url) Fields() []ent.Field {
	return []ent.Field{
		field.String("origin_url").
			NotEmpty().
			Unique(),

		field.String("shortcut_url").
			NotEmpty().
			Unique(),
	}
}

// Edges of the Url.
func (Url) Edges() []ent.Edge {
	return nil
}
