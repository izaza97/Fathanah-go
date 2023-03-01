package models

type Ab struct {
	Id   int    `form:"id" json:"id"`
	Name string `form:"nama" json:"nama"`
	Img  string `form:"img" json:"img"`
	Desc string `json:"isi"`
	Path string `json:"path"`
}

type Team struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Instagram string `json:"ig"`
	Img       string `json:"img"`
	Path      string `json:"path"`
}

type Dab struct {
	Data1 []Ab
	Data2 []Ab
	Data3 []Team
}
