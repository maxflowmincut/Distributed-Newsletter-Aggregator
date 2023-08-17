package db

import (
	"database/sql"
)

type Database interface {
    ConnectDB(filename string) (*sql.DB, error)
    InitializeDB(db *sql.DB, schema string) error
    CloseDB(db *sql.DB)
    CreateUser(user User) (sql.Result, error)
    EmailExists(email string) (bool, error)
    SaveArticle(article Article) (sql.Result, error)
    ClearPreviousArticles() error
    GetAllUsers() ([]User, error)
    GetArticlesMatchingPreferences(preferences string, limit int) ([]Article, error)
    GetRandomArticles(limit int) ([]Article, error)
}