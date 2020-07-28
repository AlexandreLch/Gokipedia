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

func (repository *Repository) GetArticleByID(id uint64) (*Article, error) {
	row := repository.Conn.QueryRow("SELECT a.id, a.title, a.header, a.authors, a.created_on, "+
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

func (repository *Repository) SaveArticle(article *Article) error {
	stmt, err := repository.Conn.Prepare("INSERT INTO article(title, header, authors, created_on," +
		"updated_on) VALUES(?,?,?,?,?)")
	if err != nil {
		return err
	}

	article.CreatedOn = time.Now()
	article.UpdatedOn = time.Now()

	res, errExec := stmt.Exec(article.Title, article.Header, article.Authors, article.CreatedOn,
		article.UpdatedOn)
	if errExec != nil {
		return fmt.Errorf("could not exec stmt: %v", errExec)
	}

	lastInsertedID, errInsert := res.LastInsertId()

	if errInsert != nil {
		return fmt.Errorf("could not retrieve last inserted ID: %v", errInsert)
	}

	article.ID = uint64(lastInsertedID)

	return nil
}

func (repository *Repository) updateArticleByID(article *Article) error {
	stmt, err := repository.Conn.Prepare("UPDATE article SET title=(?), header=(?), " +
		"authors=(?), updated_on=(?) WHERE id=(?)")
	if err != nil {
		return err
	}

	article.UpdatedOn = time.Now()

	_, errExec := stmt.Exec(article.Title, article.Header, article.Authors, article.UpdatedOn,
		article.ID)

	if errExec != nil {
		return errExec
	}

	return nil
}

func (repository *Repository) deleteArticle(id uint64) (uint64, error) {

	res, err := repository.Conn.Exec("DELETE FROM article WHERE id=(?)", id)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return uint64(rowsAffected), nil
}
