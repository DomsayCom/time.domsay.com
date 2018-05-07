package lib

import (
	"log"
)

type Project struct {
	Id          int
	SubjectId   string
	Name        string
	SubjectName string
	Created     string
	CreatedId   string
	Modified    string
	ModifiedId  string
}

func GetProjectsList() map[string]string {

	rows, err := db.Query("SELECT id, name FROM projects")
	if err != nil {
		log.Fatal("error occurred while executing select query :: ", err)
	}

	projects := make(map[string]string)

	for rows.Next() {
		var id string
		var name string
		err = rows.Scan(&id, &name)
		projects[id] = name
	}

	return projects

}

func GetProjects() []Project {

	rows, err := db.Query("SELECT p.id AS id, p.name as project_name, p.subject_id AS subject_id, s.name AS subject_name FROM projects AS p INNER JOIN subjects AS s ON p.subject_id = s.id")
	if err != nil {
		log.Fatal("error occurred while executing select query :: ", err)
	}

	projects := []Project{}
	for rows.Next() {
		var id int
		var project_name string
		var subject_id string
		var subject_name string
		err = rows.Scan(&id, &project_name, &subject_id, &subject_name)
		project := Project{Id: id, Name: project_name, SubjectId: subject_id, SubjectName: subject_name}
		projects = append(projects, project)
	}

	return projects

}

func GetProject(uid string) Project {

	var id int
	var project_name string
	var subject_id string
	var subject_name string

	err := db.QueryRow("SELECT p.id AS id, p.name as project_name, p.subject_id AS subject_id, s.name AS subject_name FROM projects AS p INNER JOIN subjects AS s ON p.subject_id = s.id WHERE p.id=? LIMIT 1", uid).Scan(&id, &project_name, &subject_id, &subject_name)

	if err != nil {
		log.Fatal("Error while reading subject", err)
	}

	project := Project{Id: id, Name: project_name, SubjectId: subject_id, SubjectName: subject_name}

	return project

}

func InsertProject(u Project) {

	stmt, err := db.Prepare("INSERT INTO projects ( name, subject_id, created_id, modified_id, created, modified  ) VALUES ( ?, ?, ?, ?, NOW(), NOW() )")
	if err != nil {
		log.Fatal("error preparing query :: ", err)
	}

	_, err = stmt.Exec(u.Name, u.SubjectId, u.CreatedId, u.ModifiedId)
	if err != nil {
		log.Fatal("error executing query :: ", err)
	}

}

func EditProject(id string, u Project) {

	stmt, err := db.Prepare("UPDATE projects SET name=?, subject_id=?, modified_id=?, modified=NOW() WHERE id=?")
	if err != nil {
		log.Fatal("error preparing query :: ", err)
	}

	_, err = stmt.Exec(u.Name, u.SubjectId, u.ModifiedId, id)
	if err != nil {
		log.Fatal("error executing query :: ", err)
	}

}
