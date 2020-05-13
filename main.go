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
)

var (
	bkgImage   *ebiten.Image
	tilesImage *ebiten.Image
	tiles      = []int{
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 16, 17, 17, 17, 17, 24, 17, 17, 17, 17, 17, 18, 0, 0, 0, 0,
		0, 0, 0, 0, 19, 17, 17, 17, 17, 29, 17, 17, 17, 17, 17, 23, 0, 0, 0, 0,
		0, 0, 0, 0, 21, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 21, 0, 0, 0, 0,
		0, 0, 0, 0, 21, 0, 0, 1, 2, 14, 2, 2, 3, 0, 0, 21, 0, 0, 0, 0,
		0, 0, 0, 0, 21, 0, 0, 6, 0, 6, 0, 0, 6, 0, 0, 21, 0, 0, 0, 0,
		0, 0, 0, 0, 21, 0, 0, 6, 0, 6, 0, 0, 6, 0, 0, 21, 0, 0, 0, 0,
		0, 0, 0, 0, 21, 0, 0, 9, 2, 7, 2, 2, 8, 0, 0, 21, 0, 0, 0, 0,
		0, 0, 0, 0, 21, 0, 0, 6, 0, 6, 0, 0, 6, 0, 0, 21, 0, 0, 0, 0,
		0, 0, 0, 0, 21, 0, 0, 6, 0, 6, 0, 0, 6, 0, 0, 21, 0, 0, 0, 0,
		0, 0, 0, 0, 21, 0, 0, 11, 2, 12, 2, 2, 13, 0, 16, 28, 0, 0, 0, 0,
		0, 0, 0, 0, 21, 0, 0, 0, 0, 0, 0, 0, 0, 16, 28, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 26, 17, 17, 17, 17, 17, 17, 17, 17, 28, 0, 0, 0, 0, 0, 0,
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

		sx, sy := getTileCoordinates(t, tilesImage, tileSize)
		sprite := tilesImage.SubImage(image.Rect(sx, sy, sx+tileSize, sy+tileSize))
		err := screen.DrawImage(sprite.(*ebiten.Image), op)
		if err != nil {
			return err
		}
	}
	return nil
}

// getTileCoordinates returns tile top-left coordinates in the image with a given
// square tile size and the ordinal number of a tile.
func getTileCoordinates(t int, img *ebiten.Image, tileSize int) (x int, y int) {
	width, _ := img.Size()
	return t % (width / tileSize) * tileSize, t / (width / tileSize) * tileSize
}

func main() {
	if err := ebiten.Run(update, screenWidth, screenHeight, 2, "Tiles"); err != nil {
		log.Fatal(err)
	}
}
