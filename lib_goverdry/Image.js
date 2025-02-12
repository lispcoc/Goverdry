var image_reminder = []

class Image {
  constructor () {
    console.log([this.constructor.name, 'constructor'].join('.'))
    image_reminder.push(this)
    this.ready = false
  }
}

function loadImage() {
  for (var e of image_reminder){
    if(!e.ready && e.onload) {
      console.log(e.src)
      e.onload()
    }
  }
}

