package main

import "C"

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/buke/quickjs-go"
	"github.com/veandco/go-sdl2/sdl"
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

func initIO(ctx *quickjs.Context) {
	console := ctx.Object()
	console.Set("readTextFile", ctx.Function(func(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value {
		f, err := os.Open(args[0].String())
		if err != nil {
			ctx.Null()
		}
		bytes, err := io.ReadAll(f)
		if err != nil {
			ctx.Null()
		}
		return ctx.String(string(bytes))
	}))
	console.Set("writeTextFile", ctx.Function(func(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value {
		f, err := os.Create(args[0].String())
		if err != nil {
			ctx.Null()
		}
		f.WriteString(args[1].String())
		f.Sync()
		return ctx.Null()
	}))
	ctx.Globals().Set("IO", console)
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

var currentFunc = ""
var use_timestamp = false

func main() {
	headless := false
	for _, v := range os.Args {
		if v == "--headless" {
			fmt.Printf("Run in headless mode.")
			headless = true
		}
	}

	// Create a new runtime
	rt := quickjs.NewRuntime(
		quickjs.WithExecuteTimeout(3000),
		quickjs.WithMemoryLimit(128*1024*1024),
		quickjs.WithGCThreshold(128*1024*1024),
		quickjs.WithMaxStackSize(65534*1024),
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

	initConsole(ctx)

	files, _ = os.ReadDir("lib_goverdry")
	for _, f := range files {
		fmt.Println(f.Name())
		ret, err := ctx.EvalFile("lib_goverdry/" + f.Name())
		if err != nil {
			println(err.Error())
		}
		fmt.Println(ret.Get("stack").String())
	}

	initDocument(ctx)
	initWindow(ctx)
	initAnimation(ctx)
	initIO(ctx)
	if headless {
		iniDummySDL(ctx)
	} else {
		iniSDL(ctx)
	}

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
	_, err = ctx.Eval("try { window.onload() } catch (error) {console.log(error); console.log(error.stack)}")
	if err != nil {
		println(err.Error())
	}

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
			_, err := ctx.Eval("try { " + f + " } catch (error) {console.log(error); console.log(error.stack)}")
			if err != nil {
				println(err.Error())
				break
			}
			s = time.Now()
		}
		ctx.Loop()

		// load image
		_, err := ctx.Eval("try { loadImage() } catch (error) {console.log(error); console.log(error.stack)}")
		if err != nil {
			println(err.Error())
			break
		}

		// update window
		resetWindow()
		_, err = ctx.Eval("try { GameBody.rootScene.getCamera().update() } catch (error) {console.log(error); console.log(error.stack)}")
		if err != nil {
			println(err.Error())
			break
		}
		applyWindow()

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			println("PollEvent")
			switch t := event.(type) {
			case *sdl.JoyButtonEvent:
				if t.Button == 1 {
					running = false
				}
			case *sdl.KeyboardEvent:
				if t.Keysym.Sym == sdl.K_q {
					running = false
				}
			case *sdl.QuitEvent: // NOTE: Please use `*sdl.QuitEvent` for `v0.4.x` (current version).
				println("Quit")
				running = false
			}
		}
		sdl.Delay(16)
	}
	ctx.Close()
	rt.Close()
}
