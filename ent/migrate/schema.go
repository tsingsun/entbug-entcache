// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// SecurityJournalColumns holds the columns for the "security_journal" table.
	SecurityJournalColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true, SchemaType: map[string]string{"mysql": "INT"}},
		{Name: "security_account_id", Type: field.TypeInt, SchemaType: map[string]string{"mysql": "INT"}},
		{Name: "account_id", Type: field.TypeInt, Nullable: true, Default: 0, SchemaType: map[string]string{"mysql": "INT"}},
		{Name: "biz_no", Type: field.TypeString, Nullable: true, Size: 45},
		{Name: "biz_type", Type: field.TypeString, Nullable: true, Size: 45},
		{Name: "org_id", Type: field.TypeInt, SchemaType: map[string]string{"mysql": "INT"}},
		{Name: "trade_code", Type: field.TypeString, Nullable: true, Size: 50},
		{Name: "change_type", Type: field.TypeString, Nullable: true, Size: 10},
		{Name: "record_method_id", Type: field.TypeInt, Nullable: true},
		{Name: "project_id", Type: field.TypeInt, Nullable: true, SchemaType: map[string]string{"mysql": "INT"}},
		{Name: "product_id", Type: field.TypeInt, Nullable: true, SchemaType: map[string]string{"mysql": "INT"}},
		{Name: "material_id", Type: field.TypeInt, Nullable: true, SchemaType: map[string]string{"mysql": "INT"}},
		{Name: "material_no", Type: field.TypeString, Nullable: true, Size: 45},
		{Name: "material_name", Type: field.TypeString, Nullable: true, Size: 45},
		{Name: "multiplier", Type: field.TypeFloat64, Nullable: true},
		{Name: "position_type", Type: field.TypeInt, Nullable: true, SchemaType: map[string]string{"mysql": "INT"}},
		{Name: "currency", Type: field.TypeString, Nullable: true, Size: 10},
		{Name: "qty", Type: field.TypeFloat64, Nullable: true, Default: 0},
		{Name: "price", Type: field.TypeFloat64, Nullable: true},
		{Name: "amount", Type: field.TypeFloat64, Nullable: true},
		{Name: "cost", Type: field.TypeFloat64, Nullable: true},
		{Name: "total", Type: field.TypeFloat64, Nullable: true},
		{Name: "unit", Type: field.TypeString, Nullable: true, Size: 10},
		{Name: "spec", Type: field.TypeString, Nullable: true, Size: 45},
		{Name: "commission", Type: field.TypeFloat64, Nullable: true},
		{Name: "fee", Type: field.TypeFloat64, Nullable: true},
		{Name: "fee_other", Type: field.TypeFloat64, Nullable: true},
		{Name: "count", Type: field.TypeInt, Nullable: true, Default: 0, SchemaType: map[string]string{"mysql": "INT"}},
		{Name: "direction", Type: field.TypeInt, Nullable: true, SchemaType: map[string]string{"mysql": "INT"}},
		{Name: "tamper", Type: field.TypeString, Nullable: true, Size: 45},
		{Name: "summary", Type: field.TypeString, Nullable: true, Size: 200},
		{Name: "pair_subject_code", Type: field.TypeString, Nullable: true, Size: 45},
		{Name: "pair_security_account_id", Type: field.TypeInt, Nullable: true, SchemaType: map[string]string{"mysql": "INT"}},
		{Name: "last_balance", Type: field.TypeFloat64, Nullable: true},
		{Name: "stl_time", Type: field.TypeTime, Nullable: true},
		{Name: "fx_rate", Type: field.TypeFloat64, Nullable: true},
		{Name: "stl_currency", Type: field.TypeString, Nullable: true, Size: 10},
		{Name: "stl_amount", Type: field.TypeFloat64, Nullable: true},
		{Name: "stl_margin_lv", Type: field.TypeFloat64, Nullable: true},
		{Name: "is_settlement", Type: field.TypeString, Nullable: true, Size: 1, Default: "N"},
		{Name: "actual_stl_time", Type: field.TypeTime, Nullable: true},
		{Name: "trans_no", Type: field.TypeString, Nullable: true, Size: 45},
		{Name: "trans_time", Type: field.TypeTime, Nullable: true},
		{Name: "accounting_no", Type: field.TypeString, Nullable: true, Size: 45},
		{Name: "accounting_time", Type: field.TypeTime, Nullable: true},
		{Name: "accounted_position_qty", Type: field.TypeFloat64, Nullable: true},
		{Name: "cd_direction", Type: field.TypeInt, Nullable: true, SchemaType: map[string]string{"mysql": "INT"}},
		{Name: "apply_fields", Type: field.TypeString, Nullable: true},
		{Name: "created_at", Type: field.TypeTime, Nullable: true},
		{Name: "ref_trans_no", Type: field.TypeString, Nullable: true, Size: 45},
		{Name: "is_day_booking", Type: field.TypeString, Nullable: true, Size: 1},
		{Name: "is_effect_fund", Type: field.TypeString, Nullable: true, Size: 1, Default: "N"},
	}
	// SecurityJournalTable holds the schema information for the "security_journal" table.
	SecurityJournalTable = &schema.Table{
		Name:       "security_journal",
		Columns:    SecurityJournalColumns,
		PrimaryKey: []*schema.Column{SecurityJournalColumns[0]},
	}
	// SecurityPositionColumns holds the columns for the "security_position" table.
	SecurityPositionColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true, SchemaType: map[string]string{"mysql": "INT"}},
		{Name: "parent_id", Type: field.TypeInt, Default: 0, SchemaType: map[string]string{"mysql": "INT"}},
		{Name: "account_id", Type: field.TypeInt, Nullable: true, SchemaType: map[string]string{"mysql": "INT"}},
		{Name: "security_account_id", Type: field.TypeInt, SchemaType: map[string]string{"mysql": "INT"}},
		{Name: "position_type", Type: field.TypeInt, Nullable: true, SchemaType: map[string]string{"mysql": "INT"}},
		{Name: "cd_direction", Type: field.TypeInt, Nullable: true, SchemaType: map[string]string{"mysql": "INT"}},
		{Name: "project_id", Type: field.TypeInt, Nullable: true, SchemaType: map[string]string{"mysql": "INT"}},
		{Name: "product_id", Type: field.TypeInt, SchemaType: map[string]string{"mysql": "INT"}},
		{Name: "material_id", Type: field.TypeInt, Nullable: true, SchemaType: map[string]string{"mysql": "INT"}},
		{Name: "material_no", Type: field.TypeString, Nullable: true, Size: 45},
		{Name: "multiplier", Type: field.TypeFloat64, Nullable: true},
		{Name: "balance", Type: field.TypeFloat64, Nullable: true, Default: 0},
		{Name: "available", Type: field.TypeFloat64, Nullable: true, Default: 0},
		{Name: "freeze", Type: field.TypeFloat64, Nullable: true, Default: 0},
		{Name: "afloat", Type: field.TypeFloat64, Nullable: true, Default: 0},
		{Name: "unit", Type: field.TypeString, Nullable: true, Size: 10},
		{Name: "spec", Type: field.TypeString, Nullable: true, Size: 45},
		{Name: "currency", Type: field.TypeString, Nullable: true, Size: 10},
		{Name: "price", Type: field.TypeFloat64, Nullable: true, Default: 0},
		{Name: "amount", Type: field.TypeFloat64, Nullable: true},
		{Name: "cost_amount", Type: field.TypeFloat64, Nullable: true},
		{Name: "cost", Type: field.TypeFloat64, Nullable: true},
		{Name: "fx_rate", Type: field.TypeFloat64, Nullable: true},
		{Name: "stl_currency", Type: field.TypeString, Nullable: true, Size: 10},
		{Name: "stl_amount", Type: field.TypeFloat64, Nullable: true},
		{Name: "stl_cost", Type: field.TypeFloat64, Nullable: true},
		{Name: "stl_margin", Type: field.TypeFloat64, Nullable: true},
		{Name: "stl_cost_amount", Type: field.TypeFloat64, Nullable: true},
		{Name: "stl_margin_lv", Type: field.TypeFloat64, Nullable: true, Default: 0},
		{Name: "stl_val_price", Type: field.TypeFloat64, Nullable: true},
		{Name: "invalid_at", Type: field.TypeTime, Nullable: true},
		{Name: "created_at", Type: field.TypeTime, Nullable: true},
		{Name: "unique_tag", Type: field.TypeString, Unique: true, Nullable: true, Size: 45},
		{Name: "org_id", Type: field.TypeInt, SchemaType: map[string]string{"mysql": "INT"}},
	}
	// SecurityPositionTable holds the schema information for the "security_position" table.
	SecurityPositionTable = &schema.Table{
		Name:       "security_position",
		Columns:    SecurityPositionColumns,
		PrimaryKey: []*schema.Column{SecurityPositionColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "age", Type: field.TypeInt},
		{Name: "name", Type: field.TypeString},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		SecurityJournalTable,
		SecurityPositionTable,
		UsersTable,
	}
)

func init() {
	SecurityJournalTable.Annotation = &entsql.Annotation{
		Table: "security_journal",
	}
	SecurityPositionTable.Annotation = &entsql.Annotation{
		Table: "security_position",
	}
}
