package database

import (
	"gorm.io/gorm"
)

//Make tsv as index later (after inserting data, it is quicker that way)
type Author struct {
	gorm.Model
	Nationality         string `gorm:"index"`
	Profession          string `gorm:"index"`
	BirthYear           int    `gorm:"index"`
	BirthMonth          string `gorm:"index"`
	BirthDate           int    `gorm:"index"`
	DeathYear           int    `gorm:"index"`
	DeathMonth          string `gorm:"index"`
	DeathDate           int    `gorm:"index"`
	Name                string `gorm:"unique"`
	NameTSV             string `gorm:"type:tsvector"`
	Quotes              []Quote
	Count               int `gorm:"index,default 0"`
	NrOfEnglishQuotes   int `gorm:"index"`
	NrOfIcelandicQuotes int `gorm:"index"`
}

type Quote struct {
	gorm.Model
	AuthorID    uint   `gorm:"index"`
	Quote       string `gorm:"index"` //indexed for when inserting the the topics!
	Count       int    `gorm:"index,default 0"`
	IsIcelandic bool
	QuoteTSV    string  `gorm:"type:tsvector"`
	Topics      []Topic `gorm:"many2many:topics_quotes;"`
}

type Topic struct {
	gorm.Model
	Name        string `gorm:"unique"`
	IsIcelandic bool
	Count       int     `gorm:"default 0"`
	Quotes      []Quote `gorm:"many2many:topics_quotes;"`
}
