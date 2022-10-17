package models

import (
	"beego_notes/lib"
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

type User struct {
	ID        uint64 `orm:"column(id)"`
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) TableName() string {
	return "users"
}

func UserCheckAvailability(email string) bool {
	fmt.Println(email)
	o := orm.NewOrm()
	user := User{}
	o.QueryTable(new(User)).Filter("username", email).RelatedSel().One(&user)
	return (user.ID == 0)
}

func UserCreate(email string, password string) *User {
	hshPasswd, _ := lib.HashPassword(password)
	o := orm.NewOrm()
	currTime := time.Now()
	user := User{Username: email, Password: hshPasswd, CreatedAt: currTime, UpdatedAt: currTime}
	id, err := o.Insert(&user)
	fmt.Println(id)
	fmt.Println(err)
	return &user
}

func UserCheck(email string, password string) *User {
	o := orm.NewOrm()
	user := User{}
	o.QueryTable(new(User)).Filter("username", email).RelatedSel().One(&user)
	if user.ID == 0 {
		return nil
	}
	match := lib.CheckPasswordHash(password, user.Password)
	if match {
		return &user
	} else {
		return nil
	}
}

func UserFind(id uint64) *User {
	o := orm.NewOrm()
	user := User{ID: id}
	err := o.Read(&user)
	if err != nil {
		return nil
	} else {
		return &user
	}
}
