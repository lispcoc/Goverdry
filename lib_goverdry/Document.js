class DocumentElement {
  constructor () {
    console.log(this.constructor.name, 'constructor')
    this.style = {}
    this.events = []
    this.children = []
  }
  setAttribute (a, b) {
    console.log(this.constructor.name, 'setAttribute', a, b)
  }
  addEventListener (id, listener) {
    console.log(this.constructor.name, 'addEventListener', id, listener.name)
    this.events.push({ id: id, listener: listener })
  }
  removeEventListener (id, listener) {
    console.log(this.constructor.name, 'removeEventListener', id, listener.name)
    this.events = this.events.filter(e => {
      return e.listener != listener
    })
  }
  emit (event) {
    for (const c of this.children) {
      c.emit(event)
    }
    for (const e of this.events) {
      if (event.id == e.id) {
        e.listener(event)
      }
    }
  }
  focus () {
    this.unfocusOther(this)
  }
  unfocus () {}
  unfocusOther (o) {
    if (o != this) {
      this.unfocus()
    }
    for (const c of this.children) {
      c.unfocusOther(this)
    }
  }
}

class Document extends DocumentElement {
  constructor () {
    super()
    console.log(this.constructor.name, 'constructor')
    this.location = {}
    this.hostname = 'fake'
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
  createElement (type) {
    console.log(this.constructor.name, 'createElement')
    var c
    if (type == 'input') {
      c = new Input()
    } else {
      c = new DocumentElement()
    }
    this.children.push(c)
    return c
  }
}

try {
  var document = new Document()
} catch (e) {
  console.log(e)
  console.log(e.stack)
  throw e
}
