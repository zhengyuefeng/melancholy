// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebook/ent/dialect/sql/schema"
	"github.com/facebook/ent/schema/field"
)

var (
	// ExitLogsColumns holds the columns for the "exit_logs" table.
	ExitLogsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "user_id", Type: field.TypeInt},
		{Name: "token", Type: field.TypeString, Unique: true},
		{Name: "date", Type: field.TypeTime},
		{Name: "user_exitlogs", Type: field.TypeInt, Nullable: true},
	}
	// ExitLogsTable holds the schema information for the "exit_logs" table.
	ExitLogsTable = &schema.Table{
		Name:       "exit_logs",
		Columns:    ExitLogsColumns,
		PrimaryKey: []*schema.Column{ExitLogsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "exit_logs_users_exitlogs",
				Columns: []*schema.Column{ExitLogsColumns[4]},

				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// FoldersColumns holds the columns for the "folders" table.
	FoldersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "parent", Type: field.TypeInt},
		{Name: "name", Type: field.TypeString},
		{Name: "owner", Type: field.TypeInt},
		{Name: "size", Type: field.TypeInt, Nullable: true},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"0", "10", "20"}, Default: "0"},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "folder_c", Type: field.TypeInt, Nullable: true},
	}
	// FoldersTable holds the schema information for the "folders" table.
	FoldersTable = &schema.Table{
		Name:       "folders",
		Columns:    FoldersColumns,
		PrimaryKey: []*schema.Column{FoldersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "folders_folders_c",
				Columns: []*schema.Column{FoldersColumns[9]},

				RefColumns: []*schema.Column{FoldersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// MfilesColumns holds the columns for the "mfiles" table.
	MfilesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "parent", Type: field.TypeInt, Unique: true},
		{Name: "name", Type: field.TypeString},
		{Name: "author", Type: field.TypeInt},
		{Name: "md5", Type: field.TypeInt, Unique: true},
		{Name: "size", Type: field.TypeInt, Nullable: true},
		{Name: "desc", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"0", "10", "20"}, Default: "0"},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "folder_mfiles", Type: field.TypeInt, Nullable: true},
	}
	// MfilesTable holds the schema information for the "mfiles" table.
	MfilesTable = &schema.Table{
		Name:       "mfiles",
		Columns:    MfilesColumns,
		PrimaryKey: []*schema.Column{MfilesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "mfiles_folders_mfiles",
				Columns: []*schema.Column{MfilesColumns[11]},

				RefColumns: []*schema.Column{FoldersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// RolesColumns holds the columns for the "roles" table.
	RolesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"0", "10", "20"}, Default: "0"},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
	}
	// RolesTable holds the schema information for the "roles" table.
	RolesTable = &schema.Table{
		Name:        "roles",
		Columns:     RolesColumns,
		PrimaryKey:  []*schema.Column{RolesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "username", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
		{Name: "phone", Type: field.TypeInt, Unique: true, Nullable: true},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "state", Type: field.TypeEnum, Enums: []string{"0", "10", "20"}},
		{Name: "salt", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:        "users",
		Columns:     UsersColumns,
		PrimaryKey:  []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// UserRolesColumns holds the columns for the "user_roles" table.
	UserRolesColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeInt},
		{Name: "role_id", Type: field.TypeInt},
	}
	// UserRolesTable holds the schema information for the "user_roles" table.
	UserRolesTable = &schema.Table{
		Name:       "user_roles",
		Columns:    UserRolesColumns,
		PrimaryKey: []*schema.Column{UserRolesColumns[0], UserRolesColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "user_roles_user_id",
				Columns: []*schema.Column{UserRolesColumns[0]},

				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:  "user_roles_role_id",
				Columns: []*schema.Column{UserRolesColumns[1]},

				RefColumns: []*schema.Column{RolesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ExitLogsTable,
		FoldersTable,
		MfilesTable,
		RolesTable,
		UsersTable,
		UserRolesTable,
	}
)

func init() {
	ExitLogsTable.ForeignKeys[0].RefTable = UsersTable
	FoldersTable.ForeignKeys[0].RefTable = FoldersTable
	MfilesTable.ForeignKeys[0].RefTable = FoldersTable
	UserRolesTable.ForeignKeys[0].RefTable = UsersTable
	UserRolesTable.ForeignKeys[1].RefTable = RolesTable
}
