/* Create users table */
CREATE TABLE IF NOT EXISTS users(
    email VARCHAR(255) PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    hashPassword VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS files(
    id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
    filename VARCHAR(255) NOT NULL,
    location VARCHAR(255) NOT NULL,
    owner VARCHAR(255) NOT NULL,
    isUploaded BOOLEAN NOT NULL,
    fileSize INT NOT NULL,
    createdAt VARCHAR(255) NOT NULL,
    updatedAt VARCHAR(255) NOT NULL
);

