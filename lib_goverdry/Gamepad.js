class Gamepad {
  constructor () {
    console.log(this.constructor.name, 'constructor')
    this.buttons = []
    for (var i = 0; i < 20; i++) {
      this.buttons.push({ pressed: false })
    }
    this.axes = []
    for (var i = 0; i < 4; i++) {
      this.axes.push(0)
    }
  }
  pressButton (i) {
    console.log(this.constructor.name, 'pressButton', i)
    this.buttons[i].pressed = true
  }
  releaseButton (i) {
    this.buttons[i].pressed = false
  }
}
