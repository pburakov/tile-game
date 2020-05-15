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
	train  = byte(0x0f)
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

var sprites = make(map[byte]image.Image)

func init() {
	log.SetFlags(0)

	var err error
	img, err := loadImage("resource/tiles.png")
	if err != nil {
		log.Fatal(err)
	}
	sprites[none] = loadSprite(none, img)
	sprites[grass] = loadSprite(grass, img)
	sprites[train] = loadSprite(train, img)
	sprites[cursor] = loadSprite(cursor, img)
	sprites[rail+ver] = loadSprite(rail+ver, img)
	sprites[rail+hor] = loadSprite(rail+hor, img)
	sprites[rail+cro] = loadSprite(rail+cro, img)
	sprites[rail+dl] = loadSprite(rail+dl, img)
	sprites[rail+dr] = loadSprite(rail+dr, img)
	sprites[rail+ul] = loadSprite(rail+ul, img)
	sprites[rail+ur] = loadSprite(rail+ur, img)
	sprites[rail+vl] = loadSprite(rail+vl, img)
	sprites[rail+vr] = loadSprite(rail+vr, img)
	sprites[rail+hu] = loadSprite(rail+hu, img)
	sprites[rail+hd] = loadSprite(rail+hd, img)

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

func loadSprite(t byte, img *ebiten.Image) image.Image {
	width, _ := img.Size()
	sx, sy := int(t)%(width/TileSize)*TileSize, int(t)/(width/TileSize)*TileSize
	return img.SubImage(image.Rect(sx, sy, sx+TileSize, sy+TileSize))
}

func getSprite(t byte) image.Image {
	sprite, ok := sprites[t]
	if !ok {
		log.Fatalf("error loading sprite %d", t)
	}
	return sprite
}
