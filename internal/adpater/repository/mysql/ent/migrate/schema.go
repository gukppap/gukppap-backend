// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ShortcutsColumns holds the columns for the "shortcuts" table.
	ShortcutsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "url", Type: field.TypeString},
		{Name: "shortcut", Type: field.TypeString},
		{Name: "exp", Type: field.TypeTime},
	}
	// ShortcutsTable holds the schema information for the "shortcuts" table.
	ShortcutsTable = &schema.Table{
		Name:       "shortcuts",
		Columns:    ShortcutsColumns,
		PrimaryKey: []*schema.Column{ShortcutsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ShortcutsTable,
	}
)

func init() {
}
