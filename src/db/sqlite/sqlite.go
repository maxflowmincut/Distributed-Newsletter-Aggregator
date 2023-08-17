package sqlite

import (
    "database/sql"
    "newsletter-aggregator/src/db"
)

type SQLiteDB struct {
    *sql.DB
}

func NewSQLiteDB(filename string) (*SQLiteDB, error) {
    db, err := ConnectDB(filename)
    if err != nil {
        return nil, err
    }
    return &SQLiteDB{db}, nil
}

func (s *SQLiteDB) Close() {
    CloseDB(s)
}

func (s *SQLiteDB) CreateUser(name, email, preferences, sendTime string) (sql.Result, error) {
    return CreateUser(s, name, email, preferences, sendTime)
}

func (s *SQLiteDB) EmailExists(email string) (bool, error) {
    return EmailExists(s, email)
}

func (s *SQLiteDB) SaveArticle(article db.Article) (sql.Result, error) {
    return SaveArticle(s, article)
}

func (s *SQLiteDB) ClearPreviousArticles() error {
    return ClearPreviousArticles(s)
}