// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/HaHadaxigua/melancholy/ent/folder"
	"github.com/HaHadaxigua/melancholy/ent/mfile"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// MFileCreate is the builder for creating a MFile entity.
type MFileCreate struct {
	config
	mutation *MFileMutation
	hooks    []Hook
}

// SetParent sets the parent field.
func (mc *MFileCreate) SetParent(i int) *MFileCreate {
	mc.mutation.SetParent(i)
	return mc
}

// SetName sets the name field.
func (mc *MFileCreate) SetName(s string) *MFileCreate {
	mc.mutation.SetName(s)
	return mc
}

// SetAuthor sets the author field.
func (mc *MFileCreate) SetAuthor(i int) *MFileCreate {
	mc.mutation.SetAuthor(i)
	return mc
}

// SetMd5 sets the md5 field.
func (mc *MFileCreate) SetMd5(i int) *MFileCreate {
	mc.mutation.SetMd5(i)
	return mc
}

// SetSize sets the size field.
func (mc *MFileCreate) SetSize(i int) *MFileCreate {
	mc.mutation.SetSize(i)
	return mc
}

// SetNillableSize sets the size field if the given value is not nil.
func (mc *MFileCreate) SetNillableSize(i *int) *MFileCreate {
	if i != nil {
		mc.SetSize(*i)
	}
	return mc
}

// SetMType sets the MType field.
func (mc *MFileCreate) SetMType(mt mfile.MType) *MFileCreate {
	mc.mutation.SetMType(mt)
	return mc
}

// SetNillableMType sets the MType field if the given value is not nil.
func (mc *MFileCreate) SetNillableMType(mt *mfile.MType) *MFileCreate {
	if mt != nil {
		mc.SetMType(*mt)
	}
	return mc
}

// SetDesc sets the desc field.
func (mc *MFileCreate) SetDesc(s string) *MFileCreate {
	mc.mutation.SetDesc(s)
	return mc
}

// SetNillableDesc sets the desc field if the given value is not nil.
func (mc *MFileCreate) SetNillableDesc(s *string) *MFileCreate {
	if s != nil {
		mc.SetDesc(*s)
	}
	return mc
}

// SetStatus sets the status field.
func (mc *MFileCreate) SetStatus(m mfile.Status) *MFileCreate {
	mc.mutation.SetStatus(m)
	return mc
}

// SetNillableStatus sets the status field if the given value is not nil.
func (mc *MFileCreate) SetNillableStatus(m *mfile.Status) *MFileCreate {
	if m != nil {
		mc.SetStatus(*m)
	}
	return mc
}

// SetCreatedAt sets the created_at field.
func (mc *MFileCreate) SetCreatedAt(t time.Time) *MFileCreate {
	mc.mutation.SetCreatedAt(t)
	return mc
}

// SetNillableCreatedAt sets the created_at field if the given value is not nil.
func (mc *MFileCreate) SetNillableCreatedAt(t *time.Time) *MFileCreate {
	if t != nil {
		mc.SetCreatedAt(*t)
	}
	return mc
}

// SetUpdatedAt sets the updated_at field.
func (mc *MFileCreate) SetUpdatedAt(t time.Time) *MFileCreate {
	mc.mutation.SetUpdatedAt(t)
	return mc
}

// SetNillableUpdatedAt sets the updated_at field if the given value is not nil.
func (mc *MFileCreate) SetNillableUpdatedAt(t *time.Time) *MFileCreate {
	if t != nil {
		mc.SetUpdatedAt(*t)
	}
	return mc
}

// SetDeletedAt sets the deleted_at field.
func (mc *MFileCreate) SetDeletedAt(t time.Time) *MFileCreate {
	mc.mutation.SetDeletedAt(t)
	return mc
}

// SetNillableDeletedAt sets the deleted_at field if the given value is not nil.
func (mc *MFileCreate) SetNillableDeletedAt(t *time.Time) *MFileCreate {
	if t != nil {
		mc.SetDeletedAt(*t)
	}
	return mc
}

// SetID sets the id field.
func (mc *MFileCreate) SetID(i int) *MFileCreate {
	mc.mutation.SetID(i)
	return mc
}

// SetFolderID sets the folder edge to Folder by id.
func (mc *MFileCreate) SetFolderID(id int) *MFileCreate {
	mc.mutation.SetFolderID(id)
	return mc
}

// SetNillableFolderID sets the folder edge to Folder by id if the given value is not nil.
func (mc *MFileCreate) SetNillableFolderID(id *int) *MFileCreate {
	if id != nil {
		mc = mc.SetFolderID(*id)
	}
	return mc
}

// SetFolder sets the folder edge to Folder.
func (mc *MFileCreate) SetFolder(f *Folder) *MFileCreate {
	return mc.SetFolderID(f.ID)
}

// Mutation returns the MFileMutation object of the builder.
func (mc *MFileCreate) Mutation() *MFileMutation {
	return mc.mutation
}

// Save creates the MFile in the database.
func (mc *MFileCreate) Save(ctx context.Context) (*MFile, error) {
	var (
		err  error
		node *MFile
	)
	mc.defaults()
	if len(mc.hooks) == 0 {
		if err = mc.check(); err != nil {
			return nil, err
		}
		node, err = mc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MFileMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = mc.check(); err != nil {
				return nil, err
			}
			mc.mutation = mutation
			node, err = mc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(mc.hooks) - 1; i >= 0; i-- {
			mut = mc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (mc *MFileCreate) SaveX(ctx context.Context) *MFile {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (mc *MFileCreate) defaults() {
	if _, ok := mc.mutation.Size(); !ok {
		v := mfile.DefaultSize
		mc.mutation.SetSize(v)
	}
	if _, ok := mc.mutation.MType(); !ok {
		v := mfile.DefaultMType
		mc.mutation.SetMType(v)
	}
	if _, ok := mc.mutation.Desc(); !ok {
		v := mfile.DefaultDesc
		mc.mutation.SetDesc(v)
	}
	if _, ok := mc.mutation.Status(); !ok {
		v := mfile.DefaultStatus
		mc.mutation.SetStatus(v)
	}
	if _, ok := mc.mutation.CreatedAt(); !ok {
		v := mfile.DefaultCreatedAt()
		mc.mutation.SetCreatedAt(v)
	}
	if _, ok := mc.mutation.UpdatedAt(); !ok {
		v := mfile.DefaultUpdatedAt()
		mc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mc *MFileCreate) check() error {
	if _, ok := mc.mutation.Parent(); !ok {
		return &ValidationError{Name: "parent", err: errors.New("ent: missing required field \"parent\"")}
	}
	if _, ok := mc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	if _, ok := mc.mutation.Author(); !ok {
		return &ValidationError{Name: "author", err: errors.New("ent: missing required field \"author\"")}
	}
	if _, ok := mc.mutation.Md5(); !ok {
		return &ValidationError{Name: "md5", err: errors.New("ent: missing required field \"md5\"")}
	}
	if _, ok := mc.mutation.MType(); !ok {
		return &ValidationError{Name: "MType", err: errors.New("ent: missing required field \"MType\"")}
	}
	if v, ok := mc.mutation.MType(); ok {
		if err := mfile.MTypeValidator(v); err != nil {
			return &ValidationError{Name: "MType", err: fmt.Errorf("ent: validator failed for field \"MType\": %w", err)}
		}
	}
	if _, ok := mc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New("ent: missing required field \"status\"")}
	}
	if v, ok := mc.mutation.Status(); ok {
		if err := mfile.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf("ent: validator failed for field \"status\": %w", err)}
		}
	}
	if _, ok := mc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New("ent: missing required field \"created_at\"")}
	}
	if _, ok := mc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New("ent: missing required field \"updated_at\"")}
	}
	return nil
}

func (mc *MFileCreate) sqlSave(ctx context.Context) (*MFile, error) {
	_node, _spec := mc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	if _node.ID == 0 {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	return _node, nil
}

func (mc *MFileCreate) createSpec() (*MFile, *sqlgraph.CreateSpec) {
	var (
		_node = &MFile{config: mc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: mfile.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: mfile.FieldID,
			},
		}
	)
	if id, ok := mc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := mc.mutation.Parent(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: mfile.FieldParent,
		})
		_node.Parent = value
	}
	if value, ok := mc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mfile.FieldName,
		})
		_node.Name = value
	}
	if value, ok := mc.mutation.Author(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: mfile.FieldAuthor,
		})
		_node.Author = value
	}
	if value, ok := mc.mutation.Md5(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: mfile.FieldMd5,
		})
		_node.Md5 = value
	}
	if value, ok := mc.mutation.Size(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: mfile.FieldSize,
		})
		_node.Size = value
	}
	if value, ok := mc.mutation.MType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: mfile.FieldMType,
		})
		_node.MType = value
	}
	if value, ok := mc.mutation.Desc(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mfile.FieldDesc,
		})
		_node.Desc = value
	}
	if value, ok := mc.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: mfile.FieldStatus,
		})
		_node.Status = value
	}
	if value, ok := mc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: mfile.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := mc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: mfile.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := mc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: mfile.FieldDeletedAt,
		})
		_node.DeletedAt = &value
	}
	if nodes := mc.mutation.FolderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   mfile.FolderTable,
			Columns: []string{mfile.FolderColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// MFileCreateBulk is the builder for creating a bulk of MFile entities.
type MFileCreateBulk struct {
	config
	builders []*MFileCreate
}

// Save creates the MFile entities in the database.
func (mcb *MFileCreateBulk) Save(ctx context.Context) ([]*MFile, error) {
	specs := make([]*sqlgraph.CreateSpec, len(mcb.builders))
	nodes := make([]*MFile, len(mcb.builders))
	mutators := make([]Mutator, len(mcb.builders))
	for i := range mcb.builders {
		func(i int, root context.Context) {
			builder := mcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MFileMutation)
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
					_, err = mutators[i+1].Mutate(root, mcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				if nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, mcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (mcb *MFileCreateBulk) SaveX(ctx context.Context) []*MFile {
	v, err := mcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
