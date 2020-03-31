package handlers

import (
	"encoding/json"
	"net/http"

	"golang-crud-api/app/models"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func GetAllSongs(db *gorm.DB, responseWriter http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	artistName := vars["artistName"]
	artist := getArtistOr404(db, artistName, responseWriter, request)
	if artist == nil {
		return
	}

	songs := []models.Song{}
	if err := db.Model(&artist).Related(&songs).Error; err != nil {
		respondError(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(responseWriter, http.StatusOK, songs)
}

func CreateSong(db *gorm.DB, responseWriter http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	artistName := vars["artistName"]
	artist := getArtistOr404(db, artistName, responseWriter, request)
	if artist == nil {
		return
	}

	song := models.Song{ArtistID: artist.ID}
	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&song); err != nil {
		respondError(responseWriter, http.StatusBadRequest, err.Error())
		return
	}
	defer request.Body.Close()

	if err := db.Save(&song).Error; err != nil {
		respondError(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(responseWriter, http.StatusCreated, song)
}

func GetSong(db *gorm.DB, responseWriter http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	artistName := vars["artistName"]
	songTitle := vars["songTitle"]
	song := getSongOr404(db, artistName, songTitle, responseWriter, request)
	if song == nil {
		return
	}

	respondJSON(responseWriter, http.StatusFound, song)
}

func UpdateSong(db *gorm.DB, responseWriter http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	artistName := vars["artistName"]
	songTitle := vars["songTitle"]
	song := getSongOr404(db, artistName, songTitle, responseWriter, request)
	if song == nil {
		return
	}

	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&song); err != nil {
		respondError(responseWriter, http.StatusBadRequest, err.Error())
		return
	}
	defer request.Body.Close()

	if err := db.Save(&song).Error; err != nil {
		respondError(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(responseWriter, http.StatusOK, song)
}

func DeleteSong(db *gorm.DB, responseWriter http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	artistName := vars["artistName"]
	songTitle := vars["songTitle"]
	song := getSongOr404(db, artistName, songTitle, responseWriter, request)
	if song == nil {
		return
	}

	if err := db.Delete(&song).Error; err != nil {
		respondError(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(responseWriter, http.StatusNoContent, nil)
}

func getSongOr404(db *gorm.DB, artistName string, title string, responseWriter http.ResponseWriter, request *http.Request) *models.Song {
	artist := getArtistOr404(db, artistName, responseWriter, request)

	song := models.Song{}
	if err := db.Find(&song, models.Song{Title: title}).Related(&artist).Error; err != nil {
		respondError(responseWriter, http.StatusNotFound, err.Error())
		return nil
	}

	return &song
}
