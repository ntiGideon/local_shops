// Code generated by ent, DO NOT EDIT.

package review

import (
	"baseProject/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Review {
	return predicate.Review(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Review {
	return predicate.Review(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Review {
	return predicate.Review(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Review {
	return predicate.Review(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Review {
	return predicate.Review(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Review {
	return predicate.Review(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Review {
	return predicate.Review(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Review {
	return predicate.Review(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Review {
	return predicate.Review(sql.FieldLTE(FieldID, id))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int) predicate.Review {
	return predicate.Review(sql.FieldEQ(FieldUserID, v))
}

// StoreID applies equality check predicate on the "store_id" field. It's identical to StoreIDEQ.
func StoreID(v int) predicate.Review {
	return predicate.Review(sql.FieldEQ(FieldStoreID, v))
}

// Rating applies equality check predicate on the "rating" field. It's identical to RatingEQ.
func Rating(v float64) predicate.Review {
	return predicate.Review(sql.FieldEQ(FieldRating, v))
}

// Comment applies equality check predicate on the "comment" field. It's identical to CommentEQ.
func Comment(v string) predicate.Review {
	return predicate.Review(sql.FieldEQ(FieldComment, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Review {
	return predicate.Review(sql.FieldEQ(FieldCreatedAt, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int) predicate.Review {
	return predicate.Review(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int) predicate.Review {
	return predicate.Review(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int) predicate.Review {
	return predicate.Review(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int) predicate.Review {
	return predicate.Review(sql.FieldNotIn(FieldUserID, vs...))
}

// StoreIDEQ applies the EQ predicate on the "store_id" field.
func StoreIDEQ(v int) predicate.Review {
	return predicate.Review(sql.FieldEQ(FieldStoreID, v))
}

// StoreIDNEQ applies the NEQ predicate on the "store_id" field.
func StoreIDNEQ(v int) predicate.Review {
	return predicate.Review(sql.FieldNEQ(FieldStoreID, v))
}

// StoreIDIn applies the In predicate on the "store_id" field.
func StoreIDIn(vs ...int) predicate.Review {
	return predicate.Review(sql.FieldIn(FieldStoreID, vs...))
}

// StoreIDNotIn applies the NotIn predicate on the "store_id" field.
func StoreIDNotIn(vs ...int) predicate.Review {
	return predicate.Review(sql.FieldNotIn(FieldStoreID, vs...))
}

// RatingEQ applies the EQ predicate on the "rating" field.
func RatingEQ(v float64) predicate.Review {
	return predicate.Review(sql.FieldEQ(FieldRating, v))
}

// RatingNEQ applies the NEQ predicate on the "rating" field.
func RatingNEQ(v float64) predicate.Review {
	return predicate.Review(sql.FieldNEQ(FieldRating, v))
}

// RatingIn applies the In predicate on the "rating" field.
func RatingIn(vs ...float64) predicate.Review {
	return predicate.Review(sql.FieldIn(FieldRating, vs...))
}

// RatingNotIn applies the NotIn predicate on the "rating" field.
func RatingNotIn(vs ...float64) predicate.Review {
	return predicate.Review(sql.FieldNotIn(FieldRating, vs...))
}

// RatingGT applies the GT predicate on the "rating" field.
func RatingGT(v float64) predicate.Review {
	return predicate.Review(sql.FieldGT(FieldRating, v))
}

// RatingGTE applies the GTE predicate on the "rating" field.
func RatingGTE(v float64) predicate.Review {
	return predicate.Review(sql.FieldGTE(FieldRating, v))
}

// RatingLT applies the LT predicate on the "rating" field.
func RatingLT(v float64) predicate.Review {
	return predicate.Review(sql.FieldLT(FieldRating, v))
}

// RatingLTE applies the LTE predicate on the "rating" field.
func RatingLTE(v float64) predicate.Review {
	return predicate.Review(sql.FieldLTE(FieldRating, v))
}

// CommentEQ applies the EQ predicate on the "comment" field.
func CommentEQ(v string) predicate.Review {
	return predicate.Review(sql.FieldEQ(FieldComment, v))
}

// CommentNEQ applies the NEQ predicate on the "comment" field.
func CommentNEQ(v string) predicate.Review {
	return predicate.Review(sql.FieldNEQ(FieldComment, v))
}

// CommentIn applies the In predicate on the "comment" field.
func CommentIn(vs ...string) predicate.Review {
	return predicate.Review(sql.FieldIn(FieldComment, vs...))
}

// CommentNotIn applies the NotIn predicate on the "comment" field.
func CommentNotIn(vs ...string) predicate.Review {
	return predicate.Review(sql.FieldNotIn(FieldComment, vs...))
}

// CommentGT applies the GT predicate on the "comment" field.
func CommentGT(v string) predicate.Review {
	return predicate.Review(sql.FieldGT(FieldComment, v))
}

// CommentGTE applies the GTE predicate on the "comment" field.
func CommentGTE(v string) predicate.Review {
	return predicate.Review(sql.FieldGTE(FieldComment, v))
}

// CommentLT applies the LT predicate on the "comment" field.
func CommentLT(v string) predicate.Review {
	return predicate.Review(sql.FieldLT(FieldComment, v))
}

// CommentLTE applies the LTE predicate on the "comment" field.
func CommentLTE(v string) predicate.Review {
	return predicate.Review(sql.FieldLTE(FieldComment, v))
}

// CommentContains applies the Contains predicate on the "comment" field.
func CommentContains(v string) predicate.Review {
	return predicate.Review(sql.FieldContains(FieldComment, v))
}

// CommentHasPrefix applies the HasPrefix predicate on the "comment" field.
func CommentHasPrefix(v string) predicate.Review {
	return predicate.Review(sql.FieldHasPrefix(FieldComment, v))
}

// CommentHasSuffix applies the HasSuffix predicate on the "comment" field.
func CommentHasSuffix(v string) predicate.Review {
	return predicate.Review(sql.FieldHasSuffix(FieldComment, v))
}

// CommentIsNil applies the IsNil predicate on the "comment" field.
func CommentIsNil() predicate.Review {
	return predicate.Review(sql.FieldIsNull(FieldComment))
}

// CommentNotNil applies the NotNil predicate on the "comment" field.
func CommentNotNil() predicate.Review {
	return predicate.Review(sql.FieldNotNull(FieldComment))
}

// CommentEqualFold applies the EqualFold predicate on the "comment" field.
func CommentEqualFold(v string) predicate.Review {
	return predicate.Review(sql.FieldEqualFold(FieldComment, v))
}

// CommentContainsFold applies the ContainsFold predicate on the "comment" field.
func CommentContainsFold(v string) predicate.Review {
	return predicate.Review(sql.FieldContainsFold(FieldComment, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Review {
	return predicate.Review(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Review {
	return predicate.Review(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Review {
	return predicate.Review(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Review {
	return predicate.Review(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Review {
	return predicate.Review(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Review {
	return predicate.Review(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Review {
	return predicate.Review(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Review {
	return predicate.Review(sql.FieldLTE(FieldCreatedAt, v))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasStore applies the HasEdge predicate on the "store" edge.
func HasStore() predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, StoreTable, StoreColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasStoreWith applies the HasEdge predicate on the "store" edge with a given conditions (other predicates).
func HasStoreWith(preds ...predicate.Store) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		step := newStoreStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Review) predicate.Review {
	return predicate.Review(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Review) predicate.Review {
	return predicate.Review(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Review) predicate.Review {
	return predicate.Review(sql.NotPredicates(p))
}
