package main

import "C"

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/buke/quickjs-go"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

func initConsole(ctx *quickjs.Context) {
	console := ctx.Object()
	console.Set("log", ctx.Function(func(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value {
		line := args[0].String()
		for i := 1; i < len(args); i++ {
			line += ", " + args[i].String()
		}
		fmt.Println(line)

		return ctx.String(line)
	}))
	ctx.Globals().Set("console", console)
}

func initDocument(ctx *quickjs.Context) {
	document := ctx.Object()

	document.Set("addEventListener", ctx.Function(func(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value {
		fmt.Println("document.addEventListener " + args[0].String())
		return ctx.String("")
	}))

	document.Set("getElementById", ctx.Function(func(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value {
		fmt.Println("document.getElementById " + args[0].String())
		element := ctx.Object()
		element.Set("innerHTML", ctx.String(":"))
		return element
	}))
	document.Set("createElement", ctx.Function(func(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value {
		fmt.Println("document.createElement " + args[0].String())
		element, _ := ctx.Eval("new DocumentElement()")
		return element
	}))

	ctx.Globals().Set("document", document)

	document_location := ctx.Object()
	document.Set("location", document_location)
	document_location.Set("hostname", ctx.String("fake"))
}

func initWindow(ctx *quickjs.Context) {
	window := ctx.Object()

	window.Set("focus", ctx.Function(func(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value {
		fmt.Println("window.focus")
		return ctx.String("")
	}))

	window.Set("addEventListener", ctx.Function(func(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value {
		fmt.Println("window.addEventListener")
		return ctx.String("")
	}))

	ctx.Globals().Set("window", window)
}

func requestAnimationFrame(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value {
	result := strings.Split(args[0].String(), "(")
	result = strings.Split(result[0], " ")
	result2 := strings.Split(args[0].String(), "\n")
	re := regexp.MustCompile(`\(.+\)`)
	fmt.Println("requestAnimationFrame " + result[1])
	use_timestamp = re.MatchString(result2[0])
	currentFunc = result[1]
	return ctx.String("")
}

func cancelAnimationFrame(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value {
	return ctx.String("")
}

func initAnimation(ctx *quickjs.Context) {
	ctx.Globals().Set("requestAnimationFrame", ctx.Function(requestAnimationFrame))
	ctx.Globals().Set("cancelAnimationFrame", ctx.Function(cancelAnimationFrame))
}

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
	SDL := ctx.Object()
	ctx.Globals().Set("SDL", SDL)
	SDL.Set("CreateWindow", ctx.Function(SDL_CreateWindow))
	SDL.Set("FillText", ctx.Function(SDL_FillText))
	SDL.Set("FillRect", ctx.Function(SDL_FillRect))
}

var currentFunc = ""
var use_timestamp = false

func main() {
	// Create a new runtime
	rt := quickjs.NewRuntime(
		quickjs.WithExecuteTimeout(30),
		quickjs.WithMemoryLimit(128*1024*1024),
		quickjs.WithGCThreshold(256*1024),
		quickjs.WithMaxStackSize(65534),
		quickjs.WithCanBlock(true),
	)
	defer rt.Close()

	// Create a new context
	ctx := rt.NewContext()
	defer ctx.Close()

	_, err := ctx.EvalFile("test_data/gameDataHTML5.js")
	if err != nil {
		println(err.Error())
	}

	files, _ := os.ReadDir("lib")
	for _, f := range files {
		fmt.Println(f.Name())
		ret, err := ctx.EvalFile("lib/" + f.Name())
		if err != nil {
			println(err.Error())
		}
		fmt.Println(ret.Get("stack").String())
	}

	files, _ = os.ReadDir("lib_goverdry")
	for _, f := range files {
		fmt.Println(f.Name())
		ret, err := ctx.EvalFile("lib_goverdry/" + f.Name())
		if err != nil {
			println(err.Error())
		}
		fmt.Println(ret.Get("stack").String())
	}

	initConsole(ctx)
	initDocument(ctx)
	initWindow(ctx)
	initAnimation(ctx)
	iniSDL(ctx)

	pre_files := []string{"SpellEffect.min.js"}
	for _, f := range pre_files {
		fmt.Println(f)
		ret, err := ctx.EvalFile("js/" + f)
		if err != nil {
			println(err.Error())
			println(ret.Error().Error())
		}
		fmt.Println(ret.String())
		fmt.Println(ret.Get("stack").String())
	}

	files, _ = os.ReadDir("js")
	for _, f := range files {
		if slices.Contains(pre_files, f.Name()) {
			continue
		}
		fmt.Println(f.Name())
		ret, err := ctx.EvalFile("js/" + f.Name())
		if err != nil {
			println(err.Error())
			println(ret.Error())
		}
		fmt.Println(ret.String())
	}
	ctx.Loop()

	// Window Events
	_, err = ctx.Eval("window.onload()")
	if err != nil {
		println(err.Error())
	}

	//
	// Window
	//
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	println("main loop")
	running := true
	s := time.Now()
	for running {
		elapsed := time.Since(s)
		if currentFunc != "" {
			f := currentFunc + "()"
			if use_timestamp {
				f = currentFunc + "(" + strconv.FormatInt(elapsed.Milliseconds(), 10) + ")"
			}
			println(f)
			_, err := ctx.Eval(f)
			if err != nil {
				println(err.Error())
				break
			}
			s = time.Now()
		}

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent: // NOTE: Please use `*sdl.QuitEvent` for `v0.4.x` (current version).
				println("Quit")
				running = false
				break
			}
		}

		sdl.Delay(33)
	}
}
