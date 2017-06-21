package tank

import (
	tl "github.com/JoelOtter/termloop"
)

type Bullet struct {
	*tl.Entity
	direction int
}

func NewBullet(x, y, d int) Bullet {
	b := Bullet{
		Entity:    tl.NewEntity(x, y, 1, 1),
		direction: d,
	}
	b.SetCell(0, 0, &tl.Cell{Fg: tl.ColorWhite, Bg: tl.ColorWhite})

	return b
}

func (b Bullet) Draw(screen *tl.Screen) {

	bX, bY := b.Position()
	screenX, screenY := screen.Size()

	if bX > screenX || bX < 0 || bY > screenY || bY < 0 {
		screen.RemoveEntity(b)
		screen.Level().RemoveEntity(b)

	}

	switch b.direction {

	case UP:
		b.SetPosition(bX, bY-1)
	case DOWN:
		b.SetPosition(bX, bY+1)
	case LEFT:
		b.SetPosition(bX-1, bY)
	case RIGHT:
		b.SetPosition(bX+1, bY)
	}
	b.Entity.Draw(screen)

}

func (b Bullet) Tick(event tl.Event) {}
