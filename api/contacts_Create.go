package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func CreateContact(w http.ResponseWriter, r *http.Request) {
	var contact Contacts
	// decode the json request to contacts
	err := json.NewDecoder(r.Body).Decode(&contact)
	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	active, err1 := strconv.ParseInt(contact.IsActive, 10, 64)
	if err1 != nil {
		log.Println(err1)
	}
	db := createConnection()
	defer db.Close()

	//INSERT
	sqlcontact := `call insertcontact($1, $2, $3, $4, $5)`
	errs2 := db.QueryRow(sqlcontact, &contact.ContactName, &contact.ContactAddress, &contact.ContactEmail, &contact.ContactMobileno, &active)
	if errs2 != nil {

		if errs2.Scan().Error() == "pq: 001" {
			msg := map[string]string{"msg": "Contact Name has been taken."}
			json.NewEncoder(w).Encode(msg)
		} else if errs2.Scan().Error() == "pq: 002" {
			msg := map[string]string{"msg": "Email Address has been taken."}
			json.NewEncoder(w).Encode(msg)
		} else if errs2.Scan().Error() == "pq: 004" {
			msg := map[string]string{"msg": "Data has been inserted successfully."}
			json.NewEncoder(w).Encode(msg)
		}

	}

}
