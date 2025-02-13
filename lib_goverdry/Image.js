var image_reminder = []

class Image {
  constructor () {
    console.log([this.constructor.name, 'constructor'].join('.'))
    image_reminder.push(this)
    this.ready = false
    this.width = 1
    this.height = 1
  }
}

class ImageData {
  constructor () {
    console.log([this.constructor.name, 'constructor'].join('.'))
    this.data = []
  }
}

function loadImage() {
  for (var e of image_reminder){
    if(!e.ready && e.onload) {
      console.log(e.src)
      e.onload()
      e.ready = true
    }
  }
}

