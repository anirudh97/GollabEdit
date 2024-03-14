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

type WootCharacter struct {
	Id        string `bson:"id"`
	Character string `bson:"character"`
	Visible   bool   `bson:"visible"`
	PrevId    string `bson:"prevId"`
	NextId    string `bson:"nextId"`
}

type WootDocument struct {
	Id         string                    `bson:"id"`
	Characters map[string]*WootCharacter `bson:"characters"`
	StartId    string                    `bson:"startId"`
	EndId      string                    `bson:"endId"`
}

func CreateNewDocument(id string) (*WootDocument, error) {
	start := &WootCharacter{Id: "start", Visible: false}
	end := &WootCharacter{Id: "end", Visible: false, PrevId: "start"}
	start.NextId = "end"

	return &WootDocument{
		Id: id,
		Characters: map[string]*WootCharacter{
			"start": start,
			"end":   end,
		},
		StartId: "start",
		EndId:   "end",
	}, nil
}

func (wd *WootDocument) InsertCharacter(character string, charId string, prevId string, nextId string, visible bool) {
	if prevId == "" {
		prevId = "start"
	}

	if nextId == "" {
		nextId = "end"
	}

	nwc := &WootCharacter{
		Id:        charId,
		Character: character,
		PrevId:    prevId,
		NextId:    nextId,
		Visible:   visible,
	}
	wd.Characters[charId] = nwc

	prevChar := wd.Characters[prevId]
	nextChar := wd.Characters[nextId]
	prevChar.NextId = nwc.Id
	nextChar.PrevId = nwc.Id

}

func (wd *WootDocument) DeleteCharacter(charID string) {
	if char, exists := wd.Characters[charID]; exists {
		char.Visible = false
	}
}
