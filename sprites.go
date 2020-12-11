package main

import (
	"image"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten"
)

// Sprite ordinal numbers
const (
	none   = byte(0x00)
	grass  = byte(0x19)
	cursor = byte(0x14)

	// To construct a tile, combine the starting tile with an offset number
	// e.g. for crossing: `road + cr` `rail + vl`
	rail = byte(0x00)
	swch = byte(0x0f)
	road = byte(0x10)

	// Tile offsets
	ver = byte(0x01) // Vertical
	hor = byte(0x02) // Horizontal
	cro = byte(0x07) // Crossing
	dl  = byte(0x03) // Down-Left curve
	dr  = byte(0x06) // Down-Right curve
	ul  = byte(0x0d) // Up-Left curve
	ur  = byte(0x0b) // Up-Right curve
	vl  = byte(0x08) // Vertical-Left T-section
	vr  = byte(0x09) // Vertical-Right T-section
	hu  = byte(0x0c) // Horizontal-Up T-section
	hd  = byte(0x0e) // Horizontal-Down T-section
)

const (
	left      = "left"
	upLeft    = "upLeft"
	up        = "up"
	upRight   = "upRight"
	right     = "right"
	downRight = "downRight"
	down      = "down"
	downLeft  = "downLeft"
)

var tileSprites = make(map[byte]image.Image)
var headSprites = make(map[string]image.Image)
var carSprites = make(map[string]image.Image)

func init() {
	log.SetFlags(0)

	var err error
	img, err := loadImage("resource/tiles.png")
	if err != nil {
		log.Fatal(err)
	}

	tileSprites[none] = loadTileSprite(none, img)
	tileSprites[grass] = loadTileSprite(grass, img)
	tileSprites[cursor] = loadTileSprite(cursor, img)

	tileSprites[rail+ver] = loadTileSprite(rail+ver, img)
	tileSprites[rail+hor] = loadTileSprite(rail+hor, img)
	tileSprites[rail+cro] = loadTileSprite(rail+cro, img)
	tileSprites[rail+dl] = loadTileSprite(rail+dl, img)
	tileSprites[rail+dr] = loadTileSprite(rail+dr, img)
	tileSprites[rail+ul] = loadTileSprite(rail+ul, img)
	tileSprites[rail+ur] = loadTileSprite(rail+ur, img)
	tileSprites[rail+vl] = loadTileSprite(rail+vl, img)
	tileSprites[rail+vr] = loadTileSprite(rail+vr, img)
	tileSprites[rail+hu] = loadTileSprite(rail+hu, img)
	tileSprites[rail+hd] = loadTileSprite(rail+hd, img)

	tileSprites[swch+ver] = loadTileSprite(swch+ver, img)
	tileSprites[swch+hor] = loadTileSprite(swch+hor, img)
	tileSprites[swch+cro] = loadTileSprite(swch+cro, img)
	tileSprites[swch+dl] = loadTileSprite(swch+dl, img)
	tileSprites[swch+dr] = loadTileSprite(swch+dr, img)
	tileSprites[swch+ul] = loadTileSprite(swch+ul, img)
	tileSprites[swch+ur] = loadTileSprite(swch+ur, img)
	tileSprites[swch+vl] = loadTileSprite(swch+vl, img)
	tileSprites[swch+vr] = loadTileSprite(swch+vr, img)
	tileSprites[swch+hu] = loadTileSprite(swch+hu, img)
	tileSprites[swch+hd] = loadTileSprite(swch+hd, img)

	tileSprites[road+ver] = loadTileSprite(road+ver, img)
	tileSprites[road+hor] = loadTileSprite(road+hor, img)
	tileSprites[road+cro] = loadTileSprite(road+cro, img)
	tileSprites[road+dl] = loadTileSprite(road+dl, img)
	tileSprites[road+dr] = loadTileSprite(road+dr, img)
	tileSprites[road+ul] = loadTileSprite(road+ul, img)
	tileSprites[road+ur] = loadTileSprite(road+ur, img)
	tileSprites[road+vl] = loadTileSprite(road+vl, img)
	tileSprites[road+vr] = loadTileSprite(road+vr, img)
	tileSprites[road+hu] = loadTileSprite(road+hu, img)
	tileSprites[road+hd] = loadTileSprite(road+hd, img)

	headSprites[right] = loadCustomSprite(55, 96, 25, 8, img)
	headSprites[left] = loadCustomSprite(55, 114, 25, 8, img)
	headSprites[up] = loadCustomSprite(0, 132, 8, 20, img)
	headSprites[down] = loadCustomSprite(9, 132, 8, 20, img)
	headSprites[upRight] = loadCustomSprite(17, 96, 18, 17, img)
	headSprites[upLeft] = loadCustomSprite(17, 114, 18, 17, img)
	headSprites[downRight] = loadCustomSprite(36, 96, 17, 17, img)
	headSprites[downLeft] = loadCustomSprite(36, 114, 17, 17, img)

	carSprites[right] = loadCustomSprite(56, 105, 24, 8, img)
	carSprites[left] = loadCustomSprite(56, 123, 24, 8, img)
	carSprites[up] = loadCustomSprite(18, 132, 8, 19, img)
	carSprites[down] = loadCustomSprite(18, 132, 8, 19, img)
	carSprites[upRight] = loadCustomSprite(0, 96, 16, 17, img)
	carSprites[upLeft] = loadCustomSprite(0, 114, 16, 17, img)
	carSprites[downRight] = loadCustomSprite(0, 114, 16, 17, img)
	carSprites[downLeft] = loadCustomSprite(0, 96, 16, 17, img)

	log.Print("loaded image assets")
}

// loadImage loads image from file
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

// loadCustomSprite loads custom-sized sprite image
func loadCustomSprite(x, y, width, height int, img *ebiten.Image) image.Image {
	return img.SubImage(image.Rect(x, y, x+width, y+height))
}

// loadTileSprite loads a squared tile sprite image
func loadTileSprite(t byte, img *ebiten.Image) image.Image {
	width, _ := img.Size()
	sx, sy := int(t)%(width/TileSize)*TileSize, int(t)/(width/TileSize)*TileSize
	return img.SubImage(image.Rect(sx, sy, sx+TileSize, sy+TileSize))
}

// GetTileSprite returns image for a given sprite offset
func GetTileSprite(t byte) image.Image {
	sprite, ok := tileSprites[t]
	if !ok {
		log.Fatalf("error loading sprite %d", t)
	}
	return sprite
}

// GetHeadSprite returns the sprite for the head of the train
func GetHeadSprite(direction string) image.Image {
	sprite, ok := headSprites[direction]
	if !ok {
		log.Fatalf("error loading head sprite %s", direction)
	}
	return sprite
}

// GetCarSprite returns the sprite for a train's car
func GetCarSprite(direction string) image.Image {
	sprite, ok := carSprites[direction]
	if !ok {
		log.Fatalf("error loading car sprite %s", direction)
	}
	return sprite
}

// SwitchOffset does simple sprite offset arithmetic, which allows to switch from one
// sprite kind to another while maintaining its direction. For example:
// 	SwitchOffset(ul, rail, road)
func SwitchOffset(cur, from, to byte) byte {
	return cur - from + to
}
