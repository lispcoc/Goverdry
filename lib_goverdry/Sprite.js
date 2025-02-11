class Sprite {
  constructor (WINDOW_WIDTH, WINDOW_HEIGHT) {
    console.log([this.constructor.name, 'constructor'].join('.'))
    this.WINDOW_WIDTH = WINDOW_WIDTH
    this.WINDOW_HEIGHT = WINDOW_HEIGHT
  }
}
