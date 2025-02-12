package main

import (
	"github.com/buke/quickjs-go"
	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

var SDL_Font *ttf.Font
var SDL_Renderer *sdl.Renderer
var SDL_Window *sdl.Window

func SDL_DrawLine(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value {
	// DrawLine(this.lines_ready, red, green, blue)
	SDL_Renderer.SetDrawColor(uint8(args[1].Uint32()), uint8(args[2].Uint32()), uint8(args[3].Uint32()), 255)

	for i := 0; i < int(args[0].ToArray().Len()); i++ {
		ret, _ := args[0].ToArray().Get(int64(i))
		x1 := ret.Get("start_x").Int32()
		y1 := ret.Get("start_y").Int32()
		x2 := ret.Get("end_x").Int32()
		y2 := ret.Get("end_y").Int32()
		SDL_Renderer.DrawLine(x1, y1, x2, y2)
	}

	return ctx.String("")
}

func SDL_Triangle(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value {
	// DrawLine(this.lines_ready, red, green, blue)
	SDL_Renderer.SetDrawColor(uint8(args[1].Uint32()), uint8(args[2].Uint32()), uint8(args[3].Uint32()), 255)

	color := sdl.Color{uint8(args[1].Uint32()), uint8(args[2].Uint32()), uint8(args[3].Uint32()), 255}
	var vt [3]sdl.Vertex
	for i := 0; i < 3; i++ {
		ret, _ := args[0].ToArray().Get(int64(i))
		x := float32(ret.Get("x").Float64())
		y := float32(ret.Get("y").Float64())
		vt[i] = sdl.Vertex{sdl.FPoint{x, y}, color, sdl.FPoint{1, 1}}
	}
	v := []sdl.Vertex{vt[0], vt[1], vt[2]}

	SDL_Renderer.RenderGeometry(nil, v, nil)

	return ctx.String("")
}

func SDL_FillText(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value {
	//FillText(Text, x, y, red, green, blue)
	text := args[0].String()
	x := args[1].Int32()
	y := args[2].Int32()
	red := uint8(args[3].Int32())
	green := uint8(args[4].Int32())
	blue := uint8(args[5].Int32())

	surface, _ := SDL_Font.RenderUTF8Solid(text, sdl.Color{R: red, G: green, B: blue, A: 255})
	texture, _ := SDL_Renderer.CreateTextureFromSurface(surface)
	_, _, width, height, _ := texture.Query()

	txtRect := sdl.Rect{0, 0, width, height}
	pasteRect := sdl.Rect{x, y, width, height}
	SDL_Renderer.Copy(texture, &txtRect, &pasteRect)

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
	colour := sdl.Color{R: red, G: green, B: blue, A: 255}
	pixel := sdl.MapRGBA(surface.Format, colour.R, colour.G, colour.B, colour.A)
	surface.FillRect(&rect, pixel)

	texture, _ := SDL_Renderer.CreateTextureFromSurface(surface)
	SDL_Renderer.Copy(texture, &rect, &rect)

	return ctx.String("")
}

func SDL_CreateWindow(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value {
	//sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 800, 600, sdl.WINDOW_SHOWN)
	sdl.GLSetAttribute(sdl.GL_CONTEXT_MAJOR_VERSION, sdl.GL_CONTEXT_PROFILE_ES);
	sdl.GLSetAttribute(sdl.GL_CONTEXT_PROFILE_MASK, 2);
	sdl.GLSetAttribute(sdl.GL_CONTEXT_MINOR_VERSION, 0);
	sdl.GLSetAttribute(sdl.GL_DOUBLEBUFFER, 1);

	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, args[0].Int32(), args[1].Int32(), sdl.WINDOW_OPENGL | sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	window.SetFullscreen(sdl.WINDOW_FULLSCREEN)
	sdl.GLSetSwapInterval(1)
	SDL_Window = window

	SDL_Renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_SOFTWARE)
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

var WavList []*mix.Chunk
var MusList []*mix.Music

func MIX_LoadWAV(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value {
	println(args[0].String())
	cnc, err := mix.LoadWAV(args[0].String())
	if err != nil {
		return ctx.Int32(-1)
	}
	WavList = append(WavList, cnc)
	println(args[0].String())
	return ctx.Int32(int32(len(WavList) - 1))
}

func Mix_LoadMUS(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value {
	println(args[0].String())
	mus, err := mix.LoadMUS(args[0].String())
	if err != nil {
		return ctx.Int32(-1)
	}
	MusList = append(MusList, mus)
	return ctx.Int32(int32(len(MusList) - 1))
}

func Mix_PlayChannel(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value {
	mus := MusList[args[0].Int32()]
	err := mus.Play(-1)
	if err != nil {
		println(err.Error())
		return ctx.Bool(false)
	}
	return ctx.Bool(true)
}

func updateWindow() {
	SDL_Renderer.Present()
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
	SDL.Set("DrawLine", ctx.Function(SDL_DrawLine))
	SDL.Set("FillText", ctx.Function(SDL_FillText))
	SDL.Set("FillRect", ctx.Function(SDL_FillRect))
	SDL.Set("Triangle", ctx.Function(SDL_Triangle))

	// sound
	mix.Init(mix.INIT_MP3 | mix.INIT_OGG)
	if err := mix.OpenAudio(48000, mix.DEFAULT_FORMAT, 1, 4096); err != nil {
		mix.Quit()
		panic(err)
	}
	mus, err := mix.LoadMUS("alarm.mp3")
	if err == nil {
		mus.Play(-1)
	}

	MIX := ctx.Object()
	ctx.Globals().Set("MIX", MIX)
	MIX.Set("LoadWAV", ctx.Function(MIX_LoadWAV))
	MIX.Set("LoadMUS", ctx.Function(Mix_LoadMUS))
	MIX.Set("PlayChannel", ctx.Function(Mix_PlayChannel))
}
