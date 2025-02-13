package main

import "C"

import (
	"encoding/base64"
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

func initBase64(ctx *quickjs.Context) {
	console := ctx.Object()
	console.Set("decode", ctx.Function(func(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value {
		str := args[0].String()
		bytes, err := base64.StdEncoding.DecodeString(str)
		if err != nil {
			panic(err)
		}
		return ctx.String(string(bytes))
	}))
	console.Set("encode", ctx.Function(func(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value {
		str := args[0].String()
		str = base64.StdEncoding.EncodeToString([]byte(str))
		return ctx.String(str)
	}))
	ctx.Globals().Set("base64", console)
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
	document_helper := ctx.Object()
	ctx.Globals().Set("DocumentHelper", document_helper)
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
			fmt.Printf("Run in headless mode.\n")
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

	initConsole(ctx)
	initBase64(ctx)
	initDocument(ctx)
	initAnimation(ctx)
	initIO(ctx)
	if headless {
		iniDummySDL(ctx)
	} else {
		var v sdl.Version
		sdl.VERSION(&v)
		fmt.Printf("SDL Version %d.%d.%d\n", v.Major, v.Minor, v.Patch)

		if err := sdl.Init(sdl.INIT_EVERYTHING | sdl.INIT_JOYSTICK); err != nil {
			panic(err)
		}
		defer sdl.Quit()

		joysticks := sdl.NumJoysticks()
		fmt.Printf("There are %d joysticks connected.\n", joysticks)
		if joysticks > 0 {
			if sdl.JoystickOpen(0) == nil {
				fmt.Printf("There was an error reading from the joystick.\n")
			}
		}

		iniSDL(ctx)
	}

	files, _ := os.ReadDir("lib_goverdry")
	for _, f := range files {
		fmt.Println(f.Name())
		_, err := ctx.EvalFile("lib_goverdry/" + f.Name())
		if err != nil {
			panic(err)
		}
	}

	files, _ = os.ReadDir("lib")
	for _, f := range files {
		fmt.Println(f.Name())
		_, err := ctx.EvalFile("lib/" + f.Name())
		if err != nil {
			panic(err)
		}
	}

	pre_files := []string{"SpellEffect.min.js", "MainPanel.min.js"}
	for _, f := range pre_files {
		fmt.Println(f)
		_, err := ctx.EvalFile("js/" + f)
		if err != nil {
			panic(err)
		}
	}

	files, _ = os.ReadDir("js")
	for _, f := range files {
		if slices.Contains(pre_files, f.Name()) {
			continue
		}
		println(f.Name())
		_, err := ctx.EvalFile("js/" + f.Name())
		if err != nil {
			panic(err)
		}
	}
	ctx.Loop()

	// read test data
	ctx.EvalFile("test_data/gameDataHTML5.js")
	ctx.EvalFile("test_data/defaultMessage_jpn.js")
	ctx.EvalFile("test_data/config.js")

	// Window Events
	_, err := ctx.Eval("try { window.onload() } catch (error) {console.log(error); console.log(error.stack)}")
	if err != nil {
		println(err.Error())
	}
	_, err = ctx.Eval("try { window.emit({ id: 'gamepadconnected',  gamepad: { index: 0 } }) } catch (error) {console.log(error); console.log(error.stack)}")
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
			_, err := ctx.Eval("try { " + f + " } catch (error) {console.log(error); console.log(error.stack)}")
			if err != nil {
				println(err.Error())
				break
			}
			ctx.Globals().Get("navigator").Get("gamepad").Call("clearButtonPressed")
		}
		ctx.Loop()

		// load image
		_, err := ctx.Eval("try { loadImage() } catch (error) {console.log(error); console.log(error.stack)}")
		if err != nil {
			println(err.Error())
			break
		}

		ret, err := ctx.Eval("GameState == 'stopStart'")
		if err != nil {
			println(err.Error())
			break
		}
		if ret.Bool() {
			ctx.Eval("try { document.emit({ id: 'mousedown',  target: { id: 'game_window' }}) } catch (error) {console.log(error); console.log(error.stack)}")
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
			switch t := event.(type) {
			case sdl.JoyButtonEvent:
				if t.State > 0 {
					ctx.Globals().Get("navigator").Get("gamepad").Call("pressButton", ctx.Int32(int32(t.Button)))
				} else {
					ctx.Globals().Get("navigator").Get("gamepad").Call("releaseButton", ctx.Int32(int32(t.Button)))
				}
			case sdl.KeyboardEvent:
				if t.Keysym.Sym == sdl.K_q {
					running = false
				}
			case sdl.QuitEvent:
				println("Quit")
				running = false
			}
		}
		sdl.Delay(33)
	}
	ctx.Close()
	rt.Close()
}
