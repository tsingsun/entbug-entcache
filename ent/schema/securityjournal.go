package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// SecurityJournal holds the schema definition for the SecurityJournal entity.
type SecurityJournal struct {
	ent.Schema
}

func (SecurityJournal) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "security_journal"},
	}
}

// Fields of the SecurityJournal
func (SecurityJournal) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").SchemaType(map[string]string{dialect.MySQL: "INT"}),
		field.Int("security_account_id").SchemaType(map[string]string{dialect.MySQL: "INT"}),
		field.Int("account_id").Default(0).Optional().SchemaType(map[string]string{dialect.MySQL: "INT"}),
		field.String("biz_no").MaxLen(45).Optional().Nillable(),
		field.String("biz_type").MaxLen(45).Optional().Nillable(),
		field.Int("org_id").Comment("运营组织").SchemaType(map[string]string{dialect.MySQL: "INT"}),
		field.String("trade_code").MaxLen(50).Optional().Nillable(),
		field.String("change_type").MaxLen(10).Optional().Nillable(),
		field.Int("record_method_id").Optional().Nillable().Comment("记账请求类型"),
		field.Int("project_id").Optional().SchemaType(map[string]string{dialect.MySQL: "INT"}),
		field.Int("product_id").Optional().SchemaType(map[string]string{dialect.MySQL: "INT"}),
		field.Int("material_id").Optional().SchemaType(map[string]string{dialect.MySQL: "INT"}),
		field.String("material_no").MaxLen(45).Optional().Nillable(),
		field.String("material_name").MaxLen(45).Optional().Nillable(),
		field.Float("multiplier").Optional(),
		field.Int("position_type").Optional().SchemaType(map[string]string{dialect.MySQL: "INT"}),
		field.String("currency").MaxLen(10).Optional().Nillable(),
		field.Float("qty").Default(0.0000).Optional(),
		field.Float("price").Optional().Nillable(),
		field.Float("amount").Optional().Nillable(),
		field.Float("cost").Optional().Nillable(),
		field.Float("total").Optional().Nillable(),
		field.String("unit").MaxLen(10).Optional().Nillable(),
		field.String("spec").MaxLen(45).Optional().Nillable(),
		field.Float("commission").Optional().Nillable(),
		field.Float("fee").Optional().Nillable(),
		field.Float("fee_other").Optional().Nillable(),
		field.Int("count").Default(0).Optional().SchemaType(map[string]string{dialect.MySQL: "INT"}),
		field.Int("direction").Optional().Nillable().SchemaType(map[string]string{dialect.MySQL: "INT"}),
		field.String("tamper").MaxLen(45).Optional().Nillable(),
		field.String("summary").MaxLen(200).Optional().Nillable(),
		field.String("pair_subject_code").MaxLen(45).Optional().Nillable(),
		field.Int("pair_security_account_id").Optional().Nillable().SchemaType(map[string]string{dialect.MySQL: "INT"}),
		field.Float("last_balance").Optional().Nillable(),
		field.Time("stl_time").Optional().Nillable(),
		field.Float("fx_rate").Optional().Nillable(),
		field.String("stl_currency").MaxLen(10).Optional().Nillable(),
		field.Float("stl_amount").Optional().Nillable(),
		field.Float("stl_margin_lv").Optional().Nillable(),
		field.String("is_settlement").Default("N").MaxLen(1).Optional(),
		field.Time("actual_stl_time").Optional().Nillable(),
		field.String("trans_no").MaxLen(45).Optional().Nillable(),
		field.Time("trans_time").Optional().Nillable(),
		field.String("accounting_no").MaxLen(45).Optional().Nillable(),
		field.Time("accounting_time").Optional().Nillable(),
		field.Float("accounted_position_qty").Optional().Nillable(),
		field.Int("cd_direction").Optional().Nillable().SchemaType(map[string]string{dialect.MySQL: "INT"}),
		field.String("apply_fields").Optional().Nillable(),
		field.Time("created_at").Optional().Nillable(),
		field.String("ref_trans_no").MaxLen(45).Optional().Nillable(),
		field.String("is_day_booking").MaxLen(1).Optional().Nillable(),
		field.String("is_effect_fund").Default("N").MaxLen(1).Optional(),
	}
}

// Edges of the SecurityJournal.
func (SecurityJournal) Edges() []ent.Edge {
	return nil
}
