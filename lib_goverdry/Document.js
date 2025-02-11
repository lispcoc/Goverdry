class DocumentElement {
  constructor () {
    console.log([this.constructor.name, 'constructor'].join('.'))
    this.style = {}
  }
  setAttribute (a, b) {
    console.log([this.constructor.name, 'setAttribute', a, b].join('.'))
  }
  addEventListener (a, b) {
    console.log([this.constructor.name, 'addEventListener', a, b].join('.'))
  }
}
