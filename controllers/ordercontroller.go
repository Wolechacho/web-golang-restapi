package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"web-golang-restapi/models"
)

type OrderController struct {
}

func (orderController *OrderController) Create(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		fmt.Println("hhhhhh", err)
		w.WriteHeader(http.StatusBadRequest)
	} else {
		var orderRequestModel models.OrderRequestModel
		err = json.Unmarshal(reqBody, &orderRequestModel)

		if err != nil {
			fmt.Println("gggggg", err)
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			err = orderRequestModel.Save()
			if err != nil {
				fmt.Println("aaaaaaaa", err)
				w.WriteHeader(http.StatusInternalServerError)
			} else {
				w.WriteHeader(http.StatusCreated)
			}
		}
	}
}
