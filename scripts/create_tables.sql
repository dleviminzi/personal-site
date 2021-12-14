CREATE TABLE experience_items (
    id INTEGER PRIMARY KEY,
    item_type TEXT,
    title TEXT,
    description TEXT,
    start_date DATETIME,
    end_date DATETIME
);

CREATE TABLE projects (
    title TEXT PRIMARY KEY,
    description TEXT,
    github_link TEXT,
    status TEXT,
    start_date DATETIME,
    end_date DATETIME
);

CREATE TABLE notes (
    title TEXT PRIMARY KEY,
    topic TEXT,
    content TEXT
);

CREATE TABLE photos (
    title TEXT PRIMARY KEY,
    description TEXT,
    location TEXT,
    created DATETIME,
    photo_path TEXT
)