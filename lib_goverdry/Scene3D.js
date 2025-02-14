class SceneNode {
  constructor () {
    console.log(this.constructor.name, 'constructor')
    this.childNodes = []
  }
  addChild (a) {
    console.log(this.constructor.name, 'addChild')
    this.childNodes.push(a)
  }
  clear () {
    console.log('[wip]', this.constructor.name, 'clear')
  }
  update () {
    console.log(this.constructor.name, 'update')
  }
  clone () {
    console.log(this.constructor.name, 'clone')
    let r = JSON.parse(JSON.stringify(this))
    r.childNodes = this.childNodes
    return r
  }
}

class Scene3D extends SceneNode {
  constructor () {
    super()
    console.log(this.constructor.name, 'constructor')
    this.camera = new Camera(this)
  }
  setAmbientLight (a) {
    console.log(this.constructor.name, 'setAmbientLight')
  }
  getCamera () {
    return this.camera
  }
}
