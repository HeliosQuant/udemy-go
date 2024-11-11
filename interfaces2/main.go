package main

import (
	"fmt"
	"math/rand"
)

type Player interface {
	KickBall()
}

type FootballPlayer struct {
	stamina int
	power   int
}

type CR7 struct {
	stamina int
	power   int
	SUI     int
}

func (fp CR7) KickBall() {
	//shot := fp.stamina + fp.power
	shot := fp.stamina + fp.power*fp.SUI
	fmt.Println("CR7 kicking  the ball", shot)
}

func (fp FootballPlayer) KickBall() {
	shot := fp.stamina + fp.power
	fmt.Println("I'm kicking  the ball", shot)
}

func main() {
	team := make([]Player, 11)

	for i := 0; i < len(team)-1; i++ {
		team[i] = FootballPlayer{
			stamina: rand.Intn(10),
			power:   rand.Intn(10),
		}
	}
	team[len(team)-1] = CR7{
		stamina: 10,
		power:   10,
		SUI:     10,
	}
	for i := 0; i < len(team); i++ {
		team[i].KickBall()
	}
}
