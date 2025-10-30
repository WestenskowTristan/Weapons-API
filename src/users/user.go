package users

import "time"

type User struct {
	Id 					string			`json:"id"`
	CreatedAt 	time.Time		`json:"createdAt"`
	UpdatedAt 	time.Time  	`json:"updatedAt"`
	FirstName  	string			`json:"firstName"`
	LastName  	string			`json:"lastName"`
	Email				string			`json:"email"`
	Password		string			`json:"password"`
}