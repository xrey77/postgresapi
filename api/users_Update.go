package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// // UpdateUser update user's detail in the postgres db
func UpdateUser(w http.ResponseWriter, r *http.Request) {

	// w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.Header().Set("Access-Control-Allow-Methods", "PUT")
	// w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// get the userid from the request params, key is "id"
	params := mux.Vars(r)

	// convert the id type from string to int
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	// create an empty user of type models.User
	var user TempUsers

	// decode the json request to user
	err = json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	db := createConnection()
	// close the db connection
	defer db.Close()

	// create the update sql query
	sqlStatement := `UPDATE users SET full_name=$2, email=$3, mobile_no=$4, username=$5, passwd=$6, created_at=$7 WHERE userid=$1`

	// execute the sql statement
	res, err := db.Exec(sqlStatement, id, user.FullName, user.Email, user.MobileNo, user.UserName, user.Password, time.Now())

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	if rowsAffected > 0 {
		msg := map[string]string{"msg": "Updated Successfully."}
		json.NewEncoder(w).Encode(msg)
	} else {
		msg := map[string]string{"msg": "Unable to Update, ID does not exists."}
		json.NewEncoder(w).Encode(msg)
	}
}
