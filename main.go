package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/katnik/spinder/defaults"
	"github.com/katnik/spinder/girls"
	"github.com/katnik/spinder/users"
	"github.com/katnik/spinder/utils/misc"
	"github.com/katnik/spinder/utils/mongo"
)

func main() {
	misc.GetInstance(mongo.GetCollectionSize("users"))

	mainRouter := mux.NewRouter()

	getSubrouter := mainRouter.Methods("GET").Subrouter()
	postSubrouter := mainRouter.Methods("POST").Subrouter()
	putSubrouter := mainRouter.Methods("PUT").Subrouter()
	deleteSubrouter := mainRouter.Methods("DELETE").Subrouter()

	getSubrouter.HandleFunc("/", defaults.HomeHandler)
	getSubrouter.HandleFunc("/girl/{id}", girls.Read)
	// getSubrouter.HandleFunc("/boy/{id}", GetBoy)
	// getSubrouter.HandleFunc("/table/{id}", GetTable)
	// getSubrouter.HandleFunc("/question{id}", GetQuestion)
	// getSubrouter.HandleFunc("/current", GetCurrent)

	postSubrouter.HandleFunc("/user", users.Create)

	postSubrouter.HandleFunc("/", defaults.NotAllowed)
	putSubrouter.HandleFunc("/", defaults.NotAllowed)
	deleteSubrouter.HandleFunc("/", defaults.NotAllowed)

	http.Handle("/", mainRouter)

	err := http.ListenAndServe("localhost:8080", nil)

	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
