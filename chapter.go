package main

type chapter struct {
	header       string    `json:"header"`
	articleArray []article `json:"articleArray"`
	chie         string
}
