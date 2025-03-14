package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Store holds the schema definition for the Store entity.
type Store struct {
	ent.Schema
}

// Fields of the Store.
func (Store) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.String("name").NotEmpty(),
		field.String("location").NotEmpty(),
		field.String("city").NotEmpty(),
		field.String("country"),
		field.String("zipcode").Optional(),
		field.String("photo").Optional(),
		field.Float("latitude").Optional(),
		field.Float("longitude").Optional(),
		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the Store.
func (Store) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("products", Product.Type),
		edge.To("prices", Price.Type),
		edge.To("reviews", Review.Type),
	}
}
