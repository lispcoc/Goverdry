class Camera {
  constructor (root) {
    console.log(this.constructor.name, 'constructor')
    this.root = root
  }
  update () {
    console.log(this.constructor.name, 'update')
    for (var n of this.root.childNodes) {
      if (n.constructor.name == 'Sprite') {
        n.update()
      }
    }
  }
}
