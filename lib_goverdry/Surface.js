class Surface extends SceneNode {
  constructor (width, height) {
    super()
    this.width = width
    this.height = height
    this.fillStyle = '#ffffff'
    this.handle = SDL.CreateRGBSurface(width, height)
    console.log(
      this.constructor.name,
      'constructor',
      this.handle,
      width,
      height
    )
    this.context = new SurfaceContext(this.handle)
    this._element = { src: 0 }
  }
  clone () {
    console.log(this.constructor.name, 'clone', this.handle)
    var r = new Surface(this.width, this.height)
    var bkup = r.handle
    var bkup2 = r.context
    Object.keys(this).forEach(k => {
      r[k] = this[k]
    })
    r.handle = bkup
    r.context = bkup2
    r.draw(this, 0, 0, this.width, this.height, 0, 0, this.width, this.height)
    // debug
    //IMG.SaveFile(this.handle, "test/" + this.handle + ".png")
    //IMG.SaveFile(r.handle, "test/" + r.handle + ".png")
    return r
  }
  clear () {
    SDL.LayerClear(this.handle)
  }
  draw (src_surface, x1, y1, w1, h1, x2, y2, w2, h2) {
    console.log(this.constructor.name, 'draw', this.handle, src_surface.handle)
    SDL.Copy(src_surface.handle, this.handle, x1, y1, w1, h1, x2, y2, w2, h2)
  }
}

class SurfaceContext {
  constructor (handle) {
    console.log(this.constructor.name, 'constructor')
    this.handle = handle
    this.x = -1
    this.y = -1
    this.points = []
    this.lines = []
    this.image_data = new ImageData()
  }
  getImageData (sx, sy, sw, sh, settings = {}) {
    console.log(this.constructor.name, 'getImageData', sx, sy, sw, sh)
    console.log('wip')
    return this.image_data
  }
  putImageData (imageData, dx, dy) {
    console.log(this.constructor.name, 'putImageData', imageData, dx, dy)
    console.log('wip')
    this.image_data = imageData
  }
  beginPath () {
    this.points = []
    this.lines = []
  }
  moveTo (x, y) {
    this.x = x
    this.y = y
    this.points = [{ x: x, y: y }]
  }
  lineTo (x, y) {
    this.lines.push({ x1: this.x, y1: this.y, x2: x, y2: y })
    this.x = x
    this.y = y
    this.points.push({ x: x, y: y })
  }
  arc (x, y, radius, startAngle, endAngle, counterclockwise = false) {
    console.log(
      this.constructor.name,
      'arc',
      x,
      y,
      radius,
      startAngle,
      endAngle,
      counterclockwise
    )
    console.log('[Todo]arc is wip.')
    if (!counterclockwise) {
      const x1 = x + radius * Math.cos(startAngle)
      const y1 = y - radius * Math.sin(startAngle)
      const x2 = x + radius * Math.cos(endAngle)
      const y2 = y - radius * Math.sin(endAngle)
      this.lines.push({ x1: this.x, y1: this.y, x2: x1, y2: y1 })
      this.lines.push({ x1: x1, y1: y1, x2: x2, y2: y2 })
      this.x = x2
      this.y = y2
      this.points.push({ x: x1, y: y1 })
      this.points.push({ x: x2, y: y2 })
    } else {
      const x1 = x + radius * Math.cos(startAngle)
      const y1 = y + radius * Math.sin(startAngle)
      const x2 = x + radius * Math.cos(endAngle)
      const y2 = y + radius * Math.sin(endAngle)
      this.lines.push({ x1: this.x, y1: this.y, x2: x1, y2: y1 })
      this.lines.push({ x1: x1, y1: y1, x2: x2, y2: y2 })
      this.x = x2
      this.y = y2
      this.points.push({ x: x1, y: y1 })
      this.points.push({ x: x2, y: y2 })
    }
  }
  arcTo (x1, y1, x2, y2, radius) {
    console.log(this.constructor.name, 'arcTo', x1, y1, x2, y2, radius)
    console.log('[Todo]arc is wip.')
    this.lines.push({ x1: this.x, y1: this.y, x2: x1, y2: y1 })
    this.x = x
    this.y = y
    this.points.push({ x: x, y: y })
  }
  closePath () {
    if (this.lines.length) {
      this.lines.push({
        x1: this.x,
        y1: this.y,
        x2: this.lines[0].x1,
        y2: this.lines[0].y1
      })
    }
  }
  fill () {
    var color = toRGB(this.fillStyle)

    var vx = this.points.map(e => e.x)
    var vy = this.points.map(e => e.y)
    SDL.FilledPolygonColor(this.handle, vx, vy, color.r, color.g, color.b)
  }
  fillTexture (img_handle, img_w, img_h, x, y, w = 0, h = 0) {
    var vx = this.points.map(e => e.x)
    var vy = this.points.map(e => e.y)

    SDL.FilledPolygonImage(
      this.handle,
      img_handle,
      0,
      0,
      img_w,
      img_h,
      x,
      y,
      w ? w : img_w,
      h ? h : img_h,
      vx,
      vy,
      255,
      255,
      255
    )
  }
  stroke () {
    var color = toRGB(this.strokeStyle)
    SDL.DrawLine(this.handle, this.lines, color.r, color.g, color.b)
  }
  fillRect (x, y, w, h) {
    var color = toRGB(this.fillStyle)
    SDL.FillRect(this.handle, x, y, w, h, color.r, color.g, color.b)
  }
  fillText (Text, x, y) {
    var color = toRGB(this.fillStyle)
    SDL.FillText(this.handle, Text, x, y, color.r, color.g, color.b)
  }
  drawImage (img, x, y, w = 0, h = 0) {
    SDL.DrawImage(
      this.handle,
      img.src,
      0,
      0,
      img.width,
      img.height,
      x,
      y,
      w ? w : img.width,
      h ? h : img.height
    )
  }
}
