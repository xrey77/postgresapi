package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func UpdateContact(w http.ResponseWriter, r *http.Request) {
	// get the userid from the request params, key is "id"
	params := mux.Vars(r)

	// convert the id type from string to int
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	var contact Contacts

	// decode the json request to user
	err = json.NewDecoder(r.Body).Decode(&contact)

	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	db := createConnection()
	defer db.Close()

	// create the update sql query
	sqlStatement := `UPDATE contacts SET contact_name=$2, contact_email=$3, contact_address=$4, contact_mobileno=$5, is_active=$6, created_at=$7 WHERE contactid=$1`

	// execute the sql statement
	res, err := db.Exec(sqlStatement, id, &contact.ContactName, &contact.ContactEmail, &contact.ContactAddress, &contact.ContactMobileno, &contact.IsActive, time.Now())

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
