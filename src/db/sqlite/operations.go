package sqlite

import (
	"database/sql"
	"time"
	"log"
    "strings"

    "newsletter-aggregator/src/curation"
	"newsletter-aggregator/src/db"
)

func CategoriesToStrings(categories []curation.Category) []string {
    strCategories := make([]string, len(categories))
    for i, category := range categories {
        strCategories[i] = category.String()
    }
    return strCategories
}

func CreateUser(db *SQLiteDB, name, email, preferences, sendTime string) (sql.Result, error) {
	statement, err := db.DB.Prepare("INSERT INTO User (Name, Email, Preferences, SendTime) VALUES (?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}
	defer statement.Close()
	return statement.Exec(name, email, preferences, sendTime)
}

func EmailExists(db *SQLiteDB, email string) (bool, error) {
	var count int
	err := db.DB.QueryRow("SELECT COUNT(*) FROM User WHERE Email = ?", email).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func SaveArticle(db *SQLiteDB, article db.Article) (sql.Result, error) {
    article.Date = time.Now().Format("2006-01-02")
    statement, err := db.DB.Prepare("INSERT INTO Article (Date, Category, Title, Description, Link) VALUES (?, ?, ?, ?, ?)")
    if err != nil {
        return nil, err
    }
    defer statement.Close()
    
	categoryStr := strings.Join(article.Category, ",")


    return statement.Exec(article.Date, categoryStr, article.Title, article.Description, article.Link)
}

func ClearPreviousArticles(db *SQLiteDB) error {
	statement, err := db.DB.Prepare("DELETE FROM Article")
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer statement.Close()
	_, err = statement.Exec()
	return err
}

func UpdateArticleCategories(db *SQLiteDB) error {
    rows, err := db.DB.Query("SELECT ArticleID, Description FROM Article WHERE Category IS NULL OR Category = ''")
    if err != nil {
        return err
    }
    defer rows.Close()

    for rows.Next() {
        var articleID int
        var description string
        err := rows.Scan(&articleID, &description)
        if err != nil {
            return err
        }
		categories := curation.Categorize(description)
		strCategories := CategoriesToStrings(categories)
		categoryStr := strings.Join(strCategories, ",")

        
        _, err = db.DB.Exec("UPDATE Article SET Category = ? WHERE ArticleID = ?", categoryStr, articleID)
        if err != nil {
            return err
        }
    }
    return rows.Err()
}

func (s *SQLiteDB) GetAllUsers() ([]db.User, error) {
    query := "SELECT UserID, Name, Email, Preferences, SendTime FROM User"

    rows, err := s.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var users []db.User
    for rows.Next() {
        var user db.User
        if err := rows.Scan(&user.UserID, &user.Name, &user.Email, &user.Preferences, &user.SendTime); err != nil {
            return nil, err
        }
        users = append(users, user)
    }
    if err := rows.Err(); err != nil {
        return nil, err
    }
    return users, nil
}

func (s *SQLiteDB) GetArticlesMatchingPreferences(preferences string, limit int) ([]db.Article, error) {
    query := "SELECT ArticleID, Date, Category, Title, Description, Link FROM Articles WHERE Category IN (?) LIMIT ?"

    prefs := strings.Join(strings.Split(preferences, ","), "','")
    rows, err := s.Query(query, prefs, limit)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var articles []db.Article
    for rows.Next() {
        var article db.Article
        var categoryStr string

        if err := rows.Scan(&article.ArticleID, &article.Date, &categoryStr, &article.Title, &article.Description, &article.Link); err != nil {
            return nil, err
        }
        
        article.Category = strings.Split(categoryStr, ",")
        articles = append(articles, article)
    }
    if err := rows.Err(); err != nil {
        return nil, err
    }
    return articles, nil
}

func (s *SQLiteDB) GetRandomArticles(limit int) ([]db.Article, error) {
    query := "SELECT ArticleID, Date, Category, Title, Description, Link FROM Articles ORDER BY RANDOM() LIMIT ?"

    rows, err := s.Query(query, limit)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var articles []db.Article
    for rows.Next() {
        var article db.Article
        var categoryStr string
        if err := rows.Scan(&article.ArticleID, &article.Date, &categoryStr, &article.Title, &article.Description, &article.Link); err != nil {
            return nil, err
        }
        
        article.Category = strings.Split(categoryStr, ",")
        articles = append(articles, article)
    }
    if err := rows.Err(); err != nil {
        return nil, err
    }
    return articles, nil
}