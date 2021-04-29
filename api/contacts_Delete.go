package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func DeleteContact(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	// convert the id in string to int
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	db := createConnection()
	defer db.Close()

	// create the delete sql query
	sqlStatement := `DELETE FROM contacts WHERE contactid=$1`

	// execute the sql statement
	res, err := db.Exec(sqlStatement, id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}
	if rowsAffected > 0 {
		msg := map[string]string{"msg": "Deleted Successfully."}
		json.NewEncoder(w).Encode(msg)
	} else {
		msg := map[string]string{"msg": "Unable to Delete, ID does not exists."}
		json.NewEncoder(w).Encode(msg)
	}

}
