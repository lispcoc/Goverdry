class Game {
  constructor (windowWidth, windowHeight) {
    console.log([this.constructor.name, 'constructor'].join('.'))
    this.windowWidth = windowWidth
    this.windowHeight = windowHeight
    this.rootScene = new Scene3D()
    this.width = 480
    this.height = 480
  }

  start () {
    console.log([this.constructor.name, 'start'].join('.'))
    SDL.CreateWindow(this.width, this.height)
  }
}
