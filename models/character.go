package models

import (
	"database/sql"
	"fmt"
	"web-golang-restapi/helpers"
	"strings"
)

//CharacterResponseModel - json character response for the calling client
type CharacterResponseModel struct {
	ID         int      `json:"_id"`
	Name       string   `json:"name"`
	Birthday   string   `json:"birthday"`
	Occupation []string `json:"occupation"`
	Img        string   `json:"img"`
	Status     string   `json:"status"`
	Nickname   string   `json:"nickname"`
	Appearance []string `json:"appearance"`
	Portrayed  string   `json:"portrayed"`
	Category   []string `json:"category"`
}

//GetCharacters fetches all characters of the breakingbad movie
func (characterResponseModel *CharacterResponseModel) GetCharacters() []CharacterResponseModel {
	characters := make([]CharacterResponseModel, 0)

	res, err := helpers.DB.Query("SELECT * FROM moviecharacter")

	defer res.Close()

	if err != nil {
		fmt.Println(err)
	}

	var occupation string
	var appearance sql.NullString
	var category string
	for res.Next() {
		var characterModel CharacterResponseModel
		err := res.Scan(&characterModel.ID, &characterModel.Name, &characterModel.Birthday,
			&occupation, &characterModel.Img, &characterModel.Status,
			&characterModel.Nickname, &appearance, &characterModel.Portrayed, &category)

		if err != nil {
			fmt.Println(err)
		}

		if appearance.String != "" {
			characterModel.Appearance = strings.Split(appearance.String, ",")
		}
		characterModel.Occupation = strings.Split(occupation, ",")
		characterModel.Category = strings.Split(category, ",")

		characters = append(characters, characterModel)
	}

	return characters

}

//GetCharacterByID fetches a character by id of the breakingbad movie
func (characterResponseModel *CharacterResponseModel) GetCharacterByID(id int64) CharacterResponseModel {

	var characterModel CharacterResponseModel
	var occupation string
	var appearance sql.NullString
	var category string

	db := helpers.CreateTestDbConnection()
	row := db.QueryRow("select * from moviecharacter where id = ?", id)

	err := row.Scan(&characterModel.ID, &characterModel.Name, &characterModel.Birthday,
		&occupation, &characterModel.Img, &characterModel.Status,
		&characterModel.Nickname, &appearance, &characterModel.Portrayed, &category)

	if err != nil {
		fmt.Println(err)
	}

	if appearance.String != "" {
		characterModel.Appearance = strings.Split(appearance.String, ",")
	}
	characterModel.Occupation = strings.Split(occupation, ",")
	characterModel.Category = strings.Split(category, ",")

	return characterModel
}

func (characterResponseModel *CharacterResponseModel) GetCharacterByName(name string) []CharacterResponseModel {
	db := helpers.CreateTestDbConnection()
	res, err := db.Query("select * from moviecharacter where name like ?", name+"%")

	defer res.Close()

	if err != nil {
		fmt.Println(err)
	}

	var occupation string
	var appearance sql.NullString
	var category string
	characters := make([]CharacterResponseModel, 0)

	for res.Next() {
		var characterModel CharacterResponseModel
		err := res.Scan(&characterModel.ID, &characterModel.Name, &characterModel.Birthday,
			&occupation, &characterModel.Img, &characterModel.Status,
			&characterModel.Nickname, &appearance, &characterModel.Portrayed, &category)

		if err != nil {
			fmt.Println(err)
		}

		if appearance.String != "" {
			characterModel.Appearance = strings.Split(appearance.String, ",")
		}
		characterModel.Occupation = strings.Split(occupation, ",")
		characterModel.Category = strings.Split(category, ",")

		characters = append(characters, characterModel)
	}
	return characters
}
