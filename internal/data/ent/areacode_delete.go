// Code generated by ent, DO NOT EDIT.

package ent

import (
	"bovs/internal/data/ent/areacode"
	"bovs/internal/data/ent/predicate"
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AreaCodeDelete is the builder for deleting a AreaCode entity.
type AreaCodeDelete struct {
	config
	hooks    []Hook
	mutation *AreaCodeMutation
}

// Where appends a list predicates to the AreaCodeDelete builder.
func (acd *AreaCodeDelete) Where(ps ...predicate.AreaCode) *AreaCodeDelete {
	acd.mutation.Where(ps...)
	return acd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (acd *AreaCodeDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, acd.sqlExec, acd.mutation, acd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (acd *AreaCodeDelete) ExecX(ctx context.Context) int {
	n, err := acd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (acd *AreaCodeDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(areacode.Table, sqlgraph.NewFieldSpec(areacode.FieldID, field.TypeInt))
	if ps := acd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, acd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	acd.mutation.done = true
	return affected, err
}

// AreaCodeDeleteOne is the builder for deleting a single AreaCode entity.
type AreaCodeDeleteOne struct {
	acd *AreaCodeDelete
}

// Where appends a list predicates to the AreaCodeDelete builder.
func (acdo *AreaCodeDeleteOne) Where(ps ...predicate.AreaCode) *AreaCodeDeleteOne {
	acdo.acd.mutation.Where(ps...)
	return acdo
}

// Exec executes the deletion query.
func (acdo *AreaCodeDeleteOne) Exec(ctx context.Context) error {
	n, err := acdo.acd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{areacode.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (acdo *AreaCodeDeleteOne) ExecX(ctx context.Context) {
	if err := acdo.Exec(ctx); err != nil {
		panic(err)
	}
}