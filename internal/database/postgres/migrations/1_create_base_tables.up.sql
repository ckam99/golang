/* Create users table */
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY,
    name VARCHAR(60) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE
);
/* Create posts table */
CREATE TABLE IF NOT EXISTS posts (
    id UUID PRIMARY KEY,
    title TEXT NOT NULL,
    content TEXT NOT NULL,
    votes INT NOT NULL,
    user_id UUID REFERENCES users(id) ON DELETE CASCADE
);
/* Create comments table */
CREATE TABLE IF NOT EXISTS comments (
    id UUID PRIMARY KEY,
    content TEXT NOT NULL,
    votes INT NOT NULL,
    post_id UUID REFERENCES posts(id) ON DELETE CASCADE
);

