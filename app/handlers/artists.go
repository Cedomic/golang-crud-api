package handlers

import (
	"encoding/json"
	"net/http"

	"golang-crud-api/app/models"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func GetAllArtists(db *gorm.DB, responseWriter http.ResponseWriter, request *http.Request) {
	artists := []models.Artist{}
	db.Find(&artists)

	respondJSON(responseWriter, http.StatusOK, artists)
}

func CreateArtist(db *gorm.DB, responseWriter http.ResponseWriter, request *http.Request) {
	artist := models.Artist{}
	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&artist); err != nil {
		respondError(responseWriter, http.StatusBadRequest, err.Error())
		return
	}
	defer request.Body.Close()

	if err := db.Save(&artist).Error; err != nil {
		respondError(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(responseWriter, http.StatusCreated, artist)
}

func GetArtist(db *gorm.DB, responseWriter http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	name := vars["artistName"]
	artist := getArtistOr404(db, name, responseWriter, request)
	if artist == nil {
		return
	}

	songs := []models.Song{}
	if err := db.Model(&artist).Related(&songs).Error; err != nil {
		respondError(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}
	artist.Songs = songs

	respondJSON(responseWriter, http.StatusFound, artist)
}

func UpdateArtist(db *gorm.DB, responseWriter http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	name := vars["artistName"]
	artist := getArtistOr404(db, name, responseWriter, request)
	if artist == nil {
		return
	}

	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&artist); err != nil {
		respondError(responseWriter, http.StatusBadRequest, err.Error())
		return
	}
	defer request.Body.Close()

	if err := db.Save(&artist).Error; err != nil {
		respondError(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(responseWriter, http.StatusOK, artist)
}

func DeleteArtist(db *gorm.DB, responseWriter http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	name := vars["artistName"]
	artist := getArtistOr404(db, name, responseWriter, request)
	if artist == nil {
		return
	}

	if err := db.Delete(&artist).Error; err != nil {
		respondError(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(responseWriter, http.StatusNoContent, nil)
}

func getArtistOr404(db *gorm.DB, name string, responseWriter http.ResponseWriter, request *http.Request) *models.Artist {
	artist := models.Artist{}
	if err := db.First(&artist, models.Artist{Name: name}).Error; err != nil {
		respondError(responseWriter, http.StatusNotFound, err.Error())
		return nil
	}

	return &artist
}
