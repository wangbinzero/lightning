package model

import "time"

type ApiToken struct {
	CommonModel
	UserId    int       `json:"user_id"`
	AccessKey string    `gorm:"type:varchar(64)" json:"access_key"`
	SecretKey string    `gorm:"type:varchar(64)"json:"secret_key"`
	Label     string    `gorm:"type:varchar(32)" json:"label"`
	Scopes    string    `gorm:"type:varchar(32)" json:"scopes"`
	ExpireAt  time.Time `gorm:"default:null" json:"expire_at"`
	DeletedAt time.Time `json:"-" gorm:"default:null"`
}
