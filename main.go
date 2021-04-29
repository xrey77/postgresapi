package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"postgresapi/api"

	"github.com/gorilla/mux"
)

//cors
func Cors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=ascii")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	// w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers,Content-Type")
}

//w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

func apiTest(w http.ResponseWriter, r *http.Request) {
	msg := map[string]string{"msg": "This is a Secured API"}
	json.NewEncoder(w).Encode(msg)
}

func main() {
	router := mux.NewRouter().PathPrefix("/api/v1/").Subrouter()
	router.HandleFunc("/test", apiTest).Methods("GET", "OPTIONS")
	router.HandleFunc("/getuser/{id}", api.GetUser).Methods(http.MethodGet)
	router.HandleFunc("/getusers", api.GetUsers).Methods("GET", "OPTIONS")
	router.HandleFunc("/register", api.Register).Methods(http.MethodPost)
	router.HandleFunc("/updateuser/{id}", api.UpdateUser).Methods("PUT", "OPTIONS")
	router.HandleFunc("/deleteuser/{id}", api.DeleteUser).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/userlogin", api.User_Login).Methods(http.MethodPost)

	router.HandleFunc("/createcontact", api.CreateContact).Methods("POST", "OPTIONS")
	router.HandleFunc("/listcontacts", api.ListContacts).Methods("GET", "OPTIONS")
	router.HandleFunc("/getcontact/{id}", api.GetContact).Methods("GET", "OPTIONS")
	router.HandleFunc("/updatecontact/{id}", api.UpdateContact).Methods("PUT", "OPTIONS")
	router.HandleFunc("/deletecontact/{id}", api.DeleteContact).Methods("DELETE", "OPTIONS")
	router.Use(api.JwtAuthentication) //attach JWT auth middleware
	router.HandleFunc("/plm/cors", Cors)

	log.Println("Server is listnening in port : 8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println(err.Error())
	}
}
