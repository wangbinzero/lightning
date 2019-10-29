package model

type User struct {
	CommonModel
	Sn             string
	PasswordDigest string
	Nickname       string
	State          int
	Activated      bool
	Disabled       bool
	ApiDisabled    bool
	Password       string
	Tokens         []Token
}
