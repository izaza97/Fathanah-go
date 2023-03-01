package models

type Art struct {
	Id       int    `json:"id"`
	Time     string `json:"time"`
	Img      string `json:"img"`
	Title    string `json:"title"`
	Category string `json:"category"`
	Desc     string `json:"isi"`
	Path     string `json:"path"`
}

type Artn struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Time  string `json:"time"`
}

type Dart struct {
	Category []Cart
	Data     []Art
}

type Vart struct {
	Category []Cart
	Data     []Art
	Data2    []Artn
}

type Cartl struct {
	PickC    []Cart
	Category []Cart
	Data     []Art
}

type Cart struct {
	Id       int    `gorm:"primaryKey" json:"id"`
	Category string `json:"category"`
}

type Kdb struct {
	Data []Cart
}
