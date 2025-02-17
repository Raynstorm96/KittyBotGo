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

var LikedTracks = newLikedTracksTable("public", "liked_tracks", "")

type likedTracksTable struct {
	postgres.Table

	//Columns
	UserID  postgres.ColumnString
	Query   postgres.ColumnString
	Title   postgres.ColumnString
	LikedAt postgres.ColumnTimestamp

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type LikedTracksTable struct {
	likedTracksTable

	EXCLUDED likedTracksTable
}

// AS creates new LikedTracksTable with assigned alias
func (a LikedTracksTable) AS(alias string) *LikedTracksTable {
	return newLikedTracksTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new LikedTracksTable with assigned schema name
func (a LikedTracksTable) FromSchema(schemaName string) *LikedTracksTable {
	return newLikedTracksTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new LikedTracksTable with assigned table prefix
func (a LikedTracksTable) WithPrefix(prefix string) *LikedTracksTable {
	return newLikedTracksTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new LikedTracksTable with assigned table suffix
func (a LikedTracksTable) WithSuffix(suffix string) *LikedTracksTable {
	return newLikedTracksTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newLikedTracksTable(schemaName, tableName, alias string) *LikedTracksTable {
	return &LikedTracksTable{
		likedTracksTable: newLikedTracksTableImpl(schemaName, tableName, alias),
		EXCLUDED:         newLikedTracksTableImpl("", "excluded", ""),
	}
}

func newLikedTracksTableImpl(schemaName, tableName, alias string) likedTracksTable {
	var (
		UserIDColumn   = postgres.StringColumn("user_id")
		QueryColumn    = postgres.StringColumn("query")
		TitleColumn    = postgres.StringColumn("title")
		LikedAtColumn  = postgres.TimestampColumn("liked_at")
		allColumns     = postgres.ColumnList{UserIDColumn, QueryColumn, TitleColumn, LikedAtColumn}
		mutableColumns = postgres.ColumnList{QueryColumn, LikedAtColumn}
	)

	return likedTracksTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		UserID:  UserIDColumn,
		Query:   QueryColumn,
		Title:   TitleColumn,
		LikedAt: LikedAtColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
