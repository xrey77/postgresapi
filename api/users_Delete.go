package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// // DeleteUser delete user's detail in the postgres db
func DeleteUser(w http.ResponseWriter, r *http.Request) {

	// w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	// w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// get the userid from the request params, key is "id"
	params := mux.Vars(r)

	// convert the id in string to int
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	// call the deleteUser, convert the int to int64
	//	deletedRows := deleteUser(int64(id))

	// format the message string
	//	msg := fmt.Sprintf("User updated successfully. Total rows/record affected %v", deletedRows)

	// format the reponse message
	// res := response{
	// 	ID:      int64(id),
	// 	Message: msg,
	// }

	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	// create the delete sql query
	sqlStatement := `DELETE FROM users WHERE userid=$1`

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

	//	fmt.Printf("Total rows/record affected %v", rowsAffected)

}
