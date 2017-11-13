package main

type legistration struct {
	id            string    `json:"id"`
	enforcer      string    `json:"enforcer"`
	effectiveDate string    `json:"effectiveDate"`
	passDate      string    `json:"passDate"`
	sign          []string  `json:"sign"`
	chapterArray  []chapter `json:"chapterArray"`
}
