package pubmedSqlStructs

import (
//	"database/sql"
)

type Article struct {
	Abstract         string
	ArticleIDs       []*ArticleID
	Authors          []Author    `gorm:"many2many:article_author;"`
	Chemicals        []*Chemical `gorm:"many2many:article_chemical;"`
	Citations        []*Citation `gorm:"many2many:article_citation;"`
	DataBanks        []*DataBank
	DateRevised      uint64 //YYYYMMDD 20140216
	Day              uint8  `sql:"type:int1"`
	ID               uint32 `gorm:"primary_key" sql:"type:int4"` // PMID
	Issue            string `sql:"size:32" sql:"type:int4"`
	Journal          *Journal
	JournalID        uint32     `sql:"type:int4"`
	Keywords         []*Keyword `gorm:"many2many:article_keyword;"`
	KeywordsOwner    string     `sql:"size:128"`
	Language         string     `sql:"size:3"`
	MeshDescriptors  []*MeshDescriptor
	Month            string `sql:"size:8"`
	OtherId          []OtherID
	PublicationTypes []*PublicationType `gorm:"many2many:article_publicationtype;"`
	Retracted        bool
	Title            string
	Version          uint8 `sql:"type:int1"`
	Volume           string
	Year             uint16 `sql:"type:int2"`
}

type PublicationType struct {
	ID   uint32 `gorm:"primary_key" sql:"type:int4"`
	UI   string `sql:"size:8"`
	Name string `sql:"size:64"`
}

type ArticleID struct {
	OtherArticleID string `sql:"size:64"`
	Type           string `sql:"size:12"`
	ID             uint32 `gorm:"primary_key" sql:"type:int4"`
	ArticleID      uint32 `sql:"type:int4"`
}

type DataBank struct {
	ID               uint32 `gorm:"primary_key" sql:"type:int4"`
	Name             string `sql:"size:32"`
	AccessionNumbers []*AccessionNumber
	ArticleID        uint32 `sql:"type:int4"`
}

type AccessionNumber struct {
	ID         uint32 `gorm:"primary_key" sql:"type:int4"`
	Number     string `sql:"size:32"`
	DataBankID uint32 `sql:"type:int2"`
}

type OtherID struct {
	ID      uint32 `gorm:"primary_key" sql:"type:int4"`
	Source  string
	OtherID string
}

type Journal struct {
	ID              uint32 `gorm:"primary_key" sql:"type:int4"`
	Articles        []Article
	IsoAbbreviation string `sql:"size:128"`
	Issn            string `sql:"size:10"`
	Title           string
}

type Author struct {
	ID             uint32 `gorm:"primary_key" sql:"type:int4"`
	LastName       string `sql:"size:48"`
	FirstName      string `sql:"size:16"`
	MiddleName     string `sql:"size:16"`
	Affiliation    string
	CollectiveName string
}

type Keyword struct {
	ID         uint32 `gorm:"primary_key" sql:"type:int4"`
	MajorTopic bool
	Name       string `sql:"size:128"`
}

type MeshDescriptor struct {
	ID         uint32 `gorm:"primary_key" sql:"type:int4"`
	Name       string
	Type       string `sql:"size:32"`
	MajorTopic bool
	Qualifiers []*MeshQualifier
	ArticleID  uint32 `sql:"type:int4"`
	UI         string
}

type MeshQualifier struct {
	ID               uint32 `gorm:"primary_key" sql:"type:int4"`
	MajorTopic       bool
	Name             string `sql:"size:128"`
	MeshDescriptorID uint32 `sql:"type:int4"`
	UI               string
}

type Gene struct {
	ID   uint32 `gorm:"primary_key" sql:"type:int4"`
	Name string
}

type Chemical struct {
	ID       uint32 `gorm:"primary_key" sql:"type:int4"`
	Name     string
	Registry string `sql:"size:32"`
}

type Citation struct {
	ID uint32 `gorm:"primary_key" sql:"type:int4"`
}
