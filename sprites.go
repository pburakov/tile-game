package main

import (
	"image"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten"
)

// sprite ordinal numbers
const (
	none   = byte(0x00)
	grass  = byte(0x19)
	cursor = byte(0x14)

	// To construct a tile, combine the starting tile with an offset number
	// e.g. for crossing: `road + cr` `rail + vl`
	road = byte(0x00)
	rail = byte(0x0f)

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

	headSprites[left] = loadCustomSprite(55, 96, 25, 8, img)
	headSprites[upLeft] = loadCustomSprite(16, 96, 17, 16, img)

	carSprites[left] = loadCustomSprite(56, 105, 25, 8, img)
	carSprites[upLeft] = loadCustomSprite(0, 96, 15, 16, img)

	log.Print("loaded image assets")
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

func loadCustomSprite(x int, y int, width int, height int, img *ebiten.Image) image.Image {
	return img.SubImage(image.Rect(x, y, x+width, y+height))
}

func loadTileSprite(t byte, img *ebiten.Image) image.Image {
	width, _ := img.Size()
	sx, sy := int(t)%(width/TileSize)*TileSize, int(t)/(width/TileSize)*TileSize
	return img.SubImage(image.Rect(sx, sy, sx+TileSize, sy+TileSize))
}

func GetTileSprite(t byte) image.Image {
	sprite, ok := tileSprites[t]
	if !ok {
		log.Fatalf("error loading sprite %d", t)
	}
	return sprite
}

func GetHeadSprite(direction string) image.Image {
	sprite, ok := headSprites[direction]
	if !ok {
		log.Fatalf("error loading head sprite %s", direction)
	}
	return sprite
}

func GetCarSprite(direction string) image.Image {
	sprite, ok := carSprites[direction]
	if !ok {
		log.Fatalf("error loading car sprite %s", direction)
	}
	return sprite
}
