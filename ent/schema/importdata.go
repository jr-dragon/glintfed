package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// ImportData holds the schema definition for the ImportData entity.
type ImportData struct {
	ent.Schema
}

// Fields of the ImportData.
func (ImportData) Fields() []ent.Field {
	return []ent.Field{
		field.Uint32("id").Unique(),
		field.Uint64("profile_id"),
		field.Uint64("job_id").Optional(),
		field.String("service").Default("instagram"),
		field.String("path").Optional(),
		field.Uint8("stage").Default(1),
		field.String("original_name").Optional(),
		field.Bool("import_accepted").Default(false).Optional(),
		field.Time("completed_at").Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the ImportData.
func (ImportData) Edges() []ent.Edge {
	return nil
}

// Annotations of the ImportData.
func (ImportData) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "import_datas"},
	}
}
