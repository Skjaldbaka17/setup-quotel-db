package database

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
	"sync"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"
const icelandicAlphabet = "aáæbcdðeéfghiíjklmnoóöpqrstuúvwxyýzþ"

func (connection *Connection) GetAuthors(language string) {
	var wg sync.WaitGroup
	var useAlpha string
	BASE_PATH := "../Quotel-Data-JSON/Authors/Authors-combined/"
	path := BASE_PATH

	if strings.ToLower(language) == "icelandic" {
		path += "Icelandic/"
		useAlpha = icelandicAlphabet
	} else {
		path += "English/"
		useAlpha = alphabet
	}

	nrOfLetters := len(useAlpha)
	for i := 0; i < nrOfLetters; i++ {
		wg.Add(1)
		go connection.GetAuthorsForLetter(string(useAlpha[i]), path, &wg)
	}
	wg.Wait()
}

func (connection *Connection) GetAuthorsForLetter(letter string, path string, wg *sync.WaitGroup) {
	// re1, _ := regexp.Compile(`.json`)
	defer wg.Done()
	info, _ := ReadDir(path + letter)
	for i, name := range info {
		if i == 10 {
			_, _ = GetJSON(fmt.Sprintf("%s/%s/%s", path, letter, name.Name()))
		}
	}
}

type AuthorJSON struct {
	Metadata Metadata `json:"metadata"`
	Name     string
	Quotes   []string
}

type Metadata struct {
	Nationality string
	Profession  string
	Days        []Date
}

type Date struct {
	Birth Day
	Death Day
}

type Day struct {
	Month string
	Day   int
	Year  int
}

func GetJSON(path string) (AuthorJSON, error) {
	// Open JSON
	jsonFile, err := os.Open(path)
	// if os.Open returns an error then handle it
	if err != nil {
		return AuthorJSON{}, err
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	//Read the opened file
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var author AuthorJSON
	//Convert the read value to json and put into the authors-var
	json.Unmarshal(byteValue, &author)
	log.Println(author.Name)
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
