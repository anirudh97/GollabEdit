/* Create users table */
CREATE TABLE IF NOT EXISTS users(
    email VARCHAR(255) PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    hashPassword VARCHAR(255) NOT NULL
);

/* Create files table */
CREATE TABLE IF NOT EXISTS files(
    id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
    filename VARCHAR(255) NOT NULL,
    location VARCHAR(255) NOT NULL,
    owner VARCHAR(255) NOT NULL,
    isUploaded BOOLEAN NOT NULL,
    fileSize INT NOT NULL,
    createdAt VARCHAR(255) NOT NULL,
    updatedAt VARCHAR(255) NOT NULL,
    FOREIGN KEY owner REFERENCES users(email) ON DELETE CASCADE
);

/* Create shared files table */
CREATE TABLE sharedFiles (
    shareId INT AUTO_INCREMENT PRIMARY KEY NOT NULL,
    fileId INT NOT NULL,
    sharedWithUserEmail VARCHAR(255) NOT NULL,
    sharedByUserEmail VARCHAR(255) NOT NULL,
    permission VARCHAR(50),
    sharedAt VARCHAR(255) NOT NULL,
    FOREIGN KEY (fileId) REFERENCES files(id) ON DELETE CASCADE,
    FOREIGN KEY (sharedWithUserEmail) REFERENCES users(email) ON DELETE CASCADE,
    FOREIGN KEY (sharedByUserEmail) REFERENCES users(email) ON DELETE CASCADE
);