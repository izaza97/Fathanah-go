package models

type AH struct {
	Id      string `form:"id" json:"id"`
	Arab    string `form:"arab" json:"arab"`
	Latin   string `form:"latin" json:"latin"`
	Meaning string `form:"arti" json:"arti"`
}

type R_ah struct {
	Data []AH
}
