// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"github.com/HaHadaxigua/melancholy/ent/folder"
	"github.com/HaHadaxigua/melancholy/ent/mfile"
	"github.com/HaHadaxigua/melancholy/ent/predicate"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// FolderUpdate is the builder for updating Folder entities.
type FolderUpdate struct {
	config
	hooks    []Hook
	mutation *FolderMutation
}

// Where adds a new predicate for the builder.
func (fu *FolderUpdate) Where(ps ...predicate.Folder) *FolderUpdate {
	fu.mutation.predicates = append(fu.mutation.predicates, ps...)
	return fu
}

// SetParent sets the parent field.
func (fu *FolderUpdate) SetParent(i int) *FolderUpdate {
	fu.mutation.ResetParent()
	fu.mutation.SetParent(i)
	return fu
}

// AddParent adds i to parent.
func (fu *FolderUpdate) AddParent(i int) *FolderUpdate {
	fu.mutation.AddParent(i)
	return fu
}

// SetName sets the name field.
func (fu *FolderUpdate) SetName(s string) *FolderUpdate {
	fu.mutation.SetName(s)
	return fu
}

// SetOwner sets the owner field.
func (fu *FolderUpdate) SetOwner(i int) *FolderUpdate {
	fu.mutation.ResetOwner()
	fu.mutation.SetOwner(i)
	return fu
}

// AddOwner adds i to owner.
func (fu *FolderUpdate) AddOwner(i int) *FolderUpdate {
	fu.mutation.AddOwner(i)
	return fu
}

// SetSize sets the size field.
func (fu *FolderUpdate) SetSize(i int) *FolderUpdate {
	fu.mutation.ResetSize()
	fu.mutation.SetSize(i)
	return fu
}

// SetNillableSize sets the size field if the given value is not nil.
func (fu *FolderUpdate) SetNillableSize(i *int) *FolderUpdate {
	if i != nil {
		fu.SetSize(*i)
	}
	return fu
}

// AddSize adds i to size.
func (fu *FolderUpdate) AddSize(i int) *FolderUpdate {
	fu.mutation.AddSize(i)
	return fu
}

// ClearSize clears the value of size.
func (fu *FolderUpdate) ClearSize() *FolderUpdate {
	fu.mutation.ClearSize()
	return fu
}

// SetStatus sets the status field.
func (fu *FolderUpdate) SetStatus(f folder.Status) *FolderUpdate {
	fu.mutation.SetStatus(f)
	return fu
}

// SetNillableStatus sets the status field if the given value is not nil.
func (fu *FolderUpdate) SetNillableStatus(f *folder.Status) *FolderUpdate {
	if f != nil {
		fu.SetStatus(*f)
	}
	return fu
}

// SetUpdatedAt sets the updated_at field.
func (fu *FolderUpdate) SetUpdatedAt(t time.Time) *FolderUpdate {
	fu.mutation.SetUpdatedAt(t)
	return fu
}

// SetDeletedAt sets the deleted_at field.
func (fu *FolderUpdate) SetDeletedAt(t time.Time) *FolderUpdate {
	fu.mutation.SetDeletedAt(t)
	return fu
}

// SetNillableDeletedAt sets the deleted_at field if the given value is not nil.
func (fu *FolderUpdate) SetNillableDeletedAt(t *time.Time) *FolderUpdate {
	if t != nil {
		fu.SetDeletedAt(*t)
	}
	return fu
}

// ClearDeletedAt clears the value of deleted_at.
func (fu *FolderUpdate) ClearDeletedAt() *FolderUpdate {
	fu.mutation.ClearDeletedAt()
	return fu
}

// AddMfileIDs adds the mfiles edge to MFile by ids.
func (fu *FolderUpdate) AddMfileIDs(ids ...int) *FolderUpdate {
	fu.mutation.AddMfileIDs(ids...)
	return fu
}

// AddMfiles adds the mfiles edges to MFile.
func (fu *FolderUpdate) AddMfiles(m ...*MFile) *FolderUpdate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return fu.AddMfileIDs(ids...)
}

// SetPID sets the p edge to Folder by id.
func (fu *FolderUpdate) SetPID(id int) *FolderUpdate {
	fu.mutation.SetPID(id)
	return fu
}

// SetNillablePID sets the p edge to Folder by id if the given value is not nil.
func (fu *FolderUpdate) SetNillablePID(id *int) *FolderUpdate {
	if id != nil {
		fu = fu.SetPID(*id)
	}
	return fu
}

// SetP sets the p edge to Folder.
func (fu *FolderUpdate) SetP(f *Folder) *FolderUpdate {
	return fu.SetPID(f.ID)
}

// AddCIDs adds the c edge to Folder by ids.
func (fu *FolderUpdate) AddCIDs(ids ...int) *FolderUpdate {
	fu.mutation.AddCIDs(ids...)
	return fu
}

// AddC adds the c edges to Folder.
func (fu *FolderUpdate) AddC(f ...*Folder) *FolderUpdate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fu.AddCIDs(ids...)
}

// Mutation returns the FolderMutation object of the builder.
func (fu *FolderUpdate) Mutation() *FolderMutation {
	return fu.mutation
}

// ClearMfiles clears all "mfiles" edges to type MFile.
func (fu *FolderUpdate) ClearMfiles() *FolderUpdate {
	fu.mutation.ClearMfiles()
	return fu
}

// RemoveMfileIDs removes the mfiles edge to MFile by ids.
func (fu *FolderUpdate) RemoveMfileIDs(ids ...int) *FolderUpdate {
	fu.mutation.RemoveMfileIDs(ids...)
	return fu
}

// RemoveMfiles removes mfiles edges to MFile.
func (fu *FolderUpdate) RemoveMfiles(m ...*MFile) *FolderUpdate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return fu.RemoveMfileIDs(ids...)
}

// ClearP clears the "p" edge to type Folder.
func (fu *FolderUpdate) ClearP() *FolderUpdate {
	fu.mutation.ClearP()
	return fu
}

// ClearC clears all "c" edges to type Folder.
func (fu *FolderUpdate) ClearC() *FolderUpdate {
	fu.mutation.ClearC()
	return fu
}

// RemoveCIDs removes the c edge to Folder by ids.
func (fu *FolderUpdate) RemoveCIDs(ids ...int) *FolderUpdate {
	fu.mutation.RemoveCIDs(ids...)
	return fu
}

// RemoveC removes c edges to Folder.
func (fu *FolderUpdate) RemoveC(f ...*Folder) *FolderUpdate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fu.RemoveCIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (fu *FolderUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	fu.defaults()
	if len(fu.hooks) == 0 {
		if err = fu.check(); err != nil {
			return 0, err
		}
		affected, err = fu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FolderMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = fu.check(); err != nil {
				return 0, err
			}
			fu.mutation = mutation
			affected, err = fu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(fu.hooks) - 1; i >= 0; i-- {
			mut = fu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (fu *FolderUpdate) SaveX(ctx context.Context) int {
	affected, err := fu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (fu *FolderUpdate) Exec(ctx context.Context) error {
	_, err := fu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fu *FolderUpdate) ExecX(ctx context.Context) {
	if err := fu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fu *FolderUpdate) defaults() {
	if _, ok := fu.mutation.UpdatedAt(); !ok {
		v := folder.UpdateDefaultUpdatedAt()
		fu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fu *FolderUpdate) check() error {
	if v, ok := fu.mutation.Status(); ok {
		if err := folder.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf("ent: validator failed for field \"status\": %w", err)}
		}
	}
	return nil
}

func (fu *FolderUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   folder.Table,
			Columns: folder.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: folder.FieldID,
			},
		},
	}
	if ps := fu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fu.mutation.Parent(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: folder.FieldParent,
		})
	}
	if value, ok := fu.mutation.AddedParent(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: folder.FieldParent,
		})
	}
	if value, ok := fu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: folder.FieldName,
		})
	}
	if value, ok := fu.mutation.Owner(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: folder.FieldOwner,
		})
	}
	if value, ok := fu.mutation.AddedOwner(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: folder.FieldOwner,
		})
	}
	if value, ok := fu.mutation.Size(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: folder.FieldSize,
		})
	}
	if value, ok := fu.mutation.AddedSize(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: folder.FieldSize,
		})
	}
	if fu.mutation.SizeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: folder.FieldSize,
		})
	}
	if value, ok := fu.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: folder.FieldStatus,
		})
	}
	if value, ok := fu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: folder.FieldUpdatedAt,
		})
	}
	if value, ok := fu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: folder.FieldDeletedAt,
		})
	}
	if fu.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: folder.FieldDeletedAt,
		})
	}
	if fu.mutation.MfilesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   folder.MfilesTable,
			Columns: []string{folder.MfilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: mfile.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fu.mutation.RemovedMfilesIDs(); len(nodes) > 0 && !fu.mutation.MfilesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   folder.MfilesTable,
			Columns: []string{folder.MfilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: mfile.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fu.mutation.MfilesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   folder.MfilesTable,
			Columns: []string{folder.MfilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: mfile.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if fu.mutation.PCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   folder.PTable,
			Columns: []string{folder.PColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: folder.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fu.mutation.PIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   folder.PTable,
			Columns: []string{folder.PColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: folder.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if fu.mutation.CCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   folder.CTable,
			Columns: []string{folder.CColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: folder.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fu.mutation.RemovedCIDs(); len(nodes) > 0 && !fu.mutation.CCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   folder.CTable,
			Columns: []string{folder.CColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: folder.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fu.mutation.CIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   folder.CTable,
			Columns: []string{folder.CColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: folder.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, fu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{folder.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// FolderUpdateOne is the builder for updating a single Folder entity.
type FolderUpdateOne struct {
	config
	hooks    []Hook
	mutation *FolderMutation
}

// SetParent sets the parent field.
func (fuo *FolderUpdateOne) SetParent(i int) *FolderUpdateOne {
	fuo.mutation.ResetParent()
	fuo.mutation.SetParent(i)
	return fuo
}

// AddParent adds i to parent.
func (fuo *FolderUpdateOne) AddParent(i int) *FolderUpdateOne {
	fuo.mutation.AddParent(i)
	return fuo
}

// SetName sets the name field.
func (fuo *FolderUpdateOne) SetName(s string) *FolderUpdateOne {
	fuo.mutation.SetName(s)
	return fuo
}

// SetOwner sets the owner field.
func (fuo *FolderUpdateOne) SetOwner(i int) *FolderUpdateOne {
	fuo.mutation.ResetOwner()
	fuo.mutation.SetOwner(i)
	return fuo
}

// AddOwner adds i to owner.
func (fuo *FolderUpdateOne) AddOwner(i int) *FolderUpdateOne {
	fuo.mutation.AddOwner(i)
	return fuo
}

// SetSize sets the size field.
func (fuo *FolderUpdateOne) SetSize(i int) *FolderUpdateOne {
	fuo.mutation.ResetSize()
	fuo.mutation.SetSize(i)
	return fuo
}

// SetNillableSize sets the size field if the given value is not nil.
func (fuo *FolderUpdateOne) SetNillableSize(i *int) *FolderUpdateOne {
	if i != nil {
		fuo.SetSize(*i)
	}
	return fuo
}

// AddSize adds i to size.
func (fuo *FolderUpdateOne) AddSize(i int) *FolderUpdateOne {
	fuo.mutation.AddSize(i)
	return fuo
}

// ClearSize clears the value of size.
func (fuo *FolderUpdateOne) ClearSize() *FolderUpdateOne {
	fuo.mutation.ClearSize()
	return fuo
}

// SetStatus sets the status field.
func (fuo *FolderUpdateOne) SetStatus(f folder.Status) *FolderUpdateOne {
	fuo.mutation.SetStatus(f)
	return fuo
}

// SetNillableStatus sets the status field if the given value is not nil.
func (fuo *FolderUpdateOne) SetNillableStatus(f *folder.Status) *FolderUpdateOne {
	if f != nil {
		fuo.SetStatus(*f)
	}
	return fuo
}

// SetUpdatedAt sets the updated_at field.
func (fuo *FolderUpdateOne) SetUpdatedAt(t time.Time) *FolderUpdateOne {
	fuo.mutation.SetUpdatedAt(t)
	return fuo
}

// SetDeletedAt sets the deleted_at field.
func (fuo *FolderUpdateOne) SetDeletedAt(t time.Time) *FolderUpdateOne {
	fuo.mutation.SetDeletedAt(t)
	return fuo
}

// SetNillableDeletedAt sets the deleted_at field if the given value is not nil.
func (fuo *FolderUpdateOne) SetNillableDeletedAt(t *time.Time) *FolderUpdateOne {
	if t != nil {
		fuo.SetDeletedAt(*t)
	}
	return fuo
}

// ClearDeletedAt clears the value of deleted_at.
func (fuo *FolderUpdateOne) ClearDeletedAt() *FolderUpdateOne {
	fuo.mutation.ClearDeletedAt()
	return fuo
}

// AddMfileIDs adds the mfiles edge to MFile by ids.
func (fuo *FolderUpdateOne) AddMfileIDs(ids ...int) *FolderUpdateOne {
	fuo.mutation.AddMfileIDs(ids...)
	return fuo
}

// AddMfiles adds the mfiles edges to MFile.
func (fuo *FolderUpdateOne) AddMfiles(m ...*MFile) *FolderUpdateOne {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return fuo.AddMfileIDs(ids...)
}

// SetPID sets the p edge to Folder by id.
func (fuo *FolderUpdateOne) SetPID(id int) *FolderUpdateOne {
	fuo.mutation.SetPID(id)
	return fuo
}

// SetNillablePID sets the p edge to Folder by id if the given value is not nil.
func (fuo *FolderUpdateOne) SetNillablePID(id *int) *FolderUpdateOne {
	if id != nil {
		fuo = fuo.SetPID(*id)
	}
	return fuo
}

// SetP sets the p edge to Folder.
func (fuo *FolderUpdateOne) SetP(f *Folder) *FolderUpdateOne {
	return fuo.SetPID(f.ID)
}

// AddCIDs adds the c edge to Folder by ids.
func (fuo *FolderUpdateOne) AddCIDs(ids ...int) *FolderUpdateOne {
	fuo.mutation.AddCIDs(ids...)
	return fuo
}

// AddC adds the c edges to Folder.
func (fuo *FolderUpdateOne) AddC(f ...*Folder) *FolderUpdateOne {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fuo.AddCIDs(ids...)
}

// Mutation returns the FolderMutation object of the builder.
func (fuo *FolderUpdateOne) Mutation() *FolderMutation {
	return fuo.mutation
}

// ClearMfiles clears all "mfiles" edges to type MFile.
func (fuo *FolderUpdateOne) ClearMfiles() *FolderUpdateOne {
	fuo.mutation.ClearMfiles()
	return fuo
}

// RemoveMfileIDs removes the mfiles edge to MFile by ids.
func (fuo *FolderUpdateOne) RemoveMfileIDs(ids ...int) *FolderUpdateOne {
	fuo.mutation.RemoveMfileIDs(ids...)
	return fuo
}

// RemoveMfiles removes mfiles edges to MFile.
func (fuo *FolderUpdateOne) RemoveMfiles(m ...*MFile) *FolderUpdateOne {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return fuo.RemoveMfileIDs(ids...)
}

// ClearP clears the "p" edge to type Folder.
func (fuo *FolderUpdateOne) ClearP() *FolderUpdateOne {
	fuo.mutation.ClearP()
	return fuo
}

// ClearC clears all "c" edges to type Folder.
func (fuo *FolderUpdateOne) ClearC() *FolderUpdateOne {
	fuo.mutation.ClearC()
	return fuo
}

// RemoveCIDs removes the c edge to Folder by ids.
func (fuo *FolderUpdateOne) RemoveCIDs(ids ...int) *FolderUpdateOne {
	fuo.mutation.RemoveCIDs(ids...)
	return fuo
}

// RemoveC removes c edges to Folder.
func (fuo *FolderUpdateOne) RemoveC(f ...*Folder) *FolderUpdateOne {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fuo.RemoveCIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (fuo *FolderUpdateOne) Save(ctx context.Context) (*Folder, error) {
	var (
		err  error
		node *Folder
	)
	fuo.defaults()
	if len(fuo.hooks) == 0 {
		if err = fuo.check(); err != nil {
			return nil, err
		}
		node, err = fuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FolderMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = fuo.check(); err != nil {
				return nil, err
			}
			fuo.mutation = mutation
			node, err = fuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(fuo.hooks) - 1; i >= 0; i-- {
			mut = fuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (fuo *FolderUpdateOne) SaveX(ctx context.Context) *Folder {
	node, err := fuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (fuo *FolderUpdateOne) Exec(ctx context.Context) error {
	_, err := fuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fuo *FolderUpdateOne) ExecX(ctx context.Context) {
	if err := fuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fuo *FolderUpdateOne) defaults() {
	if _, ok := fuo.mutation.UpdatedAt(); !ok {
		v := folder.UpdateDefaultUpdatedAt()
		fuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fuo *FolderUpdateOne) check() error {
	if v, ok := fuo.mutation.Status(); ok {
		if err := folder.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf("ent: validator failed for field \"status\": %w", err)}
		}
	}
	return nil
}

func (fuo *FolderUpdateOne) sqlSave(ctx context.Context) (_node *Folder, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   folder.Table,
			Columns: folder.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: folder.FieldID,
			},
		},
	}
	id, ok := fuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Folder.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := fuo.mutation.Parent(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: folder.FieldParent,
		})
	}
	if value, ok := fuo.mutation.AddedParent(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: folder.FieldParent,
		})
	}
	if value, ok := fuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: folder.FieldName,
		})
	}
	if value, ok := fuo.mutation.Owner(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: folder.FieldOwner,
		})
	}
	if value, ok := fuo.mutation.AddedOwner(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: folder.FieldOwner,
		})
	}
	if value, ok := fuo.mutation.Size(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: folder.FieldSize,
		})
	}
	if value, ok := fuo.mutation.AddedSize(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: folder.FieldSize,
		})
	}
	if fuo.mutation.SizeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: folder.FieldSize,
		})
	}
	if value, ok := fuo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: folder.FieldStatus,
		})
	}
	if value, ok := fuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: folder.FieldUpdatedAt,
		})
	}
	if value, ok := fuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: folder.FieldDeletedAt,
		})
	}
	if fuo.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: folder.FieldDeletedAt,
		})
	}
	if fuo.mutation.MfilesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   folder.MfilesTable,
			Columns: []string{folder.MfilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: mfile.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fuo.mutation.RemovedMfilesIDs(); len(nodes) > 0 && !fuo.mutation.MfilesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   folder.MfilesTable,
			Columns: []string{folder.MfilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: mfile.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fuo.mutation.MfilesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   folder.MfilesTable,
			Columns: []string{folder.MfilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: mfile.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if fuo.mutation.PCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   folder.PTable,
			Columns: []string{folder.PColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: folder.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fuo.mutation.PIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   folder.PTable,
			Columns: []string{folder.PColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: folder.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if fuo.mutation.CCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   folder.CTable,
			Columns: []string{folder.CColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: folder.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fuo.mutation.RemovedCIDs(); len(nodes) > 0 && !fuo.mutation.CCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   folder.CTable,
			Columns: []string{folder.CColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: folder.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fuo.mutation.CIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   folder.CTable,
			Columns: []string{folder.CColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: folder.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Folder{config: fuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues()
	if err = sqlgraph.UpdateNode(ctx, fuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{folder.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}