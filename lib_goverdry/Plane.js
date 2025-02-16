function getFunctionName (func, argsType, args) {
  var re = /^function\s*(\w*)(\s*\(.*\))/
  var arr = re.exec(func.toString())

  if (argsType === 1) {
    //  get function name and dummy arguments
    return arr[1] + arr[2]
  } else if (argsType === 2) {
    // get function name and actual arguments
    var args = args.length === 1 ? [args[0]] : Array.apply(null, args)
    return arr[1] + '(' + args + ')'
  } else {
    // get only function name
    return arr[1]
  }
}

class Plane {
  constructor () {
    console.log(this.constructor.name, 'constructor')
    this.mesh = new PlaneMesh()
    this.x = 0
    this.y = 0
    this.z = 0
    this.points = [
      [1, 1, 0],
      [-1, 1, 0],
      [-1, -1, 0],
      [1, -1, 0]
    ]
    this._scaleX = 1
    this._scaleY = 1
    this._scaleZ = 1
  }
  scale (x, y, z) {
    for (let i = 0; i < this.points.length; i++) {
      this.points[i][0] *= x
      this.points[i][1] *= y
      this.points[i][2] *= z
    }
    this._scaleX *= x
    this._scaleY *= y
    this._scaleZ *= z
  }
  rotateYaw (r) {
    console.log(JSON.stringify(this.points))
    let rotator = [
      [Math.cos(r), 0, Math.sin(r)],
      [0, 1, 0],
      [Math.sin(r), 0, Math.cos(r)]
    ]
    for (let i = 0; i < this.points.length; i++) {
      this.points[i] = [
        rotator[0][0] * this.points[i][0] + rotator[0][1] * this.points[i][1] + rotator[0][2] * this.points[i][2],
        rotator[1][0] * this.points[i][0] + rotator[1][1] * this.points[i][1] + rotator[1][2] * this.points[i][2],
        rotator[2][0] * this.points[i][0] + rotator[2][1] * this.points[i][1] + rotator[2][2] * this.points[i][2],
      ]
    }
    console.log(JSON.stringify(this.points))
  }
  rotatePitch (r) {
    console.log(JSON.stringify(this.points))
    let rotator = [
      [1, 0, 0],
      [0, Math.cos(r), Math.sin(r)],
      [0, -Math.sin(r), Math.cos(r)]
    ]
    for (let i = 0; i < this.points.length; i++) {
      this.points[i] = [
        rotator[0][0] * this.points[i][0] + rotator[0][1] * this.points[i][1] + rotator[0][2] * this.points[i][2],
        rotator[1][0] * this.points[i][0] + rotator[1][1] * this.points[i][1] + rotator[1][2] * this.points[i][2],
        rotator[2][0] * this.points[i][0] + rotator[2][1] * this.points[i][1] + rotator[2][2] * this.points[i][2],
      ]
    }
    console.log(JSON.stringify(this.points))
  }
  rotateRoll (r) {
    console.log(JSON.stringify(this.points))
    let rotator = [
      [Math.cos(r), Math.sin(r), 0],
      [-Math.sin(r), Math.cos(r), 0],
      [0, 0, 1]
    ]
    for (let i = 0; i < this.points.length; i++) {
      this.points[i] = [
        rotator[0][0] * this.points[i][0] + rotator[0][1] * this.points[i][1] + rotator[0][2] * this.points[i][2],
        rotator[1][0] * this.points[i][0] + rotator[1][1] * this.points[i][1] + rotator[1][2] * this.points[i][2],
        rotator[2][0] * this.points[i][0] + rotator[2][1] * this.points[i][1] + rotator[2][2] * this.points[i][2],
      ]
    }
    console.log(JSON.stringify(this.points))
  }
  update (ctx) {
    const offsetX = (GameBody.width * 3) / 2
    const offsetY = (GameBody.height * 3) / 2
    ctx.beginPath()
    var minX = -1,
      maxX = -1,
      minY = -1,
      maxY = -1
    for (let i = 0; i < this.points.length; i++) {
      const z = this.z + this.points[i][2] / 2
      var x =
        ((this.x + this.points[i][0] / 2) * MP.CENTER_FRAME_WIDTH) / (2.4 - z)
      var y =
        ((this.y - 0.2 + this.points[i][1] / 2) * MP.CENTER_FRAME_HEIGHT) /
        (2.4 - z)
      x += offsetX
      y += offsetY
      if (i == 0) {
        ctx.moveTo(x, y)
      } else {
        ctx.lineTo(x, y)
      }

      if (minX < 0) {
        minX = x
      }
      if (minY < 0) {
        minY = y
      }
      minX = Math.min(minX, x)
      minY = Math.min(minY, y)
      maxX = Math.max(maxX, x)
      maxY = Math.max(maxY, y)
    }
    ctx.closePath()
    
    ctx.fillTexture(
      this.mesh.texture.src.handle,
      this.mesh.texture.src.width,
      this.mesh.texture.src.height,
      minX,
      minY,
      maxX - minX,
      maxY - minY
    )
    // debug
    // ctx.fillStyle ="black"
    // ctx.fillText("" + this.mesh.texture.src.handle, (minX + maxX) / 2, (minY + maxY) / 2)
  }
  visible () {
    return this.mesh.texture.src != null
  }
}

class PlaneXZ extends Plane {
  constructor () {
    super()
    this.mesh = new PlaneMesh()
    this.x = 0
    this.y = 0
    this.z = 0
    this.points = [
      [1, 0, 1],
      [-1, 0, 1],
      [-1, 0, -1],
      [1, 0, -1]
    ]
  }
}

class PlaneMesh {
  constructor () {
    this.texture = new PlaneTexture()
  }
}

class PlaneTexture {
  constructor () {
    this.src = null
  }
}
