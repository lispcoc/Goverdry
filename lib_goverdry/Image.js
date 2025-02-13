var image_reminder = []

class Image {
  constructor () {
    console.log(this.constructor.name, 'constructor')
    console.log('[Todo]not implemented.')
    image_reminder.push(this)
    this.ready = false
    this.width = 1
    this.height = 1
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
  console.log(this.constructor.name, 'loadImage')
  console.log('[Todo]not implemented.')
  for (var e of image_reminder) {
    if (!e.ready && e.onload) {
      console.log(e.src)
      e.onerror()
      e.ready = true
    }
  }
}
