package main

import (
	"github.com/hajimehoshi/ebiten"
	"image"
	_ "image/png"
	"log"
	"os"
)

const (
	tileSize     = 16
	screenWidth  = 320
	screenHeight = 240
	tilesPerRow  = screenWidth / tileSize
	xTilesImg    = 2
)

var (
	bkgImage   *ebiten.Image
	tilesImage *ebiten.Image
	tiles      = []int{
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 1, 2, 2, 2, 2, 3, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 5, 0, 0, 0, 0, 5, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 5, 0, 0, 0, 0, 5, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 5, 0, 1, 2, 2, 7, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 5, 0, 5, 0, 0, 5, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 5, 0, 5, 0, 0, 5, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 9, 2, 10, 2, 2, 11, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	}
)

func init() {
	log.SetFlags(0)
	log.Print("initializing...")

	var err error
	bkgImage, err = loadImage("resource/bkg.png")
	if err != nil {
		log.Fatal(err)
	}
	tilesImage, err = loadImage("resource/tiles.png")
	if err != nil {
		log.Fatal(err)
	}
}

func loadImage(path string) (*ebiten.Image, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	img, format, err := image.Decode(f)
	if err != nil {
		return nil, err
	}
	log.Printf("reading %q as format %q", path, format)
	ebitenImg, err := ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	if err != nil {
		return nil, err
	}
	return ebitenImg, nil
}

func update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}

	err := drawBackground(screen)
	if err != nil {
		return err
	}

	err = drawTiles(screen)
	if err != nil {
		return err
	}

	return nil
}

func drawBackground(screen *ebiten.Image) error {
	for i := 0; i < screenWidth*screenHeight/tileSize*tileSize; i++ {
		op := &ebiten.DrawImageOptions{}
		// pixel coordinates
		x := float64(i%tilesPerRow) * tileSize
		y := float64(i/tilesPerRow) * tileSize
		op.GeoM.Translate(x, y)

		err := screen.DrawImage(bkgImage, op)
		if err != nil {
			return err
		}
	}
	return nil
}

func drawTiles(screen *ebiten.Image) error {
	for i, t := range tiles {
		op := &ebiten.DrawImageOptions{}
		// pixel coordinates
		x := float64(i%tilesPerRow) * tileSize
		y := float64(i/tilesPerRow) * tileSize
		op.GeoM.Translate(x, y)

		width, _ := tilesImage.Size()
		// sprite image rectangle top-left coordinates
		sx := t % (width / tileSize) * tileSize
		sy := t / (width / tileSize) * tileSize
		sprite := tilesImage.SubImage(image.Rect(sx, sy, sx+tileSize, sy+tileSize))
		err := screen.DrawImage(sprite.(*ebiten.Image), op)
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	if err := ebiten.Run(update, screenWidth, screenHeight, 2, "Tiles"); err != nil {
		log.Fatal(err)
	}
}
