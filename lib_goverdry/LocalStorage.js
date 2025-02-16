class LocalStorage {
  constructor() {
    this.d = new Date()
    this.lastSaveTime = this.d.getTime()
    this.SaveInterval = 10 * 60 * 1000
  }
  getItem (dataName) {
    console.log(this.constructor.name, 'getItem', dataName)
    var dataStr = IO.readTextFile("./save/" + dataName)
    return dataStr
  }
  setItem (dataName, dataStr) {
    console.log(this.constructor.name, 'setItem', dataName)
    IO.writeTextFile("./save/" + dataName, dataStr)
    this.lastSaveTime = this.d.getTime()
  }
  removeItem (dataName) {
    console.log(this.constructor.name, 'removeItem', dataName)
    IO.writeTextFile("./save/" + dataName, "")
  }
  inSaveInterval () {
    return (this.d.getTime() - this.lastSaveTime) < this.SaveInterval
  }
}

localStorage = new LocalStorage()
