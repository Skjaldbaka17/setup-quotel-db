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

	deleteTablesResp := readResponse("Delete tables? (y/n)")
	insertEnglishAuthorsResp := readResponse("Insert englishAuthors? (y/n)")
	insertIcelandicAuthorsResp := readResponse("Insert icelandicAuthors? (y/n)")
	createIndexesResp := readResponse("Run create indexes Queries? (y/n)")
	insertEnglishTopicsResp := readResponse("Insert englishTopics? (y/n)")
	insertIcelandicTopicsResp := readResponse("Insert icelandicTopics? (y/n)")
	createMatViewsResp := readResponse("Run create materialized views Queries? (y/n)")

	var connection *db.Connection
	var err error
	if strings.ToLower(deleteTablesResp) == "y" {
		log.Println("Deleting tables...")
		connection, err = db.InitializeDBConnection(true)
	} else {
		log.Println("NOT Deleting tables...")
		connection, err = db.InitializeDBConnection(false)
	}

	if err != nil {
		log.Fatalf("got error %s", err)
	}

	if strings.ToLower(insertEnglishAuthorsResp) == "y" {
		log.Println("START INSERTING ENGLISH AUTHORS")
		connection.InsertAuthors("english")
		log.Println("DONE WITH INSERTING ENGLISH AUTHORS")
	}

	if strings.ToLower(insertIcelandicAuthorsResp) == "y" {
		log.Println("START INSERTING ICELANDIC AUTHORS")
		connection.InsertAuthors("icelandic")
		log.Println("DONE WITH INSERTING ICELANDIC AUTHORS")
	}

	if strings.ToLower(createIndexesResp) == "y" {
		log.Println("START CREATING INDEXES")
		connection.CreateIndexes() //If this is done first then the rest is much quicker because the quotes.quote column has been indexed!
		log.Println("DONE CREATING INDEXES")
	}

	if strings.ToLower(insertEnglishTopicsResp) == "y" {
		log.Println("START INSERTING ENGLISH TOPICS")
		connection.InsertTopics("English")
		log.Println("DONE INSERTING ENGLISH TOPICS")
	}

	if strings.ToLower(insertIcelandicTopicsResp) == "y" {
		log.Println("START INSERTING ENGLISH TOPICS")
		connection.InsertTopics("Icelandic")
		log.Println("DONE INSERTING ICELANDIC TOPICS")
	}

	if strings.ToLower(createMatViewsResp) == "y" {
		log.Println("START CREATING INDEXES")
		connection.CreateMaterializedViews() //If this is done first then the rest is much quicker because the quotes.quote column has been indexed!
		log.Println("DONE CREATING INDEXES")
	}

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
