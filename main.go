package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"image"
	_ "image/png"
	"log"
	"os"
)

const (
	tileSize     = 16
	screenWidth  = 320
	screenHeight = 240 // 20 x 15 tiles
	tilesPerRow  = screenWidth / tileSize
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
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
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

	x, y := ebiten.CursorPosition()
	lftDown := ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
	rgtDown := ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight)
	tx, ty := positionToTile(x, y)

	if lftDown {
		placeTile(tx, ty)
	}
	if rgtDown {
		deleteTile(tx, ty)
	}

	err := drawBackground(screen)
	if err != nil {
		return err
	}

	err = drawTiles(screen)
	if err != nil {
		return err
	}

	ebitenutil.DebugPrint(screen, fmt.Sprintf(
		"left mouse down: %t\nrght mouse down: %t\npos: (%d, %d)\ntile: (%d, %d)",
		lftDown, rgtDown, x, y, tx, ty))

	return nil
}

func isEmpty(tx int, ty int) bool {
	t := tileToOrdinal(tx, ty)
	if t < 0 || t >= len(tiles) {
		return true
	} else {
		return tiles[t] == 0
	}
}

func deleteTile(tx int, ty int) {
	t := tileToOrdinal(tx, ty)
	tiles[t] = 0
}

func placeTile(tx int, ty int) {
	t := tileToOrdinal(tx, ty)
	tiles[t] = 20
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

		sx, sy := getTileCoordinates(t, tilesImage)
		sprite := tilesImage.SubImage(image.Rect(sx, sy, sx+tileSize, sy+tileSize))
		err := screen.DrawImage(sprite.(*ebiten.Image), op)
		if err != nil {
			return err
		}
	}
	return nil
}

// positionToTile returns tile coordinate encompassing pixel (x, y) coordinates
func positionToTile(x int, y int) (tx int, ty int) {
	return (x - x%tileSize) / tileSize, (y - y%tileSize) / tileSize
}

// tileToOrdinal returns ordinal number of (x, y) coordinates of a tile
func tileToOrdinal(tx int, ty int) int {
	return tilesPerRow*ty + tx
}

// getTileCoordinates returns tile top-left coordinates in the image
// from a given ordinal number of a tile.
func getTileCoordinates(t int, img *ebiten.Image) (x int, y int) {
	width, _ := img.Size()
	return t % (width / tileSize) * tileSize, t / (width / tileSize) * tileSize
}

func main() {
	if err := ebiten.Run(update, screenWidth, screenHeight, 2, "Tiles"); err != nil {
		log.Fatal(err)
	}
}
