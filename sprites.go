package main

import (
	"github.com/hajimehoshi/ebiten"
	"image"
	"log"
	"os"
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
	trainWidth  = 25
	trainHeight = 8
	trainHorX   = 0
	trainHorY   = 96
)

var tileSprites = make(map[byte]image.Image)
var trainSprite image.Image

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

	trainSprite = loadCustomSprite(trainHorX, trainHorY, trainWidth, trainHeight, img)

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

func getTileSprite(t byte) image.Image {
	sprite, ok := tileSprites[t]
	if !ok {
		log.Fatalf("error loading sprite %d", t)
	}
	return sprite
}