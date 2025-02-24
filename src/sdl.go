package main

import (
	"time"
	"unsafe"

	"github.com/buke/quickjs-go"
	"github.com/veandco/go-sdl2/gfx"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

const FONT_SIZE = 12

var SDL_Font *ttf.Font
var SDL_Renderer *sdl.Renderer
var SDL_Window *sdl.Window
var workspace *sdl.Texture
var custom sdl.BlendMode
var window_ok = false
var WINDOW_X int32
var WINDOW_Y int32
var USE_SOFTWARE_RENDER = false
var FONT_RENDER_SIZE = int32(14)
var FONT_ASPECT = 0.5

func SDL_DrawLine(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value {
	if !window_ok {
		return ctx.String("")
	}
	handle := args[0].Int32()
	layer := LayerList[handle]

	SDL_Renderer.SetRenderTarget(layer.texture)

	for i := 0; i < int(args[1].ToArray().Len()); i++ {
		ret, _ := args[1].ToArray().Get(int64(i))
		x1 := ret.Get("x1").Int32()
		y1 := ret.Get("y1").Int32()
		x2 := ret.Get("x2").Int32()
		y2 := ret.Get("y2").Int32()
		color := sdl.Color{R: uint8(args[2].Uint32()), G: uint8(args[3].Uint32()), B: uint8(args[4].Uint32()), A: 255}
		gfx.ThickLineColor(SDL_Renderer, x1, y1, x2, y2, 2, color)
	}

	return ctx.Null()
}

func SDL_Triangle(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value {
	if !window_ok {
		return ctx.String("")
	}
	handle := args[0].Int32()
	layer := LayerList[handle]

	SDL_Renderer.SetRenderTarget(layer.texture)

	color := sdl.Color{R: uint8(args[2].Uint32()), G: uint8(args[3].Uint32()), B: uint8(args[4].Uint32()), A: 255}
	var vt [3]sdl.Vertex
	for i := 0; i < 3; i++ {
		ret, _ := args[1].ToArray().Get(int64(i))
		x := float32(ret.Get("x").Float64())
		y := float32(ret.Get("y").Float64())
		vt[i] = sdl.Vertex{Position: sdl.FPoint{X: x, Y: y}, Color: color, TexCoord: sdl.FPoint{X: 1, Y: 1}}
	}
	v := []sdl.Vertex{vt[0], vt[1], vt[2]}

	SDL_Renderer.RenderGeometry(nil, v, nil)

	return ctx.Null()
}

func SDL_FilledPolygonColor(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value {
	if !window_ok {
		return ctx.String("")
	}
	handle := args[0].Int32()
	layer := LayerList[handle]

	SDL_Renderer.SetRenderTarget(layer.texture)

	var vx []int16
	var vy []int16
	for i := 0; i < int(args[1].ToArray().Len()); i++ {
		v, _ := args[1].ToArray().Get(int64(i))
		vx = append(vx, int16(v.Int32()))
	}
	for i := 0; i < int(args[2].ToArray().Len()); i++ {
		v, _ := args[2].ToArray().Get(int64(i))
		vy = append(vy, int16(v.Int32()))
	}
	color := sdl.Color{R: uint8(args[3].Uint32()), G: uint8(args[4].Uint32()), B: uint8(args[5].Uint32()), A: 255}
	gfx.FilledPolygonColor(SDL_Renderer, vx, vy, color)

	return ctx.Null()
}

func SDL_FilledPolygonImage(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value {
	if !window_ok {
		return ctx.String("")
	}
	handle := args[0].Int32()
	img_handle := args[1].Int32()
	src_x := args[2].Int32()
	src_y := args[3].Int32()
	src_w := args[4].Int32()
	src_h := args[5].Int32()
	dst_x := args[6].Int32()
	dst_y := args[7].Int32()
	dst_w := args[8].Int32()
	dst_h := args[9].Int32()
	var vx []int16
	var vy []int16
	for i := 0; i < int(args[10].ToArray().Len()); i++ {
		v, _ := args[10].ToArray().Get(int64(i))
		vx = append(vx, int16(v.Int32()))
	}
	for i := 0; i < int(args[11].ToArray().Len()); i++ {
		v, _ := args[11].ToArray().Get(int64(i))
		vy = append(vy, int16(v.Int32()))
	}
	SDL_Renderer.SetRenderTarget(workspace)
	SDL_Renderer.SetDrawColor(0, 0, 0, 0)
	SDL_Renderer.Clear()
	color := sdl.Color{R: uint8(args[12].Uint32()), G: uint8(args[13].Uint32()), B: uint8(args[14].Uint32()), A: 255}
	gfx.FilledPolygonColor(SDL_Renderer, vx, vy, color)

	src := sdl.Rect{X: src_x, Y: src_y, W: src_w, H: src_h}
	dst := sdl.Rect{X: dst_x, Y: dst_y, W: dst_w, H: dst_h}
	img := LayerList[img_handle].texture
	img.SetBlendMode(custom)
	SDL_Renderer.Copy(img, &src, &dst)

	layer := LayerList[handle]
	SDL_Renderer.SetRenderTarget(layer.texture)
	workspace.SetBlendMode(sdl.BLENDMODE_BLEND)
	SDL_Renderer.Copy(workspace, &dst, &dst)

	return ctx.Null()
}

func SDL_FillText(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value {
	if !window_ok {
		return ctx.Null()
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
	SDL_Renderer.Copy(txt, &sdl.Rect{X: 0, Y: 0, W: w, H: h}, &sdl.Rect{X: x, Y: y, W: w, H: h})

	txt.Destroy()
	surface.Free()

	return ctx.Null()
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
	rect := sdl.Rect{X: x, Y: y, W: w, H: h}
	SDL_Renderer.SetDrawBlendMode(sdl.BLENDMODE_NONE)
	SDL_Renderer.SetDrawColor(red, green, blue, 255)
	SDL_Renderer.SetRenderTarget(layer.texture)
	SDL_Renderer.FillRect(&rect)

	return ctx.Null()
}

func SDL_LayerClear(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value {
	if !window_ok {
		return ctx.String("")
	}
	handle := args[0].Int32()
	layer := LayerList[handle]
	SDL_Renderer.SetRenderTarget(layer.texture)
	SDL_Renderer.SetDrawColor(0, 0, 0, 0)
	SDL_Renderer.Clear()

	return ctx.Null()
}

func SDL_DrawSpriteToWindow(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value {
	if !window_ok {
		println("window not ready.")
		return ctx.Null()
	}
	handle := args[0].Int32()
	x1 := args[1].Int32()
	y1 := args[2].Int32()
	w1 := args[3].Int32()
	h1 := args[4].Int32()
	x2 := args[5].Int32()
	y2 := args[6].Int32()
	w2 := args[7].Int32()
	h2 := args[8].Int32()
	layer := LayerList[handle]
	src_rect := sdl.Rect{X: x1, Y: y1, W: w1, H: h1}
	dst_rect := sdl.Rect{X: x2, Y: y2, W: w2, H: h2}

	SDL_Renderer.SetRenderTarget(nil)
	if err := SDL_Renderer.Copy(layer.texture, &src_rect, &dst_rect); err != nil {
		panic(err)
	}

	return ctx.Null()
}

func SDL_Copy(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value {
	if !window_ok {
		println("window not ready.")
		return ctx.Null()
	}
	src_handle := args[0].Int32()
	dst_handle := args[1].Int32()
	x1 := args[2].Int32()
	y1 := args[3].Int32()
	w1 := args[4].Int32()
	h1 := args[5].Int32()
	x2 := args[6].Int32()
	y2 := args[7].Int32()
	w2 := args[8].Int32()
	h2 := args[9].Int32()

	src_rect := sdl.Rect{X: x1, Y: y1, W: w1, H: h1}
	dst_rect := sdl.Rect{X: x2, Y: y2, W: w2, H: h2}
	SDL_Renderer.SetRenderTarget(LayerList[dst_handle].texture)
	SDL_Renderer.Copy(LayerList[src_handle].texture, &src_rect, &dst_rect)

	return ctx.Null()
}

func SDL_CreateWindow(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value {
	if !USE_SOFTWARE_RENDER {
		sdl.GLSetAttribute(sdl.GL_CONTEXT_PROFILE_MASK, sdl.GL_CONTEXT_PROFILE_ES)
		sdl.GLSetAttribute(sdl.GL_CONTEXT_MAJOR_VERSION, 3)
		sdl.GLSetAttribute(sdl.GL_CONTEXT_MINOR_VERSION, 2)
		sdl.GLSetAttribute(sdl.GL_DOUBLEBUFFER, 1)
		sdl.GLSetAttribute(sdl.GL_RED_SIZE, 8)
		sdl.GLSetAttribute(sdl.GL_GREEN_SIZE, 8)
		sdl.GLSetAttribute(sdl.GL_BLUE_SIZE, 8)
		sdl.GLSetAttribute(sdl.GL_ALPHA_SIZE, 8)
		sdl.GLSetAttribute(sdl.GL_DEPTH_SIZE, 24)
	}
	WINDOW_X = args[0].Int32()
	WINDOW_Y = args[1].Int32()

	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, WINDOW_X, WINDOW_Y, sdl.WINDOW_OPENGL|sdl.WINDOW_SHOWN)
	if err != nil {
		println(err.Error())
		panic(err)
	}
	sdl.GLSetSwapInterval(1)
	SDL_Window = window

	flag := sdl.RENDERER_ACCELERATED | sdl.RENDERER_TARGETTEXTURE
	if USE_SOFTWARE_RENDER {
		flag = sdl.RENDERER_SOFTWARE | sdl.RENDERER_TARGETTEXTURE
	}
	SDL_Renderer, err = sdl.CreateRenderer(window, -1, flag)
	if err != nil {
		println(err.Error())
		panic(err)
	}

	workspace, err = SDL_Renderer.CreateTexture(sdl.PIXELFORMAT_RGBA8888, sdl.TEXTUREACCESS_TARGET, WINDOW_X*4, WINDOW_Y*4)
	if err != nil {
		println(err.Error())
		panic(err)
	}

	custom = sdl.ComposeCustomBlendMode(sdl.BLENDFACTOR_ONE, sdl.BLENDFACTOR_ZERO, sdl.BLENDOPERATION_ADD,
		sdl.BLENDFACTOR_ONE, sdl.BLENDFACTOR_ONE, sdl.BLENDOPERATION_MINIMUM)

	resetWindow()
	applyWindow()

	window_ok = true

	ttf.Init()
	SDL_Font, err = ttf.OpenFont("JF-Dot-MPlus12.ttf", FONT_SIZE)
	if err != nil {
		println(err.Error())
		panic(err)
	}
	ts, err := SDL_Font.RenderUTF8Solid("a", sdl.Color{R: 0, G: 0, B: 0, A: 255})
	if err != nil {
		println(err.Error())
		panic(err)
	} else {
		FONT_RENDER_SIZE = ts.H
		FONT_ASPECT = float64(ts.W) / float64(ts.H)
	}
	ts.Free()

	return ctx.Null()
}

func SDL_QueryFontSize(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value {
	return ctx.Int32(FONT_RENDER_SIZE)
}

func SDL_QueryFontAspect(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value {
	return ctx.Float64(FONT_ASPECT)
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
	rect := sdl.Rect{X: layer.x, Y: layer.y, W: layer.w, H: layer.h}
	SDL_Renderer.SetDrawColor(0, 0, 0, 0)
	SDL_Renderer.SetRenderTarget(layer.texture)
	SDL_Renderer.FillRect(&rect)

	LayerList = append(LayerList, layer)

	return ctx.Int32(int32(len(LayerList) - 1))
}

func SDL_DrawImage(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value {
	if !window_ok {
		return ctx.String("")
	}
	handle := args[0].Int32()
	src := args[1].String()
	src_x := args[2].Int32()
	src_y := args[3].Int32()
	src_w := args[4].Int32()
	src_h := args[5].Int32()
	dst_x := args[6].Int32()
	dst_y := args[7].Int32()
	dst_w := args[8].Int32()
	dst_h := args[9].Int32()
	layer := LayerList[handle]
	SDL_Renderer.SetRenderTarget(layer.texture)

	var t *sdl.Texture
	var err error
	if USE_SOFTWARE_RENDER {
		s, err := img.Load(src)
		if err != nil {
			return ctx.Null()
		}
		t, err = SDL_Renderer.CreateTextureFromSurface(s)
		if err != nil {
			return ctx.Null()
		}
		s.Free()
	} else {
		t, err = img.LoadTexture(SDL_Renderer, src)
		if err != nil {
			return ctx.Null()
		}
	}
	t.SetBlendMode(sdl.BLENDMODE_BLEND)
	SDL_Renderer.Copy(t, &sdl.Rect{X: src_x, Y: src_y, W: src_w, H: src_h}, &sdl.Rect{X: dst_x, Y: dst_y, W: dst_w, H: dst_h})
	t.Destroy()

	return ctx.Null()
}

func SDL_ApplyWindow(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value {
	applyWindow()
	return ctx.Null()
}

func IMG_SaveFileInternal(handle int32, file string) {
	layer := LayerList[handle]
	SDL_Renderer.SetRenderTarget(layer.texture)

	_, _, w, h, _ := layer.texture.Query()
	s, _ := sdl.CreateRGBSurface(0, w, h, 32, 0xFF000000, 0x00FF0000, 0x0000FF00, 0x000000FF)

	SDL_Renderer.ReadPixels(nil, s.Format.Format, unsafe.Pointer(&s.Pixels()[0]), int(s.Pitch))
	img.SavePNG(s, file)
	s.Free()
}

func IMG_SaveFile(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value {
	handle := args[0].Int32()
	file := args[1].String()
	IMG_SaveFileInternal(handle, file)

	return ctx.Null()
}

var iimgn = 0

func IMG_Load(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value {
	ret := ctx.Object()
	if USE_SOFTWARE_RENDER {
		s, err := img.Load(args[0].String())
		if err != nil {
			return ctx.Null()
		}
		ret.Set("w", ctx.Int32(s.W))
		ret.Set("h", ctx.Int32(s.H))
		iimgn++
		s.Free()
	} else {
		t, err := img.LoadTexture(SDL_Renderer, args[0].String())
		if err != nil {
			println(err.Error())
			return ctx.Null()
		}
		_, _, w, h, err := t.Query()
		if err != nil {
			return ctx.Null()
		}
		ret.Set("w", ctx.Int32(w))
		ret.Set("h", ctx.Int32(h))
		t.Destroy()

	}

	return ret
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
	volume := int(args[1].Float64() * mix.MAX_VOLUME)
	mix.VolumeMusic(volume)
	playMusic(args[0].Int32(), 0, 0)
	return ctx.Null()
}

var current_music int32

type MusicStack struct {
	handle   int32
	pos      float64
	lasttime time.Time
}

var music_stack []MusicStack

func playMusic(handle int32, loop int, pos float64) {
	if mix.PlayingMusic() {
		music_stack = append(music_stack, MusicStack{handle: current_music, pos: mix.GetMusicPosition(MusList[current_music]), lasttime: time.Now()})
	}
	MusList[handle].Play(loop)
	mix.SetMusicPosition(int64(pos))
	current_music = handle
}

func musicFinished() {
	if len(music_stack) > 0 {
		m := music_stack[len(music_stack)-1]
		pos := m.pos + time.Since(m.lasttime).Seconds()
		playMusic(m.handle, 0, pos)
		music_stack = music_stack[:len(music_stack)-1]
	}
}

func resetWindow() {
	SDL_Renderer.SetRenderTarget(nil)
	SDL_Renderer.SetDrawColor(0, 0, 0, 0)
	SDL_Renderer.Clear()
}

func applyWindow() {
	SDL_Renderer.SetRenderTarget(nil)
	SDL_Renderer.Present()
}

func DummyFunction(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value {
	return ctx.Int32(0)
}

func IMG_LoadDummy(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value {
	ret := ctx.Object()
	ret.Set("w", ctx.Int32(1))
	ret.Set("h", ctx.Int32(1))
	return ret
}

func initSDLHeadless(ctx *quickjs.Context) {
	SDL := ctx.Object()
	ctx.Globals().Set("SDL", SDL)
	SDL.Set("ApplyWindow", ctx.Function(DummyFunction))
	SDL.Set("Copy", ctx.Function(DummyFunction))
	SDL.Set("CreateRGBSurface", ctx.Function(SDL_CreateRGBSurface))
	SDL.Set("CreateWindow", ctx.Function(DummyFunction))
	SDL.Set("DrawImage", ctx.Function(DummyFunction))
	SDL.Set("DrawLine", ctx.Function(DummyFunction))
	SDL.Set("DrawSpriteToWindow", ctx.Function(DummyFunction))
	SDL.Set("FilledPolygonColor", ctx.Function(DummyFunction))
	SDL.Set("FilledPolygonImage", ctx.Function(DummyFunction))
	SDL.Set("FillRect", ctx.Function(DummyFunction))
	SDL.Set("FillText", ctx.Function(DummyFunction))
	SDL.Set("LayerClear", ctx.Function(DummyFunction))
	SDL.Set("QueryFontAspect", ctx.Function(DummyFunction))
	SDL.Set("QueryFontSize", ctx.Function(DummyFunction))
	SDL.Set("Triangle", ctx.Function(DummyFunction))

	// image
	img.Init(img.INIT_JPG | img.INIT_PNG)
	IMG := ctx.Object()
	ctx.Globals().Set("IMG", IMG)
	IMG.Set("Load", ctx.Function(IMG_LoadDummy))
	IMG.Set("SaveFile", ctx.Function(IMG_SaveFile))

	// sound
	mix.Init(mix.INIT_MP3 | mix.INIT_OGG)
	if err := mix.OpenAudio(48000, mix.DEFAULT_FORMAT, 1, 4096); err != nil {
		mix.Quit()
		panic(err)
	}
	MIX := ctx.Object()
	ctx.Globals().Set("MIX", MIX)
	MIX.Set("LoadMUS", ctx.Function(DummyFunction))
	MIX.Set("LoadWAV", ctx.Function(DummyFunction))
	MIX.Set("PlayChannel", ctx.Function(DummyFunction))
}

func initSDL(ctx *quickjs.Context) {
	SDL := ctx.Object()
	ctx.Globals().Set("SDL", SDL)
	SDL.Set("ApplyWindow", ctx.Function(SDL_ApplyWindow))
	SDL.Set("Copy", ctx.Function(SDL_Copy))
	SDL.Set("CreateRGBSurface", ctx.Function(SDL_CreateRGBSurface))
	SDL.Set("CreateWindow", ctx.Function(SDL_CreateWindow))
	SDL.Set("DrawImage", ctx.Function(SDL_DrawImage))
	SDL.Set("DrawLine", ctx.Function(SDL_DrawLine))
	SDL.Set("DrawSpriteToWindow", ctx.Function(SDL_DrawSpriteToWindow))
	SDL.Set("FilledPolygonColor", ctx.Function(SDL_FilledPolygonColor))
	SDL.Set("FilledPolygonImage", ctx.Function(SDL_FilledPolygonImage))
	SDL.Set("FillRect", ctx.Function(SDL_FillRect))
	SDL.Set("FillText", ctx.Function(SDL_FillText))
	SDL.Set("LayerClear", ctx.Function(SDL_LayerClear))
	SDL.Set("QueryFontAspect", ctx.Function(SDL_QueryFontAspect))
	SDL.Set("QueryFontSize", ctx.Function(SDL_QueryFontSize))
	SDL.Set("Triangle", ctx.Function(SDL_Triangle))

	// image
	img.Init(img.INIT_JPG | img.INIT_PNG)
	IMG := ctx.Object()
	ctx.Globals().Set("IMG", IMG)
	IMG.Set("Load", ctx.Function(IMG_Load))
	IMG.Set("SaveFile", ctx.Function(IMG_SaveFile))

	// sound
	mix.Init(mix.INIT_MP3 | mix.INIT_OGG)
	if err := mix.OpenAudio(48000, mix.DEFAULT_FORMAT, 1, 4096); err != nil {
		mix.Quit()
		panic(err)
	}
	MIX := ctx.Object()
	ctx.Globals().Set("MIX", MIX)
	MIX.Set("LoadMUS", ctx.Function(Mix_LoadMUS))
	MIX.Set("LoadWAV", ctx.Function(MIX_LoadWAV))
	MIX.Set("PlayChannel", ctx.Function(Mix_PlayChannel))
	mix.HookMusicFinished(musicFinished)
}
