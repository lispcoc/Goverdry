class Overlay {
  constructor () {
    console.log(this.constructor.name, 'constructor')
    this._surface = null
    this.context = null
  }
  setMessage (text, xp, yp) {
    this.update()
    this._surface.context.fillStyle = 'white'
    this._surface.context.fillText(
      text,
      GameBody.width * xp,
      GameBody.height * yp
    )
    this.update()
    SDL.ApplyWindow()
  }
  clearMessage () {
    this._surface.clear()
  }
  clear () {
    this._surface.clear()
  }
  update () {
    if (!this._surface) {
      this._surface = new Surface(GameBody.width, GameBody.height)
      this.context = this._surface.context
    }
    this.drawToWindow()
  }
  drawToWindow () {
    SDL.DrawSpriteToWindow(
      this._surface.handle,
      0,
      0,
      GameBody.width,
      GameBody.height,
      0,
      0,
      GameBody.width,
      GameBody.height
    )
  }
}

var OVERLAY = new Overlay()
