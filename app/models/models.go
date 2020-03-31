package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/lib/pq"
)

type Artist struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	Songs       []Song `gorm:"ForeignKey:ArtistID" json:"songs"`
}

type Song struct {
	gorm.Model
	Title    string         `json:"title"`
	ArtistID uint           `json:"artist_id"`
	Genre    pq.StringArray `gorm:"type:varchar(150)[]" json:"genre"`
	Length   int            `json:"length"`
}

func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Artist{}, &Song{})
	db.Model(&Song{}).AddForeignKey("artist_id", "artists(id)", "CASCADE", "CASCADE")
	return db
}
