class Camera {
  constructor (root) {
    console.log([this.constructor.name, 'constructor'].join('.'))
    this.root = root
  }
  update () {
    console.log([this.constructor.name, 'update'].join('.'))
    for (var n of this.root.childNodes) {
      if (n.constructor.name == 'Sprite') {
        n.update()
      }
    }
  }
}
