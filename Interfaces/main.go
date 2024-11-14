package main

import "fmt"

type PlayerFootball struct {
	power int
}

type PlayerSoccer struct {
	power int
}

type P interface {
	Kick()
}

func (p *PlayerFootball) Kick(power int) {
	fmt.Println("Kicking the small ball with power", power)
}

func (p *PlayerSoccer) Kick(power int) {
	fmt.Println("Kicking the big ball with power", power)
}

func main() {
	player1 := &PlayerFootball{
		power: 10,
	}
	player2 := &PlayerSoccer{
		power: 11,
	}
	player1.Kick(player1.power)

	player2.Kick(player2.power)
}


type Error interface {
	
}