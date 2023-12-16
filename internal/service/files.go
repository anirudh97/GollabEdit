package service

import (
	"fmt"
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
