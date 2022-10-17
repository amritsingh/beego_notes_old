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
	UserID    uint64 `orm:"column(user_id)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func (u *Note) TableName() string {
	return "notes"
}

func NotesGetAll(user *User) *[]*Note {
	o := orm.NewOrm()
	var notes []*Note
	// Documentation: https://github.com/beego/beedoc/blob/master/en-US/mvc/model/query.md
	_, err := o.QueryTable(new(Note)).Filter("user_id", user.ID).Filter("deleted_at__isnull", true).OrderBy("-updated_at").All(&notes)
	if err != nil {
		fmt.Println(err)
		return nil
	} else {
		return &notes
	}
}

func NotesCreate(user *User, name string, content string) {
	o := orm.NewOrm()
	currTime := time.Now()
	note := Note{Name: name, Content: content, UserID: user.ID, CreatedAt: currTime, UpdatedAt: currTime}
	id, err := o.Insert(&note)
	fmt.Println(id)
	fmt.Println(err)
}

func NotesFind(user *User, id uint64) *Note {
	o := orm.NewOrm()
	note := Note{ID: id, UserID: user.ID}
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

func (note *Note) MarkDelete() {
	o := orm.NewOrm()
	note.DeletedAt = time.Now()
	_, err := o.Update(note)
	if err != nil {
		fmt.Println(err)
	}
}
