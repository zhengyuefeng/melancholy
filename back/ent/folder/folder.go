// Code generated by entc, DO NOT EDIT.

package folder

import (
	"fmt"
	"time"
)

const (
	// Label holds the string label denoting the folder type in the database.
	Label = "folder"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldParent holds the string denoting the parent field in the database.
	FieldParent = "parent"
	// FieldPath holds the string denoting the path field in the database.
	FieldPath = "path"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldAuthor holds the string denoting the author field in the database.
	FieldAuthor = "author"
	// FieldSize holds the string denoting the size field in the database.
	FieldSize = "size"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"

	// EdgeMfiles holds the string denoting the mfiles edge name in mutations.
	EdgeMfiles = "mfiles"
	// EdgeP holds the string denoting the p edge name in mutations.
	EdgeP = "p"
	// EdgeC holds the string denoting the c edge name in mutations.
	EdgeC = "c"

	// Table holds the table name of the folder in the database.
	Table = "folders"
	// MfilesTable is the table the holds the mfiles relation/edge.
	MfilesTable = "mfiles"
	// MfilesInverseTable is the table name for the MFile entity.
	// It exists in this package in order to avoid circular dependency with the "mfile" package.
	MfilesInverseTable = "mfiles"
	// MfilesColumn is the table column denoting the mfiles relation/edge.
	MfilesColumn = "folder_mfiles"
	// PTable is the table the holds the p relation/edge.
	PTable = "folders"
	// PColumn is the table column denoting the p relation/edge.
	PColumn = "folder_c"
	// CTable is the table the holds the c relation/edge.
	CTable = "folders"
	// CColumn is the table column denoting the c relation/edge.
	CColumn = "folder_c"
)

// Columns holds all SQL columns for folder fields.
var Columns = []string{
	FieldID,
	FieldParent,
	FieldPath,
	FieldName,
	FieldAuthor,
	FieldSize,
	FieldStatus,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Folder type.
var ForeignKeys = []string{
	"folder_c",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultSize holds the default value on creation for the size field.
	DefaultSize int
	// DefaultCreatedAt holds the default value on creation for the created_at field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the updated_at field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	UpdateDefaultUpdatedAt func() time.Time
)

// Status defines the type for the status enum field.
type Status string

// Status0 is the default Status.
const DefaultStatus = Status0

// Status values.
const (
	Status0  Status = "0"
	Status10 Status = "10"
	Status20 Status = "20"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case Status0, Status10, Status20:
		return nil
	default:
		return fmt.Errorf("folder: invalid enum value for status field: %q", s)
	}
}
