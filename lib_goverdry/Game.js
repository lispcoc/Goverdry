class Game {
  constructor (windowWidth, windowHeight) {
    console.log(this.constructor.name, 'constructor')
    this.windowWidth = windowWidth
    this.windowHeight = windowHeight
    this.rootScene = new Scene3D()
    this.width = 480
    this.height = 480
  }

  start () {
    console.log(this.constructor.name, 'start')
    SDL.CreateWindow(this.width, this.height)
  }
}
