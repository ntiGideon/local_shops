package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Review holds the schema definition for the Review entity.
type Review struct {
	ent.Schema
}

// Fields of the Review.
func (Review) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.Int("user_id"),
		field.Int("store_id"),
		field.Float("rating").Min(1).Max(5),
		field.String("comment").Optional(),
		field.Time("created_at").Default(time.Now()),
	}
}

// Edges of the Review.
func (Review) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("reviews").Unique().Required().Field("user_id"),
		edge.From("store", Store.Type).Ref("reviews").Unique().Required().Field("store_id"),
	}
}
