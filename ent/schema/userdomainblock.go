package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// UserDomainBlock holds the schema definition for the UserDomainBlock entity.
type UserDomainBlock struct {
	ent.Schema
}

// Fields of the UserDomainBlock.
func (UserDomainBlock) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.Uint64("profile_id"),
		field.String("domain"),
	}
}

// Edges of the UserDomainBlock.
func (UserDomainBlock) Edges() []ent.Edge {
	return nil
}

// Annotations of the UserDomainBlock.
func (UserDomainBlock) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "user_domain_blocks"},
	}
}
