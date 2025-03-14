// Code generated by ent, DO NOT EDIT.

package ent

import (
	"baseProject/ent/price"
	"baseProject/ent/product"
	"baseProject/ent/store"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Price is the model entity for the Price schema.
type Price struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// ProductID holds the value of the "product_id" field.
	ProductID int `json:"product_id,omitempty"`
	// StoreID holds the value of the "store_id" field.
	StoreID int `json:"store_id,omitempty"`
	// Price holds the value of the "price" field.
	Price float64 `json:"price,omitempty"`
	// RecordedAt holds the value of the "recorded_at" field.
	RecordedAt time.Time `json:"recorded_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PriceQuery when eager-loading is set.
	Edges        PriceEdges `json:"edges"`
	selectValues sql.SelectValues
}

// PriceEdges holds the relations/edges for other nodes in the graph.
type PriceEdges struct {
	// Product holds the value of the product edge.
	Product *Product `json:"product,omitempty"`
	// Store holds the value of the store edge.
	Store *Store `json:"store,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// ProductOrErr returns the Product value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PriceEdges) ProductOrErr() (*Product, error) {
	if e.Product != nil {
		return e.Product, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: product.Label}
	}
	return nil, &NotLoadedError{edge: "product"}
}

// StoreOrErr returns the Store value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PriceEdges) StoreOrErr() (*Store, error) {
	if e.Store != nil {
		return e.Store, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: store.Label}
	}
	return nil, &NotLoadedError{edge: "store"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Price) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case price.FieldPrice:
			values[i] = new(sql.NullFloat64)
		case price.FieldID, price.FieldProductID, price.FieldStoreID:
			values[i] = new(sql.NullInt64)
		case price.FieldRecordedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Price fields.
func (pr *Price) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case price.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pr.ID = int(value.Int64)
		case price.FieldProductID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field product_id", values[i])
			} else if value.Valid {
				pr.ProductID = int(value.Int64)
			}
		case price.FieldStoreID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field store_id", values[i])
			} else if value.Valid {
				pr.StoreID = int(value.Int64)
			}
		case price.FieldPrice:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field price", values[i])
			} else if value.Valid {
				pr.Price = value.Float64
			}
		case price.FieldRecordedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field recorded_at", values[i])
			} else if value.Valid {
				pr.RecordedAt = value.Time
			}
		default:
			pr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Price.
// This includes values selected through modifiers, order, etc.
func (pr *Price) Value(name string) (ent.Value, error) {
	return pr.selectValues.Get(name)
}

// QueryProduct queries the "product" edge of the Price entity.
func (pr *Price) QueryProduct() *ProductQuery {
	return NewPriceClient(pr.config).QueryProduct(pr)
}

// QueryStore queries the "store" edge of the Price entity.
func (pr *Price) QueryStore() *StoreQuery {
	return NewPriceClient(pr.config).QueryStore(pr)
}

// Update returns a builder for updating this Price.
// Note that you need to call Price.Unwrap() before calling this method if this Price
// was returned from a transaction, and the transaction was committed or rolled back.
func (pr *Price) Update() *PriceUpdateOne {
	return NewPriceClient(pr.config).UpdateOne(pr)
}

// Unwrap unwraps the Price entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pr *Price) Unwrap() *Price {
	_tx, ok := pr.config.driver.(*txDriver)
	if !ok {
		panic("ent: Price is not a transactional entity")
	}
	pr.config.driver = _tx.drv
	return pr
}

// String implements the fmt.Stringer.
func (pr *Price) String() string {
	var builder strings.Builder
	builder.WriteString("Price(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pr.ID))
	builder.WriteString("product_id=")
	builder.WriteString(fmt.Sprintf("%v", pr.ProductID))
	builder.WriteString(", ")
	builder.WriteString("store_id=")
	builder.WriteString(fmt.Sprintf("%v", pr.StoreID))
	builder.WriteString(", ")
	builder.WriteString("price=")
	builder.WriteString(fmt.Sprintf("%v", pr.Price))
	builder.WriteString(", ")
	builder.WriteString("recorded_at=")
	builder.WriteString(pr.RecordedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Prices is a parsable slice of Price.
type Prices []*Price
