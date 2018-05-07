package lib

import (
	"log"
)

type Counter struct {
	Users    string
	Subjects string
	Projects string
	Tasks    string
}

func GetCounter() Counter {

	var users string
	var subjects string
	var projects string
	var tasks string

	query := `SELECT
            ( SELECT COUNT(id) FROM   users ) AS users,
            ( SELECT COUNT(id) FROM   subjects ) AS subjects,
            ( SELECT COUNT(id) FROM   projects ) AS projects,
            ( SELECT COUNT(id) FROM   tasks ) AS tasks`

	err := db.QueryRow(query).Scan(&users, &subjects, &projects, &tasks)

	if err != nil {
		log.Fatal("Error while reading counter", err)
	}

	project := Counter{Users: users, Subjects: subjects, Projects: projects, Tasks: tasks}

	return project

}

type ProjectDetails struct {
	Name    string
	Hours   string
	Minutes string
	Tasks   map[int]TaskDetails
}

type TaskDetails struct {
	Name        string
	Description string
	Start       string
	End         string
	Duration    string
}

func GetProjectsDetails() map[string]ProjectDetails {

	query := `SELECT
              p.name AS project,
              t.id  AS tid,
              t.name AS task,
              t.description,
              t.start,
              t.end,
              timediff(t.end , t.start) AS duration
            FROM tasks AS t
            INNER JOIN projects AS p ON t.project_id = p.id
            WHERE YEAR(t.start) = YEAR(NOW())
            AND   MONTH(t.start) = MONTH(NOW())
            ORDER BY t.start`

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal("error occurred while executing select query :: ", err)
	}

	result := make(map[string]ProjectDetails)
	projects := make(map[string]ProjectDetails)
	tasks := make(map[string]map[int]TaskDetails)

	for rows.Next() {
		var project string
		var tid int
		var task string
		var description string
		var start string
		var end string
		var duration string

		err = rows.Scan(&project, &tid, &task, &description, &start, &end, &duration)

		t := TaskDetails{
			Name:        task,
			Description: description,
			Start:       start,
			End:         end,
			Duration:    duration,
		}

		if _, ok := tasks[project]; ok {
			tasks[project][tid] = t
		} else {
			tasks[project] = make(map[int]TaskDetails)
			tasks[project][tid] = t
		}

		projects[project] = ProjectDetails{Name: project}

	}

	for key, details := range projects {

		result[key] = ProjectDetails{Name: details.Name, Tasks: tasks[key]}

	}

	return result

}
