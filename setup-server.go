package main

import (
	"log"

	db "github.com/Skjaldbaka17/setup-quotel-db/database"
)

func main() {
	connection, err := db.InitializeDBConnection()
	if err != nil {
		log.Fatalf("got error %s", err)
	}
	// connection.InsertAuthors("english")
	// connection.InsertAuthors("icelandic")
	connection.InsertTopics("english")
}
