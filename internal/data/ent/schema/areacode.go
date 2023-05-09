package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// AreaCode holds the schema definition for the AreaCode entity.
type AreaCode struct {
	ent.Schema
}

// Fields of the AreaCode.
func (AreaCode) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("code").
			NonNegative(),
		field.String("name").
			Default(""),
		field.Int8("level"),
		field.Int64("p_code").NonNegative(),
	}
}

// Edges of the AreaCode.
func (AreaCode) Edges() []ent.Edge {
	return nil
}

func (AreaCode) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "area_code"},
	}
}
