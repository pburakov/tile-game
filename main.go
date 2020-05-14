package main

import (
	"github.com/hajimehoshi/ebiten"
	"image"
	_ "image/png"
	"log"
	"os"
	"time"
)

const (
	tileSize     = 16
	screenWidth  = 320
	screenHeight = 240 // 20 x 15 tiles
	tilesPerRow  = screenWidth / tileSize
)

const (
	none        = byte(0x00)
	grass       = byte(0x19)
	trainSprite = byte(0x0f)
	cursor      = byte(0x14)

	// To construct a tile, combine the starting tile with an offset number
	// e.g. for crossing: `road + Cr` `rail + VL`
	road = byte(0x00)
	rail = byte(0x0f)

	// Tile offsets
	Ver = byte(0x01) // Vertical
	Hor = byte(0x02) // Horizontal
	DL  = byte(0x03) // Down-Left curve
	DR  = byte(0x06) // Down-Right curve
	UL  = byte(0x0d) // Up-Left curve
	UR  = byte(0x0b) // Up-Right curve
	Cro = byte(0x07) // Crossing
	VL  = byte(0x08) // Vertical-Left T-section
	VR  = byte(0x09) // Vertical-Right T-section
	HU  = byte(0x0c) // Horizontal-Up T-section
	HD  = byte(0x0e) // Horizontal-Down T-section
)

var (
	tilesImage *ebiten.Image
	tiles      = []byte{
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x15, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x15, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x17, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x18, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x1c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x11, 0x11, 0x11, 0x11, 0x11, 0x1c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	}
	clock = 0
	done  = make(chan bool) // exit signal
	train Train
)

type Train struct {
	x, y       float64
	xVelocity  float64
	yVelocity  float64
	nextUpdate int // denotes clock tick at which to move to the next pixel
}

func init() {
	log.SetFlags(0)
	log.Print("initializing...")

	var err error
	tilesImage, err = loadImage("resource/tiles.png")
	if err != nil {
		log.Fatal(err)
	}
	trainX, trainY := tileToPosition(0, 12)
	train = Train{float64(trainX), float64(trainY), 1.0, 0.0, 0}
	launchClock()
}

func launchClock() {
	ticker := time.NewTicker(100 * time.Millisecond)
	go func() {
		for {
			select {
			case <-done:
				log.Print("clock stopped")
				return
			case <-ticker.C:
				cycle()
				clock++
			}
		}
	}()
	log.Print("launched clock")
}

// cycle updates inner state
func cycle() {
	train.x += train.xVelocity
	train.y += train.yVelocity
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

	err = drawTrain(screen)
	if err != nil {
		return err
	}

	x, y := ebiten.CursorPosition()
	err = drawCursor(x, y, screen)
	if err != nil {
		return err
	}

	return nil
}

func drawCursor(x int, y int, screen *ebiten.Image) error {
	op := &ebiten.DrawImageOptions{}
	// get top-left coordinates of the nearest tile
	px, py := tileToPosition(positionToTile(x, y))
	op.GeoM.Translate(float64(px), float64(py))

	sx, sy := getTileCoordinates(cursor, tilesImage)
	sprite := tilesImage.SubImage(image.Rect(sx, sy, sx+tileSize, sy+tileSize))
	err := screen.DrawImage(sprite.(*ebiten.Image), op)
	if err != nil {
		return err
	}
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

		sx, sy := getTileCoordinates(grass, tilesImage)
		sprite := tilesImage.SubImage(image.Rect(sx, sy, sx+tileSize, sy+tileSize))
		err := screen.DrawImage(sprite.(*ebiten.Image), op)
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

func drawTrain(screen *ebiten.Image) error {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(train.x, train.y)

	sx, sy := getTileCoordinates(trainSprite, tilesImage)
	sprite := tilesImage.SubImage(image.Rect(sx, sy, sx+tileSize, sy+tileSize))
	err := screen.DrawImage(sprite.(*ebiten.Image), op)
	if err != nil {
		return err
	}
	return nil
}

// positionToTileOrdinal returns ordinal number of a tile encompassing (x, y) position
func positionToTileOrdinal(x int, y int) (t int) {
	return tileToOrdinal(positionToTile(x, y))
}

// positionToTile returns matrix coordinates of a tile encompassing (x, y) position
func positionToTile(x int, y int) (tx int, ty int) {
	return (x - x%tileSize) / tileSize, (y - y%tileSize) / tileSize
}

// tileToPosition returns position of top-left corner of a tile with (tx, ty) coordinates
func tileToPosition(tx int, ty int) (x int, y int) {
	return tx * tileSize, ty * tileSize
}

// tileToOrdinal returns ordinal number of a tile with (tx, ty) matrix coordinates
func tileToOrdinal(tx int, ty int) int {
	return tilesPerRow*ty + tx
}

// getTileCoordinates returns tile top-left coordinates in the image
// from a given ordinal number of a tile.
func getTileCoordinates(t byte, img *ebiten.Image) (x int, y int) {
	width, _ := img.Size()
	return int(t) % (width / tileSize) * tileSize, int(t) / (width / tileSize) * tileSize
}

func main() {
	if err := ebiten.Run(update, screenWidth, screenHeight, 2, "Tiles"); err != nil {
		log.Fatal(err)
	}
	done <- true
}
