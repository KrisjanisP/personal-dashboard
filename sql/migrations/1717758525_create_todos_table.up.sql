CREATE TABLE todos (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,
    job TEXT NOT NULL,
    name TEXT NOT NULL,
    description TEXT,
    deadline DATE,
    priority INTEGER DEFAULT 1,
    finished BOOLEAN DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY(user_id) REFERENCES users(id)
);
