// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/bug/ent/predicate"
	"entgo.io/bug/ent/securityposition"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SecurityPositionDelete is the builder for deleting a SecurityPosition entity.
type SecurityPositionDelete struct {
	config
	hooks    []Hook
	mutation *SecurityPositionMutation
}

// Where appends a list predicates to the SecurityPositionDelete builder.
func (spd *SecurityPositionDelete) Where(ps ...predicate.SecurityPosition) *SecurityPositionDelete {
	spd.mutation.Where(ps...)
	return spd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (spd *SecurityPositionDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(spd.hooks) == 0 {
		affected, err = spd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SecurityPositionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			spd.mutation = mutation
			affected, err = spd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(spd.hooks) - 1; i >= 0; i-- {
			if spd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = spd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, spd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (spd *SecurityPositionDelete) ExecX(ctx context.Context) int {
	n, err := spd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (spd *SecurityPositionDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: securityposition.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: securityposition.FieldID,
			},
		},
	}
	if ps := spd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, spd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// SecurityPositionDeleteOne is the builder for deleting a single SecurityPosition entity.
type SecurityPositionDeleteOne struct {
	spd *SecurityPositionDelete
}

// Exec executes the deletion query.
func (spdo *SecurityPositionDeleteOne) Exec(ctx context.Context) error {
	n, err := spdo.spd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{securityposition.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (spdo *SecurityPositionDeleteOne) ExecX(ctx context.Context) {
	spdo.spd.ExecX(ctx)
}
