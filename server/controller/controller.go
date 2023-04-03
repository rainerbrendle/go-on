package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var router *mux.Router

func Start() {

	router = mux.NewRouter()

	initHandlers()
	fmt.Printf("router initialized\n")
	log.Fatal(http.ListenAndServe(":3200", router))

}

func initHandlers() {
	// initialize HTTP handlers
	// this approach may be too limited eventually
	// will need to use closures on models instead

	// router.HandleFunc( " ... api ...", controller.GetAllPosts).Methods( "GET) ") etc.

}

func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	/*
		posts, err := model.GetAllPosts()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		} else {
			json.NewEncoder(w).Encode(posts)
		}
	*/
}
