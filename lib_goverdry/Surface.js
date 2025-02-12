function toRGB (s) {
  console.log('toRGB', s)
  if (s == undefined || s == null) {
    return { r: 255, g: 255, b: 255 }
  }
  if (s == 'red') {
    s = '#ff0000'
  }
  if (s == 'green') {
    s = '#00ff00'
  }
  if (s == 'blue') {
    s = '#0000ff'
  }
  if (s == 'pink') {
    s = '#ffc0cb'
  }
  if (s == 'yellow') {
    s = '#FFFF00'
  }
  if (s == 'black') {
    s = '#000000'
  }
  if (s == 'white') {
    s = '#ffffff'
  }
  if (s[0] != '#') {
    s = '#000000'
    console.log('Invalid colercode.')
  }
  const [r, g, b] = s.replace('#', '').match(/.{2}/g)
  const red = parseInt(r, 16)
  const green = parseInt(g, 16)
  const blue = parseInt(b, 16)
  return { r: red, g: green, b: blue }
}

class Surface extends SceneNode {
  constructor (WINDOW_WIDTH, WINDOW_HEIGHT) {
    super()
    console.log([this.constructor.name, 'constructor'].join('.'))
    this.WINDOW_WIDTH = WINDOW_WIDTH
    this.WINDOW_HEIGHT = WINDOW_HEIGHT
    this.fillStyle = '#ffffff'
    this.handle = SDL.CreateRGBSurface(this.WINDOW_WIDTH, this.WINDOW_HEIGHT)
    this.context = new SurfaceContext(this.handle)
  }
  clear () {
    SDL.LayerClear(this.handle)
  }
}

class SurfaceContext {
  constructor (handle) {
    console.log([this.constructor.name, 'constructor'].join('.'))
    this.handle = handle
    this.x = -1
    this.y = -1
    this.points = []
    this.points_ready = []
  }
  beginPath () {
    console.log([this.constructor.name, 'beginPath'].join('.'))
    this.points = []
    this.points_ready = []
  }
  moveTo (x, y) {
    console.log([this.constructor.name, 'moveTo'].join('.'))
    this.x = x
    this.y = y
    this.points = [{ x: x, y: y }]
  }
  lineTo (x, y) {
    console.log([this.constructor.name, 'lineTo'].join('.'))
    this.x = x
    this.y = y
    this.points.push({ x: x, y: y })
  }
  arc (x, y, r, a, b, c) {
    console.log([this.constructor.name, 'arc'].join('.'))
  }
  closePath () {
    console.log([this.constructor.name, 'closePath'].join('.'))
    this.points_ready = this.points
    this.points = []
  }
  fill () {
    console.log([this.constructor.name, 'fill'].join('.'))
    var color = toRGB(this.fillStyle)

    var vx = this.points_ready.map(e => e.x)
    var vy = this.points_ready.map(e => e.y)
    SDL.FilledPolygonColor(this.handle ,vx, vy, color.r, color.g, color.b)
  }
  stroke () {
    console.log([this.constructor.name, 'stroke'].join('.'))
    if(this.points_ready.length == 0) {
      this.closePath ()
    }
    if(this.points_ready.length == 0) {
      console.log("Error: no points to draw.")
      return
    }
    var color = toRGB(this.strokeStyle)
    var lines = []
    var current_x = this.points_ready[0].x
    var current_y = this.points_ready[0].y
    for (var a of this.points_ready) {
      if (a.x != current_x || a.y != current_y) {
        lines.push({
          start_x: current_x,
          start_y: current_y,
          end_x: a.x,
          end_y: a.y
        })
        current_x = a.x
        current_y = a.y
      }
    }
    lines.push({
      start_x: current_x,
      start_y: current_y,
      end_x: this.points_ready[0].x,
      end_y: this.points_ready[0].y
    })
    SDL.DrawLine(this.handle ,lines, color.r, color.g, color.b)
  }
  fillRect (x, y, w, h) {
    console.log([this.constructor.name, 'fillRect'].join('.'))
    var color = toRGB(this.fillStyle)
    SDL.FillRect(this.handle ,x, y, w, h, color.r, color.g, color.b)
  }
  fillText (Text, x, y) {
    console.log([this.constructor.name, 'fillText', Text, x, y].join('.'))
    var color = toRGB(this.fillStyle)
    SDL.FillText(this.handle ,Text, x, y, color.r, color.g, color.b)
  }
  drawImage(img, x, y) {
    console.log([this.constructor.name, 'drawImage', img, x, y].join('.'))
    console.log("[Todo]drawImage is not implemented.")
  }
}
