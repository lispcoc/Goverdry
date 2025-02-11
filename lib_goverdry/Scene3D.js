class SceneNode {
  constructor () {
    console.log([this.constructor.name, 'constructor'].join('.'))
    this.childNodes = []
  }
  addChild (a) {
    console.log(['[wip]', this.constructor.name, 'addChild'].join('.'))
    this.childNodes.push(a)
  }
  clear () {
    console.log(['[wip]', this.constructor.name, 'clear'].join('.'))
  }
}

class Scene3D extends SceneNode {
  constructor () {
    super()
    console.log([this.constructor.name, 'constructor'].join('.'))
    this.camera = new Camera()
  }
  setAmbientLight (a) {
    console.log([this.constructor.name, 'setAmbientLight'].join('.'))
  }
  getCamera () {
    return this.camera
  }
}
