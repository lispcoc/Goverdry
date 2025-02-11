class Surface extends SceneNode {
  constructor (WINDOW_WIDTH, WINDOW_HEIGHT) {
    super()
    console.log([this.constructor.name, 'constructor'].join('.'))
    this.WINDOW_WIDTH = WINDOW_WIDTH
    this.WINDOW_HEIGHT = WINDOW_HEIGHT
    this.fillStyle = '#ffffff'
    this.context = new SurfaceContext()
  }
}

class SurfaceContext {
  constructor (WINDOW_WIDTH, WINDOW_HEIGHT) {
    console.log([this.constructor.name, 'constructor'].join('.'))
  }
  fillRect (x, y, w, h) {
    console.log([this.constructor.name, 'fillRect'].join('.'))

    if (this.fillStyle[0] != '#') {
      this.fillStyle = '#ffffff'
      console.log('Invalid colercode.')
    }
    const [r, g, b] = this.fillStyle.replace('#', '').match(/.{2}/g)
    const red = parseInt(r, 16)
    const green = parseInt(g, 16)
    const blue = parseInt(b, 16)
    SDL.FillRect(x, y, w, h, red, green, blue)
  }
  fillText (Text, x, y) {
    console.log([this.constructor.name, 'fillText', Text, x, y].join('.'))

    if (this.fillStyle[0] != '#') {
      this.fillStyle = '#ffffff'
      console.log('Invalid colercode.')
    }
    const [r, g, b] = this.fillStyle.replace('#', '').match(/.{2}/g)
    const red = parseInt(r, 16)
    const green = parseInt(g, 16)
    const blue = parseInt(b, 16)
    SDL.FillText(Text, x, y, red, green, blue)
  }
}
