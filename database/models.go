package database

import (
	"time"

	"gorm.io/gorm"
)

//Make tsv as index later (after inserting data, it is quicker that way)
type Author struct {
	gorm.Model
	Nationality         string
	Profession          string
	BirthYear           int
	BirthMonth          string
	BirthDate           int
	BirthDay            time.Time
	DeathYear           int
	DeathMonth          string
	DeathDate           int
	DeathDay            time.Time
	Name                string `gorm:"unique"`
	NameTSV             string `gorm:"type:tsvector"`
	Quotes              []Quote
	Count               int `gorm:"default 0"`
	NrOfEnglishQuotes   int
	NrOfIcelandicQuotes int
	Aods                []Aod
	Aodices             []Aodice
}

type Aod struct {
	gorm.Model
	Nationality string
	Profession  string
	BirthYear   int
	BirthMonth  string
	BirthDate   int
	BirthDay    time.Time
	DeathYear   int
	DeathMonth  string
	DeathDate   int
	DeathDay    time.Time
	Name        string
	Date        string `gorm:"type:date;unique"`
	AuthorID    uint
}

type Aodice struct {
	gorm.Model
	Nationality string
	Profession  string
	BirthYear   int
	BirthMonth  string
	BirthDate   int
	BirthDay    time.Time
	DeathYear   int
	DeathMonth  string
	DeathDate   int
	DeathDay    time.Time
	Name        string
	Date        string `gorm:"type:date;unique"`
	AuthorID    uint
}

type Quote struct {
	gorm.Model
	AuthorID    uint
	Quote       string //indexed for when inserting the the topics!
	Count       int    `gorm:"default 0"`
	IsIcelandic bool
	QuoteTSV    string  `gorm:"type:tsvector"`
	TSV         string  `gorm:"type:tsvector"`
	Topics      []Topic `gorm:"many2many:topics_quotes;"`

	Name        string
	Nationality string
	Profession  string
	BirthYear   int
	BirthMonth  string
	BirthDate   int
	DeathYear   int
	DeathMonth  string
	DeathDate   int
}

type Qod struct {
	gorm.Model
	AuthorID uint
	Quote    string //indexed for when inserting the the topics!
	QuoteId  uint

	Name        string
	Nationality string
	Profession  string
	BirthYear   int
	BirthMonth  string
	BirthDate   int
	DeathYear   int
	DeathMonth  string
	DeathDate   int
	Date        string `gorm:"type:date;unique"`
}

type Qodice struct {
	gorm.Model
	AuthorID uint
	Quote    string //indexed for when inserting the the topics!
	QuoteId  uint

	Name        string
	Nationality string
	Profession  string
	BirthYear   int
	BirthMonth  string
	BirthDate   int
	BirthDay    time.Time
	DeathYear   int
	DeathMonth  string
	DeathDate   int
	DeathDay    time.Time
	Date        string `gorm:"type:date;unique"`
}

type Topic struct {
	gorm.Model
	Name        string `gorm:"unique"`
	IsIcelandic bool
	Count       int     `gorm:"default 0"`
	Quotes      []Quote `gorm:"many2many:topics_quotes;"`
}
