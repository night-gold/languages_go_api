package main

type Code struct {
	Id   int    `json:"id"`
	Code string `json:"code"`
}

type LangName struct {
	Fk     int    `json:"fk"`
	FkName int    `json:"fkname"`
	Name   string `json:"name"`
}

type LangNames []LangName

type Language struct {
	Id   int    `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

type Languages []Language
