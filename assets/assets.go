package assets

import (
	"embed"
	"image"
	_ "image/png"
	"io/fs"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed *
var assets embed.FS

var Player1Full = mustLoadImages("player1Full/*.png")
var Player1Half = mustLoadImages("player1Half/*.png")
var Player1Low = mustLoadImages("player1Low/*.png")
var Player2Full = mustLoadImages("player2Full/*.png")
var Player2Half = mustLoadImages("player2Half/*.png")
var Player2Low = mustLoadImages("player2Low/*.png")
var EbitenBullet = mustLoadImage("Sprite-00010.png")

func mustLoadImage(name string) *ebiten.Image {
	f, err := assets.Open(name)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}

	return ebiten.NewImageFromImage(img)
}

func mustLoadImages(path string) []*ebiten.Image {
	matches, err := fs.Glob(assets, path)
	if err != nil {
		panic(err)
	}

	images := make([]*ebiten.Image, len(matches))
	for i, match := range matches {
		images[i] = mustLoadImage(match)
	}

	return images
}
