package api

import "time"

type Contacts struct {
	ContactName     string `json:"contact_name"`
	ContactAddress  string `json:"contact_address"`
	ContactEmail    string `json:"contact_email"`
	ContactMobileno string `json:"contact_mobileno"`
	IsActive        string `json:"is_active"`
}

type TempContacts struct {
	Id              int64     `json:"contactid"`
	ContactName     string    `json:"contact_name"`
	ContactAddress  string    `json:"contact_address"`
	ContactEmail    string    `json:"contact_email"`
	ContactMobileno string    `json:"contact_mobileno"`
	IsActive        int64     `json:"is_active"`
	Createdat       time.Time `json:"created_time"`
}
