package models

type Vbrt struct {
	Id       int    `json:"id"`
	Img      string `json:"img"`
	Title    string `json:"judul"`
	Sinopsis string `json:"sinopsis"`
	Category string `json:"category"`
	Desc     string `json:"isi"`
	Time     string `json:"time"`
	Path     string `json:"path"`
}

type Feature struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Img  string `json:"img"`
	Path string `json:"path"`
	Url  string `json:"url"`
}

type Home struct {
	Icon  []Icon
	Logo  []Head
	Data1 []Slide
	Data2 []Feature
	Data3 []Vbrt
}

type Slide struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Img  string `json:"img"`
	Path string `json:"path"`
}

type Head struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Img  string `json:"img"`
	Path string `json:"path"`
}

type Nuser struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Img  string `json:"img"`
	Path string `json:"path"`
}

type Icon struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Img  string `json:"img"`
	Path string `json:"path"`
}

type Hd struct {
	Icon    []Icon
	Logo    []Head
	Navuser []Nuser
}
