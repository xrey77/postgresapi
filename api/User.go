package api

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type User struct {
	//ID        int64  `json:"userid"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	MobileNo string `json:"mobile_no"`
	UserName string `json:"username"`
	Password string `json:"passwd"`
}

type TempUsers struct {
	ID        int64     `json:"userid"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	MobileNo  string    `json:"mobile_no"`
	UserName  string    `json:"username"`
	Password  string    `json:"passwd"`
	Createdat time.Time `json:"created_at"`
}

type UserLogin struct {
	//	gorm.Model
	ID       uint   `json:"userid"`
	UserName string `json:"username"`
	PassWord string `json:"passwd"`
	Token    string `json:"token";sql:"-"`
}

//Token    string `json:"token";sql:"-"`

type Token struct {
	UserId uint
	jwt.StandardClaims
}

// func xLogin(email, password string) map[string]interface{} {

// 	account := &Account{}
// 	err := GetDB().Table("accounts").Where("email = ?", email).First(account).Error
// 	if err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			return u.Message(false, "Email address not found")
// 		}
// 		return u.Message(false, "Connection error. Please retry")
// 	}

// 	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))
// 	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
// 		return u.Message(false, "Invalid login credentials. Please try again")
// 	}
// 	//Worked! Logged In
// 	account.Password = ""

// 	//Create JWT token
// 	tk := &Token{UserId: account.ID}
// 	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
// 	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
// 	account.Token = tokenString //Store the token in the response

// 	resp := u.Message(true, "Logged In")
// 	resp["account"] = account
// 	return resp
// }
