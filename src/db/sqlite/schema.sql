CREATE TABLE User (
    UserID INTEGER PRIMARY KEY AUTOINCREMENT,
    Name TEXT,
    Email TEXT UNIQUE,
    Preferences TEXT,
    SendTime TEXT
);

CREATE TABLE Article (
    ArticleID INTEGER PRIMARY KEY AUTOINCREMENT,
    Date TEXT,
    Category TEXT,
    Title TEXT,
    Description TEXT,
    Link TEXT
);