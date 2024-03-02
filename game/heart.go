package game

import (
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var heartImage *ebiten.Image

func init() {
	img, _, err := ebitenutil.NewImageFromFile("resources/images/ハート.png")
	if err != nil {
		panic(err)
	}
	heartImage = img
}

type Heart struct {
	image    *ebiten.Image
	x, y     float64
	maxLimit int
	limit    int
	scale    float64
}

func newHeart(x, y float64) *Heart {
	maxLimit := int(rand.Float64()*100 + 100)
	return &Heart{
		image:    heartImage,
		x:        x,
		y:        y,
		maxLimit: maxLimit,
		limit:    maxLimit,
		scale:    rand.Float64() + 0.5,
	}
}

func (h *Heart) Update() {
	h.x += rand.Float64()*2 - 1
	h.y -= rand.Float64()
	if h.limit > 0 {
		h.limit--
	}
}

func (h *Heart) Draw(screen *ebiten.Image) {
	if h.limit <= 0 {
		return
	}
	cScale := float32(h.limit) / float32(h.maxLimit)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(h.scale, h.scale)
	op.GeoM.Translate(h.x, h.y)
	op.ColorScale.Scale(cScale+0.1, cScale, cScale, 1)
	screen.DrawImage(h.image, op)
}
