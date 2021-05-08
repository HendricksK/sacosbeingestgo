package sacos

import (
	"log"
	"os"
	"database/sql"

	_ "github.com/lib/pq"
)

type Ingest_Sacos struct {
	Id 			int `json:"id"`
	Data 		string `json:"data"`
	Source 		string `json:"source"`
	Datetime 	string 	`json:"datetime"`
	// images
}

func Ingest() {
	return true
}