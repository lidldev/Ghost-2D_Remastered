package assets

import (
	"embed"
	"image"
	"io/fs"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed *
var asset embed.FS

var ebitenBullet = mustLoadImage("Sprite-00010.png")
var player1FullSprites = mustLoadImages("player1Full/*.png")
var player1HalfSprites = mustLoadImages("player1Half/*.png")
var player1LowSprites = mustLoadImages("player1Low/*.png")
var player2FullSprites = mustLoadImages("player2Full/*.png")
var player2HalfSprites = mustLoadImages("player2Half/*.png")
var player2LowSprites = mustLoadImages("player2Half/*.png")

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
func main() {

}
