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

type Back struct {
	Id   int    `json:"id"`
	Name string `json:"surah"`
}

type Next struct {
	Id   int    `json:"id"`
	Name string `json:"surah"`
}

type Dqrn struct {
	Pickedsurah []Picked
	Surah       []Surah
	Data        []Qrn
	Back        []Back
	Next        []Next
}

type Qrns struct {
	Id      int    `json:"id"`
	Name    string `json:"surah"`
	Meaning string `json:"arti"`
}

type Dqs struct {
	Data []Qrns
}
