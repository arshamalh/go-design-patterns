package main

type Player struct {
	dress      Dress
	playerType string
}

func newPlayer(playerType, dressType string) *Player {
	dress, _ := getDressFactorySingleInstance().getDressByType(dressType)
	return &Player{
		playerType: playerType,
		dress:      dress,
	}
}