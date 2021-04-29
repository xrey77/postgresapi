package api

import (
	"encoding/json"
	"log"
	"net/http"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	db := createConnection()

	// close the db connection
	defer db.Close()
	var users []TempUsers

	// create the select sql query
	sqlStatement := `SELECT * FROM users ORDER BY userid`

	// execute the sql statement
	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// iterate over the rows
	for rows.Next() {
		var user TempUsers

		// unmarshal the row object to user
		err = rows.Scan(&user.ID, &user.FullName, &user.Email, &user.MobileNo, &user.UserName, &user.Password, &user.Createdat)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}
		users = append(users, user)
	}
	w.Header().Set("Content-Type", "application/json")
	prettyJSON, err := json.MarshalIndent(users, "", "    ")
	if err != nil {
		msg := map[string]string{"Failed to generate json": err.Error()}
		json.NewEncoder(w).Encode(msg)
	}
	w.Write(prettyJSON)
}
