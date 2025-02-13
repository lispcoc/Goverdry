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
    this.mesh = new PlaneMesh()
    this.x = 0
    this.y = 0
    this.z = 0
  }
  scale () {
    console.log(this.constructor.name, 'scale')
  }
  rotateYaw (a) {
    console.log(this.constructor.name, 'rotateYaw')
  }
  rotatePitch () {
    console.log(this.constructor.name, 'rotatePitch')
  }
  rotateRoll () {
    console.log(this.constructor.name, 'rotateRoll')
  }
}

class PlaneXZ extends Plane {}

class PlaneMesh {
  constructor () {
    this.texture = new PlaneTexture()
  }
}

class PlaneTexture {
  constructor () {}
}
