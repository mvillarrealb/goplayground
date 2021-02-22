package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func findAll(w http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	var limit = 10
	var offset = 0
	if strLimit := query.Get("limit"); strLimit != "" {
		if lmt, err := strconv.Atoi(strLimit); err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		} else {
			limit = lmt
		}
	}
	if strOffset := query.Get("offset"); strOffset != "" {
		if ofst, err := strconv.Atoi(strOffset); err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		} else {
			offset = ofst
		}
	}
	var Tasks []Task
	app.DB.Limit(limit).Offset(offset).Find(&Tasks)
	respondWithJSON(w, 200, Tasks)
}

func deleteTask(w http.ResponseWriter, req *http.Request) {
	var task Task
	vars := mux.Vars(req)
	key := vars["id"]
	err := app.DB.Delete(&task, key).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		respondWithError(w, 404, "Task Not Found")
	} else {
		respondWithJSON(w, 202, task)
	}
}

func findByID(w http.ResponseWriter, req *http.Request) {
	var task Task
	vars := mux.Vars(req)
	key := vars["id"]
	err := app.DB.First(&task, key).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		respondWithError(w, 404, "Task Not Found")
	} else {
		respondWithJSON(w, 200, task)
	}
}
func createTask(w http.ResponseWriter, r *http.Request) {
	var task Task
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&task); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := app.DB.Create(&task).Error; err != nil {
		log.Fatal(err)
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, 201, task)
}
