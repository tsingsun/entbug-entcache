// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/bug/ent/securityposition"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SecurityPositionCreate is the builder for creating a SecurityPosition entity.
type SecurityPositionCreate struct {
	config
	mutation *SecurityPositionMutation
	hooks    []Hook
}

// SetParentID sets the "parent_id" field.
func (spc *SecurityPositionCreate) SetParentID(i int) *SecurityPositionCreate {
	spc.mutation.SetParentID(i)
	return spc
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (spc *SecurityPositionCreate) SetNillableParentID(i *int) *SecurityPositionCreate {
	if i != nil {
		spc.SetParentID(*i)
	}
	return spc
}

// SetAccountID sets the "account_id" field.
func (spc *SecurityPositionCreate) SetAccountID(i int) *SecurityPositionCreate {
	spc.mutation.SetAccountID(i)
	return spc
}

// SetNillableAccountID sets the "account_id" field if the given value is not nil.
func (spc *SecurityPositionCreate) SetNillableAccountID(i *int) *SecurityPositionCreate {
	if i != nil {
		spc.SetAccountID(*i)
	}
	return spc
}

// SetSecurityAccountID sets the "security_account_id" field.
func (spc *SecurityPositionCreate) SetSecurityAccountID(i int) *SecurityPositionCreate {
	spc.mutation.SetSecurityAccountID(i)
	return spc
}

// SetPositionType sets the "position_type" field.
func (spc *SecurityPositionCreate) SetPositionType(i int) *SecurityPositionCreate {
	spc.mutation.SetPositionType(i)
	return spc
}

// SetNillablePositionType sets the "position_type" field if the given value is not nil.
func (spc *SecurityPositionCreate) SetNillablePositionType(i *int) *SecurityPositionCreate {
	if i != nil {
		spc.SetPositionType(*i)
	}
	return spc
}

// SetCdDirection sets the "cd_direction" field.
func (spc *SecurityPositionCreate) SetCdDirection(i int) *SecurityPositionCreate {
	spc.mutation.SetCdDirection(i)
	return spc
}

// SetNillableCdDirection sets the "cd_direction" field if the given value is not nil.
func (spc *SecurityPositionCreate) SetNillableCdDirection(i *int) *SecurityPositionCreate {
	if i != nil {
		spc.SetCdDirection(*i)
	}
	return spc
}

// SetProjectID sets the "project_id" field.
func (spc *SecurityPositionCreate) SetProjectID(i int) *SecurityPositionCreate {
	spc.mutation.SetProjectID(i)
	return spc
}

// SetNillableProjectID sets the "project_id" field if the given value is not nil.
func (spc *SecurityPositionCreate) SetNillableProjectID(i *int) *SecurityPositionCreate {
	if i != nil {
		spc.SetProjectID(*i)
	}
	return spc
}

// SetProductID sets the "product_id" field.
func (spc *SecurityPositionCreate) SetProductID(i int) *SecurityPositionCreate {
	spc.mutation.SetProductID(i)
	return spc
}

// SetMaterialID sets the "material_id" field.
func (spc *SecurityPositionCreate) SetMaterialID(i int) *SecurityPositionCreate {
	spc.mutation.SetMaterialID(i)
	return spc
}

// SetNillableMaterialID sets the "material_id" field if the given value is not nil.
func (spc *SecurityPositionCreate) SetNillableMaterialID(i *int) *SecurityPositionCreate {
	if i != nil {
		spc.SetMaterialID(*i)
	}
	return spc
}

// SetMaterialNo sets the "material_no" field.
func (spc *SecurityPositionCreate) SetMaterialNo(s string) *SecurityPositionCreate {
	spc.mutation.SetMaterialNo(s)
	return spc
}

// SetNillableMaterialNo sets the "material_no" field if the given value is not nil.
func (spc *SecurityPositionCreate) SetNillableMaterialNo(s *string) *SecurityPositionCreate {
	if s != nil {
		spc.SetMaterialNo(*s)
	}
	return spc
}

// SetMultiplier sets the "multiplier" field.
func (spc *SecurityPositionCreate) SetMultiplier(f float64) *SecurityPositionCreate {
	spc.mutation.SetMultiplier(f)
	return spc
}

// SetNillableMultiplier sets the "multiplier" field if the given value is not nil.
func (spc *SecurityPositionCreate) SetNillableMultiplier(f *float64) *SecurityPositionCreate {
	if f != nil {
		spc.SetMultiplier(*f)
	}
	return spc
}

// SetBalance sets the "balance" field.
func (spc *SecurityPositionCreate) SetBalance(f float64) *SecurityPositionCreate {
	spc.mutation.SetBalance(f)
	return spc
}

// SetNillableBalance sets the "balance" field if the given value is not nil.
func (spc *SecurityPositionCreate) SetNillableBalance(f *float64) *SecurityPositionCreate {
	if f != nil {
		spc.SetBalance(*f)
	}
	return spc
}

// SetAvailable sets the "available" field.
func (spc *SecurityPositionCreate) SetAvailable(f float64) *SecurityPositionCreate {
	spc.mutation.SetAvailable(f)
	return spc
}

// SetNillableAvailable sets the "available" field if the given value is not nil.
func (spc *SecurityPositionCreate) SetNillableAvailable(f *float64) *SecurityPositionCreate {
	if f != nil {
		spc.SetAvailable(*f)
	}
	return spc
}

// SetFreeze sets the "freeze" field.
func (spc *SecurityPositionCreate) SetFreeze(f float64) *SecurityPositionCreate {
	spc.mutation.SetFreeze(f)
	return spc
}

// SetNillableFreeze sets the "freeze" field if the given value is not nil.
func (spc *SecurityPositionCreate) SetNillableFreeze(f *float64) *SecurityPositionCreate {
	if f != nil {
		spc.SetFreeze(*f)
	}
	return spc
}

// SetAfloat sets the "afloat" field.
func (spc *SecurityPositionCreate) SetAfloat(f float64) *SecurityPositionCreate {
	spc.mutation.SetAfloat(f)
	return spc
}

// SetNillableAfloat sets the "afloat" field if the given value is not nil.
func (spc *SecurityPositionCreate) SetNillableAfloat(f *float64) *SecurityPositionCreate {
	if f != nil {
		spc.SetAfloat(*f)
	}
	return spc
}

// SetUnit sets the "unit" field.
func (spc *SecurityPositionCreate) SetUnit(s string) *SecurityPositionCreate {
	spc.mutation.SetUnit(s)
	return spc
}

// SetNillableUnit sets the "unit" field if the given value is not nil.
func (spc *SecurityPositionCreate) SetNillableUnit(s *string) *SecurityPositionCreate {
	if s != nil {
		spc.SetUnit(*s)
	}
	return spc
}

// SetSpec sets the "spec" field.
func (spc *SecurityPositionCreate) SetSpec(s string) *SecurityPositionCreate {
	spc.mutation.SetSpec(s)
	return spc
}

// SetNillableSpec sets the "spec" field if the given value is not nil.
func (spc *SecurityPositionCreate) SetNillableSpec(s *string) *SecurityPositionCreate {
	if s != nil {
		spc.SetSpec(*s)
	}
	return spc
}

// SetCurrency sets the "currency" field.
func (spc *SecurityPositionCreate) SetCurrency(s string) *SecurityPositionCreate {
	spc.mutation.SetCurrency(s)
	return spc
}

// SetNillableCurrency sets the "currency" field if the given value is not nil.
func (spc *SecurityPositionCreate) SetNillableCurrency(s *string) *SecurityPositionCreate {
	if s != nil {
		spc.SetCurrency(*s)
	}
	return spc
}

// SetPrice sets the "price" field.
func (spc *SecurityPositionCreate) SetPrice(f float64) *SecurityPositionCreate {
	spc.mutation.SetPrice(f)
	return spc
}

// SetNillablePrice sets the "price" field if the given value is not nil.
func (spc *SecurityPositionCreate) SetNillablePrice(f *float64) *SecurityPositionCreate {
	if f != nil {
		spc.SetPrice(*f)
	}
	return spc
}

// SetAmount sets the "amount" field.
func (spc *SecurityPositionCreate) SetAmount(f float64) *SecurityPositionCreate {
	spc.mutation.SetAmount(f)
	return spc
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (spc *SecurityPositionCreate) SetNillableAmount(f *float64) *SecurityPositionCreate {
	if f != nil {
		spc.SetAmount(*f)
	}
	return spc
}

// SetCostAmount sets the "cost_amount" field.
func (spc *SecurityPositionCreate) SetCostAmount(f float64) *SecurityPositionCreate {
	spc.mutation.SetCostAmount(f)
	return spc
}

// SetNillableCostAmount sets the "cost_amount" field if the given value is not nil.
func (spc *SecurityPositionCreate) SetNillableCostAmount(f *float64) *SecurityPositionCreate {
	if f != nil {
		spc.SetCostAmount(*f)
	}
	return spc
}

// SetCost sets the "cost" field.
func (spc *SecurityPositionCreate) SetCost(f float64) *SecurityPositionCreate {
	spc.mutation.SetCost(f)
	return spc
}

// SetNillableCost sets the "cost" field if the given value is not nil.
func (spc *SecurityPositionCreate) SetNillableCost(f *float64) *SecurityPositionCreate {
	if f != nil {
		spc.SetCost(*f)
	}
	return spc
}

// SetFxRate sets the "fx_rate" field.
func (spc *SecurityPositionCreate) SetFxRate(f float64) *SecurityPositionCreate {
	spc.mutation.SetFxRate(f)
	return spc
}

// SetNillableFxRate sets the "fx_rate" field if the given value is not nil.
func (spc *SecurityPositionCreate) SetNillableFxRate(f *float64) *SecurityPositionCreate {
	if f != nil {
		spc.SetFxRate(*f)
	}
	return spc
}

// SetStlCurrency sets the "stl_currency" field.
func (spc *SecurityPositionCreate) SetStlCurrency(s string) *SecurityPositionCreate {
	spc.mutation.SetStlCurrency(s)
	return spc
}

// SetNillableStlCurrency sets the "stl_currency" field if the given value is not nil.
func (spc *SecurityPositionCreate) SetNillableStlCurrency(s *string) *SecurityPositionCreate {
	if s != nil {
		spc.SetStlCurrency(*s)
	}
	return spc
}

// SetStlAmount sets the "stl_amount" field.
func (spc *SecurityPositionCreate) SetStlAmount(f float64) *SecurityPositionCreate {
	spc.mutation.SetStlAmount(f)
	return spc
}

// SetNillableStlAmount sets the "stl_amount" field if the given value is not nil.
func (spc *SecurityPositionCreate) SetNillableStlAmount(f *float64) *SecurityPositionCreate {
	if f != nil {
		spc.SetStlAmount(*f)
	}
	return spc
}

// SetStlCost sets the "stl_cost" field.
func (spc *SecurityPositionCreate) SetStlCost(f float64) *SecurityPositionCreate {
	spc.mutation.SetStlCost(f)
	return spc
}

// SetNillableStlCost sets the "stl_cost" field if the given value is not nil.
func (spc *SecurityPositionCreate) SetNillableStlCost(f *float64) *SecurityPositionCreate {
	if f != nil {
		spc.SetStlCost(*f)
	}
	return spc
}

// SetStlMargin sets the "stl_margin" field.
func (spc *SecurityPositionCreate) SetStlMargin(f float64) *SecurityPositionCreate {
	spc.mutation.SetStlMargin(f)
	return spc
}

// SetNillableStlMargin sets the "stl_margin" field if the given value is not nil.
func (spc *SecurityPositionCreate) SetNillableStlMargin(f *float64) *SecurityPositionCreate {
	if f != nil {
		spc.SetStlMargin(*f)
	}
	return spc
}

// SetStlCostAmount sets the "stl_cost_amount" field.
func (spc *SecurityPositionCreate) SetStlCostAmount(f float64) *SecurityPositionCreate {
	spc.mutation.SetStlCostAmount(f)
	return spc
}

// SetNillableStlCostAmount sets the "stl_cost_amount" field if the given value is not nil.
func (spc *SecurityPositionCreate) SetNillableStlCostAmount(f *float64) *SecurityPositionCreate {
	if f != nil {
		spc.SetStlCostAmount(*f)
	}
	return spc
}

// SetStlMarginLv sets the "stl_margin_lv" field.
func (spc *SecurityPositionCreate) SetStlMarginLv(f float64) *SecurityPositionCreate {
	spc.mutation.SetStlMarginLv(f)
	return spc
}

// SetNillableStlMarginLv sets the "stl_margin_lv" field if the given value is not nil.
func (spc *SecurityPositionCreate) SetNillableStlMarginLv(f *float64) *SecurityPositionCreate {
	if f != nil {
		spc.SetStlMarginLv(*f)
	}
	return spc
}

// SetStlValPrice sets the "stl_val_price" field.
func (spc *SecurityPositionCreate) SetStlValPrice(f float64) *SecurityPositionCreate {
	spc.mutation.SetStlValPrice(f)
	return spc
}

// SetNillableStlValPrice sets the "stl_val_price" field if the given value is not nil.
func (spc *SecurityPositionCreate) SetNillableStlValPrice(f *float64) *SecurityPositionCreate {
	if f != nil {
		spc.SetStlValPrice(*f)
	}
	return spc
}

// SetInvalidAt sets the "invalid_at" field.
func (spc *SecurityPositionCreate) SetInvalidAt(t time.Time) *SecurityPositionCreate {
	spc.mutation.SetInvalidAt(t)
	return spc
}

// SetNillableInvalidAt sets the "invalid_at" field if the given value is not nil.
func (spc *SecurityPositionCreate) SetNillableInvalidAt(t *time.Time) *SecurityPositionCreate {
	if t != nil {
		spc.SetInvalidAt(*t)
	}
	return spc
}

// SetCreatedAt sets the "created_at" field.
func (spc *SecurityPositionCreate) SetCreatedAt(t time.Time) *SecurityPositionCreate {
	spc.mutation.SetCreatedAt(t)
	return spc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (spc *SecurityPositionCreate) SetNillableCreatedAt(t *time.Time) *SecurityPositionCreate {
	if t != nil {
		spc.SetCreatedAt(*t)
	}
	return spc
}

// SetUniqueTag sets the "unique_tag" field.
func (spc *SecurityPositionCreate) SetUniqueTag(s string) *SecurityPositionCreate {
	spc.mutation.SetUniqueTag(s)
	return spc
}

// SetNillableUniqueTag sets the "unique_tag" field if the given value is not nil.
func (spc *SecurityPositionCreate) SetNillableUniqueTag(s *string) *SecurityPositionCreate {
	if s != nil {
		spc.SetUniqueTag(*s)
	}
	return spc
}

// SetOrgID sets the "org_id" field.
func (spc *SecurityPositionCreate) SetOrgID(i int) *SecurityPositionCreate {
	spc.mutation.SetOrgID(i)
	return spc
}

// SetID sets the "id" field.
func (spc *SecurityPositionCreate) SetID(i int) *SecurityPositionCreate {
	spc.mutation.SetID(i)
	return spc
}

// Mutation returns the SecurityPositionMutation object of the builder.
func (spc *SecurityPositionCreate) Mutation() *SecurityPositionMutation {
	return spc.mutation
}

// Save creates the SecurityPosition in the database.
func (spc *SecurityPositionCreate) Save(ctx context.Context) (*SecurityPosition, error) {
	var (
		err  error
		node *SecurityPosition
	)
	spc.defaults()
	if len(spc.hooks) == 0 {
		if err = spc.check(); err != nil {
			return nil, err
		}
		node, err = spc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SecurityPositionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = spc.check(); err != nil {
				return nil, err
			}
			spc.mutation = mutation
			if node, err = spc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(spc.hooks) - 1; i >= 0; i-- {
			if spc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = spc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, spc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*SecurityPosition)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from SecurityPositionMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (spc *SecurityPositionCreate) SaveX(ctx context.Context) *SecurityPosition {
	v, err := spc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (spc *SecurityPositionCreate) Exec(ctx context.Context) error {
	_, err := spc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (spc *SecurityPositionCreate) ExecX(ctx context.Context) {
	if err := spc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (spc *SecurityPositionCreate) defaults() {
	if _, ok := spc.mutation.ParentID(); !ok {
		v := securityposition.DefaultParentID
		spc.mutation.SetParentID(v)
	}
	if _, ok := spc.mutation.Balance(); !ok {
		v := securityposition.DefaultBalance
		spc.mutation.SetBalance(v)
	}
	if _, ok := spc.mutation.Available(); !ok {
		v := securityposition.DefaultAvailable
		spc.mutation.SetAvailable(v)
	}
	if _, ok := spc.mutation.Freeze(); !ok {
		v := securityposition.DefaultFreeze
		spc.mutation.SetFreeze(v)
	}
	if _, ok := spc.mutation.Afloat(); !ok {
		v := securityposition.DefaultAfloat
		spc.mutation.SetAfloat(v)
	}
	if _, ok := spc.mutation.Price(); !ok {
		v := securityposition.DefaultPrice
		spc.mutation.SetPrice(v)
	}
	if _, ok := spc.mutation.StlMarginLv(); !ok {
		v := securityposition.DefaultStlMarginLv
		spc.mutation.SetStlMarginLv(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (spc *SecurityPositionCreate) check() error {
	if _, ok := spc.mutation.ParentID(); !ok {
		return &ValidationError{Name: "parent_id", err: errors.New(`ent: missing required field "SecurityPosition.parent_id"`)}
	}
	if _, ok := spc.mutation.SecurityAccountID(); !ok {
		return &ValidationError{Name: "security_account_id", err: errors.New(`ent: missing required field "SecurityPosition.security_account_id"`)}
	}
	if _, ok := spc.mutation.ProductID(); !ok {
		return &ValidationError{Name: "product_id", err: errors.New(`ent: missing required field "SecurityPosition.product_id"`)}
	}
	if v, ok := spc.mutation.MaterialNo(); ok {
		if err := securityposition.MaterialNoValidator(v); err != nil {
			return &ValidationError{Name: "material_no", err: fmt.Errorf(`ent: validator failed for field "SecurityPosition.material_no": %w`, err)}
		}
	}
	if v, ok := spc.mutation.Unit(); ok {
		if err := securityposition.UnitValidator(v); err != nil {
			return &ValidationError{Name: "unit", err: fmt.Errorf(`ent: validator failed for field "SecurityPosition.unit": %w`, err)}
		}
	}
	if v, ok := spc.mutation.Spec(); ok {
		if err := securityposition.SpecValidator(v); err != nil {
			return &ValidationError{Name: "spec", err: fmt.Errorf(`ent: validator failed for field "SecurityPosition.spec": %w`, err)}
		}
	}
	if v, ok := spc.mutation.Currency(); ok {
		if err := securityposition.CurrencyValidator(v); err != nil {
			return &ValidationError{Name: "currency", err: fmt.Errorf(`ent: validator failed for field "SecurityPosition.currency": %w`, err)}
		}
	}
	if v, ok := spc.mutation.StlCurrency(); ok {
		if err := securityposition.StlCurrencyValidator(v); err != nil {
			return &ValidationError{Name: "stl_currency", err: fmt.Errorf(`ent: validator failed for field "SecurityPosition.stl_currency": %w`, err)}
		}
	}
	if v, ok := spc.mutation.UniqueTag(); ok {
		if err := securityposition.UniqueTagValidator(v); err != nil {
			return &ValidationError{Name: "unique_tag", err: fmt.Errorf(`ent: validator failed for field "SecurityPosition.unique_tag": %w`, err)}
		}
	}
	if _, ok := spc.mutation.OrgID(); !ok {
		return &ValidationError{Name: "org_id", err: errors.New(`ent: missing required field "SecurityPosition.org_id"`)}
	}
	return nil
}

func (spc *SecurityPositionCreate) sqlSave(ctx context.Context) (*SecurityPosition, error) {
	_node, _spec := spc.createSpec()
	if err := sqlgraph.CreateNode(ctx, spc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	return _node, nil
}

func (spc *SecurityPositionCreate) createSpec() (*SecurityPosition, *sqlgraph.CreateSpec) {
	var (
		_node = &SecurityPosition{config: spc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: securityposition.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: securityposition.FieldID,
			},
		}
	)
	if id, ok := spc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := spc.mutation.ParentID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: securityposition.FieldParentID,
		})
		_node.ParentID = value
	}
	if value, ok := spc.mutation.AccountID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: securityposition.FieldAccountID,
		})
		_node.AccountID = value
	}
	if value, ok := spc.mutation.SecurityAccountID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: securityposition.FieldSecurityAccountID,
		})
		_node.SecurityAccountID = value
	}
	if value, ok := spc.mutation.PositionType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: securityposition.FieldPositionType,
		})
		_node.PositionType = value
	}
	if value, ok := spc.mutation.CdDirection(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: securityposition.FieldCdDirection,
		})
		_node.CdDirection = &value
	}
	if value, ok := spc.mutation.ProjectID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: securityposition.FieldProjectID,
		})
		_node.ProjectID = value
	}
	if value, ok := spc.mutation.ProductID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: securityposition.FieldProductID,
		})
		_node.ProductID = value
	}
	if value, ok := spc.mutation.MaterialID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: securityposition.FieldMaterialID,
		})
		_node.MaterialID = value
	}
	if value, ok := spc.mutation.MaterialNo(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: securityposition.FieldMaterialNo,
		})
		_node.MaterialNo = &value
	}
	if value, ok := spc.mutation.Multiplier(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: securityposition.FieldMultiplier,
		})
		_node.Multiplier = &value
	}
	if value, ok := spc.mutation.Balance(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: securityposition.FieldBalance,
		})
		_node.Balance = value
	}
	if value, ok := spc.mutation.Available(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: securityposition.FieldAvailable,
		})
		_node.Available = value
	}
	if value, ok := spc.mutation.Freeze(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: securityposition.FieldFreeze,
		})
		_node.Freeze = value
	}
	if value, ok := spc.mutation.Afloat(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: securityposition.FieldAfloat,
		})
		_node.Afloat = value
	}
	if value, ok := spc.mutation.Unit(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: securityposition.FieldUnit,
		})
		_node.Unit = &value
	}
	if value, ok := spc.mutation.Spec(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: securityposition.FieldSpec,
		})
		_node.Spec = &value
	}
	if value, ok := spc.mutation.Currency(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: securityposition.FieldCurrency,
		})
		_node.Currency = &value
	}
	if value, ok := spc.mutation.Price(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: securityposition.FieldPrice,
		})
		_node.Price = value
	}
	if value, ok := spc.mutation.Amount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: securityposition.FieldAmount,
		})
		_node.Amount = &value
	}
	if value, ok := spc.mutation.CostAmount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: securityposition.FieldCostAmount,
		})
		_node.CostAmount = &value
	}
	if value, ok := spc.mutation.Cost(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: securityposition.FieldCost,
		})
		_node.Cost = &value
	}
	if value, ok := spc.mutation.FxRate(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: securityposition.FieldFxRate,
		})
		_node.FxRate = &value
	}
	if value, ok := spc.mutation.StlCurrency(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: securityposition.FieldStlCurrency,
		})
		_node.StlCurrency = &value
	}
	if value, ok := spc.mutation.StlAmount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: securityposition.FieldStlAmount,
		})
		_node.StlAmount = &value
	}
	if value, ok := spc.mutation.StlCost(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: securityposition.FieldStlCost,
		})
		_node.StlCost = &value
	}
	if value, ok := spc.mutation.StlMargin(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: securityposition.FieldStlMargin,
		})
		_node.StlMargin = &value
	}
	if value, ok := spc.mutation.StlCostAmount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: securityposition.FieldStlCostAmount,
		})
		_node.StlCostAmount = &value
	}
	if value, ok := spc.mutation.StlMarginLv(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: securityposition.FieldStlMarginLv,
		})
		_node.StlMarginLv = value
	}
	if value, ok := spc.mutation.StlValPrice(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: securityposition.FieldStlValPrice,
		})
		_node.StlValPrice = &value
	}
	if value, ok := spc.mutation.InvalidAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: securityposition.FieldInvalidAt,
		})
		_node.InvalidAt = &value
	}
	if value, ok := spc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: securityposition.FieldCreatedAt,
		})
		_node.CreatedAt = &value
	}
	if value, ok := spc.mutation.UniqueTag(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: securityposition.FieldUniqueTag,
		})
		_node.UniqueTag = &value
	}
	if value, ok := spc.mutation.OrgID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: securityposition.FieldOrgID,
		})
		_node.OrgID = value
	}
	return _node, _spec
}

// SecurityPositionCreateBulk is the builder for creating many SecurityPosition entities in bulk.
type SecurityPositionCreateBulk struct {
	config
	builders []*SecurityPositionCreate
}

// Save creates the SecurityPosition entities in the database.
func (spcb *SecurityPositionCreateBulk) Save(ctx context.Context) ([]*SecurityPosition, error) {
	specs := make([]*sqlgraph.CreateSpec, len(spcb.builders))
	nodes := make([]*SecurityPosition, len(spcb.builders))
	mutators := make([]Mutator, len(spcb.builders))
	for i := range spcb.builders {
		func(i int, root context.Context) {
			builder := spcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SecurityPositionMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, spcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, spcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, spcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (spcb *SecurityPositionCreateBulk) SaveX(ctx context.Context) []*SecurityPosition {
	v, err := spcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (spcb *SecurityPositionCreateBulk) Exec(ctx context.Context) error {
	_, err := spcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (spcb *SecurityPositionCreateBulk) ExecX(ctx context.Context) {
	if err := spcb.Exec(ctx); err != nil {
		panic(err)
	}
}