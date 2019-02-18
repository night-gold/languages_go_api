package main

type Code struct {
	Code string `json:"code"`
}

type LangName struct {
	Fk     string `json:"fk"`
	FkName string `json:"fkname"`
	Name   string `json:"name"`
}

type LangNames []LangName

type Language struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type Languages []Language
