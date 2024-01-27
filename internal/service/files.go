package service

import (
	"fmt"
	"os"
	"path"
	"time"

	"github.com/anirudh97/GollabEdit/internal/awsutils"
	"github.com/anirudh97/GollabEdit/internal/config"
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
type OpenFileRequest struct {
	Filename string `json:"filename"`
	Location string `json:"location"`
	Owner    string `json:"owner"`
}

func OpenFile(r *OpenFileRequest) (*CreateFileResponse, error) {
	status, err := repository.CheckForFileExistence(r.Filename, r.Location, r.Owner)
	if err != nil {
		return nil, err
	}
	if !status {
		return nil, ErrFileDoesNotExist
	}

	s3Status, err := repository.CheckIfUploaded(r.Filename, r.Location, r.Owner)
	if err != nil {
		return nil, err
	}

	s3Client, s3Err := awsutils.NewS3Client()
	if s3Err != nil {
		return nil, s3Err
	}

	conf, confErr := config.Load()
	if confErr != nil {
		return nil, confErr
	}

	bucket := conf.AWS.FilesBucket
	cleanPath := path.Clean("/" + r.Owner + "/" + r.Location + "/" + r.Filename)

	// Upload file to S3
	if !s3Status {
		content := []byte("// Empty file.")
		err := s3Client.Upload(bucket, cleanPath, content)
		if err != nil {
			return nil, err
		}

		// Update uploaded status in DB
		statusErr := repository.UpdateUploadedStatus(r.Filename, r.Location, r.Owner)
		if statusErr != nil {
			return nil, err
		}
	}

	// download file from s3
	sysPath := conf.Server.TmpPath

	fileContent, s3DownloadErr := s3Client.Download(bucket, cleanPath)
	if s3DownloadErr != nil {
		return nil, s3DownloadErr
	}

	tempFileName := r.Owner + "_" + r.Filename
	writeErr := os.WriteFile(path.Clean(sysPath+tempFileName), fileContent, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return nil, writeErr
	}

	return nil, nil

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
