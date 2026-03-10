package schema

import "entgo.io/ent"

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
