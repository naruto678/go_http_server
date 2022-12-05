package mysql

import (
	"database/sql"
	"github.com/server-practice/pkg/models"
)

type SnippetModel struct {
	DB *sql.DB
}

func (s *SnippetModel) ToString() string {
	return "somet random shit"
}

func (s *SnippetModel) Insert(title, content, expires string) (int, error) {
	stmt := `insert into snippets (title, content, created, expires) values ( ? , ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`
	result, err := s.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (s *SnippetModel) Get(id int) (*models.Snippet, error) {
	stmt := `select id, title, content, created, expires from snippets where id = ?`
	row := s.DB.QueryRow(stmt, id)

	snippetModel := &models.Snippet{}
	err := row.Scan(&snippetModel.ID, &snippetModel.Title, &snippetModel.Content, &snippetModel.Created, &snippetModel.Expires)
	if err == sql.ErrNoRows {
		return nil, models.ErrNoRecord
	}
	return snippetModel, nil
}

func (s *SnippetModel) Latest() ([]*models.Snippet, error) {
	stmt := `select id, title, content,created, expires from snippets where expires > UTC_TIMESTAMP() order by created desc limit 10`
	rows, err := s.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var snippetModels []*models.Snippet

	for rows.Next() {
		snippet := &models.Snippet{}
		err := rows.Scan(&snippet.ID, &snippet.Title, &snippet.Content, &snippet.Created, &snippet.Expires)
		if err != nil {
			return nil, err
		}
		snippetModels = append(snippetModels, snippet)

	}
	return snippetModels, nil
}

func (s *SnippetModel) GetAll() ([]*models.Snippet, error) {
	stmt := `select id, title, content, created, expires from snippets`
	rows, err := s.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var snippetModel []*models.Snippet

	for rows.Next() {
		snippet := &models.Snippet{}
		err := rows.Scan(&snippet.ID, &snippet.Title, &snippet.Content, &snippet.Created, &snippet.Expires)
		if err != nil {
			return nil, err
		}
		snippetModel = append(snippetModel, snippet)
	}
	return snippetModel, nil
}
