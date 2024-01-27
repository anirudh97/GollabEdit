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

func GetFileId(f string, l string, o string) (string, error) {
	query := "SELECT id FROM files WHERE owner = ? AND filename = ? AND location = ?"

	var fileId string

	err := database.DB.Get(&fileId, query, o, f, l)

	return fileId, err
}

func ShareFile(f *model.SharedFile) error {
	query := "INSERT INTO sharedFiles (fileId, sharedWithUserEmail, sharedByUserEmail, permission, sharedAt) VALUES (?, ?, ?, ?, ?)"

	_, err := database.DB.Exec(query, f.FileId, f.SharedWithEmail, f.SharedByEmail, f.Permission, f.SharedAt)

	return err
}

func CheckShared(sharedByEmail string, sharedWithEmail string, fileId string) (bool, error) {
	query := "SELECT EXISTS(SELECT 1 FROM sharedFiles WHERE sharedByUserEmail = ? AND sharedWithUserEmail = ? AND fileId = ?)"

	var exists bool

	err := database.DB.Get(&exists, query, sharedByEmail, sharedWithEmail, fileId)

	return exists, err
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
