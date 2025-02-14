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
      this.points[i] = math.multiply(rotator, this.points[i])
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
      this.points[i] = math.multiply(rotator, this.points[i])
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
      this.points[i] = math.multiply(rotator, this.points[i])
    }
    console.log(JSON.stringify(this.points))
  }
  update (ctx) {
    console.log(this.type)
    console.log(
      this.x,
      this.y,
      this.z,
      this._scaleX,
      this._scaleY,
      this._scaleZ
    )
    // Line
    ctx.beginPath()
    for (let i = 0; i < this.points.length; i++) {
      const z = this.z + this.points[i][2] / 2
      var x =
        ((this.x + this.points[i][0] / 2) * MP.CENTER_FRAME_WIDTH) /
        (2.4 - z)
      var y =
        ((this.y -0.2 + this.points[i][1] / 2) * MP.CENTER_FRAME_HEIGHT) /
        (2.4 - z)
      x += (GameBody.width * 3) / 2
      y += (GameBody.height * 3) / 2
      if (i == 0) {
        ctx.moveTo(x, y)
      } else {
        ctx.lineTo(x, y)
      }
    }
    ctx.closePath()
    ctx.fill()
    ctx.stroke()
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
    math.multiply(this.points, 0.5)
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
