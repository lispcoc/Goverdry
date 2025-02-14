class Document {
  constructor () {
    console.log(this.constructor.name, 'constructor')
    this.location = {}
    this.hostname = 'fake'
    this.events = []
  }
  addEventListener (id, listener) {
    console.log(this.constructor.name, 'addEventListener', id, listener.name)
    this.events.push({ id: id, listener: listener })
  }
  getElementById (id) {
    console.log(this.constructor.name, 'getElementById', id)
    if (id == 'hidden') {
      return {
        innerHTML: [
          'SaveName:Javardry',
          'Size:800',
          'Align:center',
          'Controller:on',
          'LoadingScreen:1',
          'NotLoadSaveData:off',
          'NoSoundAllowed:off',
          'GraphicPrefetch:on',
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
    console.log(this.constructor.name, 'createElement')
    return new DocumentElement()
  }
  emit (event) {
    console.log(this.constructor.name, 'emit', event)
    for (const e of this.events) {
      if (event.id == e.id) {
        e.listener(event)
      }
    }
  }
}

class DocumentElement {
  constructor () {
    console.log(this.constructor.name, 'constructor')
    this.style = {}
    this.events = []
  }
  setAttribute (a, b) {
    console.log(this.constructor.name, 'setAttribute', a, b)
  }
  addEventListener (id, listener) {
    console.log(this.constructor.name, 'addEventListener', id, listener.name)
    this.events.push({ id: id, listener: listener })
  }
  emit (event) {
    for (const e of this.events) {
      if (event.id == e.id) {
        e.listener(event)
      }
    }
  }
}

var document = new Document()
