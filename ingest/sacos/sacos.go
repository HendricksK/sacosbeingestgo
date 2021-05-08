package sacos

import (
	"fmt"
)

type Sacos struct {
	// https://echo.labstack.com/guide/binding/
	// Needed to set binding types for 
	// form data to be accepted
	Id 			int `form:"id" json:"id"`
	Data 		string `form:"data" json:"data"`
	// Source 		string `json:"source"`
	// Datetime 	string 	`json:"datetime"`
	// images
}

func Ingest(data *Sacos) Sacos {
	fmt.Println("123456")

	fmt.Println(data.Id)
	fmt.Println(data.Data)

	// For now just return data sent via postman

	return Sacos{
		Id: data.Id,
		Data: data.Data,
	}
}