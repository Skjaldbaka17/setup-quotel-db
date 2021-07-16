package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	db "github.com/Skjaldbaka17/setup-quotel-db/database"
)

func readResponse(command string) string {
	fmt.Println(command)
	var response string
	reader := bufio.NewReader(os.Stdin)
	response, _ = reader.ReadString('\n')
	response = strings.Trim(response, "\n")
	if response != "" {
		return response
	} else {

		return readResponse(command)
	}
}

func main() {

	resp := readResponse("This will delete all data in the current DB. Sure you want to start the setup? (y/n)")

	if strings.ToLower(resp) != string('y') {
		log.Println("You have decided not to setup db...")
		return
	}

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

	var author db.Author
	var count int64
	connection.DB.Where("name = ?", "Óþekktur höfundur").First(&author)
	connection.DB.Model(&db.Quote{}).Where("author_id = ?", author.ID).Count(&count)
	connection.DB.Exec("select count(*) from authors where id = ?", author.ID).Find(&count)
	err = connection.DB.Model(&author).Update("nr_of_icelandic_quotes", count).Error
	if err != nil {
		log.Fatalf("got error %s", err)
	}

}
