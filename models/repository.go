package models

import (
	"database/sql"
	"fmt"
	"time"
)

// Repository struct for db connection
type Repository struct {
	Conn *sql.DB
}

func (repository *Repository) GetArticles() ([]*Article, error) {
	rows, err := repository.Conn.Query("SELECT a.id, a.title, a.header, a.authors, a.created_on, " +
		"a.updated_on" +
		"\nFROM article a ")
	if err != nil {
		return nil, fmt.Errorf("could not prepare query: %v", err)
	}
	var articles []*Article
	var id uint64
	var title, header, authors string
	var createdOn, updatedOn time.Time
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &title, &header, &authors, &createdOn, &updatedOn)
		if err == sql.ErrNoRows {
			return nil, nil
		}
		if err != nil {
			return nil, fmt.Errorf("could not get articles : %v", err)
		}
		article := &Article{
			ID:        id,
			Title:     title,
			Header:    header,
			Authors:   authors,
			CreatedOn: createdOn,
			UpdatedOn: updatedOn,
		}
		articles = append(articles, article)
	}

	err = rows.Err()
	if err != nil {
		fmt.Print(err)
	}

	return articles, nil
}

func (repository *Repository) GetArticleById(id uint64) (*Article, error) {
	
	row := repository.Conn.QueryRow("SELECT a.id, a.title, a.header, a.authors, a.created_on, " +
	"a.updated_on FROM article a WHERE id=(?)", id)
	var title, header, authors string
	var createdOn, updatedOn time.Time
	switch err := row.Scan(&id, &title, &header, &authors, &createdOn, &updatedOn); err {
	case sql.ErrNoRows:
		return nil, nil
	case nil:
		article := &Article{
			ID:        id,
			Title:     title,
			Header:    header,
			Authors:   authors,
			CreatedOn: createdOn,
			UpdatedOn: updatedOn,
		}
		return article, nil
	default:
		return nil, err
	}
}
