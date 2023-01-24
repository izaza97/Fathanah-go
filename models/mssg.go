package models

type Mail struct {
	Datetime string `gorm:"datetime" json:"datetime"`
	Id       int64  `json:"id"`
	Subject  string `json:"subjectS"`
	Message  string `json:"message"`
	User     int64  `json:"user"`
}
