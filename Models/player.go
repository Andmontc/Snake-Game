package models

import "errors"

var (
	//EmptyName is used when the name of the player is empty
	EmptyName = errors.New("name can't be empty")
)
// Struct that represent the player data
type Player struct {
	Id	  string `json:"id"`
	Name  string `json:"name"`
	Score string `json:"score"`
}
// function that create the model
func NewPlayerCMD(id string, name string, score string) *Player {
	return &Player{
		Id: id,
		Name: name,
		Score: score,
	}
}

func (cmd *Player) Validate() error {
	if cmd.Name == "" {
		return EmptyName
	}
	return nil
}
