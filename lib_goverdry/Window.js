class Window {
  constructor () {
    console.log(this.constructor.name, 'constructor')
    this.events = []
  }
  addEventListener (id, listener) {
    console.log(this.constructor.name, 'addEventListener', id, listener.name)
    this.events.push({ id: id, listener: listener })
  }
  emit (event) {
    console.log(this.constructor.name, 'emit', event)
    for (const e of this.events) {
      if (event.id == e.id) {
        e.listener(event)
      }
    }
  }
  focus () {
    console.log(this.constructor.name, 'focus')
  }
  set onload (f) {
    this._onload = f
  }
  get onload () {
    return this._onload
  }
}

var window = new Window()
