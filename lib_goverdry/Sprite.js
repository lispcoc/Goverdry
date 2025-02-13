class Sprite extends SceneNode {
  constructor (WINDOW_WIDTH, WINDOW_HEIGHT) {
    super()
    console.log(this.constructor.name, 'constructor')
    this.WINDOW_WIDTH = WINDOW_WIDTH
    this.WINDOW_HEIGHT = WINDOW_HEIGHT
  }
  update () {
    this.drawToWindow(this.WINDOW_WIDTH, this.WINDOW_HEIGHT)
  }
  drawToWindow () {
    SDL.DrawSpriteToWindow(this.image.handle)
  }
}
