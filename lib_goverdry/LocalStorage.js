class LocalStorage {
  getItem (a) {
    console.log([this.constructor.name, 'getItem', a].join('.'))
  }
  removeItem (a) {
    console.log([this.constructor.name, 'removeItem', a].join('.'))
  }
  removeItem (a) {
    console.log([this.constructor.name, 'setItem', a].join('.'))
  }
}

localStorage = new LocalStorage()
