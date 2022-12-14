// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/bug/ent/securityposition"
	"entgo.io/ent/dialect/sql"
)

// SecurityPosition is the model entity for the SecurityPosition schema.
type SecurityPosition struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// ParentID holds the value of the "parent_id" field.
	ParentID int `json:"parent_id,omitempty"`
	// AccountID holds the value of the "account_id" field.
	AccountID int `json:"account_id,omitempty"`
	// SecurityAccountID holds the value of the "security_account_id" field.
	SecurityAccountID int `json:"security_account_id,omitempty"`
	// PositionType holds the value of the "position_type" field.
	PositionType int `json:"position_type,omitempty"`
	// CdDirection holds the value of the "cd_direction" field.
	CdDirection *int `json:"cd_direction,omitempty"`
	// ProjectID holds the value of the "project_id" field.
	ProjectID int `json:"project_id,omitempty"`
	// ProductID holds the value of the "product_id" field.
	ProductID int `json:"product_id,omitempty"`
	// MaterialID holds the value of the "material_id" field.
	MaterialID int `json:"material_id,omitempty"`
	// 标的代码
	MaterialNo *string `json:"material_no,omitempty"`
	// Multiplier holds the value of the "multiplier" field.
	Multiplier *float64 `json:"multiplier,omitempty"`
	// Balance holds the value of the "balance" field.
	Balance float64 `json:"balance,omitempty"`
	// Available holds the value of the "available" field.
	Available float64 `json:"available,omitempty"`
	// Freeze holds the value of the "freeze" field.
	Freeze float64 `json:"freeze,omitempty"`
	// Afloat holds the value of the "afloat" field.
	Afloat float64 `json:"afloat,omitempty"`
	// Unit holds the value of the "unit" field.
	Unit *string `json:"unit,omitempty"`
	// Spec holds the value of the "spec" field.
	Spec *string `json:"spec,omitempty"`
	// Currency holds the value of the "currency" field.
	Currency *string `json:"currency,omitempty"`
	// Price holds the value of the "price" field.
	Price float64 `json:"price,omitempty"`
	// Amount holds the value of the "amount" field.
	Amount *float64 `json:"amount,omitempty"`
	// CostAmount holds the value of the "cost_amount" field.
	CostAmount *float64 `json:"cost_amount,omitempty"`
	// Cost holds the value of the "cost" field.
	Cost *float64 `json:"cost,omitempty"`
	// FxRate holds the value of the "fx_rate" field.
	FxRate *float64 `json:"fx_rate,omitempty"`
	// StlCurrency holds the value of the "stl_currency" field.
	StlCurrency *string `json:"stl_currency,omitempty"`
	// StlAmount holds the value of the "stl_amount" field.
	StlAmount *float64 `json:"stl_amount,omitempty"`
	// StlCost holds the value of the "stl_cost" field.
	StlCost *float64 `json:"stl_cost,omitempty"`
	// StlMargin holds the value of the "stl_margin" field.
	StlMargin *float64 `json:"stl_margin,omitempty"`
	// StlCostAmount holds the value of the "stl_cost_amount" field.
	StlCostAmount *float64 `json:"stl_cost_amount,omitempty"`
	// StlMarginLv holds the value of the "stl_margin_lv" field.
	StlMarginLv float64 `json:"stl_margin_lv,omitempty"`
	// StlValPrice holds the value of the "stl_val_price" field.
	StlValPrice *float64 `json:"stl_val_price,omitempty"`
	// InvalidAt holds the value of the "invalid_at" field.
	InvalidAt *time.Time `json:"invalid_at,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// UniqueTag holds the value of the "unique_tag" field.
	UniqueTag *string `json:"unique_tag,omitempty"`
	// OrgID holds the value of the "org_id" field.
	OrgID int `json:"org_id,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SecurityPosition) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case securityposition.FieldMultiplier, securityposition.FieldBalance, securityposition.FieldAvailable, securityposition.FieldFreeze, securityposition.FieldAfloat, securityposition.FieldPrice, securityposition.FieldAmount, securityposition.FieldCostAmount, securityposition.FieldCost, securityposition.FieldFxRate, securityposition.FieldStlAmount, securityposition.FieldStlCost, securityposition.FieldStlMargin, securityposition.FieldStlCostAmount, securityposition.FieldStlMarginLv, securityposition.FieldStlValPrice:
			values[i] = new(sql.NullFloat64)
		case securityposition.FieldID, securityposition.FieldParentID, securityposition.FieldAccountID, securityposition.FieldSecurityAccountID, securityposition.FieldPositionType, securityposition.FieldCdDirection, securityposition.FieldProjectID, securityposition.FieldProductID, securityposition.FieldMaterialID, securityposition.FieldOrgID:
			values[i] = new(sql.NullInt64)
		case securityposition.FieldMaterialNo, securityposition.FieldUnit, securityposition.FieldSpec, securityposition.FieldCurrency, securityposition.FieldStlCurrency, securityposition.FieldUniqueTag:
			values[i] = new(sql.NullString)
		case securityposition.FieldInvalidAt, securityposition.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type SecurityPosition", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the SecurityPosition fields.
func (sp *SecurityPosition) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case securityposition.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			sp.ID = int(value.Int64)
		case securityposition.FieldParentID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field parent_id", values[i])
			} else if value.Valid {
				sp.ParentID = int(value.Int64)
			}
		case securityposition.FieldAccountID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field account_id", values[i])
			} else if value.Valid {
				sp.AccountID = int(value.Int64)
			}
		case securityposition.FieldSecurityAccountID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field security_account_id", values[i])
			} else if value.Valid {
				sp.SecurityAccountID = int(value.Int64)
			}
		case securityposition.FieldPositionType:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field position_type", values[i])
			} else if value.Valid {
				sp.PositionType = int(value.Int64)
			}
		case securityposition.FieldCdDirection:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field cd_direction", values[i])
			} else if value.Valid {
				sp.CdDirection = new(int)
				*sp.CdDirection = int(value.Int64)
			}
		case securityposition.FieldProjectID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field project_id", values[i])
			} else if value.Valid {
				sp.ProjectID = int(value.Int64)
			}
		case securityposition.FieldProductID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field product_id", values[i])
			} else if value.Valid {
				sp.ProductID = int(value.Int64)
			}
		case securityposition.FieldMaterialID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field material_id", values[i])
			} else if value.Valid {
				sp.MaterialID = int(value.Int64)
			}
		case securityposition.FieldMaterialNo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field material_no", values[i])
			} else if value.Valid {
				sp.MaterialNo = new(string)
				*sp.MaterialNo = value.String
			}
		case securityposition.FieldMultiplier:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field multiplier", values[i])
			} else if value.Valid {
				sp.Multiplier = new(float64)
				*sp.Multiplier = value.Float64
			}
		case securityposition.FieldBalance:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field balance", values[i])
			} else if value.Valid {
				sp.Balance = value.Float64
			}
		case securityposition.FieldAvailable:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field available", values[i])
			} else if value.Valid {
				sp.Available = value.Float64
			}
		case securityposition.FieldFreeze:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field freeze", values[i])
			} else if value.Valid {
				sp.Freeze = value.Float64
			}
		case securityposition.FieldAfloat:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field afloat", values[i])
			} else if value.Valid {
				sp.Afloat = value.Float64
			}
		case securityposition.FieldUnit:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field unit", values[i])
			} else if value.Valid {
				sp.Unit = new(string)
				*sp.Unit = value.String
			}
		case securityposition.FieldSpec:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field spec", values[i])
			} else if value.Valid {
				sp.Spec = new(string)
				*sp.Spec = value.String
			}
		case securityposition.FieldCurrency:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field currency", values[i])
			} else if value.Valid {
				sp.Currency = new(string)
				*sp.Currency = value.String
			}
		case securityposition.FieldPrice:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field price", values[i])
			} else if value.Valid {
				sp.Price = value.Float64
			}
		case securityposition.FieldAmount:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field amount", values[i])
			} else if value.Valid {
				sp.Amount = new(float64)
				*sp.Amount = value.Float64
			}
		case securityposition.FieldCostAmount:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field cost_amount", values[i])
			} else if value.Valid {
				sp.CostAmount = new(float64)
				*sp.CostAmount = value.Float64
			}
		case securityposition.FieldCost:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field cost", values[i])
			} else if value.Valid {
				sp.Cost = new(float64)
				*sp.Cost = value.Float64
			}
		case securityposition.FieldFxRate:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field fx_rate", values[i])
			} else if value.Valid {
				sp.FxRate = new(float64)
				*sp.FxRate = value.Float64
			}
		case securityposition.FieldStlCurrency:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field stl_currency", values[i])
			} else if value.Valid {
				sp.StlCurrency = new(string)
				*sp.StlCurrency = value.String
			}
		case securityposition.FieldStlAmount:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field stl_amount", values[i])
			} else if value.Valid {
				sp.StlAmount = new(float64)
				*sp.StlAmount = value.Float64
			}
		case securityposition.FieldStlCost:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field stl_cost", values[i])
			} else if value.Valid {
				sp.StlCost = new(float64)
				*sp.StlCost = value.Float64
			}
		case securityposition.FieldStlMargin:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field stl_margin", values[i])
			} else if value.Valid {
				sp.StlMargin = new(float64)
				*sp.StlMargin = value.Float64
			}
		case securityposition.FieldStlCostAmount:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field stl_cost_amount", values[i])
			} else if value.Valid {
				sp.StlCostAmount = new(float64)
				*sp.StlCostAmount = value.Float64
			}
		case securityposition.FieldStlMarginLv:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field stl_margin_lv", values[i])
			} else if value.Valid {
				sp.StlMarginLv = value.Float64
			}
		case securityposition.FieldStlValPrice:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field stl_val_price", values[i])
			} else if value.Valid {
				sp.StlValPrice = new(float64)
				*sp.StlValPrice = value.Float64
			}
		case securityposition.FieldInvalidAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field invalid_at", values[i])
			} else if value.Valid {
				sp.InvalidAt = new(time.Time)
				*sp.InvalidAt = value.Time
			}
		case securityposition.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				sp.CreatedAt = new(time.Time)
				*sp.CreatedAt = value.Time
			}
		case securityposition.FieldUniqueTag:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field unique_tag", values[i])
			} else if value.Valid {
				sp.UniqueTag = new(string)
				*sp.UniqueTag = value.String
			}
		case securityposition.FieldOrgID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field org_id", values[i])
			} else if value.Valid {
				sp.OrgID = int(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this SecurityPosition.
// Note that you need to call SecurityPosition.Unwrap() before calling this method if this SecurityPosition
// was returned from a transaction, and the transaction was committed or rolled back.
func (sp *SecurityPosition) Update() *SecurityPositionUpdateOne {
	return (&SecurityPositionClient{config: sp.config}).UpdateOne(sp)
}

// Unwrap unwraps the SecurityPosition entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sp *SecurityPosition) Unwrap() *SecurityPosition {
	_tx, ok := sp.config.driver.(*txDriver)
	if !ok {
		panic("ent: SecurityPosition is not a transactional entity")
	}
	sp.config.driver = _tx.drv
	return sp
}

// String implements the fmt.Stringer.
func (sp *SecurityPosition) String() string {
	var builder strings.Builder
	builder.WriteString("SecurityPosition(")
	builder.WriteString(fmt.Sprintf("id=%v, ", sp.ID))
	builder.WriteString("parent_id=")
	builder.WriteString(fmt.Sprintf("%v", sp.ParentID))
	builder.WriteString(", ")
	builder.WriteString("account_id=")
	builder.WriteString(fmt.Sprintf("%v", sp.AccountID))
	builder.WriteString(", ")
	builder.WriteString("security_account_id=")
	builder.WriteString(fmt.Sprintf("%v", sp.SecurityAccountID))
	builder.WriteString(", ")
	builder.WriteString("position_type=")
	builder.WriteString(fmt.Sprintf("%v", sp.PositionType))
	builder.WriteString(", ")
	if v := sp.CdDirection; v != nil {
		builder.WriteString("cd_direction=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	builder.WriteString("project_id=")
	builder.WriteString(fmt.Sprintf("%v", sp.ProjectID))
	builder.WriteString(", ")
	builder.WriteString("product_id=")
	builder.WriteString(fmt.Sprintf("%v", sp.ProductID))
	builder.WriteString(", ")
	builder.WriteString("material_id=")
	builder.WriteString(fmt.Sprintf("%v", sp.MaterialID))
	builder.WriteString(", ")
	if v := sp.MaterialNo; v != nil {
		builder.WriteString("material_no=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := sp.Multiplier; v != nil {
		builder.WriteString("multiplier=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	builder.WriteString("balance=")
	builder.WriteString(fmt.Sprintf("%v", sp.Balance))
	builder.WriteString(", ")
	builder.WriteString("available=")
	builder.WriteString(fmt.Sprintf("%v", sp.Available))
	builder.WriteString(", ")
	builder.WriteString("freeze=")
	builder.WriteString(fmt.Sprintf("%v", sp.Freeze))
	builder.WriteString(", ")
	builder.WriteString("afloat=")
	builder.WriteString(fmt.Sprintf("%v", sp.Afloat))
	builder.WriteString(", ")
	if v := sp.Unit; v != nil {
		builder.WriteString("unit=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := sp.Spec; v != nil {
		builder.WriteString("spec=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := sp.Currency; v != nil {
		builder.WriteString("currency=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("price=")
	builder.WriteString(fmt.Sprintf("%v", sp.Price))
	builder.WriteString(", ")
	if v := sp.Amount; v != nil {
		builder.WriteString("amount=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := sp.CostAmount; v != nil {
		builder.WriteString("cost_amount=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := sp.Cost; v != nil {
		builder.WriteString("cost=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := sp.FxRate; v != nil {
		builder.WriteString("fx_rate=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := sp.StlCurrency; v != nil {
		builder.WriteString("stl_currency=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := sp.StlAmount; v != nil {
		builder.WriteString("stl_amount=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := sp.StlCost; v != nil {
		builder.WriteString("stl_cost=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := sp.StlMargin; v != nil {
		builder.WriteString("stl_margin=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := sp.StlCostAmount; v != nil {
		builder.WriteString("stl_cost_amount=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	builder.WriteString("stl_margin_lv=")
	builder.WriteString(fmt.Sprintf("%v", sp.StlMarginLv))
	builder.WriteString(", ")
	if v := sp.StlValPrice; v != nil {
		builder.WriteString("stl_val_price=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := sp.InvalidAt; v != nil {
		builder.WriteString("invalid_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := sp.CreatedAt; v != nil {
		builder.WriteString("created_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := sp.UniqueTag; v != nil {
		builder.WriteString("unique_tag=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("org_id=")
	builder.WriteString(fmt.Sprintf("%v", sp.OrgID))
	builder.WriteByte(')')
	return builder.String()
}

// SecurityPositions is a parsable slice of SecurityPosition.
type SecurityPositions []*SecurityPosition

func (sp SecurityPositions) config(cfg config) {
	for _i := range sp {
		sp[_i].config = cfg
	}
}
