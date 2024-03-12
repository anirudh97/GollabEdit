package model

type User struct {
	Username string `db:"username"`
	Email    string `db:"email"`
	Password string `db:"hashPassword"`
}

type LoggedInUser struct {
	Username string
	Email    string
	Token    string
}

type File struct {
	Id         string `db:"id"`
	Filename   string `db:"filename"`
	Location   string `db:"location"`
	Owner      string `db:"owner"`
	IsUploaded bool   `db:"isUploaded"`
	FileSize   int64  `db:"fileSize"`
	CreatedAt  string `db:"createdAt"`
	UpdatedAt  string `db:"updatedAt"`
}

type SharedFile struct {
	ShareId         string
	Filename        string
	Location        string
	Permission      string
	SharedWithEmail string
	SharedByEmail   string
	SharedAt        string
	FileId          string
}
