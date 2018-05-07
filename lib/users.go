package lib

import (
	"log"
)

type User struct {
	Id         int
	Username   string
	Password   string
	Active     int
	Created    string
	CreatedId  string
	Modified   string
	ModifiedId string
}

func GetUsers() []User {

	rows, err := db.Query("SELECT id, username, password, active  FROM users")
	if err != nil {
		log.Fatal("error occurred while executing select query :: ", err)
	}

	users := []User{}
	for rows.Next() {
		var id int
		var username string
		var password string
		var active int
		err = rows.Scan(&id, &username, &password, &active)
		user := User{Id: id, Username: username, Password: password, Active: active}
		users = append(users, user)
	}

	return users

}

func CheckUser(username string, password string) string {

	var id string

	err := db.QueryRow("SELECT id  FROM users WHERE username=? AND password=? AND active=1 LIMIT 1", username, password).Scan(&id)

	if err != nil {
		log.Println("Error while reading user", err)
	}

	return id

}

func GetUser(uid string) User {

	var id int
	var username string
	var active int

	err := db.QueryRow("SELECT id, username, active  FROM users WHERE id=? LIMIT 1", uid).Scan(&id, &username, &active)

	if err != nil {
		log.Fatal("Error while reading user", err)
	}

	user := User{Id: id, Username: username, Active: active}

	return user

}

func InsertUser(u User) {

	stmt, err := db.Prepare("INSERT INTO users ( username, password, active, created_id, modified_id, created, modified  ) VALUES ( ?, ?, ?, ?, ?, NOW(), NOW() )")
	if err != nil {
		log.Fatal("error preparing query :: ", err)
	}

	_, err = stmt.Exec(u.Username, u.Password, u.Active, u.CreatedId, u.ModifiedId)
	if err != nil {
		log.Fatal("error executing query :: ", err)
	}

}

func EditUser(id string, u User) {

	stmt, err := db.Prepare("UPDATE users SET username=?, active=?, modified_id=?, modified=NOW() WHERE id=?")
	if err != nil {
		log.Fatal("error preparing query :: ", err)
	}

	_, err = stmt.Exec(u.Username, u.Active, u.ModifiedId, id)
	if err != nil {
		log.Fatal("error executing query :: ", err)
	}

}
