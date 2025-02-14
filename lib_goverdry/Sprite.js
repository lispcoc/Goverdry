class Sprite extends SceneNode {
  constructor (width, height) {
    super()
    console.log(this.constructor.name, 'constructor')
    this.width = width
    this.height = height
  }
  update () {
    this.drawToWindow(this.width, this.height)
  }
  drawToWindow () {
    SDL.DrawSpriteToWindow(this.image.handle, 0, 0, this.width, this.height)
  }
}
