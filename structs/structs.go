package structs

import "time"

type Song struct {
	Id       int       `json:"id"`
	Name     string    `json:"name"`
	Duration time.Time `json:"duration"`
	Genre    string    `json:"genre"`
	Artist   *Artist   `json:"artist"`
}

type Artist struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Label     string `json:"label"`
}
