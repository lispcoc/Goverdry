package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

func showFps(fps int) {
	text := fmt.Sprintf("%d", fps)
	s, _ := SDL_Font.RenderUTF8Solid(text, sdl.Color{R: 255, G: 255, B: 255, A: 255})
	t, _ := SDL_Renderer.CreateTextureFromSurface(s)
	src := sdl.Rect{X: 0, Y: 0, W: s.W, H: s.H}
	dst := sdl.Rect{X: int32(WINDOW_X * 0 / 100), Y: int32(WINDOW_Y * 0 / 100), W: s.W, H: s.H}
	SDL_Renderer.SetRenderTarget(nil)
	SDL_Renderer.Copy(t, &src, &dst)
	t.Destroy()
	s.Free()
}

func showMessage(text string) {
	s, _ := SDL_Font.RenderUTF8Solid(text, sdl.Color{R: 255, G: 255, B: 255, A: 255})
	t, _ := SDL_Renderer.CreateTextureFromSurface(s)
	src := sdl.Rect{X: 0, Y: 0, W: s.W, H: s.H}
	dst := sdl.Rect{X: int32(WINDOW_X * 80 / 100), Y: int32(WINDOW_Y * 0 / 100), W: s.W, H: s.H}
	SDL_Renderer.SetRenderTarget(nil)
	SDL_Renderer.Copy(t, &src, &dst)
	t.Destroy()
	s.Free()
}
