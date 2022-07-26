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
	return 0, nil
}

// Get: get snippet by ID
func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	return nil, nil
}

// Lates: get 10 last snippets
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
