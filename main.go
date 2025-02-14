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

func remapKey(sdl2_sym sdl.Keycode) int {
	m := map[sdl.Keycode]int{
		sdl.K_BACKSPACE:    8,
		sdl.K_RETURN:       13,
		sdl.K_LSHIFT:       16,
		sdl.K_RSHIFT:       16,
		sdl.K_RCTRL:        17,
		sdl.K_LCTRL:        17,
		sdl.K_LALT:         18,
		sdl.K_RALT:         18,
		sdl.K_PAUSE:        19,
		sdl.K_ESCAPE:       27,
		sdl.K_SPACE:        32,
		sdl.K_PAGEUP:       33,
		sdl.K_PAGEDOWN:     34,
		sdl.K_END:          35,
		sdl.K_HOME:         36,
		sdl.K_LEFT:         37,
		sdl.K_UP:           38,
		sdl.K_RIGHT:        39,
		sdl.K_DOWN:         40,
		sdl.K_INSERT:       45,
		sdl.K_DELETE:       46,
		sdl.K_0:            48,
		sdl.K_1:            49,
		sdl.K_2:            50,
		sdl.K_3:            51,
		sdl.K_4:            52,
		sdl.K_5:            53,
		sdl.K_6:            54,
		sdl.K_7:            55,
		sdl.K_8:            56,
		sdl.K_9:            57,
		sdl.K_COLON:        58,
		sdl.K_SEMICOLON:    59,
		sdl.K_AT:           64,
		sdl.K_a:            65,
		sdl.K_b:            66,
		sdl.K_c:            67,
		sdl.K_d:            68,
		sdl.K_e:            69,
		sdl.K_f:            70,
		sdl.K_g:            71,
		sdl.K_h:            72,
		sdl.K_i:            73,
		sdl.K_j:            74,
		sdl.K_k:            75,
		sdl.K_l:            76,
		sdl.K_m:            77,
		sdl.K_n:            78,
		sdl.K_o:            79,
		sdl.K_p:            80,
		sdl.K_q:            81,
		sdl.K_r:            82,
		sdl.K_s:            83,
		sdl.K_t:            84,
		sdl.K_u:            85,
		sdl.K_v:            86,
		sdl.K_w:            87,
		sdl.K_x:            88,
		sdl.K_y:            89,
		sdl.K_z:            90,
		sdl.K_KP_0:         96,
		sdl.K_KP_1:         97,
		sdl.K_KP_2:         98,
		sdl.K_KP_3:         99,
		sdl.K_KP_4:         100,
		sdl.K_KP_5:         101,
		sdl.K_KP_6:         102,
		sdl.K_KP_7:         103,
		sdl.K_KP_8:         104,
		sdl.K_KP_9:         105,
		sdl.K_KP_MULTIPLY:  106,
		sdl.K_KP_PLUS:      107,
		sdl.K_KP_MINUS:     109,
		sdl.K_KP_PERIOD:    110,
		sdl.K_KP_DIVIDE:    111,
		sdl.K_F1:           112,
		sdl.K_F2:           113,
		sdl.K_F3:           114,
		sdl.K_F4:           115,
		sdl.K_F5:           116,
		sdl.K_F6:           117,
		sdl.K_F7:           118,
		sdl.K_F8:           119,
		sdl.K_F9:           120,
		sdl.K_F10:          121,
		sdl.K_F11:          122,
		sdl.K_F12:          123,
		sdl.K_SCROLLLOCK:   145,
		sdl.K_CARET:        160,
		sdl.K_MINUS:        173,
		sdl.K_COMMA:        188,
		sdl.K_PERIOD:       190,
		sdl.K_SLASH:        191,
		sdl.K_LEFTBRACKET:  219,
		sdl.K_BACKSLASH:    220,
		sdl.K_RIGHTBRACKET: 221,
	}
	v, r := m[sdl2_sym]
	if r {
		return v
	}
	return 0
}

var saving = false
var savepost = false
var savenum = 0

func saveWorker(ctx *quickjs.Context) {
	if saving {
		println("saveWorker is busy.")
		return
	}
	saving = true
	for true {
		if savepost {
			println("save start.")
			ret := ctx.Globals().Call("getSaveDataStr")
			f, err := os.Create(fmt.Sprintf("savetest.%d.txt", savenum))
			if err != nil {
				println("something wrong.")
				f.Close()
				savepost = false
				continue
			}
			f.WriteString(ret.String())
			f.Sync()
			f.Close()
			savenum++
			println("save succeed.")
			savepost = false
		}
		sdl.Delay(1000)
	}
}

func callSaveWorker(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value {
	savepost = true
	return ctx.Null()
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

	ctx.Globals().Set("callSaveWorker", ctx.Function(callSaveWorker))
	go saveWorker(ctx)

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
	fps := 0
	chktime := time.Now()
	frametime := time.Now()
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
				} else if t.Type == sdl.KEYDOWN {
					e := fmt.Sprintf("{ id: 'keydown',  which: %d}", remapKey(t.Keysym.Sym))
					ctx.Eval("try { document.emit(" + e + ") } catch (error) {console.log(error); console.log(error.stack)}")
				} else {
					e := fmt.Sprintf("{ id: 'keyup',  which: %d}", remapKey(t.Keysym.Sym))
					ctx.Eval("try { document.emit(" + e + ") } catch (error) {console.log(error); console.log(error.stack)}")
				}
			case sdl.QuitEvent:
				println("Quit")
				running = false
			}
		}

		// fps limitter
		for time.Since(frametime).Milliseconds() < 33 {
			sdl.Delay(1)
		}
		frametime = time.Now()
		fps++
		if time.Since(chktime).Milliseconds() >= 1000 {
			fmt.Printf("fps: %d\n", fps)
			fps = 0
			chktime = time.Now()
		}
	}
	ctx.Close()
	rt.Close()
}
