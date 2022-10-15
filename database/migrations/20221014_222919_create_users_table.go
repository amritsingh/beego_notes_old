package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreateUsersTable_20221014_222919 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateUsersTable_20221014_222919{}
	m.Created = "20221014_222919"

	migration.Register("CreateUsersTable_20221014_222919", m)
}

// Run the migrations
func (m *CreateUsersTable_20221014_222919) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE users (id bigint unsigned NOT NULL AUTO_INCREMENT, username longtext, password varchar(255) DEFAULT NULL, created_at datetime(3) DEFAULT NULL, updated_at datetime(3) DEFAULT NULL, PRIMARY KEY (id))")
}

// Reverse the migrations
func (m *CreateUsersTable_20221014_222919) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE users")
}
