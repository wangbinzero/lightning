package model

import "lightning/utils"

type Device struct {
	CommonModel
	UserId    int    `json:"user_id"`
	IsUsed    bool   `json:"is_used"`
	Token     string `json:"token" gorm:"type:varchar(64)"`
	PublicKey string `json:"-"`
}

func (d *Device) InitializeToken() {
	d.Token = utils.RandStringRunes(64)
}
