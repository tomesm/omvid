package mysql

import (
	"database/sql"

	"github.com/tomesm/virtd/pkg/models"
)

// CourseModel wraps a sql.DB connection pool
type CourseModel struct {
	DB *sql.DB
}

// Insert a course into the database
func (m *CourseModel) Insert(title, content, expires string) (int, error) {
	return 0, nil
}

// Get a course based on its ID
func (m *CourseModel) Get(id int) (*models.Course, error) {
	return nil, nil
}

// Latest 10 latest
func (m *CourseModel) Latest() ([]*models.Course, error) {
	return nil, nil
}
