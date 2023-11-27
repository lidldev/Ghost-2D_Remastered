package assets

import (
	"embed"
	"image"
	_ "image/png"
	"io/fs"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed *
var asset embed.FS

var EbitenBullet = mustLoadImage("Sprite-00010.png")
var Player1FullSprites = mustLoadImages("player1Full/*.png")
var Player1HalfSprites = mustLoadImages("player1Half/*.png")
var Player1LowSprites = mustLoadImages("player1Low/*.png")
var Player2FullSprites = mustLoadImages("player2Full/*.png")
var Player2HalfSprites = mustLoadImages("player2Half/*.png")
var Player2LowSprites = mustLoadImages("player2Half/*.png")

func mustLoadImage(name string) *ebiten.Image {
	f, err := asset.Open(name)
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
	matches, err := fs.Glob(asset, path)
	if err != nil {
		panic(err)
	}

	images := make([]*ebiten.Image, len(matches))
	for i, match := range matches {
		images[i] = mustLoadImage(match)
	}

	return images
}
