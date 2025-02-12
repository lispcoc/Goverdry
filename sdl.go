package main

import (
	"github.com/buke/quickjs-go"
	"github.com/veandco/go-sdl2/gfx"
	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

var SDL_Font *ttf.Font
var SDL_Renderer *sdl.Renderer
var SDL_Window *sdl.Window
var window_ok = false

func SDL_DrawLine(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value {
	if !window_ok {
		return ctx.String("")
	}
	handle := args[0].Int32()
	layer := LayerList[handle]

	SDL_Renderer.SetRenderTarget(layer.texture)
	SDL_Renderer.SetDrawColor(uint8(args[2].Uint32()), uint8(args[3].Uint32()), uint8(args[4].Uint32()), 255)

	for i := 0; i < int(args[1].ToArray().Len()); i++ {
		ret, _ := args[1].ToArray().Get(int64(i))
		x1 := ret.Get("start_x").Int32()
		y1 := ret.Get("start_y").Int32()
		x2 := ret.Get("end_x").Int32()
		y2 := ret.Get("end_y").Int32()
		gfx.ThickLineColor(SDL_Renderer, x1, y1, x2, y2, 2, sdl.Color{uint8(args[2].Uint32()), uint8(args[3].Uint32()), uint8(args[4].Uint32()), 255})
		//SDL_Renderer.DrawLine(x1, y1, x2, y2)
	}

	return ctx.String("")
}

func SDL_Triangle(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value {
	if !window_ok {
		return ctx.String("")
	}
	handle := args[0].Int32()
	layer := LayerList[handle]

	SDL_Renderer.SetRenderTarget(layer.texture)
	SDL_Renderer.SetDrawColor(uint8(args[2].Uint32()), uint8(args[3].Uint32()), uint8(args[4].Uint32()), 255)

	color := sdl.Color{uint8(args[2].Uint32()), uint8(args[3].Uint32()), uint8(args[4].Uint32()), 255}
	var vt [3]sdl.Vertex
	for i := 0; i < 3; i++ {
		ret, _ := args[1].ToArray().Get(int64(i))
		x := float32(ret.Get("x").Float64())
		y := float32(ret.Get("y").Float64())
		vt[i] = sdl.Vertex{sdl.FPoint{x, y}, color, sdl.FPoint{1, 1}}
	}
	v := []sdl.Vertex{vt[0], vt[1], vt[2]}

	SDL_Renderer.RenderGeometry(nil, v, nil)

	return ctx.String("")
}

func SDL_FillText(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value {
	if !window_ok {
		return ctx.String("")
	}
	handle := args[0].Int32()
	layer := LayerList[handle]
	text := args[1].String()
	x := args[2].Int32()
	y := args[3].Int32()
	red := uint8(args[4].Int32())
	green := uint8(args[5].Int32())
	blue := uint8(args[6].Int32())

	surface, _ := SDL_Font.RenderUTF8Solid(text, sdl.Color{R: red, G: green, B: blue, A: 255})
	txt, _ := SDL_Renderer.CreateTextureFromSurface(surface)
	_, _, w, h, _ := txt.Query()
	SDL_Renderer.SetRenderTarget(layer.texture)
	SDL_Renderer.Copy(txt, &sdl.Rect{0, 0, w, h}, &sdl.Rect{x, y, w, h})

	return ctx.String("")
}

func SDL_FillRect(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value {
	if !window_ok {
		return ctx.String("")
	}
	handle := args[0].Int32()
	layer := LayerList[handle]
	x := args[1].Int32()
	y := args[2].Int32()
	w := args[3].Int32()
	h := args[4].Int32()
	red := uint8(args[5].Int32())
	green := uint8(args[6].Int32())
	blue := uint8(args[7].Int32())
	rect := sdl.Rect{x, y, x + w, y + h}
	SDL_Renderer.SetDrawBlendMode(sdl.BLENDMODE_NONE)
	SDL_Renderer.SetDrawColor(red, green, blue, 255)
	SDL_Renderer.SetRenderTarget(layer.texture)
	SDL_Renderer.FillRect(&rect)

	return ctx.String("")
}

func SDL_LayerClear(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value {
	if !window_ok {
		return ctx.String("")
	}
	handle := args[0].Int32()
	layer := LayerList[handle]
	rect := sdl.Rect{layer.x, layer.y, layer.x + layer.w, layer.y + layer.h}
	SDL_Renderer.SetDrawBlendMode(sdl.BLENDMODE_NONE)
	SDL_Renderer.SetDrawColor(0, 0, 0, 0)
	SDL_Renderer.SetRenderTarget(layer.texture)
	SDL_Renderer.FillRect(&rect)
	return ctx.String("")
}

func SDL_DrawSpriteToWindow(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value {
	if !window_ok {
		println("window not ready.")
		return ctx.String("")
	}
	handle := args[0].Int32()
	layer := LayerList[handle]
	surface, _ := SDL_Window.GetSurface()
	dst_rect := sdl.Rect{0, 0, surface.W, surface.H}
	src_rect := sdl.Rect{layer.x, layer.y, layer.x + layer.w, layer.y + layer.h}

	if handle > 1 {
		return ctx.String("")
	}
	SDL_Renderer.SetRenderTarget(nil)
	if err := SDL_Renderer.Copy(layer.texture, &src_rect, &dst_rect); err != nil {
		panic(err)
	}

	SDL_Renderer.Present()

	return ctx.String("")
}

func SDL_CreateWindow(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value {
	sdl.GLSetAttribute(sdl.GL_CONTEXT_PROFILE_MASK, sdl.GL_CONTEXT_PROFILE_ES)
	sdl.GLSetAttribute(sdl.GL_CONTEXT_MAJOR_VERSION, 2)
	sdl.GLSetAttribute(sdl.GL_CONTEXT_MINOR_VERSION, 0)
	sdl.GLSetAttribute(sdl.GL_DOUBLEBUFFER, 1)
	sdl.GLSetAttribute(sdl.GL_DOUBLEBUFFER, 1)

	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, args[0].Int32(), args[1].Int32(), sdl.WINDOW_OPENGL|sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	//window.SetFullscreen(sdl.WINDOW_FULLSCREEN)
	sdl.GLSetSwapInterval(1)
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

	window_ok = true

	ttf.Init()
	SDL_Font, err = ttf.OpenFont("HackGen35Console-Bold.ttf", 14)
	if err != nil {
		panic(err)
	}

	return ctx.String("")
}

type Layer struct {
	texture *sdl.Texture
	x       int32
	y       int32
	w       int32
	h       int32
}

var LayerList []Layer

func SDL_CreateRGBSurface(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value {
	w := args[0].Int32()
	h := args[1].Int32()
	t, _ := SDL_Renderer.CreateTexture(sdl.PIXELFORMAT_RGBA8888, sdl.TEXTUREACCESS_TARGET, w, h)
	t.SetBlendMode(sdl.BLENDMODE_BLEND)
	layer := Layer{t, 0, 0, w, h}
	rect := sdl.Rect{layer.x, layer.y, layer.x + layer.w, layer.y + layer.h}
	SDL_Renderer.SetDrawColor(0, 0, 0, 0)
	SDL_Renderer.SetRenderTarget(layer.texture)
	SDL_Renderer.FillRect(&rect)

	LayerList = append(LayerList, layer)

	return ctx.Int32(int32(len(LayerList) - 1))
}

var WavList []*mix.Chunk
var MusList []*mix.Music

func MIX_LoadWAV(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value {
	cnc, err := mix.LoadWAV(args[0].String())
	if err != nil {
		return ctx.Int32(-1)
	}
	WavList = append(WavList, cnc)
	return ctx.Int32(int32(len(WavList) - 1))
}

func Mix_LoadMUS(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value {
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

func resetWindow() {
	surface, _ := SDL_Window.GetSurface()
	t, _ := SDL_Renderer.CreateTexture(sdl.PIXELFORMAT_RGBA8888, sdl.TEXTUREACCESS_TARGET, surface.W, surface.H)
	t.SetBlendMode(sdl.BLENDMODE_NONE)

	rect := sdl.Rect{0, 0, surface.W, surface.H}
	SDL_Renderer.SetRenderTarget(nil)
	if err := SDL_Renderer.Copy(t, &rect, &rect); err != nil {
		panic(err)
	}
}

func DummyFunction(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value {
	return ctx.Int32(0)
}

func iniDummySDL(ctx *quickjs.Context) {
	SDL := ctx.Object()
	ctx.Globals().Set("SDL", SDL)
	SDL.Set("CreateWindow", ctx.Function(DummyFunction))
	SDL.Set("DrawSpriteToWindow", ctx.Function(DummyFunction))
	SDL.Set("DrawLine", ctx.Function(DummyFunction))
	SDL.Set("FillText", ctx.Function(DummyFunction))
	SDL.Set("FillRect", ctx.Function(DummyFunction))
	SDL.Set("Triangle", ctx.Function(DummyFunction))

	MIX := ctx.Object()
	ctx.Globals().Set("MIX", MIX)
	MIX.Set("LoadWAV", ctx.Function(DummyFunction))
	MIX.Set("LoadMUS", ctx.Function(DummyFunction))
	MIX.Set("PlayChannel", ctx.Function(DummyFunction))
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
	SDL.Set("CreateRGBSurface", ctx.Function(SDL_CreateRGBSurface))
	SDL.Set("LayerClear", ctx.Function(SDL_LayerClear))
	SDL.Set("DrawSpriteToWindow", ctx.Function(SDL_DrawSpriteToWindow))
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
