var image_reminder = []

class Image {
  constructor () {
    console.log(this.constructor.name, 'constructor')
    console.log('[Todo]not implemented.')
    image_reminder.push(this)
    this.ready = false
    this.width = 1
    this.height = 1
    this.src = null
    this.onload = null
    this.onerror = null
  }
}

class ImageData {
  constructor () {
    console.log(this.constructor.name, 'constructor')
    console.log('[Todo]not implemented.')
    this.data = []
  }
}

function loadImage () {
  for (var e of image_reminder) {
    if (!e.ready && e.onload) {
      console.log("loadImage")
      console.log(e.src)
      let ret = IMG.Load(e.src)
      if(ret == null){
        e.onerror()
      } else {
        console.log("load success.")
        console.log(ret.w, ret.h)
        e.width = ret.w
        e.height = ret.h
        e.onload()
      }
      e.ready = true
    }
  }
}
