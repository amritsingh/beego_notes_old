package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

type Note struct {
	ID        uint64 `orm:"column(id)"`
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
	_, err := o.QueryTable(new(Note)).All(&notes)
	if err != nil {
		return nil
	} else {
		return &notes
	}
}

func NotesCreate(name string, content string) {
	o := orm.NewOrm()
	currTime := time.Now()
	note := Note{Name: name, Content: content, CreatedAt: currTime, UpdatedAt: currTime}
	id, err := o.Insert(&note)
	fmt.Println(id)
	fmt.Println(err)
}

func NotesFind(id uint64) *Note {
	o := orm.NewOrm()
	note := Note{ID: id}
	err := o.Read(&note)
	if err != nil {
		return nil
	} else {
		return &note
	}
}

func (note *Note) Update(name string, content string) {
	o := orm.NewOrm()
	note.Name = name
	note.Content = content
	_, err := o.Update(note)
	if err != nil {
		fmt.Println(err)
	}
}
