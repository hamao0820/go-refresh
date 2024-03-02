package game

import (
	"image"
	"log"
	"math"
	"path"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type GopherMode int

const (
	GopherModeNormal GopherMode = iota
	GopherModeHappy
	GopherModeLovely
)

const (
	ScreenWidth  = 480
	ScreenHeight = 480
)

var (
	gophers  map[string]*ebiten.Image
	images   map[string]*ebiten.Image
	eyePos   = image.Pt(ScreenWidth/2+10, ScreenHeight/2-65)
	heartPos = image.Pt(ScreenWidth*2/3, ScreenHeight/3-30)
)

func init() {
	gophers = make(map[string]*ebiten.Image)
	imagesNames := []string{"left", "left-up", "up", "right-up", "right", "right-down", "down", "left-down", "happy-left", "happy-left-up", "happy-up", "happy-right-up", "happy-right", "happy-right-down", "happy-down", "happy-left-down"}
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
	image          *ebiten.Image
	pointer        *ebiten.Image
	rubCount       int
	gopherMode     GopherMode
	heart          *Heart
	mouseX, mouseY int
}

func newGame() (*Game, error) {
	pointer, _, err := ebitenutil.NewImageFromFile(path.Join("resources", "images", "ポインタ.png"))
	if err != nil {
		return nil, err
	}
	g := &Game{
		image:      images["left"],
		pointer:    pointer,
		gopherMode: GopherModeNormal,
	}
	return g, nil
}

func (g *Game) Update() error {
	mouseX, mouseY := ebiten.CursorPosition()

	// 視線を変える
	arg := math.Atan2(float64(mouseY-eyePos.Y), float64(mouseX-eyePos.X))
	if g.gopherMode == GopherModeNormal {
		if arg < -math.Pi*7/8 {
			g.image = images["left"]
		} else if arg < -math.Pi*5/8 {
			g.image = images["left-up"]
		} else if arg < -math.Pi*3/8 {
			g.image = images["up"]
		} else if arg < -math.Pi/8 {
			g.image = images["right-up"]
		} else if arg < math.Pi/8 {
			g.image = images["right"]
		} else if arg < math.Pi*3/8 {
			g.image = images["right-down"]
		} else if arg < math.Pi*5/8 {
			g.image = images["down"]
		} else if arg < math.Pi*7/8 {
			g.image = images["left-down"]
		} else {
			g.image = images["left"]
		}
	} else if g.gopherMode == GopherModeHappy {
		if arg < -math.Pi*7/8 {
			g.image = images["happy-left"]
		} else if arg < -math.Pi*5/8 {
			g.image = images["happy-left-up"]
		} else if arg < -math.Pi*3/8 {
			g.image = images["happy-up"]
		} else if arg < -math.Pi/8 {
			g.image = images["happy-right-up"]
		} else if arg < math.Pi/8 {
			g.image = images["happy-right"]
		} else if arg < math.Pi*3/8 {
			g.image = images["happy-right-down"]
		} else if arg < math.Pi*5/8 {
			g.image = images["happy-down"]
		} else if arg < math.Pi*7/8 {
			g.image = images["happy-left-down"]
		} else {
			g.image = images["happy-left"]
		}
	}

	if g.gopherMode == GopherModeLovely {
		if g.heart == nil || g.heart.limit <= 0 {
			g.heart = newHeart(float64(heartPos.X), float64(heartPos.Y))
		}
	}

	if g.heart != nil {
		g.heart.Update()
	}

	// gopherを撫でると、撫でた距離に応じてrubCountが増える

	if mouseX >= 175 && mouseX <= 315 && mouseY >= 120 && mouseY <= 350 {
		g.rubCount += int(math.Abs(float64(mouseX-g.mouseX))+math.Abs(float64(mouseY-g.mouseY))) / 10
	} else {
		g.rubCount -= 1
	}

	if g.rubCount < 0 {
		g.rubCount = 0
	} else if g.rubCount > 100 {
		g.rubCount = 100
	}

	if g.rubCount == 100 {
		g.gopherMode = GopherModeLovely
	} else if g.rubCount > 30 {
		g.gopherMode = GopherModeHappy
	} else {
		g.gopherMode = GopherModeNormal
	}

	g.mouseX, g.mouseY = mouseX, mouseY
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.image, nil)
	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Translate(float64(g.mouseX-g.pointer.Bounds().Dx()/2), float64(g.mouseY-g.pointer.Bounds().Dy()/2))
	screen.DrawImage(g.pointer, opt)
	if g.heart != nil {
		g.heart.Draw(screen)
	}
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
	ebiten.SetCursorMode(ebiten.CursorModeHidden)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
