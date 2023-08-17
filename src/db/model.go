package db

type User struct {
    UserID      int    `db:"UserID"`
    Name        string `db:"Name"`
    Email       string `db:"Email"`
    Preferences string `db:"Preferences"`
    SendTime    string `db:"SendTime"`
}

type Article struct {
    ArticleID   int      `db:"ArticleID"`
    Date        string   `db:"Date"`
    Category    []string `db:"Category"`
    Title       string   `db:"Title"`
    Description string   `db:"Description"`
    Link        string   `db:"Link"`
}