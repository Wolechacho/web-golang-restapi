package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"web-golang-restapi/models"
	"strconv"

	"github.com/gorilla/mux"
)

type CharacterController struct {
}

func (characterController *CharacterController) List(w http.ResponseWriter, r *http.Request) {
	characterResponseModel := &models.CharacterResponseModel{}

	characters := characterResponseModel.GetCharacters()
	fmt.Println(characters)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(characters)
}

func (characterController *CharacterController) GetCharacterById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	characterResponseModel := &models.CharacterResponseModel{}
	id, err := strconv.ParseInt(params["id"], 0, 64)

	fmt.Println("The id is", id)

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		characters := characterResponseModel.GetCharacterByID(id)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(characters)
	}
}

func (characterController *CharacterController) GetCharacterName(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")

	characterResponseModel := &models.CharacterResponseModel{}
	w.Header().Set("Content-Type", "application/json")

	characters := characterResponseModel.GetCharacterByName(name)
	fmt.Println(characters)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(characters)
}
