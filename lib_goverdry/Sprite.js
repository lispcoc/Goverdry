class Sprite extends SceneNode {
  constructor (WINDOW_WIDTH, WINDOW_HEIGHT) {
    super()
    console.log([this.constructor.name, 'constructor'].join('.'))
    this.WINDOW_WIDTH = WINDOW_WIDTH
    this.WINDOW_HEIGHT = WINDOW_HEIGHT
  }
  update () {
    console.log([this.constructor.name, 'update'].join('.'))
    this.drawToWindow(this.WINDOW_WIDTH, this.WINDOW_HEIGHT)
  }
  drawToWindow () {
    console.log([this.constructor.name, 'drawToWindow'].join('.'))
    SDL.DrawSpriteToWindow(this.image.handle)
  }
}
