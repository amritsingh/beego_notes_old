package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

type Note struct {
	ID        uint64
	Name      string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func (u *Note) TableName() string {
	return "notes"
}

func NotesGetAll() *[]*Note {
	o := orm.NewOrm()
	var notes []*Note
	num, err := o.QueryTable(new(Note)).All(&notes)
	fmt.Println(num)
	if err == nil {
		return nil
	} else {
		return &notes
	}
}
