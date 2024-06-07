-- Drop the table if it exists
DROP TABLE work_categories;

-- Create the table with columns
CREATE TABLE work_categories (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INT,
    abbreviation VARCHAR(255),
    description VARCHAR(255),
    FOREIGN KEY (user_id) REFERENCES users(id)
);