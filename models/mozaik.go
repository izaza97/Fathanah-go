package models

type Mozaikdata struct {
	Id       int    `json:"id"`
	Time     string `json:"time"`
	Img      string `json:"img"`
	Title    string `json:"title"`
	Category string `json:"category"`
	Desc     string `json:"isi"`
	Path     string `json:"path"`
}

type Dmozaik struct {
	Data []Mozaikdata
}

type Mozaikl struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Time  string `json:"time"`
}

type Vmozaik struct {
	Data     []Mozaikdata
	Sidedata []Mozaikl
}
