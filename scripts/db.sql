/* Create users table */
CREATE TABLE IF NOT EXISTS users(
    email VARCHAR(255) PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    hashPassword VARCHAR(255) NOT NULL
);

