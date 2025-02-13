class Sprite3D extends SceneNode {
  constructor () {
    super()
    console.log(this.constructor.name, 'constructor')
  }
  removeChild (node) {
    console.log(this.constructor.name, 'removeChild')
    console.log(this.childNodes.length)
    this.childNodes = this.childNodes.filter(e => {
      return e != node
    })
    console.log(this.childNodes.length)
  }
}
