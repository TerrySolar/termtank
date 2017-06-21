package main

import (
	"termtank/tank"

	tl "github.com/JoelOtter/termloop"
)

type Player struct {
	*tank.Tank
	preX  int
	preY  int
	level *tl.BaseLevel
}

type Enemy struct {
	*tank.Tank
	preX  int
	preY  int
	level *tl.BaseLevel
}

var (
	player Player
	enemy  Enemy
)

func (p Player) Tick(event tl.Event) {
	if event.Type == tl.EventKey {
		p.preX, p.preY = p.Position()

		var bulletX, bulletY, bulletDirection int
		bulletDirection = p.Tank.GetDirection()
		switch event.Key {

		case tl.KeyArrowUp:
			tank.TankUp(p.Tank)
			p.SetPosition(p.preX, p.preY-1)
		case tl.KeyArrowDown:
			tank.TankDown(p.Tank)
			p.SetPosition(p.preX, p.preY+1)
		case tl.KeyArrowRight:
			tank.TankRight(p.Tank)
			p.SetPosition(p.preX+1, p.preY)
		case tl.KeyArrowLeft:
			tank.TankLeft(p.Tank)
			p.SetPosition(p.preX-1, p.preY)

		case tl.KeySpace:
			switch bulletDirection {

			case tank.UP:
				bulletX = p.preX + 4
				bulletY = p.preY
			case tank.DOWN:
				bulletX = p.preX + 4
				bulletY = p.preY + 9
			case tank.LEFT:
				bulletX = p.preX
				bulletY = p.preY + 4
			case tank.RIGHT:
				bulletX = p.preX + 9
				bulletY = p.preY + 4
			}

			b := tank.NewBullet(bulletX, bulletY, bulletDirection)
			p.level.AddEntity(b)

		}

	}

}

func (enemy Enemy) Collide(collision tl.Physical) {
	if _, ok := collision.(tank.Bullet); ok {
		enemy.level.RemoveEntity(enemy)

	}

}

func (player Player) Draw(screen *tl.Screen) {

	tX, tY := player.Position()
	sX, sY := screen.Size()

	if tX < 0 {
		player.SetPosition(tX+1, tY)
	}
	if tX > sX-9 {
		player.SetPosition(tX-1, tY)
	}
	if tY < 0 {
		player.SetPosition(tX, tY+1)
	}
	if tY > sY-9 {
		player.SetPosition(tX, tY-1)
	}
	player.Entity.Draw(screen)
}

func main() {
	game := tl.NewGame()

	// BaseLevel
	level := tl.NewBaseLevel(tl.Cell{})

	// Initial player tank
	player := Player{
		Tank:  tank.NewTankXY(120, 120),
		level: level,
	}
	level.AddEntity(player)

	enemy := Enemy{
		Tank:  tank.NewTank(),
		level: level,
	}

	level.AddEntity(enemy)
	game.Screen().SetLevel(level)
	game.Screen().EnablePixelMode()
	game.Screen().SetFps(120)
	game.Start()

	// retrieve screen size after start(),before start(),size = 0
	sX, sY := game.Screen().Size()
	player.Tank.SetPosition(int(sX/2), int(sY/2))
}
