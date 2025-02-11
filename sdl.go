package main

import (
	"github.com/buke/quickjs-go"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

var SDL_Font *ttf.Font
var SDL_Renderer *sdl.Renderer
var SDL_Window *sdl.Window

func SDL_FillText(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value {
	//FillText(Text, x, y, red, green, blue)
	text := args[0].String()
	x := args[1].Int32()
	y := args[2].Int32()
	red := uint8(args[3].Int32())
	green := uint8(args[4].Int32())
	blue := uint8(args[5].Int32())

	surface, err := SDL_Font.RenderUTF8Solid(text, sdl.Color{R: red, G: green, B: blue, A: 255})
	if err != nil {
		panic(err)
	}
	texture, _ := SDL_Renderer.CreateTextureFromSurface(surface)
	_, _, width, height, _ := texture.Query()

	txtRect := sdl.Rect{0, 0, width, height}
	pasteRect := sdl.Rect{x, y, width, height}
	SDL_Renderer.Copy(texture, &txtRect, &pasteRect)

	SDL_Renderer.Present()

	return ctx.String("")
}

func SDL_FillRect(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value {
	// FillRect(x, y, w, h, red, green, blue)
	x := args[0].Int32()
	y := args[1].Int32()
	w := args[2].Int32()
	h := args[3].Int32()
	red := uint8(args[4].Int32())
	green := uint8(args[5].Int32())
	blue := uint8(args[6].Int32())

	surface, _ := SDL_Window.GetSurface()

	rect := sdl.Rect{x, y, x + w, y + h}
	colour := sdl.Color{R: red, G: green, B: blue, A: 255} // purple
	pixel := sdl.MapRGBA(surface.Format, colour.R, colour.G, colour.B, colour.A)
	surface.FillRect(&rect, pixel)

	texture, _ := SDL_Renderer.CreateTextureFromSurface(surface)
	SDL_Renderer.Copy(texture, &rect, &rect)

	SDL_Renderer.Present()

	return ctx.String("")
}

func SDL_CreateWindow(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value {
	//sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 800, 600, sdl.WINDOW_SHOWN)
	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, args[0].Int32(), args[1].Int32(), sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	SDL_Window = window

	SDL_Renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}

	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}
	surface.FillRect(nil, 0)

	ttf.Init()
	SDL_Font, err = ttf.OpenFont("HackGen35Console-Bold.ttf", 14)
	if err != nil {
		panic(err)
	}

	return ctx.String("")
}

func iniSDL(ctx *quickjs.Context) {
	//
	// Window
	//
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	SDL := ctx.Object()
	ctx.Globals().Set("SDL", SDL)
	SDL.Set("CreateWindow", ctx.Function(SDL_CreateWindow))
	SDL.Set("FillText", ctx.Function(SDL_FillText))
	SDL.Set("FillRect", ctx.Function(SDL_FillRect))
}
