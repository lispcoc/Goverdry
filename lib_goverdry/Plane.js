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
  }
  scale (x, y, z) {
    for (let i = 0; i < this.points.length; i++) {
      this.points[i][0] *= x
      this.points[i][1] *= y
      this.points[i][2] *= z
    }
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
    console.log(this.x, this.y, this.z)
    // Line
    ctx.beginPath()
    for (let i = 0; i < this.points.length; i++) {
      const z = this.z + this.points[i][2] + MP.CAMERA_Z
      var x =
        GameBody.width * 2 +
        ((this.x + this.points[i][0]) * GameBody.width) / 2 / z
      var y =
        GameBody.height * 2 -
        ((this.y + this.points[i][1]) * GameBody.width * 2) / z
      x = x >= 0 ? x : 0
      y = y >= 0 ? y : 0
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
  constructor () {}
}
