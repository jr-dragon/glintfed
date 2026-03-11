package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// DefaultDomainBlock holds the schema definition for the DefaultDomainBlock entity.
type DefaultDomainBlock struct {
	ent.Schema
}

// Fields of the DefaultDomainBlock.
func (DefaultDomainBlock) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.String("domain").Unique(),
		field.Text("note").Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the DefaultDomainBlock.
func (DefaultDomainBlock) Edges() []ent.Edge {
	return nil
}

// Annotations of the DefaultDomainBlock.
func (DefaultDomainBlock) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "default_domain_blocks"},
	}
}
