package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AddUserIdToNotes_20221014_223209 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AddUserIdToNotes_20221014_223209{}
	m.Created = "20221014_223209"

	migration.Register("AddUserIdToNotes_20221014_223209", m)
}

// Run the migrations
func (m *AddUserIdToNotes_20221014_223209) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("alter table notes add column user_id bigint NOT NULL")
}

// Reverse the migrations
func (m *AddUserIdToNotes_20221014_223209) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("alter table notes drop column user_id")
}
