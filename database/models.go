package database

import (
	"time"

	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	Nationality         string    `gorm:"index"`
	Profession          string    `gorm:"index"`
	BirthDay            time.Time `gorm:"index"`
	DeathDay            time.Time `gorm:"index"`
	Name                string
	NameTSV             string `gorm:"index,type:ts_vector"`
	Quotes              []Quote
	Count               int `gorm:"index"`
	NrOfEnglishQuotes   int `gorm:"index"`
	NrOfIcelandicQuotes int `gorm:"index"`
}

type Quote struct {
	gorm.Model
	AuthorID    uint `gorm:"index"`
	Quote       string
	Count       int `gorm:"index"`
	IsIcelandic bool
	QuoteTSV    string `gorm:"index,type:ts_vector"`
}
