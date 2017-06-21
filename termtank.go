package main

import (
	"fmt"
	"termtank/tank"

	tl "github.com/JoelOtter/termloop"
)

type Player struct {
	*tank.Tank
	preX int
	preY int
}

var (
	player Player
)

func (p Player) Tick(event tl.Event) {
	if event.Type == tl.EventKey {
		player.preX, player.preY = player.Position()

		switch event.Key {

		case tl.KeyArrowUp:
			tank.TankUp(player.Tank)
			player.SetPosition(player.preX, player.preY-1)
		case tl.KeyArrowDown:
			tank.TankDown(player.Tank)
			player.SetPosition(player.preX, player.preY+1)
		case tl.KeyArrowRight:
			tank.TankRight(player.Tank)
			player.SetPosition(player.preX+1, player.preY)
		case tl.KeyArrowLeft:
			tank.TankLeft(player.Tank)
			player.SetPosition(player.preX-1, player.preY)

		case tl.KeySpace:
			fmt.Println("Shoot!!!!!!")
		}

	}

}

func main() {
	game := tl.NewGame()

	// BaseLevel
	level := tl.NewBaseLevel(tl.Cell{})

	// Initial player tank
	player.Tank = tank.NewTank(*level)
	level.AddEntity(player)

	game.Screen().SetLevel(level)
	game.Screen().EnablePixelMode()
	game.Screen().SetFps(120)
	game.Start()

}
