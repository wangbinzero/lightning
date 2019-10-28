package model

import "time"

type CommonModel struct {
	Id       int       `json:"id" gorm:"primary_key"`
	CreatAt  time.Time `json:"-"`
	UpdateAt time.Time `json:"-"`
}
