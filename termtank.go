package main

import (
	"math/rand"
	"termtank/tank"
	"time"

	tl "github.com/JoelOtter/termloop"
)

// Tank directions
const (
	UP    int = 1
	DOWN  int = 2
	LEFT  int = 3
	RIGHT int = 4
)

type Player struct {
	*tank.Tank
	preX  int
	preY  int
	level *tl.BaseLevel
}

type Enemy struct {
	*tank.Tank
	preX   int
	preY   int
	level  *tl.BaseLevel
	status int // 1:normal 0:dead
}

var (
	player    Player
	enemy     Enemy
	countTime float64
)

func (p *Player) Tick(event tl.Event) {
	if event.Type == tl.EventKey {
		p.preX, p.preY = p.Position()

		var bulletX, bulletY, bulletDirection int
		bulletDirection = p.Tank.GetDirection()
		cell := tl.Cell{Fg: tl.ColorRed, Bg: tl.ColorRed}

		switch event.Key {

		case tl.KeyArrowUp:
			tank.TankUp(p.Tank, cell)
			p.SetPosition(p.preX, p.preY-1)
		case tl.KeyArrowDown:
			tank.TankDown(p.Tank, cell)
			p.SetPosition(p.preX, p.preY+1)
		case tl.KeyArrowRight:
			tank.TankRight(p.Tank, cell)
			p.SetPosition(p.preX+1, p.preY)
		case tl.KeyArrowLeft:
			tank.TankLeft(p.Tank, cell)
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

func (enemy *Enemy) Collide(collision tl.Physical) {

	if _, ok := collision.(tank.Bullet); ok {

		// set dead status
		enemy.status = 0
		// remove from screen
		enemy.level.RemoveEntity(enemy)

	} else if _, ok := collision.(tank.Tank); ok {
		enemy.SetPosition(enemy.preX, enemy.preY)
	}

}

func (enemy *Enemy) Draw(screen *tl.Screen) {

	countTime += screen.TimeDelta()

	enemy.preX, enemy.preY = enemy.Position()
	rand.Seed(time.Now().UnixNano())

	step := 3

	if countTime > 0.8 {

		direction := rand.Intn(4)
		cell := tl.Cell{Bg: tl.ColorBlue}

		switch direction + 1 {

		case UP:
			tank.TankUp(enemy.Tank, cell)
			enemy.SetPosition(enemy.preX, enemy.preY-step)
		case DOWN:
			tank.TankDown(enemy.Tank, cell)
			enemy.SetPosition(enemy.preX, enemy.preY+step)
		case LEFT:
			tank.TankLeft(enemy.Tank, cell)
			enemy.SetPosition(enemy.preX-step, enemy.preY)
		case RIGHT:
			tank.TankRight(enemy.Tank, cell)
			enemy.SetPosition(enemy.preX+step, enemy.preY)
		}

		// reset countTime
		countTime = 0.0

		tX, tY := enemy.Position()
		sX, sY := screen.Size()

		if tX < 0 {
			enemy.SetPosition(tX+step, tY)
		}
		if tX > sX-9 {
			enemy.SetPosition(tX-step, tY)
		}
		if tY < 0 {
			enemy.SetPosition(tX, tY+step)
		}
		if tY > sY-9 {
			enemy.SetPosition(tX, tY-step)
		}

	}
	enemy.Entity.Draw(screen)

}

func (player *Player) Draw(screen *tl.Screen) {

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
		Tank:  tank.NewTankXY(120, 120, tl.Cell{Bg: tl.ColorRed}),
		level: level,
	}
	level.AddEntity(&player)

	enemy := Enemy{
		Tank:   tank.NewTankXY(120, 60, tl.Cell{Bg: tl.ColorBlue}),
		level:  level,
		status: 1,
	}

	enemy1 := Enemy{
		Tank:   tank.NewTankXY(60, 60, tl.Cell{Bg: tl.ColorBlue}),
		level:  level,
		status: 1,
	}

	level.AddEntity(&enemy1)
	level.AddEntity(&enemy)

	game.Screen().SetLevel(level)
	game.Screen().EnablePixelMode()
	game.Screen().SetFps(120)
	game.Start()

}
