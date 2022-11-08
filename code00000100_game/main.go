package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"image"
	_ "image/png"
	"os"
)

func main() {
	pixelgl.Run(runGame)
}

func runGame() {
	cfg := pixelgl.WindowConfig{
		Title:     "MMO",
		Bounds:    pixel.R(0, 0, 1024, 768),
		VSync:     true,
		Resizable: true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	win.SetSmooth(false)
	manSprite, err := getSprite("./img/img.png")
	if err != nil {
		panic(err)
	}
	manPosition := win.Bounds().Center()

	for !win.JustPressed(pixelgl.KeyEscape) {
		win.Clear(pixel.RGB(0, 0, 0))
		manSprite.Draw(win, pixel.IM.Scaled(pixel.ZV, 2.0).Moved(manPosition))
		win.Update()
	}
}

func getSprite(path string) (*pixel.Sprite, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	//defer File.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	pic := pixel.PictureDataFromImage(img)
	return pixel.NewSprite(pic, pic.Bounds()), nil
}
