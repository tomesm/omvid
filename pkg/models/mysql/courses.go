package mysql

import (
	"database/sql"

	"github.com/tomesm/virtd/pkg/models"
)

// CourseModel wraps a sql.DB connection pool
type CourseModel struct {
	DB *sql.DB
}

// Insert a course into the database and return its ID
func (m *CourseModel) Insert(title, content, expires string) (int, error) {
	statement := `INSERT INTO courses (title, content, created, expires)
	VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	res, err := m.DB.Exec(statement, title, content, expires)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

// Get a course based on its ID
func (m *CourseModel) Get(id int) (*models.Course, error) {
	statement := `SELECT id, title, content, created, expires FROM courses
	WHERE expires > UTC_TIMESTAMP() AND id = ?`

	row := m.DB.QueryRow(statement, id)
	c := &models.Course{}

	err := row.Scan(&c.ID, &c.Title, &c.Content, &c.Created, &c.Expires)
	if err == sql.ErrNoRows {
		return nil, models.ErrNoRecord
	} else if err != nil {
		return nil, err
	}

	return c, nil
}

// Latest 10 latest
func (m *CourseModel) Latest() ([]*models.Course, error) {
	return nil, nil
}
