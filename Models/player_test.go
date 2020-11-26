package models

import "testing"
// test for the creation of a new player model
func Test_NewPlayerCorrect(t *testing.T) {
	p := NewPlayerCMD("dam", 200)
	err := p.Validate()

	if err != nil {
		t.Error("Validation Fail")
		t.Fail()
	}
}
// test for the failure creating a player
func TestNewPlayerFail(t *testing.T) {
	p := NewPlayerCMD("", 100)
	err := p.Validate()

	if err == nil {
		t.Error("Validation Fail")
		t.Fail()
	}
}
