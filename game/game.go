package game

import (
	"log"
	"path"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	ScreenWidth  = 480
	ScreenHeight = 480
)

var (
	gophers map[string]*ebiten.Image
	images  map[string]*ebiten.Image
)

func init() {
	gophers = make(map[string]*ebiten.Image)
	imagesNames := []string{"left", "left-up", "up", "right-up", "right", "right-down", "down", "left-down"}
	for _, name := range imagesNames {
		img, _, err := ebitenutil.NewImageFromFile(path.Join("resources", "images", name+".png"))
		if err != nil {
			log.Fatal(err)
		}
		gophers[name] = img
	}
	images = make(map[string]*ebiten.Image)
	for _, name := range imagesNames {
		gopher := gophers[name]
		img := ebiten.NewImage(ScreenWidth, ScreenHeight)
		opt := &ebiten.DrawImageOptions{}
		opt.GeoM.Translate(float64(ScreenWidth/2-gopher.Bounds().Dx()/2), float64(ScreenHeight/2-gopher.Bounds().Dy()/2))
		img.DrawImage(gophers[name], opt)
		images[name] = img
	}
}

type Game struct {
	image *ebiten.Image
}

func newGame() (*Game, error) {
	g := &Game{
		image: images["left"],
	}
	return g, nil
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.image, nil)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}

func RunGame() {
	g, err := newGame()
	if err != nil {
		log.Fatal(err)
	}
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowTitle("Go Refresh")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
