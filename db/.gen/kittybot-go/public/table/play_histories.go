//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/postgres"
)

var PlayHistories = newPlayHistoriesTable("public", "play_histories", "")

type playHistoriesTable struct {
	postgres.Table

	//Columns
	UserID     postgres.ColumnString
	Query      postgres.ColumnString
	Title      postgres.ColumnString
	LastUsedAt postgres.ColumnTimestamp

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type PlayHistoriesTable struct {
	playHistoriesTable

	EXCLUDED playHistoriesTable
}

// AS creates new PlayHistoriesTable with assigned alias
func (a PlayHistoriesTable) AS(alias string) *PlayHistoriesTable {
	return newPlayHistoriesTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new PlayHistoriesTable with assigned schema name
func (a PlayHistoriesTable) FromSchema(schemaName string) *PlayHistoriesTable {
	return newPlayHistoriesTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new PlayHistoriesTable with assigned table prefix
func (a PlayHistoriesTable) WithPrefix(prefix string) *PlayHistoriesTable {
	return newPlayHistoriesTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new PlayHistoriesTable with assigned table suffix
func (a PlayHistoriesTable) WithSuffix(suffix string) *PlayHistoriesTable {
	return newPlayHistoriesTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newPlayHistoriesTable(schemaName, tableName, alias string) *PlayHistoriesTable {
	return &PlayHistoriesTable{
		playHistoriesTable: newPlayHistoriesTableImpl(schemaName, tableName, alias),
		EXCLUDED:           newPlayHistoriesTableImpl("", "excluded", ""),
	}
}

func newPlayHistoriesTableImpl(schemaName, tableName, alias string) playHistoriesTable {
	var (
		UserIDColumn     = postgres.StringColumn("user_id")
		QueryColumn      = postgres.StringColumn("query")
		TitleColumn      = postgres.StringColumn("title")
		LastUsedAtColumn = postgres.TimestampColumn("last_used_at")
		allColumns       = postgres.ColumnList{UserIDColumn, QueryColumn, TitleColumn, LastUsedAtColumn}
		mutableColumns   = postgres.ColumnList{QueryColumn, LastUsedAtColumn}
	)

	return playHistoriesTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		UserID:     UserIDColumn,
		Query:      QueryColumn,
		Title:      TitleColumn,
		LastUsedAt: LastUsedAtColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
