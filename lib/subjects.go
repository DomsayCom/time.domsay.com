package lib

import (
	"log"
)

type Subject struct {
	Id         int
	Name       string
	VAT        string
	Address    string
	Created    string
	CreatedId  string
	Modified   string
	ModifiedId string
}

func GetSubjectsList() map[string]string {

	rows, err := db.Query("SELECT id, name FROM subjects")
	if err != nil {
		log.Fatal("error occurred while executing select query :: ", err)
	}

	subjects := make(map[string]string)

	for rows.Next() {
		var id string
		var name string
		err = rows.Scan(&id, &name)
		subjects[id] = name
	}

	return subjects

}

func GetSubjects() []Subject {

	rows, err := db.Query("SELECT id, name, vat, address  FROM subjects")
	if err != nil {
		log.Fatal("error occurred while executing select query :: ", err)
	}

	subjects := []Subject{}
	for rows.Next() {
		var id int
		var name string
		var vat string
		var address string
		err = rows.Scan(&id, &name, &vat, &address)
		subject := Subject{Id: id, Name: name, VAT: vat, Address: address}
		subjects = append(subjects, subject)
	}

	return subjects

}

func GetSubject(uid string) Subject {

	var id int
	var name string
	var vat string
	var address string

	err := db.QueryRow("SELECT id, name, vat, address  FROM subjects WHERE id=? LIMIT 1", uid).Scan(&id, &name, &vat, &address)

	if err != nil {
		log.Fatal("Error while reading subject", err)
	}

	subject := Subject{Id: id, Name: name, VAT: vat, Address: address}

	return subject

}

func InsertSubject(u Subject) {

	stmt, err := db.Prepare("INSERT INTO subjects ( name, vat, address, created_id, modified_id, created, modified  ) VALUES ( ?, ?, ?, ?, ?, NOW(), NOW() )")
	if err != nil {
		log.Fatal("error preparing query :: ", err)
	}

	_, err = stmt.Exec(u.Name, u.VAT, u.Address, u.CreatedId, u.ModifiedId)
	if err != nil {
		log.Fatal("error executing query :: ", err)
	}

}

func EditSubject(id string, u Subject) {

	stmt, err := db.Prepare("UPDATE subjects SET name=?, vat=?, address=?, modified_id=?, modified=NOW() WHERE id=?")
	if err != nil {
		log.Fatal("error preparing query :: ", err)
	}

	_, err = stmt.Exec(u.Name, u.VAT, u.Address, u.ModifiedId, id)
	if err != nil {
		log.Fatal("error executing query :: ", err)
	}

}
