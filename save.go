package main

import (
	"github.com/buke/quickjs-go"
)

var SaveRuntime quickjs.Runtime
var SaveContext *quickjs.Context

func initSaveWorker(rt quickjs.Runtime, mainctx *quickjs.Context) {
	SaveRuntime = rt
	SaveContext = SaveRuntime.NewContext()
	defer SaveContext.Close()

	initConsole(SaveContext)
	initBase64(SaveContext)
	initIO(SaveContext)

	_, err := SaveContext.EvalFile("lib_goverdry/LocalStorage.js")
	if err != nil {
		panic(err)
	}

	save := mainctx.Object()
	save.Set("saveWorker", mainctx.Function(func(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value {
		file := args[0].String()
		json_str := args[1].String()
		go saveWorker(file, json_str)
		return ctx.Null()
	}))
	mainctx.Globals().Set("Save", save)
	println("saveWorker")
}

var Saving = false

func saveWorker(file, json_str string) {
	if Saving {
		println("saveWorker locked")
		return
	}
	Saving = true
	println("saveWorker start")
	SaveContext.Globals().Set("json_str", SaveContext.String(json_str))
	SaveContext.Globals().Set("file", SaveContext.String(file))
	ret, err := SaveContext.EvalFile("js_coworker/save.js")
	if err != nil {
		println(err.Error())
	}
	ret.Free()
	Saving = false
	println("saveWorker end")
}
