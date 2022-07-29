package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// SecurityPosition holds the schema definition for the SecurityPosition entity.
type SecurityPosition struct {
	ent.Schema
}

func (SecurityPosition) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "security_position"},
	}
}

// Fields of the SecurityPosition
func (SecurityPosition) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").SchemaType(map[string]string{dialect.MySQL: "INT"}),
		field.Int("parent_id").Default(0).SchemaType(map[string]string{dialect.MySQL: "INT"}),
		field.Int("account_id").Optional().SchemaType(map[string]string{dialect.MySQL: "INT"}),
		field.Int("security_account_id").SchemaType(map[string]string{dialect.MySQL: "INT"}),
		field.Int("position_type").Optional().SchemaType(map[string]string{dialect.MySQL: "INT"}),
		field.Int("cd_direction").Optional().Nillable().SchemaType(map[string]string{dialect.MySQL: "INT"}),
		field.Int("project_id").Optional().SchemaType(map[string]string{dialect.MySQL: "INT"}),
		field.Int("product_id").SchemaType(map[string]string{dialect.MySQL: "INT"}),
		field.Int("material_id").Optional().SchemaType(map[string]string{dialect.MySQL: "INT"}),
		field.String("material_no").MaxLen(45).Optional().Nillable().Comment("标的代码"),
		field.Float("multiplier").Optional().Nillable(),
		field.Float("balance").Default(0.0000).Optional(),
		field.Float("available").Default(0.0000).Optional(),
		field.Float("freeze").Default(0.0000).Optional(),
		field.Float("afloat").Default(0.0000).Optional(),
		field.String("unit").MaxLen(10).Optional().Nillable(),
		field.String("spec").MaxLen(45).Optional().Nillable(),
		field.String("currency").MaxLen(10).Optional().Nillable(),
		field.Float("price").Default(0.000000).Optional(),
		field.Float("amount").Optional().Nillable(),
		field.Float("cost_amount").Optional().Nillable(),
		field.Float("cost").Optional().Nillable(),
		field.Float("fx_rate").Optional().Nillable(),
		field.String("stl_currency").MaxLen(10).Optional().Nillable(),
		field.Float("stl_amount").Optional().Nillable(),
		field.Float("stl_cost").Optional().Nillable(),
		field.Float("stl_margin").Optional().Nillable(),
		field.Float("stl_cost_amount").Optional().Nillable(),
		field.Float("stl_margin_lv").Default(0.000000).Optional(),
		field.Float("stl_val_price").Optional().Nillable(),
		field.Time("invalid_at").Optional().Nillable(),
		field.Time("created_at").Optional().Nillable(),
		field.String("unique_tag").MaxLen(45).Optional().Nillable().Unique(),
		field.Int("org_id").SchemaType(map[string]string{dialect.MySQL: "INT"}),
	}
}

// Edges of the SecurityPosition.
func (SecurityPosition) Edges() []ent.Edge {
	return nil
}
