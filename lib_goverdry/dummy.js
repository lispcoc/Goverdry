GoverdryTodo = false
GoverdryPatch = true
GoverdrySaveTweak = true

class DummyEnchant {
  constructor () {
    this.Sound = {
      enabledInMobileSafari: true
    }
  }
}

class DummyWebFont {
  constructor (a) {}
  load (a) {
    console.log('[wip]', this.constructor.name, 'load')
    a.active()
  }
}
WebFont = new DummyWebFont({})

function printDebugMessage (a) {
  console.log(a)
}
