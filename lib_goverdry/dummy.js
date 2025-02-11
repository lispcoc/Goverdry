GoverdryTodo = false
GoverdryPatch = true

class DummyEnchant {
  constructor () {
    this.Sound = {
      enabledInMobileSafari: true
    }
  }
}

class DummyWebFont {
  /*
  WebFont.load({
  classes: false,
  custom: { families: ['GameFont'] },
  loading: function () {},
  active: function () {
    TempVariable['loadWebFont'] = true
  },
  inactive: function () {
    TempVariable['loadWebFont'] = true
  }
})
  */
  constructor (a) {}
  load (a) {
    console.log(['[wip]', this.constructor.name, 'load'].join('.'))
    a.active()
  }
}
WebFont = new DummyWebFont({})

class DummyNavigator {
  constructor () {
    this.userAgent = {
      toLowerCase: () => {
        return 'linux'
      }
    }
  }
  load () {}
}
navigator = new DummyNavigator()

function printDebugMessage (a) {
  console.log(a)
}
