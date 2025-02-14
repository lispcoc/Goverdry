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
    this.points = math.multiply(this.points, 0.5)
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
    this.x = 200
    this.y = 200
    this.scale(16, 16, 1)
    // Line
    ctx.beginPath()
    ctx.moveTo(this.x + this.points[0][0], this.y + this.points[0][1])
    for (let i = 1; i < this.points.length; i++) {
      const x = this.x + this.points[i][0]
      const y = this.y + this.points[i][1]
      ctx.lineTo(x, y)
    }
    ctx.closePath()
    ctx.fill()
    ctx.stroke()
    this.scale(1/16, 1/16, 1)
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
    this.points = math.multiply(this.points, 0.5)
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
