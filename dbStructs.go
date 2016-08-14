package pubmedSqlStructs

import (
	"database/sql"
)

type Article struct {
	Abstract        string
	Authors         []Author    `gorm:"many2many:article_author;"`
	Chemicals       []*Chemical `gorm:"many2many:article_chemical;"`
	Citations       []*Citation `gorm:"many2many:article_citation;"`
	Day             int
	Genes           []Gene `gorm:"many2many:article_gene;"`
	ID              int64  `gorm:"primary_key"` // PMID
	Issue           string
	Journal         *Journal
	JournalID       sql.NullInt64
	Keywords        []*Keyword `gorm:"many2many:article_keyword;"`
	KeywordsOwner   string
	Language        string
	MeshDescriptors []*MeshDescriptor
	Month           string `sql:"size:8`
	OtherId         []OtherID
	Title           string
	Volume          string
	Year            int   `sql:"size:4`
	DateRevised     int64 //YYYYMMDD 20140216
}

type OtherID struct {
	ID      int `gorm:"primary_key"`
	Source  string
	OtherID string
}

type Journal struct {
	ID              int `gorm:"primary_key"`
	Articles        []Article
	IsoAbbreviation string `sql:"size:128`
	Issn            string `sql:"size:10`
	Title           string
}

type Author struct {
	ID          int    `gorm:"primary_key"`
	LastName    string `sql:"size:48`
	FirstName   string `sql:"size:16`
	MiddleName  string `sql:"size:16`
	Affiliation string
}

type Keyword struct {
	ID         int `gorm:"primary_key"`
	MajorTopic bool
	Name       string `sql:"size:128`
}

type MeshDescriptor struct {
	ID         int `gorm:"primary_key"`
	Name       string
	Type       string `sql:"size:32`
	MajorTopic bool
	Qualifiers []*MeshQualifier
	ArticleID  int64
	UI         string
}

type MeshQualifier struct {
	ID               int `gorm:"primary_key"`
	MajorTopic       bool
	Name             string `sql:"size:128`
	MeshDescriptorID int
	UI               string
}

type Gene struct {
	ID   int `gorm:"primary_key"`
	Name string
}

type Chemical struct {
	ID       int `gorm:"primary_key"`
	Name     string
	Registry string `sql:"size:32`
}

type Citation struct {
	ID int64 `gorm:"primary_key"`
}
