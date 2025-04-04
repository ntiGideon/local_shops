// Code generated by ent, DO NOT EDIT.

package store

import (
	"baseProject/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Store {
	return predicate.Store(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Store {
	return predicate.Store(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Store {
	return predicate.Store(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Store {
	return predicate.Store(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Store {
	return predicate.Store(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Store {
	return predicate.Store(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Store {
	return predicate.Store(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Store {
	return predicate.Store(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Store {
	return predicate.Store(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Store {
	return predicate.Store(sql.FieldEQ(FieldName, v))
}

// Location applies equality check predicate on the "location" field. It's identical to LocationEQ.
func Location(v string) predicate.Store {
	return predicate.Store(sql.FieldEQ(FieldLocation, v))
}

// City applies equality check predicate on the "city" field. It's identical to CityEQ.
func City(v string) predicate.Store {
	return predicate.Store(sql.FieldEQ(FieldCity, v))
}

// Country applies equality check predicate on the "country" field. It's identical to CountryEQ.
func Country(v string) predicate.Store {
	return predicate.Store(sql.FieldEQ(FieldCountry, v))
}

// Zipcode applies equality check predicate on the "zipcode" field. It's identical to ZipcodeEQ.
func Zipcode(v string) predicate.Store {
	return predicate.Store(sql.FieldEQ(FieldZipcode, v))
}

// Photo applies equality check predicate on the "photo" field. It's identical to PhotoEQ.
func Photo(v string) predicate.Store {
	return predicate.Store(sql.FieldEQ(FieldPhoto, v))
}

// Latitude applies equality check predicate on the "latitude" field. It's identical to LatitudeEQ.
func Latitude(v float64) predicate.Store {
	return predicate.Store(sql.FieldEQ(FieldLatitude, v))
}

// Longitude applies equality check predicate on the "longitude" field. It's identical to LongitudeEQ.
func Longitude(v float64) predicate.Store {
	return predicate.Store(sql.FieldEQ(FieldLongitude, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Store {
	return predicate.Store(sql.FieldEQ(FieldCreatedAt, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Store {
	return predicate.Store(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Store {
	return predicate.Store(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Store {
	return predicate.Store(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Store {
	return predicate.Store(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Store {
	return predicate.Store(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Store {
	return predicate.Store(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Store {
	return predicate.Store(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Store {
	return predicate.Store(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Store {
	return predicate.Store(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Store {
	return predicate.Store(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Store {
	return predicate.Store(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Store {
	return predicate.Store(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Store {
	return predicate.Store(sql.FieldContainsFold(FieldName, v))
}

// LocationEQ applies the EQ predicate on the "location" field.
func LocationEQ(v string) predicate.Store {
	return predicate.Store(sql.FieldEQ(FieldLocation, v))
}

// LocationNEQ applies the NEQ predicate on the "location" field.
func LocationNEQ(v string) predicate.Store {
	return predicate.Store(sql.FieldNEQ(FieldLocation, v))
}

// LocationIn applies the In predicate on the "location" field.
func LocationIn(vs ...string) predicate.Store {
	return predicate.Store(sql.FieldIn(FieldLocation, vs...))
}

// LocationNotIn applies the NotIn predicate on the "location" field.
func LocationNotIn(vs ...string) predicate.Store {
	return predicate.Store(sql.FieldNotIn(FieldLocation, vs...))
}

// LocationGT applies the GT predicate on the "location" field.
func LocationGT(v string) predicate.Store {
	return predicate.Store(sql.FieldGT(FieldLocation, v))
}

// LocationGTE applies the GTE predicate on the "location" field.
func LocationGTE(v string) predicate.Store {
	return predicate.Store(sql.FieldGTE(FieldLocation, v))
}

// LocationLT applies the LT predicate on the "location" field.
func LocationLT(v string) predicate.Store {
	return predicate.Store(sql.FieldLT(FieldLocation, v))
}

// LocationLTE applies the LTE predicate on the "location" field.
func LocationLTE(v string) predicate.Store {
	return predicate.Store(sql.FieldLTE(FieldLocation, v))
}

// LocationContains applies the Contains predicate on the "location" field.
func LocationContains(v string) predicate.Store {
	return predicate.Store(sql.FieldContains(FieldLocation, v))
}

// LocationHasPrefix applies the HasPrefix predicate on the "location" field.
func LocationHasPrefix(v string) predicate.Store {
	return predicate.Store(sql.FieldHasPrefix(FieldLocation, v))
}

// LocationHasSuffix applies the HasSuffix predicate on the "location" field.
func LocationHasSuffix(v string) predicate.Store {
	return predicate.Store(sql.FieldHasSuffix(FieldLocation, v))
}

// LocationEqualFold applies the EqualFold predicate on the "location" field.
func LocationEqualFold(v string) predicate.Store {
	return predicate.Store(sql.FieldEqualFold(FieldLocation, v))
}

// LocationContainsFold applies the ContainsFold predicate on the "location" field.
func LocationContainsFold(v string) predicate.Store {
	return predicate.Store(sql.FieldContainsFold(FieldLocation, v))
}

// CityEQ applies the EQ predicate on the "city" field.
func CityEQ(v string) predicate.Store {
	return predicate.Store(sql.FieldEQ(FieldCity, v))
}

// CityNEQ applies the NEQ predicate on the "city" field.
func CityNEQ(v string) predicate.Store {
	return predicate.Store(sql.FieldNEQ(FieldCity, v))
}

// CityIn applies the In predicate on the "city" field.
func CityIn(vs ...string) predicate.Store {
	return predicate.Store(sql.FieldIn(FieldCity, vs...))
}

// CityNotIn applies the NotIn predicate on the "city" field.
func CityNotIn(vs ...string) predicate.Store {
	return predicate.Store(sql.FieldNotIn(FieldCity, vs...))
}

// CityGT applies the GT predicate on the "city" field.
func CityGT(v string) predicate.Store {
	return predicate.Store(sql.FieldGT(FieldCity, v))
}

// CityGTE applies the GTE predicate on the "city" field.
func CityGTE(v string) predicate.Store {
	return predicate.Store(sql.FieldGTE(FieldCity, v))
}

// CityLT applies the LT predicate on the "city" field.
func CityLT(v string) predicate.Store {
	return predicate.Store(sql.FieldLT(FieldCity, v))
}

// CityLTE applies the LTE predicate on the "city" field.
func CityLTE(v string) predicate.Store {
	return predicate.Store(sql.FieldLTE(FieldCity, v))
}

// CityContains applies the Contains predicate on the "city" field.
func CityContains(v string) predicate.Store {
	return predicate.Store(sql.FieldContains(FieldCity, v))
}

// CityHasPrefix applies the HasPrefix predicate on the "city" field.
func CityHasPrefix(v string) predicate.Store {
	return predicate.Store(sql.FieldHasPrefix(FieldCity, v))
}

// CityHasSuffix applies the HasSuffix predicate on the "city" field.
func CityHasSuffix(v string) predicate.Store {
	return predicate.Store(sql.FieldHasSuffix(FieldCity, v))
}

// CityEqualFold applies the EqualFold predicate on the "city" field.
func CityEqualFold(v string) predicate.Store {
	return predicate.Store(sql.FieldEqualFold(FieldCity, v))
}

// CityContainsFold applies the ContainsFold predicate on the "city" field.
func CityContainsFold(v string) predicate.Store {
	return predicate.Store(sql.FieldContainsFold(FieldCity, v))
}

// CountryEQ applies the EQ predicate on the "country" field.
func CountryEQ(v string) predicate.Store {
	return predicate.Store(sql.FieldEQ(FieldCountry, v))
}

// CountryNEQ applies the NEQ predicate on the "country" field.
func CountryNEQ(v string) predicate.Store {
	return predicate.Store(sql.FieldNEQ(FieldCountry, v))
}

// CountryIn applies the In predicate on the "country" field.
func CountryIn(vs ...string) predicate.Store {
	return predicate.Store(sql.FieldIn(FieldCountry, vs...))
}

// CountryNotIn applies the NotIn predicate on the "country" field.
func CountryNotIn(vs ...string) predicate.Store {
	return predicate.Store(sql.FieldNotIn(FieldCountry, vs...))
}

// CountryGT applies the GT predicate on the "country" field.
func CountryGT(v string) predicate.Store {
	return predicate.Store(sql.FieldGT(FieldCountry, v))
}

// CountryGTE applies the GTE predicate on the "country" field.
func CountryGTE(v string) predicate.Store {
	return predicate.Store(sql.FieldGTE(FieldCountry, v))
}

// CountryLT applies the LT predicate on the "country" field.
func CountryLT(v string) predicate.Store {
	return predicate.Store(sql.FieldLT(FieldCountry, v))
}

// CountryLTE applies the LTE predicate on the "country" field.
func CountryLTE(v string) predicate.Store {
	return predicate.Store(sql.FieldLTE(FieldCountry, v))
}

// CountryContains applies the Contains predicate on the "country" field.
func CountryContains(v string) predicate.Store {
	return predicate.Store(sql.FieldContains(FieldCountry, v))
}

// CountryHasPrefix applies the HasPrefix predicate on the "country" field.
func CountryHasPrefix(v string) predicate.Store {
	return predicate.Store(sql.FieldHasPrefix(FieldCountry, v))
}

// CountryHasSuffix applies the HasSuffix predicate on the "country" field.
func CountryHasSuffix(v string) predicate.Store {
	return predicate.Store(sql.FieldHasSuffix(FieldCountry, v))
}

// CountryEqualFold applies the EqualFold predicate on the "country" field.
func CountryEqualFold(v string) predicate.Store {
	return predicate.Store(sql.FieldEqualFold(FieldCountry, v))
}

// CountryContainsFold applies the ContainsFold predicate on the "country" field.
func CountryContainsFold(v string) predicate.Store {
	return predicate.Store(sql.FieldContainsFold(FieldCountry, v))
}

// ZipcodeEQ applies the EQ predicate on the "zipcode" field.
func ZipcodeEQ(v string) predicate.Store {
	return predicate.Store(sql.FieldEQ(FieldZipcode, v))
}

// ZipcodeNEQ applies the NEQ predicate on the "zipcode" field.
func ZipcodeNEQ(v string) predicate.Store {
	return predicate.Store(sql.FieldNEQ(FieldZipcode, v))
}

// ZipcodeIn applies the In predicate on the "zipcode" field.
func ZipcodeIn(vs ...string) predicate.Store {
	return predicate.Store(sql.FieldIn(FieldZipcode, vs...))
}

// ZipcodeNotIn applies the NotIn predicate on the "zipcode" field.
func ZipcodeNotIn(vs ...string) predicate.Store {
	return predicate.Store(sql.FieldNotIn(FieldZipcode, vs...))
}

// ZipcodeGT applies the GT predicate on the "zipcode" field.
func ZipcodeGT(v string) predicate.Store {
	return predicate.Store(sql.FieldGT(FieldZipcode, v))
}

// ZipcodeGTE applies the GTE predicate on the "zipcode" field.
func ZipcodeGTE(v string) predicate.Store {
	return predicate.Store(sql.FieldGTE(FieldZipcode, v))
}

// ZipcodeLT applies the LT predicate on the "zipcode" field.
func ZipcodeLT(v string) predicate.Store {
	return predicate.Store(sql.FieldLT(FieldZipcode, v))
}

// ZipcodeLTE applies the LTE predicate on the "zipcode" field.
func ZipcodeLTE(v string) predicate.Store {
	return predicate.Store(sql.FieldLTE(FieldZipcode, v))
}

// ZipcodeContains applies the Contains predicate on the "zipcode" field.
func ZipcodeContains(v string) predicate.Store {
	return predicate.Store(sql.FieldContains(FieldZipcode, v))
}

// ZipcodeHasPrefix applies the HasPrefix predicate on the "zipcode" field.
func ZipcodeHasPrefix(v string) predicate.Store {
	return predicate.Store(sql.FieldHasPrefix(FieldZipcode, v))
}

// ZipcodeHasSuffix applies the HasSuffix predicate on the "zipcode" field.
func ZipcodeHasSuffix(v string) predicate.Store {
	return predicate.Store(sql.FieldHasSuffix(FieldZipcode, v))
}

// ZipcodeIsNil applies the IsNil predicate on the "zipcode" field.
func ZipcodeIsNil() predicate.Store {
	return predicate.Store(sql.FieldIsNull(FieldZipcode))
}

// ZipcodeNotNil applies the NotNil predicate on the "zipcode" field.
func ZipcodeNotNil() predicate.Store {
	return predicate.Store(sql.FieldNotNull(FieldZipcode))
}

// ZipcodeEqualFold applies the EqualFold predicate on the "zipcode" field.
func ZipcodeEqualFold(v string) predicate.Store {
	return predicate.Store(sql.FieldEqualFold(FieldZipcode, v))
}

// ZipcodeContainsFold applies the ContainsFold predicate on the "zipcode" field.
func ZipcodeContainsFold(v string) predicate.Store {
	return predicate.Store(sql.FieldContainsFold(FieldZipcode, v))
}

// PhotoEQ applies the EQ predicate on the "photo" field.
func PhotoEQ(v string) predicate.Store {
	return predicate.Store(sql.FieldEQ(FieldPhoto, v))
}

// PhotoNEQ applies the NEQ predicate on the "photo" field.
func PhotoNEQ(v string) predicate.Store {
	return predicate.Store(sql.FieldNEQ(FieldPhoto, v))
}

// PhotoIn applies the In predicate on the "photo" field.
func PhotoIn(vs ...string) predicate.Store {
	return predicate.Store(sql.FieldIn(FieldPhoto, vs...))
}

// PhotoNotIn applies the NotIn predicate on the "photo" field.
func PhotoNotIn(vs ...string) predicate.Store {
	return predicate.Store(sql.FieldNotIn(FieldPhoto, vs...))
}

// PhotoGT applies the GT predicate on the "photo" field.
func PhotoGT(v string) predicate.Store {
	return predicate.Store(sql.FieldGT(FieldPhoto, v))
}

// PhotoGTE applies the GTE predicate on the "photo" field.
func PhotoGTE(v string) predicate.Store {
	return predicate.Store(sql.FieldGTE(FieldPhoto, v))
}

// PhotoLT applies the LT predicate on the "photo" field.
func PhotoLT(v string) predicate.Store {
	return predicate.Store(sql.FieldLT(FieldPhoto, v))
}

// PhotoLTE applies the LTE predicate on the "photo" field.
func PhotoLTE(v string) predicate.Store {
	return predicate.Store(sql.FieldLTE(FieldPhoto, v))
}

// PhotoContains applies the Contains predicate on the "photo" field.
func PhotoContains(v string) predicate.Store {
	return predicate.Store(sql.FieldContains(FieldPhoto, v))
}

// PhotoHasPrefix applies the HasPrefix predicate on the "photo" field.
func PhotoHasPrefix(v string) predicate.Store {
	return predicate.Store(sql.FieldHasPrefix(FieldPhoto, v))
}

// PhotoHasSuffix applies the HasSuffix predicate on the "photo" field.
func PhotoHasSuffix(v string) predicate.Store {
	return predicate.Store(sql.FieldHasSuffix(FieldPhoto, v))
}

// PhotoIsNil applies the IsNil predicate on the "photo" field.
func PhotoIsNil() predicate.Store {
	return predicate.Store(sql.FieldIsNull(FieldPhoto))
}

// PhotoNotNil applies the NotNil predicate on the "photo" field.
func PhotoNotNil() predicate.Store {
	return predicate.Store(sql.FieldNotNull(FieldPhoto))
}

// PhotoEqualFold applies the EqualFold predicate on the "photo" field.
func PhotoEqualFold(v string) predicate.Store {
	return predicate.Store(sql.FieldEqualFold(FieldPhoto, v))
}

// PhotoContainsFold applies the ContainsFold predicate on the "photo" field.
func PhotoContainsFold(v string) predicate.Store {
	return predicate.Store(sql.FieldContainsFold(FieldPhoto, v))
}

// LatitudeEQ applies the EQ predicate on the "latitude" field.
func LatitudeEQ(v float64) predicate.Store {
	return predicate.Store(sql.FieldEQ(FieldLatitude, v))
}

// LatitudeNEQ applies the NEQ predicate on the "latitude" field.
func LatitudeNEQ(v float64) predicate.Store {
	return predicate.Store(sql.FieldNEQ(FieldLatitude, v))
}

// LatitudeIn applies the In predicate on the "latitude" field.
func LatitudeIn(vs ...float64) predicate.Store {
	return predicate.Store(sql.FieldIn(FieldLatitude, vs...))
}

// LatitudeNotIn applies the NotIn predicate on the "latitude" field.
func LatitudeNotIn(vs ...float64) predicate.Store {
	return predicate.Store(sql.FieldNotIn(FieldLatitude, vs...))
}

// LatitudeGT applies the GT predicate on the "latitude" field.
func LatitudeGT(v float64) predicate.Store {
	return predicate.Store(sql.FieldGT(FieldLatitude, v))
}

// LatitudeGTE applies the GTE predicate on the "latitude" field.
func LatitudeGTE(v float64) predicate.Store {
	return predicate.Store(sql.FieldGTE(FieldLatitude, v))
}

// LatitudeLT applies the LT predicate on the "latitude" field.
func LatitudeLT(v float64) predicate.Store {
	return predicate.Store(sql.FieldLT(FieldLatitude, v))
}

// LatitudeLTE applies the LTE predicate on the "latitude" field.
func LatitudeLTE(v float64) predicate.Store {
	return predicate.Store(sql.FieldLTE(FieldLatitude, v))
}

// LatitudeIsNil applies the IsNil predicate on the "latitude" field.
func LatitudeIsNil() predicate.Store {
	return predicate.Store(sql.FieldIsNull(FieldLatitude))
}

// LatitudeNotNil applies the NotNil predicate on the "latitude" field.
func LatitudeNotNil() predicate.Store {
	return predicate.Store(sql.FieldNotNull(FieldLatitude))
}

// LongitudeEQ applies the EQ predicate on the "longitude" field.
func LongitudeEQ(v float64) predicate.Store {
	return predicate.Store(sql.FieldEQ(FieldLongitude, v))
}

// LongitudeNEQ applies the NEQ predicate on the "longitude" field.
func LongitudeNEQ(v float64) predicate.Store {
	return predicate.Store(sql.FieldNEQ(FieldLongitude, v))
}

// LongitudeIn applies the In predicate on the "longitude" field.
func LongitudeIn(vs ...float64) predicate.Store {
	return predicate.Store(sql.FieldIn(FieldLongitude, vs...))
}

// LongitudeNotIn applies the NotIn predicate on the "longitude" field.
func LongitudeNotIn(vs ...float64) predicate.Store {
	return predicate.Store(sql.FieldNotIn(FieldLongitude, vs...))
}

// LongitudeGT applies the GT predicate on the "longitude" field.
func LongitudeGT(v float64) predicate.Store {
	return predicate.Store(sql.FieldGT(FieldLongitude, v))
}

// LongitudeGTE applies the GTE predicate on the "longitude" field.
func LongitudeGTE(v float64) predicate.Store {
	return predicate.Store(sql.FieldGTE(FieldLongitude, v))
}

// LongitudeLT applies the LT predicate on the "longitude" field.
func LongitudeLT(v float64) predicate.Store {
	return predicate.Store(sql.FieldLT(FieldLongitude, v))
}

// LongitudeLTE applies the LTE predicate on the "longitude" field.
func LongitudeLTE(v float64) predicate.Store {
	return predicate.Store(sql.FieldLTE(FieldLongitude, v))
}

// LongitudeIsNil applies the IsNil predicate on the "longitude" field.
func LongitudeIsNil() predicate.Store {
	return predicate.Store(sql.FieldIsNull(FieldLongitude))
}

// LongitudeNotNil applies the NotNil predicate on the "longitude" field.
func LongitudeNotNil() predicate.Store {
	return predicate.Store(sql.FieldNotNull(FieldLongitude))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Store {
	return predicate.Store(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Store {
	return predicate.Store(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Store {
	return predicate.Store(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Store {
	return predicate.Store(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Store {
	return predicate.Store(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Store {
	return predicate.Store(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Store {
	return predicate.Store(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Store {
	return predicate.Store(sql.FieldLTE(FieldCreatedAt, v))
}

// HasProducts applies the HasEdge predicate on the "products" edge.
func HasProducts() predicate.Store {
	return predicate.Store(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ProductsTable, ProductsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProductsWith applies the HasEdge predicate on the "products" edge with a given conditions (other predicates).
func HasProductsWith(preds ...predicate.Product) predicate.Store {
	return predicate.Store(func(s *sql.Selector) {
		step := newProductsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPrices applies the HasEdge predicate on the "prices" edge.
func HasPrices() predicate.Store {
	return predicate.Store(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PricesTable, PricesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPricesWith applies the HasEdge predicate on the "prices" edge with a given conditions (other predicates).
func HasPricesWith(preds ...predicate.Price) predicate.Store {
	return predicate.Store(func(s *sql.Selector) {
		step := newPricesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasReviews applies the HasEdge predicate on the "reviews" edge.
func HasReviews() predicate.Store {
	return predicate.Store(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ReviewsTable, ReviewsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasReviewsWith applies the HasEdge predicate on the "reviews" edge with a given conditions (other predicates).
func HasReviewsWith(preds ...predicate.Review) predicate.Store {
	return predicate.Store(func(s *sql.Selector) {
		step := newReviewsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Store) predicate.Store {
	return predicate.Store(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Store) predicate.Store {
	return predicate.Store(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Store) predicate.Store {
	return predicate.Store(sql.NotPredicates(p))
}
