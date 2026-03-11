package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// RemoteAuthInstance holds the schema definition for the RemoteAuthInstance entity.
type RemoteAuthInstance struct {
	ent.Schema
}

// Fields of the RemoteAuthInstance.
func (RemoteAuthInstance) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.String("domain").Unique().Optional(),
		field.Uint32("instance_id").Optional(),
		field.String("client_id").Optional(),
		field.String("client_secret").Optional(),
		field.String("redirect_uri").Optional(),
		field.String("root_domain").Optional(),
		field.Bool("allowed").Optional(),
		field.Bool("banned").Default(false),
		field.Bool("active").Default(true),
		field.Time("last_refreshed_at").Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the RemoteAuthInstance.
func (RemoteAuthInstance) Edges() []ent.Edge {
	return nil
}

// Annotations of the RemoteAuthInstance.
func (RemoteAuthInstance) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "remote_auth_instances"},
	}
}
