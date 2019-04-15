// Reusable service/data access layer
package models

import (
	"errors"
	"time"
)
// ErrNoRecord is a database service error informing about no matching records
var ErrNoRecord = errors.New("models: no matching record found")
// Course struct corresponds to database table
type Course struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}
