package api

import (
	"encoding/json"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

// CreateUser create a user in the postgres db
func Register(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.Header().Set("Access-Control-Allow-Methods", "POST")
	// w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// create an empty user of type models.User
	var user User
	// decode the json request to user
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	db := createConnection()
	defer db.Close()

	xbyte := getPwd(user.Password)
	hashPwd := hashAndSalt(xbyte)
	//INSERT
	sqlStatement := `call insertuser($1, $2, $3, $4, $5)`
	errs := db.QueryRow(sqlStatement, &user.FullName, &user.Email, &user.MobileNo, &user.UserName, &hashPwd)
	if errs != nil {

		if errs.Scan().Error() == "pq: 001" {
			msg := map[string]string{"msg": "Fullname has been taken."}
			json.NewEncoder(w).Encode(msg)
		} else if errs.Scan().Error() == "pq: 002" {
			msg := map[string]string{"msg": "Email Address has been taken."}
			json.NewEncoder(w).Encode(msg)
		} else if errs.Scan().Error() == "pq: 003" {
			msg := map[string]string{"msg": "Username has been taken."}
			json.NewEncoder(w).Encode(msg)
		} else if errs.Scan().Error() == "pq: 004" {
			msg := map[string]string{"msg": "Data has been inserted successfully."}
			json.NewEncoder(w).Encode(msg)
		}

	}
}

func getPwd(pwd string) []byte {
	return []byte(pwd)
}

func hashAndSalt(pwd []byte) string {

	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}
