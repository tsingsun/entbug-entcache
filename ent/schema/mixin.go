package schema

import (
	"entgo.io/ent/dialect"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type EditorMixin struct {
	mixin.Schema
}

func (e EditorMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Int("created_by").Default(0).SchemaType(map[string]string{dialect.MySQL: "INT"}),
		field.Time("created_at").Immutable().Default(time.Now),
		field.Int("updated_by").Default(0).SchemaType(map[string]string{dialect.MySQL: "INT"}),
		field.Time("updated_at").Default(time.Now),
	}
}
