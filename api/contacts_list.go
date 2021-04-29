package api

import (
	"encoding/json"
	"log"
	"net/http"
)

func ListContacts(w http.ResponseWriter, r *http.Request) {
	db := createConnection()
	defer db.Close()

	var contactx []TempContacts

	// create the select sql query
	sqlStatement := `SELECT contactid,contact_name,contact_email,contact_address,contact_mobileno,is_active,created_at FROM contacts ORDER BY contactid`

	// execute the sql statement
	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// close the statement

	// iterate over the rows
	for rows.Next() {
		var contact TempContacts

		// unmarshal the row object to user
		err = rows.Scan(&contact.Id, &contact.ContactName, &contact.ContactEmail, &contact.ContactAddress, &contact.ContactMobileno, &contact.IsActive, &contact.Createdat)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}
		contactx = append(contactx, contact)
	}
	w.Header().Set("Content-Type", "application/json")
	prettyJSON, err := json.MarshalIndent(contactx, "", "    ")
	if err != nil {
		msg := map[string]string{"Failed to generate json": err.Error()}
		json.NewEncoder(w).Encode(msg)
	}
	w.Write(prettyJSON)
}
