package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetContact(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	// convert the id type from string to int
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}
	db := createConnection()
	defer db.Close()

	var contact TempContacts

	// create the select sql query
	sqlStatement := `SELECT contactid,contact_name,contact_email,contact_address,contact_mobileno,is_active,created_at FROM contacts WHERE contactid=$1`

	// execute the sql statement
	row := db.QueryRow(sqlStatement, id)

	// unmarshal the row object to user
	errs := row.Scan(&contact.Id, &contact.ContactName, &contact.ContactEmail, &contact.ContactAddress, &contact.ContactMobileno, &contact.IsActive, &contact.Createdat)
	if errs != nil {
		log.Println("error", errs)
	}
	if contact.ContactName == "" {
		msg := map[string]string{"msg": "Contacts does not exists."}
		json.NewEncoder(w).Encode(msg)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	prettyJSON, err := json.MarshalIndent(contact, "", "    ")
	if err != nil {
		msg := map[string]string{"Failed to generate json": err.Error()}
		json.NewEncoder(w).Encode(msg)
	}
	w.Write(prettyJSON)
}
