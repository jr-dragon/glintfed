package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
)

// UIKit holds the schema definition for the UIKit entity.
type UIKit struct {
	ent.Schema
}

// Fields of the UIKit.
func (UIKit) Fields() []ent.Field {
	return nil
}

// Edges of the UIKit.
func (UIKit) Edges() []ent.Edge {
	return nil
}

// Annotations of the UIKit.
func (UIKit) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "uikit"},
	}
}
