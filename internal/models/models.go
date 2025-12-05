package models

import "time"

type User struct {
	ID			 int64	   `json:"id"`
	Username	 string	   `json:"username"`
	PasswordHash string	   `json:"-"`	// Omit password hash from JSON
	DisplayName	 string	   `json:"display_name"`
	CreatedAt	 time.Time `json:"created_at"`
	UpdatedAt	 time.Time `json:"updated_at"`
}

type Category struct {
	ID			int64	`json:"id"`
	ParentID	*int64	`json:"parent_id"`	// Use pointer for nullable foreign key
	Name		string	`json:"name"`
	Description	string	`json:"description"`
}

type Expense struct {
	ID				int64	  `json:"id"`
	UserID			int64	  `json:"user_id"`
	TransacionTime	time.Time `json:"transaction_time"`
	Amount			float64	  `json:"amount"`
	CategoryID		int64	  `json:"category_id"`
	Description		string	  `json:"description"`
	CreatedAt		time.Time `json:"created_at"`
	UpdatedAt		time.Time `json:"updated_at"`
}

type Setting struct {
	ID				int64	  `json:"id"`
	Key				string	  `json:"key"`
	Value			string	  `json:"value"`
	ForUserID		int64	  `json:"for_user_id"`
	UpdatedByUserID	int64	  `json:"updated_by_user_id"`
	UpdatedAt		time.Time `json:"updated_at"`
}