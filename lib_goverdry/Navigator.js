class Navigator {
  constructor () {
    console.log(this.constructor.name, 'constructor')
    this.userAgent = {
      toLowerCase: () => {
        return 'linux'
      }
    }
    this.gamepad = new Gamepad()
  }

  load () {
    console.log(this.constructor.name, 'load')
  }

  getGamepads () {
    console.log(this.constructor.name, 'getGamepads')
    return [this.gamepad]
  }
}

navigator = new Navigator()
