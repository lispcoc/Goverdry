class SDLSoundData {
  constructor (handle) {
    console.log([this.constructor.name, 'constructor', handle].join('.'))
    this.handle = handle
  }
  clone() {
    console.log([this.constructor.name, 'clone'].join('.'))
    return this
  }
  play() {
    console.log([this.constructor.name, 'play'].join('.'))
    MIX.PlayChannel(this.handle)
  }
}
