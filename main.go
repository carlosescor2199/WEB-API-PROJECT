package main

import (
	"./Database"
	"./Controller/Contact"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	client, ctx := Database.GetConnectionMongo()
	defer client.Disconnect(ctx)
	fmt.Println("Staring the application on port:4000")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/contacts", Contact.GetContacts).Methods("GET")
	router.HandleFunc("/contacts/{id}", Contact.GetContact).Methods("GET")
	router.HandleFunc("/contacts", Contact.CreateContact).Methods("POST")
	router.HandleFunc("/contacts/{id}", Contact.UpdateContact).Methods("PUT")
	router.HandleFunc("/contacts/{id}", Contact.DeleteContact).Methods("DELETE")
	_ = http.ListenAndServe(":4000", router)

}
