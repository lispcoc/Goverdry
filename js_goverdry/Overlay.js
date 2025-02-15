class Overlay {
  constructor (){
    console.log(
      this.constructor.name,
      'constructor'
    )
    this._surface = null
  }
  setMessage (text, xp, yp) {
    console.log(this.constructor.name, "setMessage")
    this.update ()
    this._surface.context.fillStyle = "white"
    this._surface.context.fillText(text, GameBody.width * xp, GameBody.height * yp)
    SDL.ApplyWindow()
  }
  clearMessage () {
    console.log(this.constructor.name, "clearMessage")
    this._surface.clear()
  }
  update () {
    if (!this._surface) {
      this._surface = new Surface (GameBody.width, GameBody.height)
    }
    this.drawToWindow()
  }
  drawToWindow () {
    SDL.DrawSpriteToWindow(this._surface.handle, 0, 0, GameBody.width, GameBody.height, 0, 0, GameBody.width, GameBody.height)
  }
}

var OVERLAY = new Overlay ()
