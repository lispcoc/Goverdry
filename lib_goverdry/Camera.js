class Camera {
  constructor (root) {
    console.log(this.constructor.name, 'constructor')
    this.root = root

    this.fillStyle = '#ffffff'
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
      this.handle = SDL.CreateRGBSurface(GameBody.width, GameBody.height)
      this.context = new SurfaceContext(this.handle)
    } else {
      console.log(this.handle)
      for (let i = 0; i < MP.Floor3D.length; i++) {
        for (let j = 0; j < MP.Floor3D[i].length; j++) {
          for (let k = 0; k < MP.Floor3D[i][j].length; k++) {
            for (let l = 0; l < MP.Floor3D[i][j][k].length; l++) {
              for (let m = 0; m < MP.Floor3D[i][j][k][l].length; m++) {
                MP.Floor3D[i][j][k][l][m].update(this.context)
              }
            }
          }
        }
      }
    }
    this.drawToWindow ()
  }
  drawToWindow () {
    SDL.DrawSpriteToWindow(this.handle)
  }
}
