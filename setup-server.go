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
	log.Println("START INSERTING ENGLISH AUTHORS")
	connection.InsertAuthors("english")
	log.Println("DONE WITH INSERTING ENGLISH AUTHORS")
	log.Println("START INSERTING ICELANDIC AUTHORS")
	connection.InsertAuthors("icelandic")
	log.Println("DONE WITH INSERTING ICELANDIC AUTHORS")
	log.Println("START WRAPPING IT UP")
	connection.WrapItUp() //If this is done first then the rest is much quicker because the quotes.quote column has been indexed!
	log.Println("DONE WRAPPING IT UP")
	log.Println("START INSERTING ENGLISH TOPICS")
	connection.InsertTopics("English")
	log.Println("DONE INSERTING ENGLISH TOPICS")
	log.Println("START INSERTING ENGLISH TOPICS")
	connection.InsertTopics("Icelandic")
	log.Println("DONE INSERTING ICELANDIC TOPICS")

}
