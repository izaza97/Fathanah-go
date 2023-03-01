package models

type Diary struct {
	No      int64  `gorm:"primaryKey" json:"no"`
	Time    string `json:"time"`
	User    int64  `json:"user"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

type Diaryscn struct {
	No      int64  `gorm:"primaryKey" json:"no"`
	Time    string `json:"time"`
	Subject string `json:"subject"`
}

type Dd struct {
	Data []Diaryscn
}

type Diarys struct {
	No      int64  `gorm:"primaryKey" json:"no"`
	Time    string `json:"time"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

type Dd2 struct {
	Data []Diarys
}
