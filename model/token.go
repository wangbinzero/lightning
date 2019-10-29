package model

import (
	"lightning/utils"
	"time"
)

type Token struct {
	CommonModel
	Token        string    `gorm:"type:varchar(64)" json:"token"`
	UserId       int       `json:"user_id"`
	IsUsed       bool      `json:"is_used"`
	ExpireAt     time.Time `gorm:"default:null" json:"expire_at"`
	LastVerifyAt time.Time `gorm:"detault:null" json:"last_verify_at"`
}

func (token *Token) InitializeLoginToken() {
	token.Token = utils.RandStringRunes(64)
	secondsEastOfUtc := int((8 * time.Hour).Seconds())
	beijing := time.FixedZone("Beijing Time", secondsEastOfUtc)
	now := time.Now()
	token.ExpireAt = time.Date(now.Year(), now.Month(), now.Day()+7, now.Hour(), now.Minute(), now.Second(), 0, beijing)

}
