package game

import (
	"image"
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
	eyePos  = image.Pt(ScreenWidth/2+10, ScreenHeight/2-65)
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
	mouseX, mouseY := ebiten.CursorPosition()
	if mouseX < eyePos.X {
		if mouseY < eyePos.Y {
			g.image = images["left-up"]
		} else if mouseY > eyePos.Y {
			g.image = images["left-down"]
		} else {
			g.image = images["left"]
		}
	} else if mouseX > eyePos.X {
		if mouseY < eyePos.Y {
			g.image = images["right-up"]
		} else if mouseY > eyePos.Y {
			g.image = images["right-down"]
		} else {
			g.image = images["right"]
		}
	} else {
		if mouseY < eyePos.Y {
			g.image = images["up"]
		} else if mouseY > eyePos.Y {
			g.image = images["down"]
		}
	}
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
