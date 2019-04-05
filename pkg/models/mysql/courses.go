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
	var err error
	defer func() {
		if err != nil {
			return 0, err
		}
	}()

	statement := `INSERT INTO courses (title, content, created, expires)
	VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	res, err := m.DB.Exec(statement, title, content, expires)
	if err != nil {
		return
	}

	id, err := res.LastInsertId()
	if err != nil {
		return
	}

	return int(id), nil
}

// Get a course based on its ID
func (m *CourseModel) Get(id int) (*models.Course, error) {
	return nil, nil
}

// Latest 10 latest
func (m *CourseModel) Latest() ([]*models.Course, error) {
	return nil, nil
}
