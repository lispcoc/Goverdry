class Sprite3D extends SceneNode {
  constructor () {
    super()
    console.log([this.constructor.name, 'constructor'].join('.'))
  }
  removeChild (node) {
    console.log([this.constructor.name, 'removeChild'].join('.'))
    console.log(childNodes.length)
    this.childNodes = this.childNodes.filter(e => {
      return e != node
    })
    console.log(childNodes.length)
  }
}
