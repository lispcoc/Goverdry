soundhandle = 0

class DOMSoundClass {
  constructor () {
    console.log(this.constructor.name, 'constructor')
  }
  load (src, type, f_success, f_error) {
    console.log(
      [
        this.constructor.name,
        'load',
        src,
        type,
        f_success.name,
        f_error.name
      ].join('.')
    )

    var handle = -1
    handle = MIX.LoadMUS(src)
    if (handle < 0) {
      f_error(null)
    } else {
      f_success(new SDLSoundData(handle))
    }
    return null
  }
}

DOMSound = new DOMSoundClass()
