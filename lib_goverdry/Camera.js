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
      }
    }

    return
    // 3D (test)
    if (this.handle == null) {
      this.handle = SDL.CreateRGBSurface(
        GameBody.width * 4,
        GameBody.height * 4
      )
      this.context = new SurfaceContext(this.handle)
      this.context.fillStyle = '#00ff00'
      this.context.strokeStyle = '#ffffff'
    } else {
      console.log(this.handle)

      this.clear()
      MP.Sprite3D.childNodes.sort((a, b) => {
        if (a.z > b.z) {
          return -1
        } else if (a.z < b.z) {
          return 1
        }
        return 0
      })
      MP.Sprite3D.childNodes.forEach(e => {
        e.update(this.context)
      })
      /*
      for (let i = 0; i < MP.Ladder3D[0].length; i++) {
        if (isObject(MP.Ladder3D[0][i][0])) {
          console.log(JSON.stringify(MP.Ladder3D[0][i][0]))
          MP.Ladder3D[0][i][0].update(this.context)
        }
      }
      for (let i = 0; i < MP.Wall3D[0].length; i++) {
        if (isObject(MP.Wall3D[0][i][0])) {
          console.log(JSON.stringify(MP.Wall3D[0][i][0]))
          MP.Wall3D[0][i][0].update(this.context)
        }
      }
      for (let i = 0; i < MP.Floor3D[0].length; i++) {
        for (let j = 0; j < MP.Floor3D[0][i].length; j++) {
          for (let k = 0; k < MP.Floor3D[0][i][j].length; k++) {
            if (isObject(MP.Floor3D[0][i][j][k][0])) {
              console.log(JSON.stringify(MP.Floor3D[0][i][j][k][0]))
              MP.Floor3D[0][i][j][k][0].update(this.context)
            }
          }
        }
      }
      */
    }
    this.drawToWindow()
  }
  drawToWindow () {
    SDL.DrawSpriteToWindow(
      this.handle,
      GameBody.width * 1.5,
      GameBody.height * 1.5,
      GameBody.width,
      GameBody.height
    )
  }
  clear () {
    SDL.LayerClear(this.handle)
  }
}
