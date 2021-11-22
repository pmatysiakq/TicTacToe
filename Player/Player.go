package Player

type Player struct {
	Name string
	Sign string
}

func (p *Player) getPlayerName() string {
	return p.Name
}

func (p *Player) getPlayerSign() string {
	return p.Sign
}

func (p *Player) setPlayerName(playerName string) {
	p.Name = playerName
}

func (p *Player) setPlayerSign(playerSign string) {
	p.Sign = playerSign
}