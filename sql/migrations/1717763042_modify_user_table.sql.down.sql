-- Drop the table if it exists
DROP TABLE work_categories;

-- Create the table with columns
CREATE TABLE work_categories (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL UNIQUE
);
