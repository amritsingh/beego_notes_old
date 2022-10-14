package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreateNotesTable_20221012_133000 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateNotesTable_20221012_133000{}
	m.Created = "20221012_133000"

	migration.Register("CreateNotesTable_20221012_133000", m)
}

// Run the migrations
func (m *CreateNotesTable_20221012_133000) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE notes (id bigint unsigned NOT NULL AUTO_INCREMENT, name varchar(255) DEFAULT NULL, content text, created_at datetime(3) DEFAULT NULL, updated_at datetime(3) DEFAULT NULL, deleted_at datetime(3) DEFAULT NULL, PRIMARY KEY (id), KEY idx_notes_updated_at (updated_at), KEY idx_notes_deleted_at (deleted_at))")
}

// Reverse the migrations
func (m *CreateNotesTable_20221012_133000) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE notes")
}
