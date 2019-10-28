package model

type Currency struct {
	CommonModel
	Key         string `json:"key"`
	Code        string `json:"code"`
	Symbol      string `json:"-"`
	Coin        bool   `json:"coin"`
	Precision   int    `json:"precision"`
	Erc20       bool   `json:"erc20"`
	Erc23       bool   `json:"erc23"`
	Visible     bool   `json:"visible"`
	Tradable    bool   `json:"tradable"`
	Depositable bool   `json:"depositable"`
}
