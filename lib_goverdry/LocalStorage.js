class LocalStorage {
  getItem (dataName) {
    console.log([this.constructor.name, 'getItem', dataName].join('.'))
    var dataStr = IO.readTextFile("./save/" + dataName)
    return dataStr
  }
  setItem (dataName, dataStr) {
    console.log([this.constructor.name, 'setItem', dataName, dataStr].join('.'))
    IO.writeTextFile("./save/" + dataName, dataStr)
  }
  removeItem (dataName) {
    console.log([this.constructor.name, 'removeItem', dataName].join('.'))
    IO.writeTextFile("./save/" + dataName, "")
  }
}

localStorage = new LocalStorage()
