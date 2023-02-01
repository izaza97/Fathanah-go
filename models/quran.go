package models

type Qrn struct {
	Id      int    `json:"id"`
	Name    string `json:"surah"`
	Arab    string `json:"arab"`
	Latin   string `json:"latin"`
	Meaning string `json:"arti"`
}

type Picked struct {
	Id   int    `json:"id"`
	Name string `json:"surah"`
}

type Surah struct {
	Id   int    `json:"id"`
	Name string `json:"surah"`
}

type Dqrn struct {
	Icon        []Icon
	Logo        []Head
	Pickedsurah []Picked
	Surah       []Surah
	Data        []Qrn
}

type Qrns struct {
	Id      int    `json:"id"`
	Name    string `json:"surah"`
	Meaning string `json:"arti"`
}

type Dqs struct {
	Icon []Icon
	Logo []Head
	Data []Qrns
}
