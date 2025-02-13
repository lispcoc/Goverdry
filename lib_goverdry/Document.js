class Document {
  constructor () {
    console.log([this.constructor.name, 'constructor'].join('.'))
    this.location = {}
    this.hostname = 'fake'
  }
  addEventListener (event) {
    console.log([this.constructor.name, 'addEventListener', event].join('.'))
  }
  getElementById (id) {
    console.log([this.constructor.name, 'getElementById', id].join('.'))
    if (id == 'hidden') {
      return {
        innerHTML: [
          'SaveName:Javardry',
          'Size:800',
          'Align:center',
          'Controller:off',
          'LoadingScreen:1',
          'NotLoadSaveData:off',
          'NoSoundAllowed:off',
          'GraphicPrefetch:off',
          'SoundPrefetch:off',
          'GraphicLimit:off',
          'GraphicCache:on',
          'SoundCache:on'
        ].join(',')
      }
    }
    return { innerHTML: ':' }
  }
  createElement () {
    console.log([this.constructor.name, 'createElement'].join('.'))
    return new DocumentElement()
  }
}

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

var document = new Document()
