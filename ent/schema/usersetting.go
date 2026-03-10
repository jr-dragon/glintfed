package schema

import "entgo.io/ent"

// UserSetting holds the schema definition for the UserSetting entity.
type UserSetting struct {
	ent.Schema
}

// Fields of the UserSetting.
func (UserSetting) Fields() []ent.Field {
	return nil
}

// Edges of the UserSetting.
func (UserSetting) Edges() []ent.Edge {
	return nil
}
