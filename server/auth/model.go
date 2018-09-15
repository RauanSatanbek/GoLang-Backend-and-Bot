package auth


import (
	"database/sql"
	"fmt"
	"log"
)

type User struct {
	ID 			int 	`json:"id"`
	Username 	string 	`json:"username"`
	FirstName 	string 	`json:"first_name"`
	TelegramID 	int 	`json:"telegram_id"`
	PhoneNumber string 	`json:"phone_number"`
}

func Migration(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		telegram_id INTEGER NOT NULL,
		username VARCHAR(255),
		first_name VARCHAR(255),
		phone_number VARCHAR(255)
	)`

	if _, err := db.Exec(query); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Auth: table 'Users' - OK")
}

func (u *User) Get(db *sql.DB) error {
	query := `SELECT telegram_id, 
					 username, 
					 first_name, 
					 phone_number 
			  FROM users 
			  WHERE id=$1`
	row := db.QueryRow(query, u.ID)

	return row.Scan(&u.TelegramID, &u.Username, &u.FirstName, &u.PhoneNumber)
}

func (u *User) Create(db *sql.DB) error {
	query := `INSERT INTO 
			  	  users (
			  	  	  telegram_id, 
			  	  	  username, 
			  		  first_name, 
			  		  phone_number
			  	  ) 

			  VALUES ($1, $2, $3, $4) 
			  RETURNING id`

	err := db.QueryRow(query, u.TelegramID, u.Username, u.FirstName, u.PhoneNumber).Scan(&u.ID)

	return err
}

// TODO: define functions: [Update, Delete, GetAll, ...].
func (u *User) Update(db *sql.DB) {

}

func (u *User) Delete(db *sql.DB) {

}

func (u *User) GetAll(db *sql.DB, count, limit int) ([]User){

	return []User{}
}