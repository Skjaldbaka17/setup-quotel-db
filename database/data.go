package database

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"sort"
	"strings"
	"sync"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TopicJSON struct {
	Topic  string
	Quotes []QuotesJSON
}

type QuotesJSON struct {
	Quote  string
	Author string
}

type AuthorJSON struct {
	Metadata Metadata `json:"metadata"`
	Name     string
	Quotes   []string
}

type Metadata struct {
	Nationality string
	Profession  string
	Days        Days `json:"days"`
}

type Days struct {
	Birth Day
	Death Day
}

type Day struct {
	Month string
	Day   int
	Year  int
}

const alphabet = "abcdefghijklmnopqrstuvwxyz"
const icelandicAlphabet = "aáæbcdðeéfghiíjklmnoóöpqrstuúvwxyýzþ"

var count = 0
var errCount = 0

func (connection *Connection) InsertTopics(language string) {
	var wg sync.WaitGroup
	BASE_PATH := "../Quotel-Data-JSON/Topics/Topics-combined/"
	path := BASE_PATH
	isIcelandic := false

	if strings.ToLower(language) == "icelandic" {
		path += "Icelandic/"
		isIcelandic = true
	} else {
		path += "English/"
	}

	info, _ := ReadDir(path)
	for _, name := range info {
		topicJSON, _ := GetTopicJSON(fmt.Sprintf("%s/%s", path, name.Name()))
		wg.Add(1)
		go connection.InsertTopic(topicJSON, isIcelandic, &wg)
	}
	wg.Wait()
}

func (connection *Connection) InsertTopic(topicJSON TopicJSON, isIcelandic bool, wg *sync.WaitGroup) {
	defer wg.Done()
	log.Printf("Creating topic %s, nr of quotes %d", topicJSON.Topic, len(topicJSON.Quotes))
	topic := Topic{
		Name:   topicJSON.Topic,
		Quotes: []Quote{},
	}

	tenPercent := math.Floor(float64(len(topicJSON.Quotes)) * 0.1)
	for idx, quote := range topicJSON.Quotes {
		quoteFromDB := connection.GetQuote(quote.Quote)
		topicQuote := Quote{
			AuthorID:    quoteFromDB.AuthorID,
			Quote:       quote.Quote,
			IsIcelandic: isIcelandic,
		}
		topicQuote.ID = quoteFromDB.ID
		topic.Quotes = append(topic.Quotes, topicQuote)

		if idx%int(tenPercent) == 0 {
			log.Printf("%s, %.f%%", topicJSON.Topic, 100*float64(idx)/float64(len(topicJSON.Quotes)))
		}
	}

	err := connection.DB.
		Create(&topic).Error

	if err != nil {
		log.Printf("GOT ERROR IN CREATING TOPIC: %s", err)
	}

	log.Printf("Topic %s created", topicJSON.Topic)
}

func (connection *Connection) GetQuote(quoteString string) Quote {
	var quote Quote
	connection.DB.Where("quote = ?", quoteString).First(&quote)
	return quote
}

func (connection *Connection) GetAuthor(name string) Author {
	var author Author
	connection.DB.Where("name = ?", name).First(&author)
	return author
}

func (connection *Connection) InsertAuthors(language string) {
	var wg sync.WaitGroup
	var useAlpha string
	BASE_PATH := "../Quotel-Data-JSON/Authors/Authors-combined/"
	path := BASE_PATH
	isIcelandic := false

	if strings.ToLower(language) == "icelandic" {
		path += "Icelandic/"
		useAlpha = icelandicAlphabet
		isIcelandic = true
	} else {
		path += "English/"
		useAlpha = alphabet
	}

	for _, letter := range useAlpha {
		wg.Add(1)
		go connection.InsertAuthorsForLetter(isIcelandic, strings.ToUpper(string(letter)), path, &wg)
	}
	wg.Wait()
	log.Println(count)
	log.Println(errCount)
}

func (connection *Connection) InsertAuthorsForLetter(isIcelandic bool, letter string, path string, wg *sync.WaitGroup) {
	// re1, _ := regexp.Compile(`.json`)
	defer wg.Done()
	info, _ := ReadDir(path + letter)
	log.Printf("Starting on: %s\n", path+letter)
	for _, name := range info {

		authorJSON, _ := GetAuthorJSON(fmt.Sprintf("%s/%s/%s", path, letter, name.Name()))
		author := Author{
			Nationality: authorJSON.Metadata.Nationality,
			Profession:  authorJSON.Metadata.Profession,
			Name:        authorJSON.Name,
			BirthYear:   authorJSON.Metadata.Days.Birth.Year,
			BirthMonth:  authorJSON.Metadata.Days.Birth.Month,
			BirthDate:   authorJSON.Metadata.Days.Birth.Day,
			DeathYear:   authorJSON.Metadata.Days.Death.Year,
			DeathMonth:  authorJSON.Metadata.Days.Death.Month,
			DeathDate:   authorJSON.Metadata.Days.Death.Day,
		}

		if isIcelandic {
			author.NrOfIcelandicQuotes = len(authorJSON.Quotes)
		} else {
			author.NrOfEnglishQuotes = len(authorJSON.Quotes)
		}

		quotes := []Quote{}
		for _, quote := range authorJSON.Quotes {
			quotes = append(quotes, Quote{
				Quote:       quote,
				IsIcelandic: isIcelandic,
			})
		}
		author.Quotes = quotes
		connection.InsertAuthor(author, isIcelandic)

	}
	log.Printf("Done with: %s\n", path+letter)
}

func (connection *Connection) InsertAuthor(author Author, isIcelandic bool) {
	count += len(author.Quotes)
	var err error
	// Insert into Authors -- Omitting the Quotes (because of error, unique constraint on quote-column, when the quote is the same as another quote)
	if isIcelandic {
		err = connection.DB.Clauses(clause.OnConflict{
			Columns: []clause.Column{{Name: "name"}},
			DoUpdates: clause.Assignments(map[string]interface{}{
				"nr_of_icelandic_quotes": gorm.Expr("?", len(author.Quotes)),
			}),
		}).Omit("Quotes").Select("*").
			Create(&author).Error
	} else {
		err = connection.DB.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "name"}},
			DoUpdates: clause.Assignments(map[string]interface{}{"nr_of_english_quotes": gorm.Expr("?", len(author.Quotes))}),
		}).Omit("Quotes").Select("*").Create(&author).Error
	}
	if err != nil {
		log.Printf("Got error: %s when inserting author %+v", err, author)
	}

	//Setting the authorID for each quotes -- otherwise we get foreign key constraint error from postgres
	for idx := range author.Quotes {
		author.Quotes[idx].AuthorID = author.ID
		//Inserting the quotes for the author -- one at a time so that if the quote is already there then the 'DoNothing' onConflict statement
		// will not stop us from inserting the other quotes in the batch
		err = connection.DB.Create(&author.Quotes[idx]).Error
		if err != nil {
			log.Printf("Got error: %s when inserting quotes for authorID %d, %s", err, author.ID, author.Quotes[idx].Quote)
			errCount += 1
		}
	}
}

func GetTopicJSON(path string) (TopicJSON, error) {
	// Open JSON
	jsonFile, err := os.Open(path)
	// if os.Open returns an error then handle it
	if err != nil {
		log.Println("ERROR OPENING", path)
		return TopicJSON{}, err
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	//Read the opened file
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var topic TopicJSON
	//Convert the read value to json and put into the topic-var
	json.Unmarshal(byteValue, &topic)
	return topic, nil
}

func GetAuthorJSON(path string) (AuthorJSON, error) {
	// Open JSON
	jsonFile, err := os.Open(path)
	// if os.Open returns an error then handle it
	if err != nil {
		log.Println("ERROR OPENING", path)
		return AuthorJSON{}, err
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	//Read the opened file
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var author AuthorJSON
	//Convert the read value to json and put into the authors-var
	json.Unmarshal(byteValue, &author)
	return author, nil
}

func ReadDir(dirname string) ([]os.FileInfo, error) {
	f, err := os.Open(dirname)
	if err != nil {
		return nil, err
	}
	list, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		return nil, err
	}
	sort.Slice(list, func(i, j int) bool { return list[i].Name() < list[j].Name() })
	return list, nil
}
