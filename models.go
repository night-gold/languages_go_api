package main

type Code struct {
	id   int    `json:"id"`
	code string `json:"code"`
}

type Codes []Code

type LangName struct {
	fk     int    `json:"fk"`
	fkName int    `json:"fkname"`
	name   string `json:"name"`
}

type LangNames []LangName

type Language struct {
	id   int    `json:"id"`
	code string `json:"code"`
	name string `json:"name"`
}

type Languages []Language
