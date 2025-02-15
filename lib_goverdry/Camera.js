function isObject (value) {
  return value !== null && typeof value === 'object' && !Array.isArray(value)
}

class Camera {
  constructor (root) {
    console.log(this.constructor.name, 'constructor')
    this.root = root

    this.handle = null
    this.context = null
  }
  update () {
    for (var n of this.root.childNodes) {
      if (n.constructor.name == 'Sprite') {
        n.update()
        if (n.image.handle == MP.MainSurface.handle) {
          this.update3D()
        }
      }
    }
  }
  update3D () {
    // 3D (test)
    if (this.handle == null) {
      this.handle = SDL.CreateRGBSurface(
        GameBody.width * 3,
        GameBody.height * 3
      )
      this.context = new SurfaceContext(this.handle)
    } else {
      console.log(this.handle)

      this.clear()
      MP.Sprite3D.childNodes.sort((a, b) => {
        if (a.z < b.z) {
          return -1
        } else if (a.z > b.z) {
          return 1
        }
        return 0
      })
      MP.Sprite3D.childNodes.forEach(e => {
        if (e.visible()) {
          this.context.fillStyle = '#ffffff'
          this.context.strokeStyle = '#ffffff'
          e.update(this.context)
        }
      })
    }
    this.drawToWindow()
  }
  drawToWindow () {
    SDL.DrawSpriteToWindow(
      this.handle,
      GameBody.width + MP.CENTER_FRAME_X,
      GameBody.height + MP.CENTER_FRAME_Y,
      MP.CENTER_FRAME_WIDTH,
      MP.CENTER_FRAME_HEIGHT,
      MP.CENTER_FRAME_X,
      MP.CENTER_FRAME_Y,
      MP.CENTER_FRAME_WIDTH,
      MP.CENTER_FRAME_HEIGHT
    )
  }
  clear () {
    SDL.LayerClear(this.handle)
  }
}
