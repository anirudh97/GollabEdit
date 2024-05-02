package service

import (
	"fmt"
	"log"
	"time"

	"github.com/anirudh97/GollabEdit/internal/model"
	"github.com/anirudh97/GollabEdit/internal/repository"
)

type CreateFileRequest struct {
	Filename string `json:"filename"`
	Location string `json:"location"`
	Owner    string `json:"owner"`
}
type CreateFileResponse struct {
	Filename  string `json:"filename"`
	Location  string `json:"location"`
	CreatedAt string
}
type OpenFileResponse struct {
	Filename        string             `json:"filename"`
	Location        string             `json:"location"`
	DocumentID      string             `json:"documentId"`
	DocumentContent model.WootDocument `json:"documentContent"`
	WebSocketURL    string             `json:"webSocketUrl"`
}
type OpenFileRequest struct {
	Filename string `json:"filename"`
	Location string `json:"location"`
	Owner    string `json:"owner"`
}
type ShareFileRequest struct {
	Filename        string `json:"filename"`
	Location        string `json:"location"`
	SharedByEmail   string `json:"sharedByEmail"`
	SharedWithEmail string `json:"sharedWithEmail"`
	Permission      string `json:"permission"`
}
type ShareFileResponse struct {
	Filename        string
	Location        string
	Permission      string
	SharedWithEmail string
}

type CharacterRequest struct {
	RequestType string `json:"typeRequest"`
	DocumentID  string `json:"documentId"`
	UserID      string `json:"userId"`
	PrevID      string `json:"prevId"`
	NextID      string `json:"nextId"`
	Character   string `json:"character"`
	CharID      string `json:"charId"`
	Visible     bool   `json:"visible"`
}

// type InsertCharacterRequest struct {
// 	DocumentID string `json:"documentId"`
// 	UserID     string `json:"userId"`
// 	PrevID     string `json:"prevId"`
// 	NextID     string `json:"nextId"`
// 	Character  string `json:"character"`
// 	CharID     string `json:"charId"`
// 	Visible    bool   `json:"visible"`
// }

// type DeleteCharacterRequest struct {
// 	DocumentID string `json:"documentId"`
// 	UserID     string `json:"userId"`
// 	CharID     string `json:"charId"`
// }

func InsertCharacter(r *CharacterRequest) (*model.WootDocument, error) {

	wd, err := repository.ReterieveDocument(r.DocumentID)
	if err != nil {
		return nil, err
	}

	if wd == nil {
		// Make a new woot document
		log.Println("Document not present, creating new one")
		wd, err = model.CreateNewDocument(r.DocumentID)
		if err != nil {
			return nil, err
		}
	}

	wd.InsertCharacter(r.Character, r.CharID, r.PrevID, r.NextID, r.Visible)
	log.Println("After insert charater!!!")
	log.Println(wd)
	err = repository.UpdateDocument(wd)
	if err != nil {
		return nil, err
	}

	return wd, nil

}
func DeleteCharacter(r *CharacterRequest) (*model.WootDocument, error) {

	wd, err := repository.ReterieveDocument(r.DocumentID)
	if err != nil {
		return nil, err
	}

	// Handle case when document not present

	wd.DeleteCharacter(r.CharID)

	err = repository.UpdateDocument(wd)
	if err != nil {
		return nil, err
	}

	return wd, nil

}
func ShareFile(r *ShareFileRequest) (*ShareFileResponse, error) {
	status, err := repository.CheckForFileExistence(r.Filename, r.Location, r.SharedByEmail)
	if err != nil {
		return nil, err
	}
	if !status {
		return nil, ErrFileDoesNotExist
	}

	sharedByEmailStatus, sharedByEmailErr := repository.CheckForUserExistence(r.SharedByEmail)
	if sharedByEmailErr != nil {
		return nil, sharedByEmailErr
	}
	if !sharedByEmailStatus {
		return nil, ErrUserDoesNotExist
	}

	sharedWithEmailStatus, sharedWithEmailErr := repository.CheckForUserExistence(r.SharedWithEmail)
	if sharedWithEmailErr != nil {
		return nil, sharedWithEmailErr
	}
	if !sharedWithEmailStatus {
		return nil, ErrUserDoesNotExist
	}

	loc, err := time.LoadLocation("UTC")
	if err != nil {
		fmt.Println("Error loading location:", err)
		return nil, err
	}

	fileId, fileIdErr := repository.GetFileId(r.Filename, r.Location, r.SharedByEmail)
	if fileIdErr != nil {
		return nil, fileIdErr
	}

	shareStatus, shareStatusErr := repository.CheckShared(r.SharedByEmail, r.SharedWithEmail, fileId)
	if shareStatusErr != nil {
		return nil, shareStatusErr
	}
	if shareStatus {
		return nil, ErrAlreadyShared
	}

	currentTime := time.Now().In(loc)

	sharedFile := &model.SharedFile{
		Filename:        r.Filename,
		Location:        r.Location,
		Permission:      r.Permission,
		SharedWithEmail: r.SharedWithEmail,
		SharedByEmail:   r.SharedByEmail,
		SharedAt:        currentTime.Format("2006-01-02 15:04:05"),
		FileId:          fileId,
	}

	shareFileErr := repository.ShareFile(sharedFile)
	if shareFileErr != nil {
		return nil, shareFileErr
	}

	sfr := &ShareFileResponse{
		Filename:        sharedFile.Filename,
		Location:        sharedFile.Location,
		Permission:      sharedFile.Permission,
		SharedWithEmail: sharedFile.SharedWithEmail,
	}

	return sfr, nil

}

func OpenFile(r *OpenFileRequest) (*OpenFileResponse, error) {
	status, err := repository.CheckForFileExistence(r.Filename, r.Location, r.Owner)
	if err != nil {
		return nil, err
	}
	if !status {
		return nil, ErrFileDoesNotExist
	}

	fileId, fileIdErr := repository.GetFileId(r.Filename, r.Location, r.Owner)
	if fileIdErr != nil {
		return nil, fileIdErr
	}

	wd, documentErr := repository.ReterieveDocument(fileId)
	if documentErr != nil {
		return nil, documentErr
	}

	if wd == nil {
		// Make a new woot document
		log.Println("Document not present, creating new one")
		wd, err = model.CreateNewDocument(fileId)
		if err != nil {
			return nil, err
		}
	}
	webSocketURL := fmt.Sprintf("ws://gollabedit.com/ws?docId=%s", fileId)

	response := &OpenFileResponse{
		Filename:        r.Filename,
		Location:        r.Location,
		DocumentID:      fileId,
		DocumentContent: *wd,
		WebSocketURL:    webSocketURL,
	}

	return response, nil
	// s3Status, err := repository.CheckIfUploaded(r.Filename, r.Location, r.Owner)
	// if err != nil {
	// 	return nil, err
	// }

	// s3Client, s3Err := awsutils.NewS3Client()
	// if s3Err != nil {
	// 	return nil, s3Err
	// }

	// conf, confErr := config.Load()
	// if confErr != nil {
	// 	return nil, confErr
	// }

	// bucket := conf.AWS.FilesBucket
	// cleanPath := path.Clean("/" + r.Owner + "/" + r.Location + "/" + r.Filename)

	// // Upload file to S3
	// if !s3Status {
	// 	content := []byte("// Empty file.")
	// 	err := s3Client.Upload(bucket, cleanPath, content)
	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	// Update uploaded status in DB
	// 	statusErr := repository.UpdateUploadedStatus(r.Filename, r.Location, r.Owner)
	// 	if statusErr != nil {
	// 		return nil, err
	// 	}
	// }

	// // download file from s3
	// sysPath := conf.Server.TmpPath

	// fileContent, s3DownloadErr := s3Client.Download(bucket, cleanPath)
	// if s3DownloadErr != nil {
	// 	return nil, s3DownloadErr
	// }

	// tempFileName := r.Owner + "_" + r.Filename
	// writeErr := os.WriteFile(path.Clean(sysPath+tempFileName), fileContent, 0644)
	// if err != nil {
	// 	fmt.Println("Error writing to file:", err)
	// 	return nil, writeErr
	// }

	// return nil, nil

}

func CreateFile(r *CreateFileRequest) (*CreateFileResponse, error) {

	status, err := repository.CheckForFileExistence(r.Filename, r.Location, r.Owner)
	if err != nil {
		return nil, err
	}
	if status {
		return nil, ErrFileAlreadyExists
	}

	loc, err := time.LoadLocation("UTC")
	if err != nil {
		fmt.Println("Error loading location:", err)
		return nil, err
	}
	currentTime := time.Now().In(loc)

	file := &model.File{
		Filename:   r.Filename,
		Location:   r.Location,
		Owner:      r.Owner,
		CreatedAt:  currentTime.Format("2006-01-02 15:04:05"),
		UpdatedAt:  currentTime.Format("2006-01-02 15:04:05"),
		IsUploaded: false,
		FileSize:   0,
	}

	// create file

	createFileErr := repository.CreateFile(file)
	if createFileErr != nil {
		return nil, createFileErr
	}

	cfr := &CreateFileResponse{
		Filename:  file.Filename,
		Location:  file.Location,
		CreatedAt: file.CreatedAt,
	}

	return cfr, nil
}
