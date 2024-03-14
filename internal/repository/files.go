package repository

import (
	"context"

	"github.com/anirudh97/GollabEdit/internal/database"
	"github.com/anirudh97/GollabEdit/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func UpdateDocument(wd *model.WootDocument) error {
	collection := database.MongoDB.Database("gollabedit").Collection("documents")

	// Prepare the filter and replacement document
	filter := bson.M{"id": wd.Id}
	replacement := wd

	// Prepare the upsert option
	opts := options.Replace().SetUpsert(true)

	// Perform the replace operation with upsert true
	_, err := collection.ReplaceOne(context.Background(), filter, replacement, opts)
	if err != nil {
		return err
	}

	return nil
}
func ReterieveDocument(d string) (*model.WootDocument, error) {
	var result model.WootDocument

	collection := database.MongoDB.Database("gollabedit").Collection("documents")

	err := collection.FindOne(context.Background(), bson.M{"id": d}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return &result, nil
}

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
