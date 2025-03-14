package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Price holds the schema definition for the Price entity.
type Price struct {
	ent.Schema
}

// Fields of the Price.
func (Price) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.Int("product_id"),
		field.Int("store_id"),
		field.Float("price"),
		field.Time("recorded_at").Default(time.Now),
	}
}

// Edges of the Price.
func (Price) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("product", Product.Type).Ref("prices").Unique().Required().Field("product_id"),
		edge.From("store", Store.Type).Ref("prices").Unique().Required().Field("store_id"),
	}
}
