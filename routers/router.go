package routers

import (
	"fmt"
	"net/http"

	"web-golang-restapi/controllers"

	"github.com/gorilla/mux"
)

//RegisterRoutes - Configure routes for all incoming routes
func RegisterRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)

	pc := &controllers.CharacterController{}
	router.HandleFunc("/api/characters", pc.GetCharacterName).Methods("GET").Queries(
		"name", "{name}",
	)
	router.HandleFunc("/api/characters/{id:[0-9]+}", pc.GetCharacterById).Methods("GET")
	router.HandleFunc("/api/characters", pc.List).Methods("GET")

	cc := &controllers.OrderController{}
	router.HandleFunc("/api/orders", cc.Create).Methods("POST")

	return router
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}
