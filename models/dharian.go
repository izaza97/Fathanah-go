package models

type Dh struct {
	Id      int    `json:"id"`
	Menu    string `json:"menu"`
	Title   string `json:"judul"`
	Arab    string `json:"arab"`
	Latin   string `json:"latin"`
	Meaning string `json:"arti"`
}

type Dh2 struct {
	Id   int    `json:"id"`
	Menu string `json:"menu"`
}

type Ddh struct {
	Menu []Dh2
	Data []Dh
}

type Kdh struct {
	Id   int    `gorm:"primaryKey" json:"id"`
	Menu string `json:"menu"`
}

type Ddh2 struct {
	Data []Kdh
}
