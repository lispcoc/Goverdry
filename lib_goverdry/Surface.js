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
    this.lines = []
  }
  beginPath () {
    console.log([this.constructor.name, 'beginPath'].join('.'))
    this.points = []
    this.lines = []
  }
  moveTo (x, y) {
    console.log([this.constructor.name, 'moveTo'].join('.'))
    this.x = x
    this.y = y
    this.points = [{ x: x, y: y }]
  }
  lineTo (x, y) {
    console.log([this.constructor.name, 'lineTo'].join('.'))
    this.lines.push({ x1: this.x, y1: this.y, x2: x, y2: y })
    this.x = x
    this.y = y
    this.points.push({ x: x, y: y })
  }
  arc (x, y, r, a, b, c) {
    console.log([this.constructor.name, 'arc'].join('.'))
  }
  arcTo(x1, y1, x2, y2, radius){
    console.log([this.constructor.name, 'arcTo', x1, y1, x2, y2, radius].join('.'))
  }
  closePath () {
    console.log([this.constructor.name, 'closePath'].join('.'))
    if(this.lines.length){
      this.lines.push({ x1: this.x, y1: this.y, x2: this.lines[0].x1, y2: this.lines[0].y1 })
    }
  }
  fill () {
    console.log([this.constructor.name, 'fill'].join('.'))
    var color = toRGB(this.fillStyle)

    var vx = this.points.map(e => e.x)
    var vy = this.points.map(e => e.y)
    SDL.FilledPolygonColor(this.handle ,vx, vy, color.r, color.g, color.b)
  }
  stroke () {
    console.log([this.constructor.name, 'stroke'].join('.'))
    var color = toRGB(this.strokeStyle)
    SDL.DrawLine(this.handle ,this.lines, color.r, color.g, color.b)
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
