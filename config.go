package main

type config struct {
	DRIVER   string `toml:"DRIVER"`
	HOST     string `toml:"HOST"`
	PORT     string `toml:"PORT"`
	USER     string `toml:"USER"`
	PASSWORD string `toml:"PASSWORD"`
	DBNAME   string `toml:"DBNAME"`
}
