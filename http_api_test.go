package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"web-golang-restapi/controllers"
	"web-golang-restapi/models"

	"github.com/gorilla/mux"
)

func TestGetCharacterByName(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/characters?name=Walter", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	pc := &controllers.CharacterController{}
	handler := http.HandlerFunc(pc.GetCharacterName)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestGetCharacterByID(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/characters/", nil)
	params := map[string]string{"id": "2"}
	req = mux.SetURLVars(req, params)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	pc := &controllers.CharacterController{}
	handler := http.HandlerFunc(pc.GetCharacterById)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var character models.CharacterResponseModel
	json.Unmarshal(rr.Body.Bytes(), &character)

	expected := "Jesse Pinkman"
	if character.Name != expected {
		t.Errorf("handler returned character name : got (%s) want %s",
			character.Name, expected)
	}
}
