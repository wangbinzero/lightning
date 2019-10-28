package model

import "time"

type Token struct {
	CommonModel
	Token        string    `gorm:"type:varchar(64)" json:"token"`
	UserId       int       `json:"user_id"`
	IsUsed       bool      `json:"is_used"`
	ExpireAt     time.Time `gorm:"default:null" json:"expire_at"`
	LastVerifyAt time.Time `gorm:"detault:null" json:"last_verify_at"`
}
