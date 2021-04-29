package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/xrey77/postgresapi/utils"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

//var ctx *gin.Context

type XUser struct {
	Username string `json:"username"`
	Passwd   string `json:"passwd"`
}

func User_Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var xuser XUser
	err := json.NewDecoder(r.Body).Decode(&xuser)
	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	huser, err := Getlogin(xuser.Username, xuser.Passwd)
	if err != nil {
		log.Fatalf("Unable to get user. %v", err)
	}

	if comparePassword(huser.PassWord, getPassword(xuser.Passwd)) == true {
		tk := &Token{UserId: huser.ID}
		token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
		tokenString, _ := token.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
		huser.Token = tokenString

		msg := map[string]string{"token": huser.Token}
		json.NewEncoder(w).Encode(msg)
		resp := utils.Message(true, "Logged In")
		resp["user"] = huser

	} else {
		msg := map[string]string{"msg": "Access Denied."}
		json.NewEncoder(w).Encode(msg)
	}

	//	json.NewEncoder(w).Encode(user)
}

func Getlogin(uname string, pwd string) (UserLogin, error) {
	db := createConnection()
	defer db.Close()

	var user UserLogin
	sql := `SELECT userid, username, passwd FROM users WHERE username=$1`
	row := db.QueryRow(sql, uname)

	err := row.Scan(&user.ID, &user.UserName, &user.PassWord)
	if err != nil {
		fmt.Println(err.Error())
	}
	return user, nil
	// if errs != nil {
	// 	return
	// } else {

	// if comparePassword(user.PassWord, getPassword(pwd)) == true {
	// 	tk := &model.Token{UserId: user.ID}
	// 	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	// 	tokenString, _ := token.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	// 	user.Token = tokenString

	// 	msg := map[string]string{"token": user.Token}
	// 	json.NewEncoder(w).Encode(msg)
	// 	resp := u.Message(true, "Logged In")
	// 	resp["user"] = user

	// } else {
	// 	msg := map[string]string{"msg": "Access Denied."}
	// 	json.NewEncoder(w).Encode(msg)
	// }
	//	}

}

func comparePassword(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func getPassword(pwd string) []byte {
	return []byte(pwd)
}
