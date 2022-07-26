package mysql

import (
	"database/sql"

	"github.com/AndreyUA/go-snippetbox/pkg/models"
)

type SnippetModel struct {
	DB *sql.DB
}

// Insert: create a new snippet
func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	stmt := `INSERT INTO snippets (title, content, created, expires)
    VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	result, err := m.DB.Exec(stmt, title, content, expires)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return int(id), nil
}

// Get: get snippet by ID
func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	return nil, nil
}

// Lates: get 10 last snippets
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
