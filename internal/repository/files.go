package repository

import (
	"github.com/anirudh97/GollabEdit/internal/database"
	"github.com/anirudh97/GollabEdit/internal/model"
)

func CheckForFileExistence(f string, l string, o string) (bool, error) {
	query := "SELECT EXISTS(SELECT 1 FROM files WHERE owner = ? AND filename = ? AND location = ?)"

	var exists bool

	err := database.DB.Get(&exists, query, o, f, l)

	return exists, err
}

func CreateFile(f *model.File) error {
	query := "INSERT INTO files (filename, location, owner, isUploaded, fileSize, createdAt, updatedAt) VALUES (?, ?, ?, ?, ?, ?, ?)"

	_, err := database.DB.Exec(query, f.Filename, f.Location, f.Owner, f.IsUploaded, f.FileSize, f.CreatedAt, f.UpdatedAt)

	return err
}

func CheckIfUploaded(f string, l string, o string) (bool, error) {
	query := "SELECT EXISTS(SELECT 1 FROM files WHERE owner = ? AND filename = ? AND location = ? AND isUploaded = 1)"

	var uploaded bool

	err := database.DB.Get(&uploaded, query, o, f, l)

	return uploaded, err
}

func UpdateUploadedStatus(f string, l string, o string) error {
	query := "UPDATE files SET isUploaded = 1 WHERE owner = ? AND filename = ? AND location = ?"

	_, err := database.DB.Exec(query, o, f, l)

	return err
}
