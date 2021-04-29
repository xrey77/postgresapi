package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// GetUser will return a single user by its id
func GetUser(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	// get the userid from the request params, key is "id"
	params := mux.Vars(r)

	// convert the id type from string to int
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}
	db := createConnection()

	// close the db connection
	defer db.Close()

	// create a user of models.User type
	var user TempUsers

	// create the select sql query
	sqlStatement := `SELECT * FROM users WHERE userid=$1`

	// execute the sql statement
	row := db.QueryRow(sqlStatement, id)

	// unmarshal the row object to user
	errs := row.Scan(&user.ID, &user.FullName, &user.Email, &user.MobileNo, &user.UserName, &user.Password, &user.Createdat)
	if errs != nil {
		log.Println("error", errs)
	}
	w.Header().Set("Content-Type", "application/json")
	prettyJSON, err := json.MarshalIndent(user, "", "    ")
	if err != nil {
		msg := map[string]string{"Failed to generate json": err.Error()}
		json.NewEncoder(w).Encode(msg)
	}
	w.Write(prettyJSON)
}
