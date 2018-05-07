package lib

import (
	"log"
)

type Task struct {
	Id          int
	ProjectId   string
	ProjectName string
	Name        string
	Description string
	Start       string
	End         string
	Duration    string
	Created     string
	CreatedId   string
	Modified    string
	ModifiedId  string
}

func GetTasks() []Task {

	query := `SELECT
              t.id as id,
              p.id AS project_id,
              p.name AS project_name,
              t.name AS task_name,
              t.description AS description,
              t.start AS start, t.end AS end,
              timediff(t.end , t.start) AS duration
            FROM tasks AS t
            INNER JOIN projects AS p ON t.project_id = p.id`

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal("error occurred while executing select query :: ", err)
	}

	tasks := []Task{}
	for rows.Next() {
		var id int
		var project_id string
		var project_name string
		var task_name string
		var description string
		var start string
		var end string
		var duration string

		err = rows.Scan(
			&id,
			&project_id,
			&project_name,
			&task_name,
			&description,
			&start,
			&end,
			&duration,
		)

		task := Task{
			Id:          id,
			Name:        task_name,
			ProjectId:   project_id,
			ProjectName: project_name,
			Description: description,
			Start:       start,
			End:         end,
			Duration:    duration,
		}
		tasks = append(tasks, task)
	}

	return tasks

}

func GetTask(uid string) Task {

	var id int
	var project_id string
	var project_name string
	var task_name string
	var description string
	var start string
	var end string
	var duration string

	query := `SELECT
              t.id as id,
              p.id AS project_id,
              p.name AS project_name,
              t.name AS task_name,
              t.description AS description,
              t.start AS start, t.end AS end,
              timediff(t.end , t.start) AS duration
            FROM tasks AS t
            INNER JOIN projects AS p ON t.project_id = p.id
            WHERE p.id=?
            LIMIT 1`

	err := db.QueryRow(query, uid).Scan(
		&id,
		&project_id,
		&project_name,
		&task_name,
		&description,
		&start,
		&end,
		&duration,
	)

	if err != nil {
		log.Fatal("Error while reading subject", err)
	}

	task := Task{
		Id:          id,
		Name:        task_name,
		ProjectId:   project_id,
		ProjectName: project_name,
		Description: description,
		Start:       start,
		End:         end,
		Duration:    duration,
	}

	return task

}

func InsertTask(u Task) {

	query := `INSERT INTO tasks
            ( name, project_id, description, start, end, created_id, modified_id, created, modified  )
            VALUES
            ( ?, ?, ?, ?, ?, ?, ?, NOW(), NOW() )`

	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal("error preparing query :: ", err)
	}

	_, err = stmt.Exec(u.Name, u.ProjectId, u.Description, u.Start, u.End, u.CreatedId, u.ModifiedId)
	if err != nil {
		log.Fatal("error executing query :: ", err)
	}

}

func EditTask(id string, u Task) {

	query := `UPDATE tasks
            SET
              name=?,
              project_id=?,
              description=?,
              start=?,
              end=?,
              modified_id=?,
              modified=NOW()
            WHERE id=?`

	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal("error preparing query :: ", err)
	}

	_, err = stmt.Exec(u.Name, u.ProjectId, u.Description, u.Start, u.End, u.ModifiedId, id)
	if err != nil {
		log.Fatal("error executing query :: ", err)
	}

}
